package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Bnb is the BNB Smart Chain mainnet configuration.
var Bnb = chains.Chain{
	ID:   big.NewInt(56),
	Name: "BNB Smart Chain",
	NativeCurrency: chains.NativeCurrency{
		Name:     "BNB",
		Symbol:   "BNB",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://bsc-dataseed.bnbchain.org",
				"https://bsc-dataseed1.defibit.io",
				"https://bsc-dataseed1.ninicoin.io",
				"https://bsc.publicnode.com",
				"https://rpc.ankr.com/bsc",
			},
			WebSocket: []string{"wss://bsc.publicnode.com"},
		},
		"public": {
			Http: []string{
				"https://bsc-dataseed.bnbchain.org",
				"https://bsc-dataseed1.defibit.io",
				"https://bsc-dataseed1.ninicoin.io",
				"https://bsc.publicnode.com",
				"https://rpc.ankr.com/bsc",
			},
			WebSocket: []string{"wss://bsc.publicnode.com"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "BscScan",
			URL:  "https://bscscan.com",
		},
		"bscscan": {
			Name: "BscScan",
			URL:  "https://bscscan.com",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 15921452,
		},
	},
	IsTestnet: false,
}
