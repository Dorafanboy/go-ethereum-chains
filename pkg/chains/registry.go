package chains

import (
	"fmt"
	"math/big"
	"sync"
)

// registryByID stores chains keyed by their chain ID (int64).
var registryByID sync.Map // map[int64]Chain

// registryByName stores chains keyed by their name (string).
var registryByName sync.Map // map[string]Chain

// userRPCs stores user-defined RPC endpoints keyed by chain ID (int64).
var userRPCs sync.Map // map[int64][]string

// RegisterChain adds or updates a chain definition in the global registries.
func RegisterChain(chain Chain) {
	if chain.ID == nil {
		// Maybe log a warning here? For now, just return.
		return
	}
	registryByID.Store(chain.ID.Int64(), chain)
	registryByName.Store(chain.Name, chain)
}

// GetChainByID retrieves a chain definition from the registry by its ID.
func GetChainByID(id *big.Int) (Chain, bool) {
	if id == nil {
		return Chain{}, false
	}
	val, ok := registryByID.Load(id.Int64())
	if !ok {
		return Chain{}, false
	}
	chain, ok := val.(Chain)
	if !ok {
		// This should ideally not happen if RegisterChain is used correctly.
		// Log an error or panic?
		return Chain{}, false
	}
	return chain, true
}

// GetChainByName retrieves a chain definition from the registry by its name.
func GetChainByName(name string) (Chain, bool) {
	if name == "" {
		return Chain{}, false
	}
	val, ok := registryByName.Load(name)
	if !ok {
		return Chain{}, false
	}
	chain, ok := val.(Chain)
	if !ok {
		// Log an error or panic?
		return Chain{}, false
	}
	return chain, true
}

// SetChainRPCs sets or overrides the RPC endpoints for a specific chain
func SetChainRPCs(identifier any, rpcs []string) error {
	var chainID int64
	var found bool

	switch id := identifier.(type) {
	case *big.Int:
		if id == nil {
			return fmt.Errorf("identifier (big.Int) cannot be nil")
		}
		_, found = GetChainByID(id)
		if !found {
			return fmt.Errorf("chain with ID %s not found in registry", id.String())
		}
		chainID = id.Int64()
	case string:
		if id == "" {
			return fmt.Errorf("identifier (string) cannot be empty")
		}
		chain, found := GetChainByName(id)
		if !found {
			return fmt.Errorf("chain with name '%s' not found in registry", id)
		}
		if chain.ID == nil { // Should not happen if registered correctly
			return fmt.Errorf("found chain '%s' but its ID is nil", id)
		}
		chainID = chain.ID.Int64()
	default:
		return fmt.Errorf("invalid identifier type: %T, expected *big.Int or string", identifier)
	}

	if len(rpcs) == 0 {
		userRPCs.Delete(chainID)
	} else {
		userRPCs.Store(chainID, rpcs)
	}

	return nil
}

// GetChainRPCs retrieves the RPC endpoints for a specific chain
func GetChainRPCs(identifier any) ([]string, error) {
	var chain Chain
	var found bool

	switch id := identifier.(type) {
	case *big.Int:
		if id == nil {
			return nil, fmt.Errorf("identifier (big.Int) cannot be nil")
		}
		chain, found = GetChainByID(id)
		if !found {
			return nil, fmt.Errorf("chain with ID %s not found in registry", id.String())
		}
	case string:
		if id == "" {
			return nil, fmt.Errorf("identifier (string) cannot be empty")
		}
		chain, found = GetChainByName(id)
		if !found {
			return nil, fmt.Errorf("chain with name '%s' not found in registry", id)
		}
	default:
		return nil, fmt.Errorf("invalid identifier type: %T, expected *big.Int or string", identifier)
	}

	if chain.ID == nil { // Defensive check
		return nil, fmt.Errorf("found chain '%s' but its ID is nil", chain.Name)
	}

	val, userrpcsOk := userRPCs.Load(chain.ID.Int64())
	if userrpcsOk {
		rpcs, typeOK := val.([]string)
		if typeOK {
			rpcsCopy := make([]string, len(rpcs))
			copy(rpcsCopy, rpcs)
			return rpcsCopy, nil
		}
	}

	defaultTarget, ok := chain.RPCUrls["default"]
	if !ok || len(defaultTarget.Http) == 0 {
		return []string{}, nil
	}

	defaultRPCs := defaultTarget.Http
	rpcsCopy := make([]string, len(defaultRPCs))
	copy(rpcsCopy, defaultRPCs)
	return rpcsCopy, nil
}
