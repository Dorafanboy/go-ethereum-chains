package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Polygon is the Polygon PoS mainnet configuration.
var Polygon = chains.Chain{
	ID:   big.NewInt(137),
	Name: "Polygon",
	NativeCurrency: chains.NativeCurrency{
		Name:     "MATIC",
		Symbol:   "MATIC",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://polygon-rpc.com",
				"https://rpc.ankr.com/polygon",
				"https://polygon.llamarpc.com",
			},
			WebSocket: []string{"wss://polygon-bor-rpc.publicnode.com"},
		},
		"public": {
			Http: []string{
				"https://polygon-rpc.com",
				"https://rpc.ankr.com/polygon",
				"https://polygon.llamarpc.com",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "PolygonScan",
			URL:  "https://polygonscan.com",
		},
		"polygonscan": {
			Name: "PolygonScan",
			URL:  "https://polygonscan.com",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 25770160,
		},
	},
	IsTestnet: false,
}
