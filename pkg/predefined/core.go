package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Core is the Core DAO mainnet configuration.
var Core = types.Chain{
	ID:   big.NewInt(1116),
	Name: "Core",
	NativeCurrency: types.NativeCurrency{
		Name:     "Core",
		Symbol:   "CORE",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
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
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "CoreScan",
			URL:  "https://scan.coredao.org",
		},
		"corescan": {
			Name: "CoreScan",
			URL:  "https://scan.coredao.org",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 5608481,
		},
	},
	IsTestnet: false,
}
