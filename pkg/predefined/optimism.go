package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// TODO: завтра продолжить завершить фикс ошибок рефактор и выкатку

// Optimism (now OP Mainnet) is the Optimism mainnet configuration.
var Optimism = chains.Chain{
	ID:   big.NewInt(10),
	Name: "OP Mainnet", // Official name
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http:      []string{"https://mainnet.optimism.io", "https://optimism.publicnode.com", "https://rpc.ankr.com/optimism"},
			WebSocket: []string{"wss://mainnet.optimism.io"},
		},
		"public": {
			Http:      []string{"https://mainnet.optimism.io", "https://optimism.publicnode.com", "https://rpc.ankr.com/optimism"},
			WebSocket: []string{"wss://mainnet.optimism.io"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "OP Etherscan",
			URL:  "https://optimistic.etherscan.io",
		},
		"etherscan": {
			Name: "OP Etherscan",
			URL:  "https://optimistic.etherscan.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 4286263,
		},
	},
	IsTestnet: false,
}
