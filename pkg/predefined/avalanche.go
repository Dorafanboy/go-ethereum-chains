package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Avalanche is the Avalanche C-Chain mainnet configuration.
var Avalanche = chains.Chain{
	ID:   big.NewInt(43114),
	Name: "Avalanche",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Avalanche",
		Symbol:   "AVAX",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://api.avax.network/ext/bc/C/rpc",
				"https://avalanche.public-rpc.com",
				"https://rpc.ankr.com/avalanche",
			},
			WebSocket: []string{
				"wss://api.avax.network/ext/bc/C/ws",
				"wss://avalanche-c-chain-rpc.publicnode.com",
			},
		},
		"public": {
			Http: []string{
				"https://api.avax.network/ext/bc/C/rpc",
				"https://avalanche.public-rpc.com",
				"https://rpc.ankr.com/avalanche",
			},
			WebSocket: []string{
				"wss://api.avax.network/ext/bc/C/ws",
				"wss://avalanche-c-chain-rpc.publicnode.com",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Snowtrace",
			URL:  "https://snowtrace.io",
		},
		"snowtrace": {
			Name: "Snowtrace",
			URL:  "https://snowtrace.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 11907934,
		},
	},
	IsTestnet: false,
}
