package keeper_test

import (
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow/types"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "dredd-secure/testutil/keeper"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestOraclePriceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOraclePrice(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetOraclePriceRequest
		response *types.QueryGetOraclePriceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetOraclePriceRequest{
				Symbol: msgs[0].Symbol,
			},
			response: &types.QueryGetOraclePriceResponse{OraclePrice: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetOraclePriceRequest{
				Symbol: msgs[1].Symbol,
			},
			response: &types.QueryGetOraclePriceResponse{OraclePrice: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetOraclePriceRequest{
				Symbol: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.OraclePrice(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestOraclePriceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EscrowKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNOraclePrice(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllOraclePriceRequest {
		return &types.QueryAllOraclePriceRequest{
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
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OraclePriceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OraclePrice), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OraclePrice),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.OraclePriceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.OraclePrice), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.OraclePrice),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.OraclePriceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.OraclePrice),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.OraclePriceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
