package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// ArbitrumNova is the Arbitrum Nova mainnet configuration.
var ArbitrumNova = chains.Chain{
	ID:   big.NewInt(42170),
	Name: "Arbitrum Nova",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http:      []string{"https://nova.arbitrum.io/rpc", "https://arbitrum-nova.publicnode.com"},
			WebSocket: []string{"wss://nova.arbitrum.io/feed"},
		},
		"public": {
			Http:      []string{"https://nova.arbitrum.io/rpc", "https://arbitrum-nova.publicnode.com"},
			WebSocket: []string{"wss://nova.arbitrum.io/feed"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Arbitrum Nova Explorer",
			URL:  "https://nova-explorer.arbitrum.io",
		},
		"explorer": {
			Name: "Arbitrum Nova Explorer",
			URL:  "https://nova-explorer.arbitrum.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 1746963,
		},
	},
	IsTestnet: false,
}
