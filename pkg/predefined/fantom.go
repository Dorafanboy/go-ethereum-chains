package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Fantom is the Fantom Opera mainnet configuration.
var Fantom = chains.Chain{
	ID:   big.NewInt(250),
	Name: "Fantom Opera",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Fantom",
		Symbol:   "FTM",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.ftm.tools",
				"https://fantom.publicnode.com",
				"https://rpc.ankr.com/fantom",
				"https://rpcapi.fantom.network",
			},
			WebSocket: []string{"wss://fantom-rpc.publicnode.com"},
		},
		"public": { // Alias for default public access
			Http: []string{
				"https://rpc.ftm.tools",
				"https://fantom.publicnode.com",
				"https://rpc.ankr.com/fantom",
				"https://rpcapi.fantom.network",
			},
			WebSocket: []string{"wss://fantom-rpc.publicnode.com"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "FTMScan",
			URL:  "https://ftmscan.com",
		},
		"ftmscan": {
			Name: "FTMScan",
			URL:  "https://ftmscan.com",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 33001987,
		},
	},
	IsTestnet: false,
}
