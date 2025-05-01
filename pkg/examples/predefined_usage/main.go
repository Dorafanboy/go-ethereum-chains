package main

import (
	"fmt"
	"math/big"

	"go-ethereum-chains/pkg/chains"

	_ "go-ethereum-chains/pkg/predefined"
)

func main() {
	fmt.Println("--- Example: Accessing Predefined Chains ---")

	// Access via registry functions (after importing predefined)
	mainnetByID, ok := chains.GetChainByID(big.NewInt(1))
	if ok {
		fmt.Printf("Found via GetChainByID(1): %s (ID: %d)\n", mainnetByID.Name, mainnetByID.ID)
		rpcs, err := chains.GetChainRPCs(mainnetByID.ID)
		if err != nil {
			fmt.Printf("  Error getting RPCs: %v\n", err)
		} else {
			fmt.Printf("  Default RPCs: %v\n", rpcs)
		}
	} else {
		fmt.Println("Mainnet (ID 1) not found in registry!")
	}

	optimismByName, ok := chains.GetChainByName("OP Mainnet")
	if ok {
		fmt.Printf("\nFound via GetChainByName('OP Mainnet'): %s (ID: %d)\n", optimismByName.Name, optimismByName.ID)
		rpcs, err := chains.GetChainRPCs(optimismByName.Name)
		if err != nil {
			fmt.Printf("  Error getting RPCs: %v\n", err)
		} else {
			fmt.Printf("  Default RPCs: %v\n", rpcs)
		}
	} else {
		fmt.Println("OP Mainnet not found in registry!")
	}

	// Example: Get RPCs for Sepolia directly
	sepoliaID := big.NewInt(11155111)
	sepoliaRPCs, err := chains.GetChainRPCs(sepoliaID)
	if err != nil {
		fmt.Printf("\nError getting Sepolia RPCs: %v\n", err)
	} else {
		fmt.Printf("\nSepolia (ID: %d) RPCs: %v\n", sepoliaID, sepoliaRPCs)
	}

	// Example: Non-existent chain
	_, ok3 := chains.GetChainByID(big.NewInt(999999))
	if !ok3 {
		fmt.Println("\nChain with ID 999999 correctly not found.")
	}

	fmt.Println("\nNote: You can also access predefined chains directly via variables,")
	fmt.Println("e.g., by importing 'your_module/predefined' and using 'predefined.Mainnet'.")
	fmt.Println("This example uses the registry access method after blank import.")
}
