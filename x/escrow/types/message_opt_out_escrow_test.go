package types_test

import (
	"dredd-secure/testutil/sample"
	"dredd-secure/x/escrow/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stretchr/testify/require"
)

func TestMsgOptOutEscrow_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgOptOutEscrow
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgOptOutEscrow{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgOptOutEscrow{
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
