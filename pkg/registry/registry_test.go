package registry_test

import (
	"errors"
	"math/big"
	"reflect"
	"testing"

	"go-ethereum-chains/internal/types"
	"go-ethereum-chains/pkg/registry"
)

// TestRegisterChain tests the registration and overwriting of chains.
func TestRegisterChain(t *testing.T) {
	// Test case 1: Register a new chain
	chain1 := types.Chain{
		ID:             big.NewInt(1),
		Name:           "TestChain1",
		NativeCurrency: types.NativeCurrency{Name: "TestCoin", Symbol: "TC1", Decimals: 18},
		RPCUrls:        map[string]types.RpcTarget{"default": {Http: []string{"http://localhost:8545"}}},
	}
	registry.RegisterChain(chain1)

	retrievedByID, okByID := registry.GetChainByID(chain1.ID)
	if !okByID || !reflect.DeepEqual(retrievedByID, chain1) {
		t.Errorf("TestRegisterChain case 1 (by ID): expected %v, got %v (found: %v)", chain1, retrievedByID, okByID)
	}

	retrievedByName, okByName := registry.GetChainByName(chain1.Name)
	if !okByName || !reflect.DeepEqual(retrievedByName, chain1) {
		t.Errorf("TestRegisterChain case 1 (by Name): expected %v, got %v (found: %v)", chain1, retrievedByName, okByName)
	}

	// Test case 2: Register another chain
	chain2 := types.Chain{
		ID:             big.NewInt(2),
		Name:           "TestChain2",
		NativeCurrency: types.NativeCurrency{Name: "AnotherCoin", Symbol: "AC2", Decimals: 6},
		RPCUrls:        map[string]types.RpcTarget{"default": {Http: []string{"http://localhost:8546"}}},
	}
	registry.RegisterChain(chain2)

	retrieved2ByID, ok2ByID := registry.GetChainByID(chain2.ID)
	if !ok2ByID || !reflect.DeepEqual(retrieved2ByID, chain2) {
		t.Errorf("TestRegisterChain case 2 (by ID): expected %v, got %v (found: %v)", chain2, retrieved2ByID, ok2ByID)
	}

	retrieved1AgainByID, ok1AgainByID := registry.GetChainByID(chain1.ID)
	if !ok1AgainByID || !reflect.DeepEqual(retrieved1AgainByID, chain1) {
		t.Errorf("TestRegisterChain case 2 (verify chain1): expected %v, got %v (found: %v)", chain1, retrieved1AgainByID, ok1AgainByID)
	}

	// Test case 3: Overwrite chain1 with new data (same ID and Name)
	chain1Overwrite := types.Chain{
		ID:             big.NewInt(1),
		Name:           "TestChain1",
		NativeCurrency: types.NativeCurrency{Name: "UpdatedCoin", Symbol: "TC1_UPD", Decimals: 18},
		RPCUrls:        map[string]types.RpcTarget{"default": {Http: []string{"http://new-rpc:8545"}}},
	}
	registry.RegisterChain(chain1Overwrite)

	retrievedOverwriteByID, okOverwriteByID := registry.GetChainByID(chain1Overwrite.ID)
	if !okOverwriteByID || !reflect.DeepEqual(retrievedOverwriteByID, chain1Overwrite) {
		t.Errorf("TestRegisterChain case 3 (overwrite by ID): expected %v, got %v (found: %v)", chain1Overwrite, retrievedOverwriteByID, okOverwriteByID)
	}

	retrievedOverwriteByName, okOverwriteByName := registry.GetChainByName(chain1Overwrite.Name)
	if !okOverwriteByName || !reflect.DeepEqual(retrievedOverwriteByName, chain1Overwrite) {
		t.Errorf("TestRegisterChain case 3 (overwrite by Name): expected %v, got %v (found: %v)", chain1Overwrite, retrievedOverwriteByName, okOverwriteByName)
	}

	// Test case 4: Register chain with nil ID (should be ignored)
	nilIDChain := types.Chain{ID: nil, Name: "NilIDChain"}
	registry.RegisterChain(nilIDChain)
	_, okNilID := registry.GetChainByName("NilIDChain")
	if okNilID {
		t.Errorf("TestRegisterChain case 4: chain with nil ID should not have been registered")
	}

	// Test case 5: Register chain with empty name (should register by ID)
	emptyNameChain := types.Chain{ID: big.NewInt(99), Name: ""}
	registry.RegisterChain(emptyNameChain)
	retrievedEmptyNameByID, okEmptyNameByID := registry.GetChainByID(emptyNameChain.ID)
	if !okEmptyNameByID || !reflect.DeepEqual(retrievedEmptyNameByID, emptyNameChain) {
		t.Errorf("TestRegisterChain case 5 (empty name by ID): expected %v, got %v (found: %v)", emptyNameChain, retrievedEmptyNameByID, okEmptyNameByID)
	}
	_, okEmptyNameByName := registry.GetChainByName("")
	if okEmptyNameByName {
		t.Errorf("TestRegisterChain case 5: chain with empty name should not be retrievable by empty name")
	}
}

