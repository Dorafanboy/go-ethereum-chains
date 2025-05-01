package chains

import "math/big"

// NativeCurrency represents the native currency of the chain.
// Example: Ether for Ethereum Mainnet (Symbol: ETH, Decimals: 18).
type NativeCurrency struct {
	Name     string `json:"name" yaml:"name"`
	Symbol   string `json:"symbol" yaml:"symbol"`
	Decimals uint   `json:"decimals" yaml:"decimals"`
}

// RpcTarget holds the RPC endpoints for a network provider (e.g., default, infura).
type RpcTarget struct {
	Http      []string `json:"http,omitempty" yaml:"http,omitempty"`
	WebSocket []string `json:"webSocket,omitempty" yaml:"webSocket,omitempty"`
}

// BlockExplorer represents a block explorer for the chain.
type BlockExplorer struct {
	Name string `json:"name" yaml:"name"`
	URL  string `json:"url" yaml:"url"`
}

// Contract represents a known contract address on the chain.
type Contract struct {
	Address      string `json:"address" yaml:"address"`
	BlockCreated uint64 `json:"blockCreated,omitempty" yaml:"blockCreated,omitempty"`
}

// Contracts holds known contract addresses for the chain.
type Contracts struct {
	Multicall3 *Contract `json:"multicall3,omitempty" yaml:"multicall3,omitempty"`
}

// Chain represents an Ethereum compatible network.
type Chain struct {
	ID                   *big.Int                 `json:"id" yaml:"id"`
	Name                 string                   `json:"name" yaml:"name"`
	NativeCurrency       NativeCurrency           `json:"nativeCurrency" yaml:"nativeCurrency"`
	RPCUrls              map[string]RpcTarget     `json:"rpcUrls" yaml:"rpcUrls"`
	BlockExplorers       map[string]BlockExplorer `json:"blockExplorers,omitempty" yaml:"blockExplorers,omitempty"`
	Contracts            *Contracts               `json:"contracts,omitempty" yaml:"contracts,omitempty"`
	IsTestnet            bool                     `json:"isTestnet,omitempty" yaml:"isTestnet,omitempty"`
	EnsRegistry          *Contract                `json:"ensRegistry,omitempty" yaml:"ensRegistry,omitempty"`
	EnsUniversalResolver *Contract                `json:"ensUniversalResolver,omitempty" yaml:"ensUniversalResolver,omitempty"`
}
