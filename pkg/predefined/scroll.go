package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Scroll is the Scroll mainnet configuration.
var Scroll = types.Chain{
	ID:   big.NewInt(534352),
	Name: "Scroll",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Scrollscan",
			URL:  "https://scrollscan.com",
		},
		"scrollscan": {
			Name: "Scrollscan",
			URL:  "https://scrollscan.com",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 14,
		},
	},
	IsTestnet: false,
}
