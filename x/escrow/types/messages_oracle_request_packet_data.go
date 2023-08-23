package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendOracleRequestPacketData = "send_oracle_request_packet_data"

var _ sdk.Msg = &MsgSendOracleRequestPacketData{}

func NewMsgSendOracleRequestPacketData(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	clientId string,
	oracleScriptId uint64,
	calldata []byte,
	askCount uint64,
	minCount uint64,
	feeLimit sdk.Coins,
	prepareGas uint64,
	executeGas uint64,
) *MsgSendOracleRequestPacketData {
	return &MsgSendOracleRequestPacketData{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		ClientId:         clientId,
		OracleScriptId:   oracleScriptId,
		Calldata:         calldata,
		AskCount:         askCount,
		MinCount:         minCount,
		FeeLimit:         feeLimit,
		PrepareGas:       prepareGas,
		ExecuteGas:       executeGas,
	}
}

func (msg *MsgSendOracleRequestPacketData) Route() string {
	return RouterKey
}

func (msg *MsgSendOracleRequestPacketData) Type() string {
	return TypeMsgSendOracleRequestPacketData
}

func (msg *MsgSendOracleRequestPacketData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendOracleRequestPacketData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendOracleRequestPacketData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
