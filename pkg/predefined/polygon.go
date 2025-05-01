package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Polygon is the Polygon PoS mainnet configuration.
var Polygon = types.Chain{
	ID:   big.NewInt(137),
	Name: "Polygon",
	NativeCurrency: types.NativeCurrency{
		Name:     "MATIC",
		Symbol:   "MATIC",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "PolygonScan",
			URL:  "https://polygonscan.com",
		},
		"polygonscan": {
			Name: "PolygonScan",
			URL:  "https://polygonscan.com",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 25770160,
		},
	},
	IsTestnet: false,
}
