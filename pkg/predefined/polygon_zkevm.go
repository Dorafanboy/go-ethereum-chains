package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// PolygonZkEvm is the Polygon zkEVM mainnet configuration.
var PolygonZkEvm = types.Chain{
	ID:   big.NewInt(1101),
	Name: "Polygon zkEVM",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://zkevm-rpc.com",
				"https://rpc.ankr.com/polygon_zkevm",
			},
			WebSocket: []string{"wss://polygon-zkevm.drpc.org"},
		},
		"public": {
			Http: []string{
				"https://zkevm-rpc.com",
				"https://rpc.ankr.com/polygon_zkevm",
			},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Polygon zkEVM Scan",
			URL:  "https://zkevm.polygonscan.com",
		},
		"polygonscan": {
			Name: "Polygon zkEVM Scan",
			URL:  "https://zkevm.polygonscan.com",
		},
	},
	IsTestnet: false,
}
