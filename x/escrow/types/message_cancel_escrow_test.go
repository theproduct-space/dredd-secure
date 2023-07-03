package types_test

import (
	"testing"

	"dredd-secure/testutil/sample"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"

	keepertest "dredd-secure/testutil/keeper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"
)

func TestMsgCancelEscrow_ValidateBasic(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EscrowList: []types.Escrow{
			{
				Initiator: "1",
				Id:        0,
			},
			{
				Initiator: "2",
				Id:        1,
			},
		},
		EscrowCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EscrowKeeper(t)
	escrow.InitGenesis(ctx, *k, genesisState)
	msgServer := keeper.NewMsgServerImpl(*k)
	_, err := msgServer.CancelEscrow(ctx, &types.MsgCancelEscrow{
		Creator: "1",
		Id:      0,
	})
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgCancelEscrow
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCancelEscrow{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgCancelEscrow{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
