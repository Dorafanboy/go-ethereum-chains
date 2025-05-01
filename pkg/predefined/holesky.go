package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Holesky is the Holesky testnet configuration.
var Holesky = chains.Chain{
	ID:   big.NewInt(17000),
	Name: "Holesky",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Holesky Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http:      []string{"https://rpc.holesky.ethpandaops.io", "https://ethereum-holesky.publicnode.com"},
			WebSocket: []string{"wss://ethereum-holesky.publicnode.com"},
		},
		"public": {
			Http:      []string{"https://rpc.holesky.ethpandaops.io", "https://ethereum-holesky.publicnode.com"},
			WebSocket: []string{"wss://ethereum-holesky.publicnode.com"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Etherscan",
			URL:  "https://holesky.etherscan.io",
		},
		"etherscan": {
			Name: "Etherscan",
			URL:  "https://holesky.etherscan.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 49461,
		},
	},
	IsTestnet: true,
}
