package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Gnosis is the Gnosis Chain mainnet configuration.
var Gnosis = chains.Chain{
	ID:   big.NewInt(100),
	Name: "Gnosis",
	NativeCurrency: chains.NativeCurrency{
		Name:     "xDai",
		Symbol:   "xDAI",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.gnosischain.com",
				"https://gnosis.publicnode.com",
				"https://rpc.ankr.com/gnosis",
				"https://gnosis.drpc.org",
			},
			WebSocket: []string{
				"wss://rpc.gnosischain.com/wss",
				"wss://gnosis-mainnet.public.blastapi.io",
			},
		},
		"public": {
			Http: []string{
				"https://rpc.gnosischain.com",
				"https://gnosis.publicnode.com",
				"https://rpc.ankr.com/gnosis",
				"https://gnosis.drpc.org",
			},
			WebSocket: []string{
				"wss://rpc.gnosischain.com/wss",
				"wss://gnosis-mainnet.public.blastapi.io",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "GnosisScan",
			URL:  "https://gnosisscan.io",
		},
		"gnosisscan": {
			Name: "GnosisScan",
			URL:  "https://gnosisscan.io",
		},
		"blockscout": {
			Name: "Blockscout",
			URL:  "https://gnosis.blockscout.com",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 21022491,
		},
	},
	IsTestnet: false,
}
