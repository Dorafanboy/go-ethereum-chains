package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// Base is the Base mainnet configuration.
var Base = types.Chain{
	ID:   big.NewInt(8453),
	Name: "Base",
	NativeCurrency: types.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{"https://mainnet.base.org", "https://base-mainnet.public.blastapi.io", "https://base.gateway.tenderly.co"},
			// TODO: Add official node?
		},
		"public": {
			Http: []string{"https://mainnet.base.org", "https://base-mainnet.public.blastapi.io", "https://base.gateway.tenderly.co"},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Basescan",
			URL:  "https://basescan.org",
		},
		"basescan": {
			Name: "Basescan",
			URL:  "https://basescan.org",
		},
	},
	Contracts: &types.Contracts{
		Multicall3: &types.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 5022,
		},
	},
	IsTestnet: false,
}
