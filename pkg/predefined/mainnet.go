package predefined

import (
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

// Mainnet is the official Ethereum Mainnet chain configuration.
var Mainnet = chains.Chain{
	ID:   big.NewInt(1),
	Name: "Ethereum Mainnet",
	NativeCurrency: chains.NativeCurrency{
		Name:     "Ether",
		Symbol:   "ETH",
		Decimals: 18,
	},
	RPCUrls: map[string]chains.RpcTarget{
		"default": {
			Http: []string{
				"https://cloudflare-eth.com",
				"https://rpc.ankr.com/eth",
			},
			WebSocket: []string{"wss://ethereum-rpc.publicnode.com"},
		},
		"public": {
			Http: []string{
				"https://cloudflare-eth.com",
				"https://rpc.ankr.com/eth",
			},
		},
	},
	BlockExplorers: map[string]chains.BlockExplorer{
		"default": {
			Name: "Etherscan",
			URL:  "https://etherscan.io",
		},
		"etherscan": {
			Name: "Etherscan",
			URL:  "https://etherscan.io",
		},
	},
	Contracts: &chains.Contracts{
		Multicall3: &chains.Contract{
			Address:      "0xcA11bde05977b3631167028862bE2a173976CA11",
			BlockCreated: 14353601,
		},
		// TODO: Add more contracts
	},
	EnsRegistry: &chains.Contract{
		Address: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
	},
	EnsUniversalResolver: &chains.Contract{
		Address:      "0xc0497E381f536Be9ce14B0dD3817cBcAe57d2F62",
		BlockCreated: 16966585,
	},
	IsTestnet: false,
}
