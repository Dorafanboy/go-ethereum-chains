package chains

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
)

// DefaultRPCCriteria returns default criteria (HTTP only, default/public providers).
func DefaultRPCCriteria() RPCCriteria {
	return RPCCriteria{
		AllowHTTP:      true,
		AllowWebSocket: false,
		Providers:      []string{"default", "public"},
	}
}

// GetRandomRPC selects a random configured RPC URL matching criteria using crypto/rand (no availability check).
func GetRandomRPC(identifier any, criteria RPCCriteria) (string, error) {
	chain, err := getChain(identifier)
	if err != nil {
		return "", fmt.Errorf("failed to get chain %v: %w", identifier, err)
	}

	var candidates []string
	providersToCheck := criteria.Providers
	if len(providersToCheck) == 0 { // If empty, check all providers
		providersToCheck = make([]string, 0, len(chain.RPCUrls))
		for provider := range chain.RPCUrls {
			providersToCheck = append(providersToCheck, provider)
		}
	}

	for _, provider := range providersToCheck {
		if target, ok := chain.RPCUrls[provider]; ok {
			if criteria.AllowHTTP {
				candidates = append(candidates, target.Http...)
			}
			if criteria.AllowWebSocket {
				candidates = append(candidates, target.WebSocket...)
			}
		}
	}

	if len(candidates) == 0 {
		return "", fmt.Errorf("no RPC URL found matching criteria for %v", identifier)
	}

	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(candidates))))
	if err != nil {
		return "", fmt.Errorf("failed to generate random index: %w", err)
	}

	return candidates[idx.Int64()], nil
}

// GetFirstRPC finds the first configured RPC URL matching criteria (no availability check).
func GetFirstRPC(identifier any, criteria RPCCriteria) (string, error) {
	chain, err := getChain(identifier)
	if err != nil {
		return "", fmt.Errorf("failed to get chain %v: %w", identifier, err)
	}

	providersToCheck := criteria.Providers
	if len(providersToCheck) == 0 {
		providersToCheck = make([]string, 0, len(chain.RPCUrls))
		for provider := range chain.RPCUrls {
			providersToCheck = append(providersToCheck, provider)
		}
		sort.Strings(providersToCheck)
	}

	for _, provider := range providersToCheck {
		if target, ok := chain.RPCUrls[provider]; ok {
			if criteria.AllowHTTP {
				if len(target.Http) > 0 {
					return target.Http[0], nil
				}
			}
			if criteria.AllowWebSocket {
				if len(target.WebSocket) > 0 {
					return target.WebSocket[0], nil
				}
			}
		}
	}

	return "", fmt.Errorf("no RPC URL found matching criteria for %v", identifier)
}
