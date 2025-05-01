package chains

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: Add comprehensive tests for CheckRPCs using httptest for HTTP and potentially custom WS server/mocking library.

func setupHTTPServer(t *testing.T, handler func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(handler))
	t.Cleanup(ts.Close)
	return ts
}

func setupWSServer(t *testing.T, handler func(conn *websocket.Conn)) (*httptest.Server, string) {
	t.Helper()
	upgrader := websocket.Upgrader{}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Logf("ws upgrade error: %v", err)
			return
		}
		defer conn.Close()
		handler(conn)
	}))
	wsURL := "ws" + s.URL[len("http"):]
	t.Cleanup(s.Close)
	return s, wsURL
}

func TestCheckRPCs_Basic(t *testing.T) {
	httpServer := setupHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		var req jsonRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		require.Equal(t, "eth_blockNumber", req.Method)

		resp := jsonRPCResponse{
			Version: "2.0",
			ID:      req.ID,
			Result:  json.RawMessage(`"0x123abc"`),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	_, wsServerURL := setupWSServer(t, func(conn *websocket.Conn) {
		var req jsonRPCRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			return // Ошибки соединения обрабатываются в тесте
		}
		require.Equal(t, "eth_blockNumber", req.Method)

		resp := jsonRPCResponse{
			Version: "2.0",
			ID:      req.ID,
			Result:  json.RawMessage(`"0x456def"`),
		}
		_ = conn.WriteJSON(resp)
	})

	mockChainID := big.NewInt(7777)
	mockChain := Chain{
		ID:   mockChainID,
		Name: "Mock RPC Test Chain",
		RPCUrls: map[string]RpcTarget{
			"default": {
				Http:      []string{httpServer.URL, "http://invalid-http-url.test"},
				WebSocket: []string{wsServerURL, "ws://invalid-ws-url.test"},
			},
		},
	}
	RegisterChain(mockChain)

	opts := DefaultCheckOptions()
	opts.TimeoutPerCheck = 2 * time.Second

	statuses, err := CheckRPCs(context.Background(), mockChainID, opts)
	require.NoError(t, err)
	require.Len(t, statuses, 4, "Should have checked 4 URLs")

	var httpOK, httpFail, wsOK, wsFail int
	blockNumHttp := big.NewInt(0)
	blockNumHttp.SetString("123abc", 16)
	blockNumWs := big.NewInt(0)
	blockNumWs.SetString("456def", 16)

	for _, s := range statuses {
		if s.URL == httpServer.URL {
			assert.True(t, s.IsHTTP)
			assert.True(t, s.IsAvailable)
			assert.NoError(t, s.Error)
			assert.Less(t, s.Latency, opts.TimeoutPerCheck)
			assert.NotNil(t, s.BlockNumber)
			if s.BlockNumber != nil {
				assert.Equal(t, 0, s.BlockNumber.Cmp(blockNumHttp))
			}
			httpOK++
		} else if s.URL == "http://invalid-http-url.test" {
			assert.True(t, s.IsHTTP)
			assert.False(t, s.IsAvailable)
			assert.Error(t, s.Error)
			httpFail++
		} else if s.URL == wsServerURL {
			assert.True(t, s.IsWebSocket)
			assert.True(t, s.IsAvailable)
			assert.NoError(t, s.Error)
			assert.Less(t, s.Latency, opts.TimeoutPerCheck)
			assert.NotNil(t, s.BlockNumber)
			if s.BlockNumber != nil {
				assert.Equal(t, 0, s.BlockNumber.Cmp(blockNumWs))
			}
			wsOK++
		} else if s.URL == "ws://invalid-ws-url.test" {
			assert.True(t, s.IsWebSocket)
			assert.False(t, s.IsAvailable)
			assert.Error(t, s.Error)
			wsFail++
		}
	}

	assert.Equal(t, 1, httpOK, "Expected 1 successful HTTP check")
	assert.Equal(t, 1, httpFail, "Expected 1 failed HTTP check")
	assert.Equal(t, 1, wsOK, "Expected 1 successful WS check")
	assert.Equal(t, 1, wsFail, "Expected 1 failed WS check")
}

var originalRPCChecker = rpcChecker

func mockCheckRPCs(results []RPCStatus, expectedIdentifier any, expectedOpts CheckRPCOptions) func(ctx context.Context, identifier any, opts CheckRPCOptions) ([]RPCStatus, error) {
	return func(ctx context.Context, identifier any, opts CheckRPCOptions) ([]RPCStatus, error) {
		if fmt.Sprintf("%v", identifier) != fmt.Sprintf("%v", expectedIdentifier) {
			return nil, fmt.Errorf("mockCheckRPCs: unexpected identifier: got %v, want %v", identifier, expectedIdentifier)
		}
		return results, nil
	}
}

