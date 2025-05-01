package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Linea is the Linea mainnet configuration.
var Linea = types.Chain{
	ID:   big.NewInt(59144),
	Name: "Linea",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.linea.build",
				"https://linea.blockpi.network/v1/rpc/public",
				"https://linea.drpc.org",
			},
			WebSocket: []string{"wss://linea-rpc.publicnode.com"},
		},
		"public": {
			Http: []string{
				"https://rpc.linea.build",
				"https://linea.blockpi.network/v1/rpc/public",
				"https://linea.drpc.org",
			},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "LineaScan",
			URL:  "https://lineascan.build",
		},
		"lineascan": {
			Name: "LineaScan",
			URL:  "https://lineascan.build",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 42,
		},
	},
	IsTestnet: false,
}
