package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// ZkSync is the zkSync Era mainnet configuration.
var ZkSync = types.Chain{
	ID:   big.NewInt(324),
	Name: "zkSync Era",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http:      []string{"https://mainnet.era.zksync.io"},
			WebSocket: []string{"wss://mainnet.era.zksync.io/ws"},
		},
		"public": {
			Http:      []string{"https://mainnet.era.zksync.io"},
			WebSocket: []string{"wss://mainnet.era.zksync.io/ws"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "zkSync Era Explorer",
			URL:  "https://explorer.zksync.io",
		},
	},
	IsTestnet: false,
}
