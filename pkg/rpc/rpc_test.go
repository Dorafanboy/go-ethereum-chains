package rpc_test

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	chainstypes "go-ethereum-chains/internal/types"
	"go-ethereum-chains/pkg/registry"
	"go-ethereum-chains/pkg/rpc"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCheckRPCs_Basic performs a basic test of the CheckRPCs function using mock servers.
func TestCheckRPCs_Basic(t *testing.T) {
	httpServer := setupHTTPServer(t, func(w http.ResponseWriter, r *http.Request) {
		var req chainstypes.JsonRPCRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		require.NoError(t, err)
		require.Equal(t, "eth_blockNumber", req.Method)

		resp := chainstypes.JsonRPCResponse{
			Version: "2.0",
			ID:      req.ID,
			Result:  json.RawMessage(`"0x123abc"`),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	_, wsServerURL := setupWSServer(t, func(conn *websocket.Conn) {
		var req chainstypes.JsonRPCRequest
		err := conn.ReadJSON(&req)
		if err != nil {
			// Log or handle? Test might fail if conn closes unexpectedly.
			t.Logf("ws read error: %v", err)
			return
		}
		require.Equal(t, "eth_blockNumber", req.Method)

		resp := chainstypes.JsonRPCResponse{
			Version: "2.0",
			ID:      req.ID,
			Result:  json.RawMessage(`"0x456def"`),
		}
		err = conn.WriteJSON(resp)
		if err != nil {
			t.Logf("ws write error: %v", err)
		}
	})

	mockChainID := big.NewInt(7777)
	mockChain := chainstypes.Chain{
		ID:   mockChainID,
		Name: "Mock RPC Test Chain",
		RPCUrls: map[string]chainstypes.RpcTarget{
			"default": {
				Http:      []string{httpServer.URL, "http://invalid-http-url.test"},
				WebSocket: []string{wsServerURL, "ws://invalid-ws-url.test"},
			},
		},
	}
	registry.RegisterChain(mockChain) // Use registry package to register

	opts := rpc.DefaultCheckOptions() // Use rpc package
	opts.TimeoutPerCheck = 2 * time.Second

	statuses, err := rpc.CheckRPCs(context.Background(), mockChainID, opts) // Use rpc package
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
