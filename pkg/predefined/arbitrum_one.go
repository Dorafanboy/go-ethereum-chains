package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// ArbitrumOne is the Arbitrum One mainnet configuration.
var ArbitrumOne = types.Chain{
	ID:   big.NewInt(42161),
	Name: "Arbitrum One",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://arb1.arbitrum.io/rpc",
				"https://arbitrum.publicnode.com",
				"https://rpc.ankr.com/arbitrum",
			},
			WebSocket: []string{"wss://arb1.arbitrum.io/rpc"},
		},
		"public": {
			Http: []string{
				"https://arb1.arbitrum.io/rpc",
				"https://arbitrum.publicnode.com",
				"https://rpc.ankr.com/arbitrum",
			},
			WebSocket: []string{"wss://arb1.arbitrum.io/rpc"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Arbiscan",
			URL:  "https://arbiscan.io",
		},
		"arbiscan": {
			Name: "Arbiscan",
			URL:  "https://arbiscan.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 7654707,
		},
	},
	IsTestnet: false,
}
