package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EscrowList: []Escrow{},
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
		PendingEscrows: []uint64{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// Check for duplicated ID in escrow
	escrowIDMap := make(map[uint64]bool)
	escrowCount := gs.GetEscrowCount()
	for _, elem := range gs.EscrowList {
		if _, ok := escrowIDMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for escrow")
		}
		if elem.Id >= escrowCount {
			return fmt.Errorf("escrow id should be lower or equal than the last id")
		}
		escrowIDMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