// TestGetChainByID tests retrieving chains by their ID.
func TestGetChainByID(t *testing.T) {
	// Setup: Register a known chain for testing retrieval
	testChain := types.Chain{ID: big.NewInt(101), Name: "GetByID_Test"}
	registry.RegisterChain(testChain)

	tests := []struct {
		name      string
		id        *big.Int
		want      types.Chain
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
			want:      types.Chain{},
			wantFound: false,
		},
		{
			name:      "Nil ID",
			id:        nil,
			want:      types.Chain{},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFound := registry.GetChainByID(tt.id)
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
	testChain := types.Chain{ID: big.NewInt(102), Name: "GetByName_Test"}
	registry.RegisterChain(testChain)

	tests := []struct {
		name      string
		chainName string
		want      types.Chain
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
			want:      types.Chain{},
			wantFound: false,
		},
		{
			name:      "Empty Name",
			chainName: "",
			want:      types.Chain{},
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotFound := registry.GetChainByName(tt.chainName)
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
	testChain := types.Chain{
		ID:      chainID,
		Name:    chainName,
		RPCUrls: map[string]types.RpcTarget{"default": {Http: defaultRPCs}},
	}
	registry.RegisterChain(testChain)

	_ = registry.SetChainRPCs(chainID, []string{})

	tests := []struct {
		name            string
		identifier      any
		setRPCs         []string
		setExpectedErr  error
		getExpectedRPCs []string
		getExpectedErr  error
	}{
		{
			name:            "Get default RPCs (by ID)",
			identifier:      chainID,
			setRPCs:         nil,
			setExpectedErr:  nil,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  nil,
		},
		{
			name:            "Get default RPCs (by Name)",
			identifier:      chainName,
			setRPCs:         nil,
			setExpectedErr:  nil,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  nil,
		},
		{
			name:            "Set and Get custom RPCs (by ID)",
			identifier:      chainID,
			setRPCs:         []string{"http://custom1.local", "http://custom2.local"},
			setExpectedErr:  nil,
			getExpectedRPCs: []string{"http://custom1.local", "http://custom2.local"},
			getExpectedErr:  nil,
		},
		{
			name:            "Get previously set custom RPCs (by Name)",
			identifier:      chainName,
			setRPCs:         nil,
			setExpectedErr:  nil,
			getExpectedRPCs: []string{"http://custom1.local", "http://custom2.local"},
			getExpectedErr:  nil,
		},
		{
			name:            "Override custom RPCs with new list (by Name)",
			identifier:      chainName,
			setRPCs:         []string{"http://override.local"},
			setExpectedErr:  nil,
			getExpectedRPCs: []string{"http://override.local"},
			getExpectedErr:  nil,
		},
		{
			name:            "Remove override by setting empty list (by ID)",
			identifier:      chainID,
			setRPCs:         []string{},
			setExpectedErr:  nil,
			getExpectedRPCs: defaultRPCs,
			getExpectedErr:  nil,
		},
		{
			name:            "Set RPCs for non-existent ID",
			identifier:      big.NewInt(999),
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  registry.ErrChainNotFound,
			getExpectedRPCs: nil,
			getExpectedErr:  registry.ErrChainNotFound,
		},
		{
			name:            "Set RPCs for non-existent Name",
			identifier:      "NonExistent",
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  registry.ErrChainNotFound,
			getExpectedRPCs: nil,
			getExpectedErr:  registry.ErrChainNotFound,
		},
		{
			name:            "Set RPCs with invalid identifier type",
			identifier:      123,
			setRPCs:         []string{"http://shouldfail.local"},
			setExpectedErr:  errors.New(""),
			getExpectedRPCs: nil,
			getExpectedErr:  errors.New(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setRPCs != nil {
				err := registry.SetChainRPCs(tt.identifier, tt.setRPCs)
				if tt.setExpectedErr == nil && err != nil {
					t.Errorf("SetChainRPCs() unexpected error = %v", err)
					return
				}
				if tt.setExpectedErr != nil {
					if err == nil {
						t.Errorf("SetChainRPCs() expected error '%v', got nil", tt.setExpectedErr)
						return
					}
					if !errors.Is(err, tt.setExpectedErr) && tt.setExpectedErr.Error() != "" {
						t.Errorf("SetChainRPCs() error type = %T, want target type %T (or specific error check failed: %v vs %v)", err, tt.setExpectedErr, err, tt.setExpectedErr)
					}
				}
			}

			gotRPCs, err := registry.GetChainRPCs(tt.identifier)
			if tt.getExpectedErr == nil && err != nil {
				t.Errorf("GetChainRPCs() unexpected error = %v", err)
				return
			}
			if tt.getExpectedErr != nil {
				if err == nil {
					t.Errorf("GetChainRPCs() expected error '%v', got nil", tt.getExpectedErr)
					return
				}
				if !errors.Is(err, tt.getExpectedErr) && tt.getExpectedErr.Error() != "" {
					t.Errorf("GetChainRPCs() error type = %T, want target type %T (or specific error check failed: %v vs %v)", err, tt.getExpectedErr, err, tt.getExpectedErr)
				}
			}

			if !reflect.DeepEqual(gotRPCs, tt.getExpectedRPCs) {
				t.Errorf("GetChainRPCs() = %v, want %v", gotRPCs, tt.getExpectedRPCs)
			}
		})
	}

	_ = registry.SetChainRPCs(chainID, []string{})
}
