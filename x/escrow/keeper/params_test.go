package keeper_test

import (
	"testing"

	testkeeper "dredd-secure/testutil/keeper"
	"dredd-secure/x/escrow/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EscrowKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
