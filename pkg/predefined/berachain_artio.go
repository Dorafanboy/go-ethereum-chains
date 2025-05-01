package predefined

import (
	"go-ethereum-chains/internal/types"
	"math/big"
)

// BerachainArtio is the Berachain Artio testnet configuration.
// Note: Artio v1 (ID 80085) was deprecated in favor of bArtio (ID 80084) and then Bepolia.
// Keeping 80085 for historical reference if needed, but it's likely inactive.
var BerachainArtio = types.Chain{
	ID:   big.NewInt(80085), // Original Artio ID, may be inactive.
	Name: "Berachain Artio (Deprecated)",
	NativeCurrency: types.NativeCurrency{
		Name:     "BERA",
		Symbol:   "BERA",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http: []string{
				"https://artio.rpc.berachain.com", // Might be inactive or point to bArtio/Bepolia now
			},
			// WebSocket: []string{"wss://..."}, // Public WS not available for Artio/bArtio
		},
		"public": {
			Http: []string{
				"https://artio.rpc.berachain.com",
			},
		},
	},
	BlockExplorers: map[string]types.BlockExplorer{
		"default": {
			Name: "Beratrail (Artio)",
			URL:  "https://artio.beratrail.io", // Might redirect or be inactive
		},
		"beratrail": {
			Name: "Beratrail (Artio)",
			URL:  "https://artio.beratrail.io",
		},
	},
	// Contracts: Specific contracts for Berachain might differ.
	IsTestnet: true,
}
