# go-ethereum-chains

A Go library for managing Ethereum chain configurations, inspired by viem's chains module.

Provides easy access to predefined chain configurations (Mainnet, Sepolia, Base, Optimism, etc.) and allows registration of custom chains.

## Installation

```bash
go get go-ethereum-chains // Assuming it will be hosted at this path
// OR go get github.com/Dorafanboy/go-ethereum-chains if using GitHub path
```

## Usage

### Using Predefined Chains

Predefined chains are automatically registered when you import the `predefined` package using a blank identifier (`_`). You can then retrieve them using functions from the `chains` package.

```go
package main

import (
	"fmt"
	"math/big"

	"go-ethereum-chains/pkg/chains"
	_ "go-ethereum-chains/pkg/predefined"
)

func main() {
	// Get Mainnet by ID
	mainnet, ok := chains.GetChainByID(big.NewInt(1))
	if ok {
		fmt.Printf("Chain: %s, ID: %d\n", mainnet.Name, mainnet.ID)
		rpcs, _ := chains.GetChainRPCs(mainnet.ID)
		fmt.Printf("RPCs: %v\n", rpcs) // Note: Returns default HTTP RPCs
		// Access specific RPC providers or WS:
		// fmt.Println(mainnet.RPCUrls["default"].Http)
		// fmt.Println(mainnet.RPCUrls["infura"].WebSocket) // If defined
	}

	// Get Optimism by Name
	opMainnet, ok := chains.GetChainByName("OP Mainnet")
	if ok {
		fmt.Printf("Chain: %s, ID: %d\n", opMainnet.Name, opMainnet.ID)
		// Access specific block explorers:
		// fmt.Println(opMainnet.BlockExplorers["default"].URL)
	}
}
```

*(See `pkg/examples/predefined_usage/main.go` for more details)*

Alternatively, you can import the `predefined` package directly and use the exported variables:

```go
import (
	"fmt"
	
	"go-ethereum-chains/pkg/predefined"
)

func main() {
    fmt.Println(predefined.Mainnet.Name)
    fmt.Println(predefined.Sepolia.RPCUrls["default"].Http[0])
}
```

### Using Custom Chains

You can define and register your own chains.

```go
package main

import (
	"fmt"
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

func main() {
	myChain := chains.Chain{
		ID:   big.NewInt(31337),
		Name: "MyLocalTestnet",
		NativeCurrency: chains.NativeCurrency{Symbol: "LET", Decimals: 18},
		RPCUrls: map[string]chains.RpcTarget{
			"default": {Http: []string{"http://127.0.0.1:8545"}},
		},
		BlockExplorers: map[string]chains.BlockExplorer{
		    "default": {Name: "LocalScan", URL: "http://localhost:4000"},
		},
		IsTestnet: true,
	}

	// Register
	chains.RegisterChain(myChain)

	// Retrieve
	retrieved, ok := chains.GetChainByName("MyLocalTestnet")
	if ok {
		fmt.Printf("Found: %s\n", retrieved.Name)
	}

	// Set/Get Custom HTTP RPCs (overrides default HTTP only)
	customRPCs := []string{"http://localhost:9000"}
	_ = chains.SetChainRPCs(myChain.ID, customRPCs)
	currentRPCs, _ := chains.GetChainRPCs(myChain.ID)
	fmt.Printf("Current HTTP RPCs (via GetChainRPCs): %v\n", currentRPCs) // Output: [http://localhost:9000]

    // Access full RPC target map directly:
    fmt.Println(retrieved.RPCUrls["default"].Http)
}
```

*(See `pkg/examples/custom_usage/main.go` for more details)*


## API

### Package `pkg/chains`

*   `type Chain struct { ... }` - Represents an Ethereum compatible network.
    *   `RPCUrls map[string]RpcTarget`
    *   `BlockExplorers map[string]BlockExplorer`
*   `type NativeCurrency struct { ... }`
*   `type RpcTarget struct { Http, WebSocket []string }`
*   `type BlockExplorer struct { Name, URL string }`
*   `type Contracts struct { ... }`
*   `type Contract struct { ... }`
*   `func RegisterChain(chain Chain)` - Registers or updates a chain in the registry.
*   `func GetChainByID(id *big.Int) (Chain, bool)` - Retrieves a chain by its ID.
*   `func GetChainByName(name string) (Chain, bool)` - Retrieves a chain by its name.
*   `func SetChainRPCs(identifier any, rpcs []string) error` - Sets custom *HTTP* RPC endpoints override (for `GetChainRPCs`).
*   `func GetChainRPCs(identifier any) ([]string, error)` - Gets default *HTTP* RPCs (or custom override if set via `SetChainRPCs`). For WS or other providers, access `Chain.RPCUrls` map directly.
*   **NEW:** `type RPCStatus struct { ... }` - Holds the result of checking a single RPC endpoint (URL, Type, Availability, Latency, BlockNumber, Error).
*   **NEW:** `type CheckRPCOptions struct { ... }` - Options for checking RPCs (Timeout, CheckHTTP, CheckWS, Providers).
*   **NEW:** `func DefaultCheckOptions() CheckRPCOptions` - Returns default options for checking RPCs.
*   **NEW:** `func CheckRPCs(ctx context.Context, identifier any, opts CheckRPCOptions) ([]RPCStatus, error)` - Checks availability and latency of RPC endpoints for a given chain.
*   **NEW:** `type RPCCriteria struct { ... }` - Criteria for selecting an RPC (AllowHTTP, AllowWS, Providers).
*   **NEW:** `func DefaultRPCCriteria() RPCCriteria` - Returns default criteria for selecting RPCs.
*   **NEW:** `func GetFirstRPC(identifier any, criteria RPCCriteria) (string, error)` - Gets the first configured RPC URL matching the criteria (no availability check).
*   **NEW:** `func GetRandomRPC(identifier any, criteria RPCCriteria) (string, error)` - Gets a random configured RPC URL matching the criteria (no availability check).
*   **NEW:** `func GetFastestRPC(ctx context.Context, identifier any, criteria RPCCriteria) (string, error)` - Checks RPCs matching criteria and returns the fastest available one (uses a 1-minute cache).

*(See `pkg/examples/rpc_selection/main.go` for usage examples of RPC checking and selection)*

### Package `pkg/predefined`

This package exports variables for commonly used chains. Importing this package with `_` automatically registers these chains.

Exported variables include:
`Mainnet`, `Sepolia`, `Holesky`, `Base`, `Optimism`, `ArbitrumOne`, `ZkSync`, `Scroll`, `ArbitrumNova`, `Polygon`, `BerachainArtio`, `Avalanche`, `Bnb`, `Gnosis`, `Celo`, `Core`, `Linea`, `Fantom`, `PolygonZkEvm`, `Blast`.

## License

MIT License - see the [LICENSE](LICENSE) file for details.
