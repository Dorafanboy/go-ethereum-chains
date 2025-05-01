package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Gnosis is the Gnosis Chain mainnet configuration.
var Gnosis = types.Chain{
	ID:   big.NewInt(100),
	Name: "Gnosis",
	NativeCurrency: types.NativeCurrency{
		Name:     "xDai",
		Symbol:   "xDAI",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
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
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 21022491,
		},
	},
	IsTestnet: false,
}
