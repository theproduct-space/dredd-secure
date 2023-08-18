package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgOptOutEscrow = "opt_out_escrow"

var _ sdk.Msg = &MsgOptOutEscrow{}

func NewMsgOptOutEscrow(creator string, id uint64) *MsgOptOutEscrow {
	return &MsgOptOutEscrow{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgOptOutEscrow) Route() string {
	return RouterKey
}

func (msg *MsgOptOutEscrow) Type() string {
	return TypeMsgOptOutEscrow
}

func (msg *MsgOptOutEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOptOutEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOptOutEscrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
