package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Avalanche is the Avalanche C-Chain mainnet configuration.
var Avalanche = types.Chain{
	ID:   big.NewInt(43114),
	Name: "Avalanche",
	NativeCurrency: types.NativeCurrency{
		Name:     "Avalanche",
		Symbol:   "AVAX",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Snowtrace",
			URL:  "https://snowtrace.io",
		},
		"snowtrace": {
			Name: "Snowtrace",
			URL:  "https://snowtrace.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 11907934,
		},
	},
	IsTestnet: false,
}
