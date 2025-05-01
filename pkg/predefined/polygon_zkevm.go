package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// PolygonZkEvm is the Polygon zkEVM mainnet configuration.
var PolygonZkEvm = chains.Chain{
	ID:   big.NewInt(1101),
	Name: "Polygon zkEVM",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://zkevm-rpc.com",
				"https://rpc.ankr.com/polygon_zkevm",
			},
			// WebSocket: []string{"wss://..."}, // Public WS endpoint requires provider API key
		},
		"public": { // Alias for default public access
			Http: []string{
				"https://zkevm-rpc.com",
				"https://rpc.ankr.com/polygon_zkevm",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Polygon zkEVM Scan",
			URL:  "https://zkevm.polygonscan.com",
		},
		"polygonscan": {
			Name: "Polygon zkEVM Scan",
			URL:  "https://zkevm.polygonscan.com",
		},
	},
	// Contracts: Specific contracts for Polygon zkEVM might differ.
	// Multicall3 seems not deployed or has a different address.
	IsTestnet: false,
}
