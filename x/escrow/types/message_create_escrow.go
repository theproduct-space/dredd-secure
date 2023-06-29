package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateEscrow = "create_escrow"

var _ sdk.Msg = &MsgCreateEscrow{}

func NewMsgCreateEscrow(creator string, initiatorCoins sdk.Coins, fulfillerCoins sdk.Coins, startDate string, endDate string) *MsgCreateEscrow {
	return &MsgCreateEscrow{
		Creator:        creator,
		InitiatorCoins: initiatorCoins,
		FulfillerCoins: fulfillerCoins,
		StartDate:      startDate,
		EndDate:        endDate,
	}
}

func (msg *MsgCreateEscrow) Route() string {
	return RouterKey
}

func (msg *MsgCreateEscrow) Type() string {
	return TypeMsgCreateEscrow
}

func (msg *MsgCreateEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateEscrow) ValidateBasic() error {
	// PSTODO: validate the start_date and end_date
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
   
	return nil
}
