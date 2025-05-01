package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Fantom is the Fantom Opera mainnet configuration.
var Fantom = types.Chain{
	ID:   big.NewInt(250),
	Name: "Fantom Opera",
	NativeCurrency: types.NativeCurrency{
		Name:     "Fantom",
		Symbol:   "FTM",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.ftm.tools",
				"https://fantom.publicnode.com",
				"https://rpc.ankr.com/fantom",
				"https://rpcapi.fantom.network",
			},
			WebSocket: []string{"wss://fantom-rpc.publicnode.com"},
		},
		"public": {
			Http: []string{
				"https://rpc.ftm.tools",
				"https://fantom.publicnode.com",
				"https://rpc.ankr.com/fantom",
				"https://rpcapi.fantom.network",
			},
			WebSocket: []string{"wss://fantom-rpc.publicnode.com"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "FTMScan",
			URL:  "https://ftmscan.com",
		},
		"ftmscan": {
			Name: "FTMScan",
			URL:  "https://ftmscan.com",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 33001987,
		},
	},
	IsTestnet: false,
}
