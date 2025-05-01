package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// ZkSync is the zkSync Era mainnet configuration.
var ZkSync = chains.Chain{
	ID:   big.NewInt(324),
	Name: "zkSync Era",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http:      []string{"https://mainnet.era.zksync.io"},
			WebSocket: []string{"wss://mainnet.era.zksync.io/ws"},
		},
		"public": {
			Http:      []string{"https://mainnet.era.zksync.io"},
			WebSocket: []string{"wss://mainnet.era.zksync.io/ws"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "zkSync Era Explorer",
			URL:  "https://explorer.zksync.io",
		},
	},
	IsTestnet: false,
}
