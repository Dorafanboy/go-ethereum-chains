package chains

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"
)

// DefaultRPCCriteria returns default criteria (HTTP only, default/public providers).
func DefaultRPCCriteria() RPCCriteria {
	return RPCCriteria{
		AllowHTTP:      true,
		AllowWebSocket: false,
		Providers:      []string{"default", "public"},
	}
}

type rpcCacheKey struct {
	identifierKey string
	providersKey  string
	checkOptsKey  string
}

type rpcCacheEntry struct {
	statuses []RPCStatus
	expiry   time.Time
}

var (
	rpcCheckCache      = make(map[rpcCacheKey]rpcCacheEntry)
	rpcCheckCacheMutex sync.RWMutex
	rpcCacheTTL        = 1 * time.Minute
)

var rpcChecker = CheckRPCs

// getIdentifierKey создает строковый ключ для идентификатора сети.
func getIdentifierKey(identifier any) (string, error) {
	switch v := identifier.(type) {
	case *big.Int:
		return v.String(), nil
	case int:
		return big.NewInt(int64(v)).String(), nil
	case int64:
		return big.NewInt(v).String(), nil
	case uint:
		return new(big.Int).SetUint64(uint64(v)).String(), nil
	case uint64:
		return new(big.Int).SetUint64(v).String(), nil
	case string:
		if id, ok := new(big.Int).SetString(v, 0); ok {
			return id.String(), nil
		}
		return v, nil
	default:
		return "", fmt.Errorf("unsupported identifier type: %T", identifier)
	}
}

// getCacheKey создает ключ для кэша на основе идентификатора и опций.
func getCacheKey(identifier any, opts CheckRPCOptions) (rpcCacheKey, error) {
	idKey, err := getIdentifierKey(identifier)
	if err != nil {
		return rpcCacheKey{}, err
	}

	// Сортируем провайдеров для консистентности ключа
	providers := make([]string, len(opts.Providers))
	copy(providers, opts.Providers)
	sort.Strings(providers)
	providersKey := fmt.Sprintf("%v", providers)

	checkOptsKey := fmt.Sprintf("http:%t-ws:%t", opts.CheckHTTP, opts.CheckWebSocket)

	return rpcCacheKey{
		identifierKey: idKey,
		providersKey:  providersKey,
		checkOptsKey:  checkOptsKey,
	}, nil
}

// GetFastestRPC checks RPCs matching criteria and returns the fastest available one (uses 1-minute cache).
func GetFastestRPC(ctx context.Context, identifier any, criteria RPCCriteria) (string, error) {
	_, err := getChain(identifier)
	if err != nil {
		return "", fmt.Errorf("failed to get chain %v: %w", identifier, err)
	}

	identifierKey, err := getIdentifierKey(identifier)
	if err != nil {
		return "", err
	}

	sortedProviders := make([]string, len(criteria.Providers))
	copy(sortedProviders, criteria.Providers)
	sort.Strings(sortedProviders)
	providersKey := fmt.Sprintf("%v", sortedProviders)

	cacheKey := rpcCacheKey{
		identifierKey: identifierKey,
		providersKey:  providersKey,
		checkOptsKey:  fmt.Sprintf("http:%t-ws:%t", criteria.AllowHTTP, criteria.AllowWebSocket),
	}

	rpcCheckCacheMutex.RLock()
	entry, foundInCache := rpcCheckCache[cacheKey]
	rpcCheckCacheMutex.RUnlock()

	var statuses []RPCStatus
	if foundInCache && time.Now().Before(entry.expiry) {
		statuses = entry.statuses
	} else {
		checkOpts := CheckRPCOptions{
			TimeoutPerCheck: DefaultCheckOptions().TimeoutPerCheck,
			CheckHTTP:       criteria.AllowHTTP,
			CheckWebSocket:  criteria.AllowWebSocket,
			Providers:       criteria.Providers,
		}

		checkedStatuses, checkErr := rpcChecker(ctx, identifier, checkOpts)
		if checkErr != nil {
			return "", fmt.Errorf("failed to check RPCs for %v: %w", identifier, checkErr)
		}

		rpcCheckCacheMutex.Lock()
		rpcCheckCache[cacheKey] = rpcCacheEntry{
			statuses: checkedStatuses,
			expiry:   time.Now().Add(rpcCacheTTL),
		}
		rpcCheckCacheMutex.Unlock()
		statuses = checkedStatuses
	}

	var bestURL string
	minLatency := time.Duration(-1)

	for _, s := range statuses {
		if s.IsAvailable {
			if (criteria.AllowHTTP && s.IsHTTP) || (criteria.AllowWebSocket && s.IsWebSocket) {
				if minLatency == -1 || s.Latency < minLatency {
					minLatency = s.Latency
					bestURL = s.URL
				}
			}
		}
	}

	if bestURL == "" {
		return "", fmt.Errorf("no available RPC found matching criteria for %v", identifier)
	}

	return bestURL, nil
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

// TODO: Add cache realized for GetFastestRPC
