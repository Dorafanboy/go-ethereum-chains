package chains

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var ErrChainNotFound = errors.New("chain not found")

// RPCStatus holds the result of checking a single RPC endpoint.
type RPCStatus struct {
	URL         string
	IsHTTP      bool
	IsWebSocket bool
	IsAvailable bool
	Latency     time.Duration
	BlockNumber *big.Int
	Error       error
}

// CheckRPCOptions defines parameters for checking RPC endpoints.
type CheckRPCOptions struct {
	TimeoutPerCheck time.Duration
	CheckHTTP       bool
	CheckWebSocket  bool
	Providers       []string
}

// RPCCriteria defines criteria for selecting an RPC endpoint.
type RPCCriteria struct {
	AllowHTTP      bool
	AllowWebSocket bool
	Providers      []string
}

// DefaultCheckOptions returns default options for CheckRPCs.
func DefaultCheckOptions() CheckRPCOptions {
	return CheckRPCOptions{
		TimeoutPerCheck: 5 * time.Second,
		CheckHTTP:       true,
		CheckWebSocket:  true,
		Providers:       []string{"default", "public"},
	}
}

// CheckRPCs checks availability and latency of RPCs for a chain.
func CheckRPCs(ctx context.Context, identifier any, opts CheckRPCOptions) ([]RPCStatus, error) {
	chain, err := getChain(identifier)
	if err != nil {
		return nil, err
	}

	var urlsToCheck []struct {
		url  string
		isWS bool
	}

	providersToCheck := opts.Providers
	if len(providersToCheck) == 0 {
		providersToCheck = make([]string, 0, len(chain.RPCUrls))
		for k := range chain.RPCUrls {
			providersToCheck = append(providersToCheck, k)
		}
	}

	for _, provider := range providersToCheck {
		if target, ok := chain.RPCUrls[provider]; ok {
			if opts.CheckHTTP {
				for _, u := range target.Http {
					if u != "" {
						urlsToCheck = append(urlsToCheck, struct {
							url  string
							isWS bool
						}{u, false})
					}
				}
			}
			if opts.CheckWebSocket {
				for _, u := range target.WebSocket {
					if u != "" {
						urlsToCheck = append(urlsToCheck, struct {
							url  string
							isWS bool
						}{u, true})
					}
				}
			}
		}
	}

	if len(urlsToCheck) == 0 {
		return []RPCStatus{}, nil
	}

	results := make([]RPCStatus, len(urlsToCheck))
	var wg sync.WaitGroup
	wg.Add(len(urlsToCheck))

	for i, u := range urlsToCheck {
		go func(index int, urlToCheck struct {
			url  string
			isWS bool
		}) {
			defer wg.Done()
			checkCtx, cancel := context.WithTimeout(ctx, opts.TimeoutPerCheck)
			defer cancel()
			if urlToCheck.isWS {
				results[index] = checkWebSocket(checkCtx, urlToCheck.url, opts.TimeoutPerCheck)
			} else {
				results[index] = checkHTTP(checkCtx, urlToCheck.url, opts.TimeoutPerCheck)
			}
		}(i, u)
	}

	wg.Wait()

	if ctx.Err() != nil {
		return results, ctx.Err()
	}

	return results, nil
}

// jsonRPCRequest represents a JSON-RPC request object.
type jsonRPCRequest struct {
	Version string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// jsonRPCResponse represents a JSON-RPC response object.
type jsonRPCResponse struct {
	Version string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *jsonRPCError   `json:"error,omitempty"`
}

// jsonRPCError represents a JSON-RPC error object.
type jsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *jsonRPCError) Error() string {
	return fmt.Sprintf("RPC error %d: %s", e.Code, e.Message)
}

// checkHTTP performs the eth_blockNumber check against an HTTP endpoint.
func checkHTTP(ctx context.Context, url string, timeout time.Duration) RPCStatus {
	start := time.Now()
	status := RPCStatus{URL: url, IsHTTP: true}

	client := http.Client{
		Timeout: timeout,
	}

	reqBody := jsonRPCRequest{
		Version: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		status.Error = fmt.Errorf("failed to marshal json-rpc request: %w", err)
		return status
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		status.Error = fmt.Errorf("failed to create http request: %w", err)
		return status
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		status.Error = fmt.Errorf("http request failed: %w", err)
		return status
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		status.Error = fmt.Errorf("failed to read response body: %w", err)
		return status
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		status.Error = fmt.Errorf("http status %d: %s", resp.StatusCode, string(bodyBytes))
		return status
	}

	var rpcResp jsonRPCResponse
	if err := json.Unmarshal(bodyBytes, &rpcResp); err != nil {
		status.Error = fmt.Errorf("failed to unmarshal json-rpc response: %w", err)
		return status
	}

	if rpcResp.Error != nil {
		status.Error = rpcResp.Error
		return status
	}

	if rpcResp.ID != reqBody.ID {
		status.Error = fmt.Errorf("rpc response id mismatch (got %d, expected %d)", rpcResp.ID, reqBody.ID)
		return status
	}

	var blockNumberHex string
	if err := json.Unmarshal(rpcResp.Result, &blockNumberHex); err != nil {
		status.Error = fmt.Errorf("failed to unmarshal block number result: %w", err)
		return status
	}

	blockNumber := new(big.Int)
	if _, ok := blockNumber.SetString(strings.TrimPrefix(blockNumberHex, "0x"), 16); !ok {
		status.Error = fmt.Errorf("failed to parse block number hex: %s", blockNumberHex)
		return status
	}

	status.Latency = time.Since(start)
	status.IsAvailable = true
	status.BlockNumber = blockNumber
	return status
}

