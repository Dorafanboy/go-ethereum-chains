package chains_test

import (
	"math/big"
	"reflect"
	"testing"

	"go-ethereum-chains/pkg/chains"
)

// TestRegisterChain tests the registration and overwriting of chains.
func TestRegisterChain(t *testing.T) {
	// Test case 1: Register a new chain
	chain1 := chains.Chain{
		ID:             big.NewInt(1),
		Name:           "TestChain1",
		NativeCurrency: chains.NativeCurrency{Name: "TestCoin", Symbol: "TC1", Decimals: 18},
		RPCUrls:        map[string]chains.RpcTarget{"default": {Http: []string{"http://localhost:8545"}}},
	}
	chains.RegisterChain(chain1)

	retrievedByID, okByID := chains.GetChainByID(chain1.ID)
	if !okByID || !reflect.DeepEqual(retrievedByID, chain1) {
		t.Errorf("TestRegisterChain case 1 (by ID): expected %v, got %v (found: %v)", chain1, retrievedByID, okByID)
	}

	retrievedByName, okByName := chains.GetChainByName(chain1.Name)
	if !okByName || !reflect.DeepEqual(retrievedByName, chain1) {
		t.Errorf("TestRegisterChain case 1 (by Name): expected %v, got %v (found: %v)", chain1, retrievedByName, okByName)
	}

	// Test case 2: Register another chain
	chain2 := chains.Chain{
		ID:             big.NewInt(2),
		Name:           "TestChain2",
		NativeCurrency: chains.NativeCurrency{Name: "AnotherCoin", Symbol: "AC2", Decimals: 6},
		RPCUrls:        map[string]chains.RpcTarget{"default": {Http: []string{"http://localhost:8546"}}},
	}
	chains.RegisterChain(chain2)

	retrieved2ByID, ok2ByID := chains.GetChainByID(chain2.ID)
	if !ok2ByID || !reflect.DeepEqual(retrieved2ByID, chain2) {
		t.Errorf("TestRegisterChain case 2 (by ID): expected %v, got %v (found: %v)", chain2, retrieved2ByID, ok2ByID)
	}

	// Verify chain1 still exists
	retrieved1AgainByID, ok1AgainByID := chains.GetChainByID(chain1.ID)
	if !ok1AgainByID || !reflect.DeepEqual(retrieved1AgainByID, chain1) {
		t.Errorf("TestRegisterChain case 2 (verify chain1): expected %v, got %v (found: %v)", chain1, retrieved1AgainByID, ok1AgainByID)
	}

	// Test case 3: Overwrite chain1 with new data (same ID and Name)
	chain1Overwrite := chains.Chain{
		ID:             big.NewInt(1),
		Name:           "TestChain1",
		NativeCurrency: chains.NativeCurrency{Name: "UpdatedCoin", Symbol: "TC1_UPD", Decimals: 18},
		RPCUrls:        map[string]chains.RpcTarget{"default": {Http: []string{"http://new-rpc:8545"}}},
	}
	chains.RegisterChain(chain1Overwrite)

	retrievedOverwriteByID, okOverwriteByID := chains.GetChainByID(chain1Overwrite.ID)
	if !okOverwriteByID || !reflect.DeepEqual(retrievedOverwriteByID, chain1Overwrite) {
		t.Errorf("TestRegisterChain case 3 (overwrite by ID): expected %v, got %v (found: %v)", chain1Overwrite, retrievedOverwriteByID, okOverwriteByID)
	}

	retrievedOverwriteByName, okOverwriteByName := chains.GetChainByName(chain1Overwrite.Name)
	if !okOverwriteByName || !reflect.DeepEqual(retrievedOverwriteByName, chain1Overwrite) {
		t.Errorf("TestRegisterChain case 3 (overwrite by Name): expected %v, got %v (found: %v)", chain1Overwrite, retrievedOverwriteByName, okOverwriteByName)
	}

	// Test case 4: Register chain with nil ID (should be ignored)
	nilIDChain := chains.Chain{ID: nil, Name: "NilIDChain"}
	chains.RegisterChain(nilIDChain)
	_, okNilID := chains.GetChainByName("NilIDChain")
	if okNilID {
		t.Errorf("TestRegisterChain case 4: chain with nil ID should not have been registered")
	}
}

