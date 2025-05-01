package predefined

import "go-ethereum-chains/pkg/chains"

// init automatically registers all predefined chains in the central registry.
func init() {
	// Register chains one by one
	chains.RegisterChain(Mainnet)
	chains.RegisterChain(Sepolia)
	chains.RegisterChain(Holesky)
	chains.RegisterChain(Base)
	chains.RegisterChain(Optimism)
	chains.RegisterChain(ArbitrumOne)
	chains.RegisterChain(ZkSync) // Note: Variable name matches the file/concept
	chains.RegisterChain(Scroll)
	chains.RegisterChain(ArbitrumNova)
	chains.RegisterChain(Polygon)
	chains.RegisterChain(BerachainArtio)
	chains.RegisterChain(Avalanche)
	chains.RegisterChain(Bnb)
	chains.RegisterChain(Gnosis)
	chains.RegisterChain(Celo)
	chains.RegisterChain(Core)
	chains.RegisterChain(Linea)
	chains.RegisterChain(Fantom)
	chains.RegisterChain(PolygonZkEvm) // Note: Variable name matches the file/concept
	chains.RegisterChain(Blast)
}
