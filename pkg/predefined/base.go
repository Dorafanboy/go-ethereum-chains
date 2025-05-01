package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Base is the Base mainnet configuration.
var Base = chains.Chain{
	ID:   big.NewInt(8453),
	Name: "Base",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://mainnet.base.org",
				//"https://developer-access-mainnet.base.org", // Requires auth
				"https://base.publicnode.com",
			},
			WebSocket: []string{"wss://mainnet.base.org"},
		},
		"public": {
			Http: []string{
				"https://mainnet.base.org",
				"https://base.publicnode.com",
			},
			WebSocket: []string{"wss://mainnet.base.org"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Basescan",
			URL:  "https://basescan.org",
		},
		"basescan": {
			Name: "Basescan",
			URL:  "https://basescan.org",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 5022,
		},
	},
	IsTestnet: false,
}