// TestGetChainByID tests retrieving chains by their ID.
func TestGetChainByID(t *testing.T) {
	// Setup: Register a known chain for testing retrieval
	testChain := chains.Chain{ID: big.NewInt(101), Name: "GetByID_Test"}
	chains.RegisterChain(testChain)

	tests := []struct {
		name      string
		id        *big.Int
		want      chains.Chain
		wantFound bool
	}{
		{
			name:      "Found",
			id:        big.NewInt(101),
			want:      testChain,
			wantFound: true,
		},
		{
			name:      "Not Found",
			id:        big.NewInt(999),
			want:      chains.Chain{},
			wantFound: false,
		},
		{
			name:      "Nil ID",
			id:        nil,
			want:      chains.Chain{},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFound := chains.GetChainByID(tt.id)
			if gotFound != tt.wantFound {
				t.Errorf("GetChainByID() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChainByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetChainByName tests retrieving chains by their name.
func TestGetChainByName(t *testing.T) {
	testChain := chains.Chain{ID: big.NewInt(102), Name: "GetByName_Test"}
	chains.RegisterChain(testChain)

	tests := []struct {
		name      string
		chainName string
		want      chains.Chain
		wantFound bool
	}{
		{
			name:      "Found",
			chainName: "GetByName_Test",
			want:      testChain,
			wantFound: true,
		},
		{
			name:      "Not Found",
			chainName: "NonExistentChain",
			want:      chains.Chain{},
			wantFound: false,
		},
		{
			name:      "Empty Name",
			chainName: "",
			want:      chains.Chain{},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFound := chains.GetChainByName(tt.chainName)
			if gotFound != tt.wantFound {
				t.Errorf("GetChainByName() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChainByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSetAndGetChainRPCs tests setting and getting custom HTTP RPC overrides.
func TestSetAndGetChainRPCs(t *testing.T) {
	chainID := big.NewInt(202)
	chainName := "RpcTestChain"
	defaultRPCs := []string{"http://default.local"}
	testChain := chains.Chain{
		ID:      chainID,
		Name:    chainName,
		RPCUrls: map[string]chains.RpcTarget{"default": {Http: defaultRPCs}},
	}
	chains.RegisterChain(testChain)

	tests := []struct {
		name            string
		identifier      any
		setRPCs         []string
		setExpectedErr  bool
		getExpectedRPCs []string
		getExpectedErr  bool
	}{
		{
			name:            "Get default RPCs (by ID)",
			identifier:      chainID,
			setRPCs:         nil,
			setExpectedErr:  false,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  false,
		},
		{
			name:            "Get default RPCs (by Name)",
			identifier:      chainName,
			setRPCs:         nil,
			setExpectedErr:  false,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  false,
		},
		{
			name:            "Set and Get custom RPCs (by ID)",
			identifier:      chainID,
			setRPCs:         []string{"http://custom1.local", "http://custom2.local"},
			setExpectedErr:  false,
			getExpectedRPCs: []string{"http://custom1.local", "http://custom2.local"},
			getExpectedErr:  false,
		},
		{
			name:            "Get previously set custom RPCs (by Name)",
			identifier:      chainName,
			setRPCs:         nil,
			setExpectedErr:  false,
			getExpectedRPCs: []string{"http://custom1.local", "http://custom2.local"},
			getExpectedErr:  false,
		},
		{
			name:            "Override custom RPCs with new list (by Name)",
			identifier:      chainName,
			setRPCs:         []string{"http://override.local"},
			setExpectedErr:  false,
			getExpectedRPCs: []string{"http://override.local"},
			getExpectedErr:  false,
		},
		{
			name:            "Remove override by setting empty list (by ID)",
			identifier:      chainID,
			setRPCs:         []string{},
			setExpectedErr:  false,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  false,
		},
		{
			name:            "Set RPCs for non-existent ID",
			identifier:      big.NewInt(999),
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  true,
			getExpectedRPCs: nil,
			getExpectedErr:  true,
		},
		{
			name:            "Set RPCs for non-existent Name",
			identifier:      "NonExistent",
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  true,
			getExpectedRPCs: nil,
			getExpectedErr:  true,
		},
		{
			name:            "Set RPCs with invalid identifier type",
			identifier:      123,
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  true,
			getExpectedRPCs: nil,
			getExpectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setRPCs != nil {
				err := chains.SetChainRPCs(tt.identifier, tt.setRPCs)
				if (err != nil) != tt.setExpectedErr {
					t.Errorf("SetChainRPCs() error = %v, wantErr %v", err, tt.setExpectedErr)
					return
				}
			}

			gotRPCs, err := chains.GetChainRPCs(tt.identifier)
			if (err != nil) != tt.getExpectedErr {
				t.Errorf("GetChainRPCs() error = %v, wantErr %v", err, tt.getExpectedErr)
				return
			}
			if !reflect.DeepEqual(gotRPCs, tt.getExpectedRPCs) {
				t.Errorf("GetChainRPCs() = %v, want %v", gotRPCs, tt.getExpectedRPCs)
			}
		})
	}

	_ = chains.SetChainRPCs(chainID, []string{})
}
