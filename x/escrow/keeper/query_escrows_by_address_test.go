package keeper_test

import (
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow/types"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "dredd-secure/testutil/keeper"
)

func TestEscrowQueryByAddress(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	addresses := [3]string{"alice", "bob", ""}
	escrowsResults := [3]int{0, 0, 0}

	items := make([]types.Escrow, 20)
	for i := range items {
		r1 := rand.Intn(3)
		items[i].Initiator = addresses[r1]
		r2 := rand.Intn(3)
		items[i].Fulfiller = addresses[r2]
		items[i].Id = keeper.AppendEscrow(ctx, items[i])

		escrowsResults[r1] += 1
		if r1 != r2 {
			escrowsResults[r2] += 1
		}
	}

	tests := []struct {
		desc     string
		request  *types.QueryEscrowsByAddressRequest
		response int
		err      error
	}{
		{
			desc:     "Alice",
			request:  &types.QueryEscrowsByAddressRequest{Address: addresses[0]},
			response: escrowsResults[0],
		},
		{
			desc:     "Bob",
			request:  &types.QueryEscrowsByAddressRequest{Address: addresses[1]},
			response: escrowsResults[1],
		},
		{
			desc:     "InvalidAddress",
			request:  &types.QueryEscrowsByAddressRequest{Address: "invalid"},
			response: 0,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.EscrowsByAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				count := len(response.Escrow)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(count),
				)
			}
		})
	}
}

func TestEscrowQueryByAddressPaginated(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	// msgs := createNEscrow(keeper, ctx, 5)

	addresses := [3]string{"alice", "bob", ""}
	escrowsResults := [3]int{0, 0, 0}

	items := make([]types.Escrow, 20)
	var items_test []types.Escrow
	for i := range items {
		r1 := rand.Intn(3)
		items[i].Initiator = addresses[r1]
		r2 := rand.Intn(3)
		items[i].Fulfiller = addresses[r2]
		items[i].Id = keeper.AppendEscrow(ctx, items[i])

		escrowsResults[r1] += 1
		if r1 != r2 {
			escrowsResults[r2] += 1
		}

		if r1 == 0 || r2 == 0 {
			items_test = append(items_test, items[i])
		}
	}

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryEscrowsByAddressRequest {
		return &types.QueryEscrowsByAddressRequest{
			Address: addresses[0],
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(items); i += step {
			resp, err := keeper.EscrowsByAddress(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			count := len(resp.Escrow)
			require.LessOrEqual(t, count, step)
			require.Subset(t,
				nullify.Fill(items_test),
				nullify.Fill(resp.Escrow),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(items); i += step {
			resp, err := keeper.EscrowsByAddress(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Escrow), step)
			require.Subset(t,
				nullify.Fill(items_test),
				nullify.Fill(resp.Escrow),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.EscrowsByAddress(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(items), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(items_test),
			nullify.Fill(resp.Escrow),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.EscrowAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
