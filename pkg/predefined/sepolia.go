package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Sepolia is the official Sepolia testnet configuration.
var Sepolia = types.Chain{
	ID:   big.NewInt(11155111),
	Name: "Sepolia",
	NativeCurrency: types.NativeCurrency{
		Name:     "Sepolia Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.sepolia.org",
				"https://rpc2.sepolia.org",
				"https://rpc.ankr.com/eth_sepolia",
			},
			WebSocket: []string{"wss://rpc.sepolia.org"},
		},
		"public": {
			Http: []string{
				"https://rpc.sepolia.org",
				"https://rpc2.sepolia.org",
				"https://rpc.ankr.com/eth_sepolia",
			},
			WebSocket: []string{"wss://rpc.sepolia.org"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Etherscan",
			URL:  "https://sepolia.etherscan.io",
		},
		"etherscan": {
			Name: "Etherscan",
			URL:  "https://sepolia.etherscan.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 650767,
		},
	},
	IsTestnet: true,
}
