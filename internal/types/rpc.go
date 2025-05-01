package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// RPCStatus holds the result of checking a single RPC endpoint.
type RPCStatus struct {
	URL         string        `json:"url"`
	IsHTTP      bool          `json:"isHttp"`
	IsWebSocket bool          `json:"isWebSocket"`
	IsAvailable bool          `json:"isAvailable"`
	Latency     time.Duration `json:"latency"`
	BlockNumber *big.Int      `json:"blockNumber,omitempty"`
	Error       error         `json:"error,omitempty"`
}

// JsonRPCRequest represents a JSON-RPC request object.
type JsonRPCRequest struct {
	Version string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// JsonRPCResponse represents a JSON-RPC response object.
type JsonRPCResponse struct {
	Version string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *JsonRPCError   `json:"error,omitempty"`
}

// JsonRPCError represents a JSON-RPC error object.
type JsonRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *JsonRPCError) Error() string {
	return fmt.Sprintf("RPC error %d: %s", e.Code, e.Message)
}
