package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Holesky is the Holesky testnet configuration.
var Holesky = types.Chain{
	ID:   big.NewInt(17000),
	Name: "Holesky",
	NativeCurrency: types.NativeCurrency{
		Name:     "Holesky Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http:      []string{"https://rpc.holesky.ethpandaops.io", "https://ethereum-holesky.publicnode.com"},
			WebSocket: []string{"wss://ethereum-holesky.publicnode.com"},
		},
		"public": {
			Http:      []string{"https://rpc.holesky.ethpandaops.io", "https://ethereum-holesky.publicnode.com"},
			WebSocket: []string{"wss://ethereum-holesky.publicnode.com"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Etherscan",
			URL:  "https://holesky.etherscan.io",
		},
		"etherscan": {
			Name: "Etherscan",
			URL:  "https://holesky.etherscan.io",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 49461,
		},
	},
	IsTestnet: true,
}
