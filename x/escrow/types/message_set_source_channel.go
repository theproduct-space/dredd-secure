package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetSourceChannel = "set_source_channel"

var _ sdk.Msg = &MsgSetSourceChannel{}

func NewMsgSetSourceChannel(creator string, channel string) *MsgSetSourceChannel {
	return &MsgSetSourceChannel{
		Creator:        creator,
		Channel: 		channel,
	}
}

func (msg *MsgSetSourceChannel) Route() string {
	return RouterKey
}

func (msg *MsgSetSourceChannel) Type() string {
	return TypeMsgSetSourceChannel
}

func (msg *MsgSetSourceChannel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetSourceChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetSourceChannel) ValidateBasic() error {
	// Account validation
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