// checkWebSocket performs the eth_blockNumber check against a WebSocket endpoint.
func checkWebSocket(ctx context.Context, url string, timeout time.Duration) RPCStatus {
	start := time.Now()
	status := RPCStatus{URL: url, IsWebSocket: true}

	dialer := websocket.Dialer{
		HandshakeTimeout: timeout,
	}

	checkCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, resp, err := dialer.DialContext(checkCtx, url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("websocket dial failed: %v", err)
		if resp != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			errMsg = fmt.Sprintf("%s (status: %s, body: %s)", errMsg, resp.Status, string(bodyBytes))
		}
		status.Error = fmt.Errorf(errMsg)
		return status
	}
	defer conn.Close()

	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	_ = conn.SetWriteDeadline(time.Now().Add(timeout))

	reqBody := jsonRPCRequest{
		Version: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}
	if err := conn.WriteJSON(reqBody); err != nil {
		status.Error = fmt.Errorf("websocket write json failed: %w", err)
		return status
	}

	var rpcResp jsonRPCResponse
	if err := conn.ReadJSON(&rpcResp); err != nil {
		status.Error = fmt.Errorf("websocket read json failed: %w", err)
		return status
	}

	if rpcResp.Error != nil {
		status.Error = rpcResp.Error
		return status
	}
	if rpcResp.ID != reqBody.ID {
		status.Error = fmt.Errorf("rpc response id mismatch (got %d, expected %d)", rpcResp.ID, reqBody.ID)
		return status
	}
	var blockNumberHex string
	if err := json.Unmarshal(rpcResp.Result, &blockNumberHex); err != nil {
		status.Error = fmt.Errorf("failed to unmarshal block number result: %w", err)
		return status
	}
	blockNumber := new(big.Int)
	if _, ok := blockNumber.SetString(strings.TrimPrefix(blockNumberHex, "0x"), 16); !ok {
		status.Error = fmt.Errorf("failed to parse block number hex: %s", blockNumberHex)
		return status
	}

	status.Latency = time.Since(start)
	status.IsAvailable = true
	status.BlockNumber = blockNumber
	return status
}

// getChain is a helper to retrieve a chain by ID or name.
func getChain(identifier any) (Chain, error) {
	switch v := identifier.(type) {
	case *big.Int:
		chain, found := GetChainByID(v)
		if !found {
			return Chain{}, fmt.Errorf("%w: ID %s", ErrChainNotFound, v.String())
		}
		return chain, nil
	case int:
		chain, found := GetChainByID(big.NewInt(int64(v)))
		if !found {
			return Chain{}, fmt.Errorf("%w: ID %d", ErrChainNotFound, v)
		}
		return chain, nil
	case int64:
		chain, found := GetChainByID(big.NewInt(v))
		if !found {
			return Chain{}, fmt.Errorf("%w: ID %d", ErrChainNotFound, v)
		}
		return chain, nil
	case uint:
		chain, found := GetChainByID(new(big.Int).SetUint64(uint64(v)))
		if !found {
			return Chain{}, fmt.Errorf("%w: ID %d", ErrChainNotFound, v)
		}
		return chain, nil
	case uint64:
		chain, found := GetChainByID(new(big.Int).SetUint64(v))
		if !found {
			return Chain{}, fmt.Errorf("%w: ID %d", ErrChainNotFound, v)
		}
		return chain, nil
	case string:
		if id, ok := new(big.Int).SetString(v, 0); ok {
			chain, found := GetChainByID(id)
			if found {
				return chain, nil
			}
		}
		chain, found := GetChainByName(v)
		if !found {
			return Chain{}, fmt.Errorf("%w: Name '%s'", ErrChainNotFound, v)
		}
		return chain, nil
	default:
		return Chain{}, fmt.Errorf("unsupported identifier type: %T", identifier)
	}
}
