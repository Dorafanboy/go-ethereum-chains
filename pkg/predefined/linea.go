package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Linea is the Linea mainnet configuration.
var Linea = chains.Chain{
	ID:   big.NewInt(59144),
	Name: "Linea",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.linea.build",
				"https://linea.blockpi.network/v1/rpc/public",
				"https://linea.drpc.org",
			},
			// WebSocket: []string{"wss://..."}, // Public WS endpoint requires provider API key
		},
		"public": { // Alias for default public access
			Http: []string{
				"https://rpc.linea.build",
				"https://linea.blockpi.network/v1/rpc/public",
				"https://linea.drpc.org",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "LineaScan",
			URL:  "https://lineascan.build",
		},
		"lineascan": {
			Name: "LineaScan",
			URL:  "https://lineascan.build",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 42, // Early deployment
		},
	},
	IsTestnet: false,
}
