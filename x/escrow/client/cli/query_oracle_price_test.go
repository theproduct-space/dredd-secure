package cli_test

import (
	"dredd-secure/testutil/network"
	"dredd-secure/testutil/nullify"
	"dredd-secure/x/escrow/client/cli"
	"dredd-secure/x/escrow/types"
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithOraclePriceObjects(t *testing.T, n int) (*network.Network, []types.OraclePrice) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{PortId: types.PortID}
	for i := 0; i < n; i++ {
		oraclePrice := types.OraclePrice{
			Symbol: strconv.Itoa(i),
		}
		nullify.Fill(&oraclePrice)
		state.OraclePriceList = append(state.OraclePriceList, oraclePrice)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.OraclePriceList
}

func TestShowOraclePrice(t *testing.T) {
	net, objs := networkWithOraclePriceObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc     string
		idSymbol string

		args []string
		err  error
		obj  types.OraclePrice
	}{
		{
			desc:     "found",
			idSymbol: objs[0].Symbol,

			args: common,
			obj:  objs[0],
		},
		{
			desc:     "not found",
			idSymbol: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idSymbol,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowOraclePrice(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetOraclePriceResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.OraclePrice)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.OraclePrice),
				)
			}
		})
	}
}

func TestListOraclePrice(t *testing.T) {
	net, objs := networkWithOraclePriceObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOraclePrice(), args)
			require.NoError(t, err)
			var resp types.QueryAllOraclePriceResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.OraclePrice), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.OraclePrice),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOraclePrice(), args)
			require.NoError(t, err)
			var resp types.QueryAllOraclePriceResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.OraclePrice), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.OraclePrice),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListOraclePrice(), args)
		require.NoError(t, err)
		var resp types.QueryAllOraclePriceResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.OraclePrice),
		)
	})
}
