package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EscrowList: []Escrow{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in escrow
	escrowIdMap := make(map[uint64]bool)
	escrowCount := gs.GetEscrowCount()
	for _, elem := range gs.EscrowList {
		if _, ok := escrowIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for escrow")
		}
		if elem.Id >= escrowCount {
			return fmt.Errorf("escrow id should be lower or equal than the last id")
		}
		escrowIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
