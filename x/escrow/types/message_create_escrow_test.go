package types

import (
	"dredd-secure/testutil/sample"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateEscrow_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateEscrow
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateEscrow{
				Creator: "invalid_address",
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid request, missing IniatiatorCoins",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid request, missing FulfillerCoins",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "invalid InitiatorCoins denom",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "*this_is_an_invalid_denom*",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidCoins,
		},
		{
			name: "invalid FulfillerCoins denom",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "*this_is_an_invalid_denom*",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidCoins,
		},
		{
			name: "invalid InitiatorCoins amount",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(-1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidCoins,
		},
		{
			name: "invalid FulfillerCoins amount",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(-9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "2188148578",
			},
			err: sdkerrors.ErrInvalidCoins,
		},
		{
			name: "End date is not in the future",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1188148578",
				EndDate:   "1688148884",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "End date is before start date",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1888148578",
				EndDate:   "1788148978",
			},
			err: sdkerrors.ErrInvalidRequest,
		},
		{
			name: "Everything is Valid",
			msg: MsgCreateEscrow{
				Creator: sample.AccAddress(),
				InitiatorCoins: []sdk.Coin{{
					Denom:  "token",
					Amount: sdk.NewInt(1000),
				}},
				FulfillerCoins: []sdk.Coin{{
					Denom:  "stake",
					Amount: sdk.NewInt(9000),
				}},
				Tips:      nil,
				StartDate: "1588148578",
				EndDate:   "2788148978",
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
