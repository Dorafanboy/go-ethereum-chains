package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Optimism (now OP Mainnet) is the Optimism mainnet configuration.
var Optimism = types.Chain{
	ID:   big.NewInt(10),
	Name: "OP Mainnet",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://mainnet.optimism.io",
				"https://optimism.publicnode.com",
				"https://rpc.ankr.com/optimism",
			},
			WebSocket: []string{"wss://mainnet.optimism.io"},
		},
		"public": {
			Http: []string{
				"https://mainnet.optimism.io",
				"https://optimism.publicnode.com",
				"https://rpc.ankr.com/optimism",
			},
			WebSocket: []string{"wss://mainnet.optimism.io"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "OP Etherscan",
			URL:  "https://optimistic.etherscan.io",
		},
		"etherscan": {
			Name: "OP Etherscan",
			URL:  "https://optimistic.etherscan.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 4286263,
		},
	},
	IsTestnet: false,
}
