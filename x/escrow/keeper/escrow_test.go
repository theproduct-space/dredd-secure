package keeper_test

import (
	"testing"

	keepertest "dredd-secure/testutil/keeper"
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNEscrow(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Escrow {
	items := make([]types.Escrow, n)
	for i := range items {
		items[i].Id = keeper.AppendEscrow(ctx, items[i])
	}
	return items
}

func TestEscrowGet(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNEscrow(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetEscrow(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestEscrowRemove(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNEscrow(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEscrow(ctx, item.Id)
		_, found := keeper.GetEscrow(ctx, item.Id)
		require.False(t, found)
	}
}

func TestEscrowGetAll(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNEscrow(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEscrow(ctx)),
	)
}

func TestEscrowCount(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNEscrow(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetEscrowCount(ctx))
}
