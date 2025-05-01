package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Blast is the Blast mainnet configuration.
var Blast = chains.Chain{
	ID:   big.NewInt(81457),
	Name: "Blast",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.blast.io",
				"https://blast.blockpi.network/v1/rpc/public",
				"https://blastl2-mainnet.public.blastapi.io",
				"https://rpc.ankr.com/blast",
			},
			WebSocket: []string{"wss://blast.drpc.org"},
		},
		"public": {
			Http: []string{
				"https://rpc.blast.io",
				"https://blast.blockpi.network/v1/rpc/public",
				"https://blastl2-mainnet.public.blastapi.io",
				"https://rpc.ankr.com/blast",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Blastscan",
			URL:  "https://blastscan.io",
		},
		"blastscan": {
			Name: "Blastscan",
			URL:  "https://blastscan.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 88,
		},
	},
	IsTestnet: false,
}
