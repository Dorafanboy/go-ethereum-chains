package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Celo is the Celo mainnet configuration.
var Celo = types.Chain{
	ID:   big.NewInt(42220),
	Name: "Celo",
	NativeCurrency: types.NativeCurrency{
		Name:     "CELO",
		Symbol:   "CELO",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://forno.celo.org",
				"https://rpc.ankr.com/celo",
			},
			WebSocket: []string{"wss://forno.celo.org/ws"},
		},
		"public": {
			Http: []string{
				"https://forno.celo.org",
				"https://rpc.ankr.com/celo",
			},
			WebSocket: []string{"wss://forno.celo.org/ws"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "CeloScan",
			URL:  "https://celoscan.io",
		},
		"celoscan": {
			Name: "CeloScan",
			URL:  "https://celoscan.io",
		},
		"blockscout": {
			Name: "Blockscout",
			URL:  "https://explorer.celo.org/mainnet",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 13112599,
		},
	},
	IsTestnet: false,
}
