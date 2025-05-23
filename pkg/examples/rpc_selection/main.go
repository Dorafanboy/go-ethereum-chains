package main

import (
	"context"
	"fmt"
	"go-ethereum-chains/pkg/selector"
	"math/big"
	"time"

	"go-ethereum-chains/pkg/chains"

	_ "go-ethereum-chains/pkg/predefined"
)

func main() {
	fmt.Println("--- Example: Checking and Selecting RPCs ---")

	ctx := context.Background()
	fmt.Println("\n--- Checking Sepolia RPCs (Default Providers, HTTP & WS) ---")
	sepoliaID := big.NewInt(11155111)
	ctxCheck, cancelCheck := context.WithTimeout(ctx, 15*time.Second)
	defer cancelCheck()

	checkOpts := chains.DefaultCheckOptions()
	checkOpts.TimeoutPerCheck = 4 * time.Second

	statuses, err := chains.CheckRPCs(ctxCheck, sepoliaID, checkOpts)
	if err != nil {
		fmt.Printf("Error checking Sepolia RPCs: %v\n", err)
	} else {
		fmt.Printf("Checked %d RPC endpoints for Sepolia:\n", len(statuses))
		for _, s := range statuses {
			fmt.Printf("  URL: %s\n", s.URL)
			fmt.Printf("    Type: %s\n", map[bool]string{true: "HTTP", false: "WebSocket"}[s.IsHTTP])
			fmt.Printf("    Available: %v\n", s.IsAvailable)
			if s.IsAvailable {
				fmt.Printf("    Latency: %v\n", s.Latency)
				fmt.Printf("    Block Number: %v\n", s.BlockNumber)
			} else {
				fmt.Printf("    Error: %v\n", s.Error)
			}
		}
	}

	fmt.Println("\n--- Selecting Sepolia RPCs ---")

	// 1. Get the first available HTTP RPC from default/public providers
	firstCrit := selector.DefaultRPCCriteria()
	firstRPC, err := selector.GetFirstRPC(sepoliaID, firstCrit)
	if err != nil {
		fmt.Printf("Error getting first RPC: %v\n", err)
	} else {
		fmt.Printf("First matching HTTP RPC: %s\n", firstRPC)
	}

	// 2. Get a random HTTP RPC from default/public providers
	randomCrit := selector.DefaultRPCCriteria()
	randomRPC, err := selector.GetRandomRPC(sepoliaID, randomCrit)
	if err != nil {
		fmt.Printf("Error getting random RPC: %v\n", err)
	} else {
		fmt.Printf("Random matching HTTP RPC: %s\n", randomRPC)
	}
}
