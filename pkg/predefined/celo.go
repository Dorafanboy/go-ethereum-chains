package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Celo is the Celo mainnet configuration.
var Celo = chains.Chain{
	ID:   big.NewInt(42220),
	Name: "Celo",
	NativeCurrency: chains.NativeCurrency{
		Name:     "CELO",
		Symbol:   "CELO",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://forno.celo.org",
				"https://rpc.ankr.com/celo",
			},
			WebSocket: []string{"wss://forno.celo.org/ws"},
		},
		"public": { // Alias for default public access
			Http: []string{
				"https://forno.celo.org",
				"https://rpc.ankr.com/celo",
			},
			WebSocket: []string{"wss://forno.celo.org/ws"},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
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
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 13112599,
		},
	},
	IsTestnet: false,
}
