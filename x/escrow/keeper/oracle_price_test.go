package keeper_test

import (
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"
	"strconv"
	"testing"

	keepertest "dredd-secure/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNOraclePrice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.OraclePrice {
	items := make([]types.OraclePrice, n)
	for i := range items {
		items[i].Symbol = strconv.Itoa(i)

		keeper.SetOraclePrice(ctx, items[i])
	}
	return items
}

func TestOraclePriceGet(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNOraclePrice(keeper, ctx, 10)
	for i := range items {
		rst, found := keeper.GetOraclePrice(ctx,
			items[i].Symbol,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&items[i]),
			nullify.Fill(&rst),
		)
	}
}

func TestOraclePriceRemove(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNOraclePrice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOraclePrice(ctx,
			item.Symbol,
		)
		_, found := keeper.GetOraclePrice(ctx,
			item.Symbol,
		)
		require.False(t, found)
	}
}

func TestOraclePriceGetAll(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	items := createNOraclePrice(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOraclePrice(ctx)),
	)
}
