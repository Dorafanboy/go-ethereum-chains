package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// ArbitrumOne is the Arbitrum One mainnet configuration.
var ArbitrumOne = chains.Chain{
	ID:   big.NewInt(42161),
	Name: "Arbitrum One",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
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
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Arbiscan",
			URL:  "https://arbiscan.io",
		},
		"arbiscan": {
			Name: "Arbiscan",
			URL:  "https://arbiscan.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 7654707,
		},
	},
	IsTestnet: false,
}
