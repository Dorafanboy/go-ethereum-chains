package selector

import (
	"math/big"
	"slices"
	"testing"

	"go-ethereum-chains/internal/types"
	"go-ethereum-chains/pkg/registry"
)

// Mocks and test data setup
var testChainID = big.NewInt(9999)
var testChainName = "Test Chain Selector"
var testChain = types.Chain{
	ID:   testChainID,
	Name: testChainName,
	NativeCurrency: types.NativeCurrency{
		Name:     "Test",
		Symbol:   "TST",
		Decimals: 18,
	},
	RPCUrls: map[string]types.RpcTarget{
		"default": {
			Http:      []string{"http://default1.com", "http://default2.com"},
			WebSocket: []string{"ws://default1.com", "ws://default2.com"},
		},
		"public": {
			Http: []string{"http://public1.com"},
		},
		"specific": {
			WebSocket: []string{"ws://specific1.com"},
		},
	},
	IsTestnet: true,
}

// TestGetFirstRPC tests the GetFirstRPC function.
func TestGetFirstRPC(t *testing.T) {
	setupSelectorTest()

	tests := []struct {
		name       string
		identifier any
		criteria   RPCCriteria
		want       string
		wantErr    bool
	}{
		{
			name:       "Default HTTP First",
			identifier: testChainID,
			criteria:   DefaultRPCCriteria(),
			want:       "http://default1.com",
			wantErr:    false,
		},
		{
			name:       "Default WS First (HTTP disallowed)",
			identifier: testChainName,
			criteria:   RPCCriteria{AllowHTTP: false, AllowWebSocket: true, Providers: []types.ProviderName{types.ProviderDefault}},
			want:       "ws://default1.com",
			wantErr:    false,
		},
		{
			name:       "Public HTTP First (Default has no HTTP allowed)",
			identifier: testChainID,
			criteria:   RPCCriteria{AllowHTTP: true, AllowWebSocket: false, Providers: []types.ProviderName{types.ProviderName("specific"), types.ProviderPublic}},
			want:       "http://public1.com",
			wantErr:    false,
		},
		{
			name:       "Specific WS First",
			identifier: testChainID,
			criteria:   RPCCriteria{AllowHTTP: false, AllowWebSocket: true, Providers: []types.ProviderName{types.ProviderName("specific")}},
			want:       "ws://specific1.com",
			wantErr:    false,
		},
		{
			name:       "No Match Found (Specific provider, only HTTP allowed)",
			identifier: testChainID,
			criteria:   RPCCriteria{AllowHTTP: true, AllowWebSocket: false, Providers: []types.ProviderName{types.ProviderName("specific")}},
			want:       "",
			wantErr:    true,
		},
		{
			name:       "No Match Found (WS disallowed)",
			identifier: testChainID,
			criteria:   RPCCriteria{AllowHTTP: false, AllowWebSocket: false, Providers: []types.ProviderName{types.ProviderDefault}},
			want:       "",
			wantErr:    true,
		},
		{
			name:       "Chain Not Found",
			identifier: big.NewInt(12345),
			criteria:   DefaultRPCCriteria(),
			want:       "",
			wantErr:    true,
		},
		{
			name:       "Invalid Identifier Type",
			identifier: 123,
			criteria:   DefaultRPCCriteria(),
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFirstRPC(tt.identifier, tt.criteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirstRPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFirstRPC() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetRandomRPC tests the GetRandomRPC function.
func TestGetRandomRPC(t *testing.T) {
	setupSelectorTest()

	tests := []struct {
		name           string
		identifier     any
		criteria       RPCCriteria
		wantCandidates []string
		wantErr        bool
	}{
		{
			name:           "Default and Public HTTP",
			identifier:     testChainID,
			criteria:       DefaultRPCCriteria(),
			wantCandidates: []string{"http://default1.com", "http://default2.com", "http://public1.com"},
			wantErr:        false,
		},
		{
			name:           "Default WS Only",
			identifier:     testChainName,
			criteria:       RPCCriteria{AllowHTTP: false, AllowWebSocket: true, Providers: []types.ProviderName{types.ProviderDefault}},
			wantCandidates: []string{"ws://default1.com", "ws://default2.com"},
			wantErr:        false,
		},
		{
			name:           "Specific Provider WS",
			identifier:     testChainID,
			criteria:       RPCCriteria{AllowHTTP: false, AllowWebSocket: true, Providers: []types.ProviderName{types.ProviderName("specific")}},
			wantCandidates: []string{"ws://specific1.com"},
			wantErr:        false,
		},
		{
			name:       "All Providers, All Types",
			identifier: testChainID,
			criteria:   RPCCriteria{AllowHTTP: true, AllowWebSocket: true, Providers: []types.ProviderName{types.ProviderDefault, types.ProviderPublic, types.ProviderName("specific")}},
			wantCandidates: []string{
				"http://default1.com", "http://default2.com",
				"ws://default1.com", "ws://default2.com",
				"http://public1.com",
				"ws://specific1.com",
			},
			wantErr: false,
		},
		{
			name:           "No Match Found (Specific provider, only HTTP allowed)",
			identifier:     testChainID,
			criteria:       RPCCriteria{AllowHTTP: true, AllowWebSocket: false, Providers: []types.ProviderName{types.ProviderName("specific")}},
			wantCandidates: nil,
			wantErr:        true,
		},
		{
			name:           "Chain Not Found",
			identifier:     big.NewInt(12345),
			criteria:       DefaultRPCCriteria(),
			wantCandidates: nil,
			wantErr:        true,
		},
		{
			name:           "Invalid Identifier Type",
			identifier:     123,
			criteria:       DefaultRPCCriteria(),
			wantCandidates: nil,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := make(map[string]bool)
			for i := 0; i < 20; i++ {
				got, err := GetRandomRPC(tt.identifier, tt.criteria)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetRandomRPC() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if err != nil { // If error was expected, stop here for this iteration
					break
				}
				if !slices.Contains(tt.wantCandidates, got) {
					t.Errorf("GetRandomRPC() got = %v, which is not in wantCandidates %v", got, tt.wantCandidates)
					return
				}
				results[got] = true
			}

			if !tt.wantErr && len(tt.wantCandidates) > 1 && len(results) < 2 {
				t.Logf("GetRandomRPC() produced the same result multiple times: %v. This might be okay, but indicates low randomness in this run.", results)
			}
			if !tt.wantErr && len(results) > len(tt.wantCandidates) {
				t.Errorf("GetRandomRPC() produced more unique results (%d) than candidates (%d)", len(results), len(tt.wantCandidates))
			}
		})
	}
}

func setupSelectorTest() {
	registry.RegisterChain(testChain)
}
