package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// ArbitrumNova is the Arbitrum Nova mainnet configuration.
var ArbitrumNova = types.Chain{
	ID:   big.NewInt(42170),
	Name: "Arbitrum Nova",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http:      []string{"https://nova.arbitrum.io/rpc", "https://arbitrum-nova.publicnode.com"},
			WebSocket: []string{"wss://nova.arbitrum.io/feed"},
		},
		"public": {
			Http:      []string{"https://nova.arbitrum.io/rpc", "https://arbitrum-nova.publicnode.com"},
			WebSocket: []string{"wss://nova.arbitrum.io/feed"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Arbitrum Nova Explorer",
			URL:  "https://nova-explorer.arbitrum.io",
		},
		"explorer": {
			Name: "Arbitrum Nova Explorer",
			URL:  "https://nova-explorer.arbitrum.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 1746963,
		},
	},
	IsTestnet: false,
}
