package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Bnb is the BNB Smart Chain mainnet configuration.
var Bnb = types.Chain{
	ID:   big.NewInt(56),
	Name: "BNB Smart Chain",
	NativeCurrency: types.NativeCurrency{
		Name:     "BNB",
		Symbol:   "BNB",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "BscScan",
			URL:  "https://bscscan.com",
		},
		"bscscan": {
			Name: "BscScan",
			URL:  "https://bscscan.com",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 15921452,
		},
	},
	IsTestnet: false,
}
