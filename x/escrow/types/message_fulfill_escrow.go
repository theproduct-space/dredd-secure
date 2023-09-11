package types

import (
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFulfillEscrow = "fulfill_escrow"

var _ sdk.Msg = &MsgFulfillEscrow{}

func NewMsgFulfillEscrow(creator string, id uint64, denomMap map[string]string) *MsgFulfillEscrow {
	return &MsgFulfillEscrow{
		Creator:  creator,
		Id:       id,
		DenomMap: denomMap,
	}
}

func (msg *MsgFulfillEscrow) Route() string {
	return RouterKey
}

func (msg *MsgFulfillEscrow) Type() string {
	return TypeMsgFulfillEscrow
}

func (msg *MsgFulfillEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFulfillEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFulfillEscrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
