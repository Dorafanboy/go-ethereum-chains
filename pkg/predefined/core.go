package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Core is the Core DAO mainnet configuration.
var Core = chains.Chain{
	ID:   big.NewInt(1116),
	Name: "Core",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Core",
		Symbol:   "CORE",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://rpc.coredao.org",
				"https://rpc.ankr.com/core",
				"https://core.public.infstones.com",
				"https://rpc-core.icecreamswap.com",
				"https://core.drpc.org",
			},
			WebSocket: []string{
				"wss://ws.coredao.org",
				"wss://core.drpc.org",
			},
		},
		"public": {
			Http: []string{
				"https://rpc.coredao.org",
				"https://rpc.ankr.com/core",
				"https://core.public.infstones.com",
				"https://rpc-core.icecreamswap.com",
				"https://core.drpc.org",
			},
			WebSocket: []string{
				"wss://ws.coredao.org",
				"wss://core.drpc.org",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "CoreScan",
			URL:  "https://scan.coredao.org",
		},
		"corescan": {
			Name: "CoreScan",
			URL:  "https://scan.coredao.org",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 5608481,
		},
	},
	IsTestnet: false,
}
