package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Scroll is the Scroll mainnet configuration.
var Scroll = chains.Chain{
	ID:   big.NewInt(534352),
	Name: "Scroll",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http:      []string{"https://rpc.scroll.io", "https://scroll.blockpi.network/v1/rpc/public"},
			WebSocket: []string{"wss://wss-rpc.scroll.io/ws"},
		},
		"public": {
			Http: []string{
				"https://rpc.scroll.io",
				"https://scroll.blockpi.network/v1/rpc/public",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Scrollscan",
			URL:  "https://scrollscan.com",
		},
		"scrollscan": {
			Name: "Scrollscan",
			URL:  "https://scrollscan.com",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 14,
		},
	},
	IsTestnet: false,
}
