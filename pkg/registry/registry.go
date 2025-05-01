package registry

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"go-ethereum-chains/internal/types"
)

// ErrChainNotFound is returned when a chain is not found in the registry.
var ErrChainNotFound = errors.New("chain not found")

// registryByID stores chains keyed by their chain ID (int64).
var registryByID sync.Map

// registryByName stores chains keyed by their name (string).
var registryByName sync.Map

// userRPCs stores user-defined RPC endpoints keyed by chain ID (int64).
var userRPCs sync.Map

// RegisterChain adds or updates a chain definition in the global registries.
func RegisterChain(chain types.Chain) {
	if chain.ID == nil {
		// Maybe log a warning here? For now, just return.
		return
	}
	registryByID.Store(chain.ID.Int64(), chain)
	if chain.Name != "" {
		registryByName.Store(chain.Name, chain)
	}
}

// GetChainByID retrieves a chain definition from the registry by its ID.
func GetChainByID(id *big.Int) (types.Chain, bool) {
	if id == nil {
		return types.Chain{}, false
	}
	val, ok := registryByID.Load(id.Int64())
	if !ok {
		return types.Chain{}, false
	}
	chain, ok := val.(types.Chain)
	if !ok {
		// This should ideally not happen if RegisterChain is used correctly.
		// Log an error or panic?
		return types.Chain{}, false
	}
	return chain, true
}

// GetChainByName retrieves a chain definition from the registry by its name.
func GetChainByName(name string) (types.Chain, bool) {
	if name == "" {
		return types.Chain{}, false
	}
	val, ok := registryByName.Load(name)
	if !ok {
		return types.Chain{}, false
	}
	chain, ok := val.(types.Chain)
	if !ok {
		// Log an error or panic?
		return types.Chain{}, false
	}
	return chain, true
}

// SetChainRPCs sets or overrides the RPC endpoints for a specific chain
func SetChainRPCs(identifier any, rpcs []string) error {
	chain, err := FindChain(identifier)
	if err != nil {
		return err
	}

	if chain.ID == nil {
		return fmt.Errorf("found chain '%s' but its ID is nil", chain.Name)
	}
	chainID := chain.ID.Int64()

	if len(rpcs) == 0 {
		userRPCs.Delete(chainID)
	} else {
		rpcsCopy := make([]string, len(rpcs))
		copy(rpcsCopy, rpcs)
		userRPCs.Store(chainID, rpcsCopy)
	}

	return nil
}

// GetChainRPCs retrieves the RPC endpoints for a specific chain
func GetChainRPCs(identifier any) ([]string, error) {
	chain, err := FindChain(identifier)
	if err != nil {
		return nil, err
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
		return []string{}, nil // Return empty slice, not nil
	}

	defaultRPCs := defaultTarget.Http
	rpcsCopy := make([]string, len(defaultRPCs))
	copy(rpcsCopy, defaultRPCs)
	return rpcsCopy, nil
}

// FindChain is an internal helper to retrieve a chain by ID or name.
func FindChain(identifier any) (types.Chain, error) {
	switch id := identifier.(type) {
	case *big.Int:
		chain, found := GetChainByID(id)
		if !found {
			return types.Chain{}, fmt.Errorf("%w: ID %s", ErrChainNotFound, id.String())
		}
		return chain, nil
	case int:
		return FindChain(big.NewInt(int64(id)))
	case int64:
		return FindChain(big.NewInt(id))
	case uint:
		return FindChain(new(big.Int).SetUint64(uint64(id)))
	case uint64:
		return FindChain(new(big.Int).SetUint64(id))
	case string:
		if id == "" {
			return types.Chain{}, fmt.Errorf("identifier (string) cannot be empty")
		}
		if parsedID, ok := new(big.Int).SetString(id, 0); ok {
			chain, found := GetChainByID(parsedID)
			if found {
				return chain, nil
			}
		}
		chain, found := GetChainByName(id)
		if !found {
			return types.Chain{}, fmt.Errorf("%w: Name '%s' (or ID parse failed)", ErrChainNotFound, id)
		}
		return chain, nil
	default:
		return types.Chain{}, fmt.Errorf("invalid identifier type: %T, expected *big.Int or string", identifier)
	}
}
