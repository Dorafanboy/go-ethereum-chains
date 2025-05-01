package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"go-ethereum-chains/internal/types"
	"go-ethereum-chains/pkg/registry"

	"github.com/gorilla/websocket"
)

// CheckRPCOptions defines parameters for checking RPC endpoints.
type CheckRPCOptions struct {
	TimeoutPerCheck time.Duration
	CheckHTTP       bool
	CheckWebSocket  bool
	Providers       []types.ProviderName
}

// DefaultCheckOptions returns default options for CheckRPCs.
func DefaultCheckOptions() CheckRPCOptions {
	return CheckRPCOptions{
		TimeoutPerCheck: 5 * time.Second,
		CheckHTTP:       true,
		CheckWebSocket:  true,
		Providers:       []types.ProviderName{types.ProviderDefault, types.ProviderPublic},
	}
}

// CheckRPCs checks availability and latency of RPCs for a chain identified by ID or name.
func CheckRPCs(ctx context.Context, identifier any, opts CheckRPCOptions) ([]types.RPCStatus, error) {
	chain, err := registry.FindChain(identifier)
	if err != nil {
		return nil, err // Error already includes ErrChainNotFound info
	}

	var urlsToCheck []struct {
		url  string
		isWS bool
	}

	providersToCheck := opts.Providers
	if len(providersToCheck) == 0 {
		providersToCheck = make([]types.ProviderName, 0, len(chain.RPCUrls))
		for k := range chain.RPCUrls {
			providersToCheck = append(providersToCheck, types.ProviderName(k))
		}
	}

	for _, provider := range providersToCheck {
		if target, ok := chain.RPCUrls[string(provider)]; ok {
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
		return []types.RPCStatus{}, nil // Return empty slice if no URLs found
	}

	results := make([]types.RPCStatus, len(urlsToCheck))
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
		// Return partial results along with the context error (e.g., timeout)
		return results, ctx.Err()
	}

	return results, nil
}

// checkHTTP performs the eth_blockNumber check against an HTTP endpoint.
func checkHTTP(ctx context.Context, url string, timeout time.Duration) types.RPCStatus {
	start := time.Now()
	status := types.RPCStatus{URL: url, IsHTTP: true}

	client := http.Client{
		Timeout: timeout,
	}

	reqBody := types.JsonRPCRequest{
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		status.Error = fmt.Errorf("failed to create http request: %w", err)
		return status
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

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

	var rpcResp types.JsonRPCResponse
	if err := json.Unmarshal(bodyBytes, &rpcResp); err != nil {
		status.Error = fmt.Errorf("failed to unmarshal json-rpc response (body: %s): %w", string(bodyBytes), err)
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
		status.Error = fmt.Errorf("failed to unmarshal block number result (%s): %w", string(rpcResp.Result), err)
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
func checkWebSocket(ctx context.Context, url string, timeout time.Duration) types.RPCStatus {
	start := time.Now()
	status := types.RPCStatus{URL: url, IsWebSocket: true}

	dialer := websocket.Dialer{
		HandshakeTimeout: timeout / 2,
	}

	conn, resp, err := dialer.DialContext(ctx, url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("websocket dial failed: %v", err)
		if resp != nil {
			bodyBytes, readErr := io.ReadAll(resp.Body)
			resp.Body.Close()
			if readErr == nil {
				errMsg = fmt.Sprintf("%s (status: %s, body: %s)", errMsg, resp.Status, string(bodyBytes))
			}
		}
		status.Error = fmt.Errorf(errMsg)
		return status
	}
	defer conn.Close()

	deadline := time.Now().Add(timeout - time.Since(start))
	_ = conn.SetReadDeadline(deadline)
	_ = conn.SetWriteDeadline(deadline)

	reqBody := types.JsonRPCRequest{
		Version: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		ID:      1,
	}
	if err := conn.WriteJSON(reqBody); err != nil {
		status.Error = fmt.Errorf("websocket write json failed: %w", err)
		return status
	}

	var rpcResp types.JsonRPCResponse
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
		status.Error = fmt.Errorf("failed to unmarshal block number result (%s): %w", string(rpcResp.Result), err)
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
