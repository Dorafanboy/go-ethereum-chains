package main

import (
	"fmt"
	"math/big"

	"go-ethereum-chains/pkg/chains"
)

func main() {
	fmt.Println("--- Example: Registering and Using Custom Chains ---")

	// 1. Define a custom chain
	myChainID := big.NewInt(31337)
	myChainName := "MyLocalTestnet"
	myChain := chains.Chain{
		ID:   myChainID,
		Name: myChainName,
		NativeCurrency: chains.NativeCurrency{
			Name:     "Local Ether",
			Symbol:   "LET",
			Decimals: 18,
		},
		RPCUrls: map[string]chains.RpcTarget{
			"default": {
				Http: []string{"http://127.0.0.1:8545"},
			},
		},
		IsTestnet: true,
	}

	// 2. Register the custom chain
	chains.RegisterChain(myChain)
	fmt.Printf("Registered custom chain: %s (ID: %d)\n", myChainName, myChainID)

	// 3. Retrieve the custom chain
	retrievedByID, ok1 := chains.GetChainByID(myChainID)
	if ok1 {
		fmt.Printf("Retrieved custom chain by ID: %s\n", retrievedByID.Name)
	} else {
		fmt.Printf("Failed to retrieve custom chain by ID %d\n", myChainID)
	}

	retrievedByName, ok2 := chains.GetChainByName(myChainName)
	if ok2 {
		fmt.Printf("Retrieved custom chain by Name: %s\n", retrievedByName.Name)
	} else {
		fmt.Printf("Failed to retrieve custom chain by Name %s\n", myChainName)
	}

	// 4. Get default RPCs for the custom chain
	defaultRPCs, err := chains.GetChainRPCs(myChainID)
	if err != nil {
		fmt.Printf("Error getting default RPCs for %s: %v\n", myChainName, err)
	} else {
		fmt.Printf("Default RPCs for %s: %v\n", myChainName, defaultRPCs)
	}

	// 5. Set custom RPCs for the custom chain
	customRPCs := []string{"http://localhost:9000", "ws://localhost:9001"}
	err = chains.SetChainRPCs(myChainName, customRPCs)
	if err != nil {
		fmt.Printf("Error setting custom RPCs for %s: %v\n", myChainName, err)
	} else {
		fmt.Printf("Set custom RPCs for %s: %v\n", myChainName, customRPCs)
	}

	// 6. Get RPCs again (should now be the custom ones)
	currentRPCs, err := chains.GetChainRPCs(myChainID)
	if err != nil {
		fmt.Printf("Error getting RPCs for %s after setting custom: %v\n", myChainName, err)
	} else {
		fmt.Printf("Current RPCs for %s: %v (should be custom)\n", myChainName, currentRPCs)
	}

	// 7. Remove custom RPC override by setting an empty list
	err = chains.SetChainRPCs(myChainID, []string{})
	if err != nil {
		fmt.Printf("Error removing custom RPCs for %s: %v\n", myChainName, err)
	} else {
		fmt.Println("Removed custom RPC override.")
	}

	// 8. Get RPCs one last time (should be default again)
	finalRPCs, err := chains.GetChainRPCs(myChainName)
	if err != nil {
		fmt.Printf("Error getting RPCs for %s after removing custom: %v\n", myChainName, err)
	} else {
		fmt.Printf("Final RPCs for %s: %v (should be default again)\n", myChainName, finalRPCs)
	}
}