func TestGetFastestRPC(t *testing.T) {
	testID := big.NewInt(8888)
	RegisterChain(Chain{ID: testID, Name: "Fastest Test"})

	t.Cleanup(func() {
		rpcChecker = originalRPCChecker
		rpcCheckCache = make(map[rpcCacheKey]rpcCacheEntry)
	})

	status1 := RPCStatus{URL: "http://fast.com", IsHTTP: true, IsAvailable: true, Latency: 100 * time.Millisecond}
	status2 := RPCStatus{URL: "http://slow.com", IsHTTP: true, IsAvailable: true, Latency: 500 * time.Millisecond}
	status3 := RPCStatus{URL: "ws://fast.com", IsWebSocket: true, IsAvailable: true, Latency: 150 * time.Millisecond}
	status4 := RPCStatus{URL: "http://unavailable.com", IsHTTP: true, IsAvailable: false, Error: fmt.Errorf("timeout")}
	status5 := RPCStatus{URL: "ws://unavailable.com", IsWebSocket: true, IsAvailable: false, Error: fmt.Errorf("dial failed")}

	tests := []struct {
		name           string
		mockResults    []RPCStatus
		criteria       RPCCriteria
		wantURL        string
		wantErr        bool
		checkCacheHits int
	}{
		{
			name:           "Select fastest HTTP",
			mockResults:    []RPCStatus{status1, status2, status3, status4, status5},
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: false},
			wantURL:        "http://fast.com",
			wantErr:        false,
			checkCacheHits: 1,
		},
		{
			name:           "Select fastest WS",
			mockResults:    []RPCStatus{status1, status2, status3, status4, status5},
			criteria:       RPCCriteria{AllowHTTP: false, AllowWebSocket: true},
			wantURL:        "ws://fast.com",
			wantErr:        false,
			checkCacheHits: 1,
		},
		{
			name:           "Select fastest overall (HTTP is faster)",
			mockResults:    []RPCStatus{status1, status2, status3, status4, status5},
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: true},
			wantURL:        "http://fast.com",
			wantErr:        false,
			checkCacheHits: 1,
		},
		{
			name:           "Select fastest overall (WS is faster)",
			mockResults:    []RPCStatus{status2, status3, status4, status5},
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: true},
			wantURL:        "ws://fast.com",
			wantErr:        false,
			checkCacheHits: 1,
		},
		{
			name:           "No available matching criteria",
			mockResults:    []RPCStatus{status4, status5},
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: true},
			wantURL:        "",
			wantErr:        true,
			checkCacheHits: 0,
		},
		{
			name:           "Only unavailable HTTP matches criteria",
			mockResults:    []RPCStatus{status3, status4, status5},
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: false},
			wantURL:        "",
			wantErr:        true,
			checkCacheHits: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedOpts := CheckRPCOptions{
				TimeoutPerCheck: DefaultCheckOptions().TimeoutPerCheck,
				CheckHTTP:       tt.criteria.AllowHTTP,
				CheckWebSocket:  tt.criteria.AllowWebSocket,
				Providers:       tt.criteria.Providers,
			}
			rpcChecker = mockCheckRPCs(tt.mockResults, testID, expectedOpts)
			rpcCheckCache = make(map[rpcCacheKey]rpcCacheEntry)

			gotURL, err := GetFastestRPC(context.Background(), testID, tt.criteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFastestRPC() first call error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("GetFastestRPC() first call gotURL = %v, want %v", gotURL, tt.wantURL)
			}

			hitCount := 0
			rpcChecker = func(ctx context.Context, identifier any, opts CheckRPCOptions) ([]RPCStatus, error) {
				hitCount++
				return tt.mockResults, nil
			}

			gotURL2, err2 := GetFastestRPC(context.Background(), testID, tt.criteria)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("GetFastestRPC() second call error = %v, wantErr %v", err2, tt.wantErr)
			}
			if gotURL2 != tt.wantURL {
				t.Errorf("GetFastestRPC() second call gotURL = %v, want %v", gotURL2, tt.wantURL)
			}

			if tt.checkCacheHits > 0 {
				assert.Equal(t, 0, hitCount, "CheckRPCs should not have been called on the second run (cache hit expected)")
			} else {
				assert.GreaterOrEqual(t, hitCount, 1, "CheckRPCs should have been called on the second run (cache miss expected)")
			}
		})
	}
}
