package keeper

import (
	"dredd-secure/x/escrow/types"
	"errors"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

// TransmitOracleRequestPacketDataPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitOracleRequestPacketDataPacket(
	ctx sdk.Context,
	packetData types.OracleRequestPacketDataPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	oracleRequestPacket bandtypes.OracleRequestPacketData,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvOracleRequestPacketDataPacket processes packet reception
func (k Keeper) OnRecvOracleRequestPacketDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.OracleResponsePacketDataPacketData) (packetAck types.OracleRequestPacketDataPacketAck, err error) {
	// validate packet data upon receiving
	/*if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}*/

	// TODO: packet reception logic, can be delete as we will never receive these types of packet..???

	return packetAck, nil
}

// OnAcknowledgementOracleRequestPacketDataPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementOracleRequestPacketDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.OracleRequestPacketDataPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.OracleRequestPacketDataPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutOracleRequestPacketDataPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutOracleRequestPacketDataPacket(ctx sdk.Context, packet channeltypes.Packet, data types.OracleRequestPacketDataPacketData) error {
	// TODO: packet timeout logic

	return nil
}
