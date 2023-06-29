package escrow_test

import (
	"testing"

	keepertest "dredd-secure/testutil/keeper"
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EscrowList: []types.Escrow{
			{
				Id: 1,
			},
			{
				Id: 0,
			},
		},
		EscrowCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EscrowKeeper(t)
	escrow.InitGenesis(ctx, *k, genesisState)
	got := escrow.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EscrowList, got.EscrowList)
	require.Equal(t, genesisState.EscrowCount, got.EscrowCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
