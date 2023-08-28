package escrow_test

import (
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/types"
	"testing"

	keepertest "dredd-secure/testutil/keeper"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		EscrowList: []types.Escrow{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		EscrowCount: 2,
		OraclePriceList: []types.OraclePrice{
			{
				Symbol: "0",
			},
			{
				Symbol: "1",
			},
		},
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
	require.Equal(t, genesisState.PortId, got.PortId)
	require.ElementsMatch(t, genesisState.OraclePriceList, got.OraclePriceList)
	// this line is used by starport scaffolding # genesis/test/assert
}
