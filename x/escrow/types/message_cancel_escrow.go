package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelEscrow = "cancel_escrow"

var _ sdk.Msg = &MsgCancelEscrow{}

func NewMsgCancelEscrow(creator string, id uint64) *MsgCancelEscrow {
	return &MsgCancelEscrow{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelEscrow) Route() string {
	return RouterKey
}

func (msg *MsgCancelEscrow) Type() string {
	return TypeMsgCancelEscrow
}

func (msg *MsgCancelEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelEscrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
