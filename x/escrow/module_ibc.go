package escrow

import (
	"fmt"
	"strings"

	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v7/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
)

type IBCModule struct {
	keeper keeper.Keeper
}

func NewIBCModule(k keeper.Keeper) IBCModule {
	return IBCModule{
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCModule) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {

	// Require portID is the portID module is bound to
	boundPort := im.keeper.GetPort(ctx)
	if boundPort != portID {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if strings.TrimSpace(version) == "" {
		version = types.Version
	}

	if version != types.Version {
		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "got %s, expected %s", version, types.Version)
	}

	// Claim channel capability passed back by IBC module
	if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
		return "", err
	}

	return version, nil
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCModule) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {

	// Require portID is the portID module is bound to
	boundPort := im.keeper.GetPort(ctx)
	if boundPort != portID {
		return "", sdkerrors.Wrapf(porttypes.ErrInvalidPort, "invalid port: %s, expected %s", portID, boundPort)
	}

	if counterpartyVersion != types.Version {
		return "", sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: got: %s, expected %s", counterpartyVersion, types.Version)
	}

	// Module may have already claimed capability in OnChanOpenInit in the case of crossing hellos
	// (ie chainA and chainB both call ChanOpenInit before one of them calls ChanOpenTry)
	// If module can already authenticate the capability then module already owns it so we don't need to claim
	// Otherwise, module does not have channel capability and we must claim it from IBC
	if !im.keeper.AuthenticateCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)) {
		// Only claim channel capability passed back by IBC module if we do not already own it
		if err := im.keeper.ClaimCapability(ctx, chanCap, host.ChannelCapabilityPath(portID, channelID)); err != nil {
			return "", err
		}
	}

	return types.Version, nil
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCModule) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	_,
	counterpartyVersion string,
) error {
	if counterpartyVersion != types.Version {
		return sdkerrors.Wrapf(types.ErrInvalidVersion, "invalid counterparty version: %s, expected %s", counterpartyVersion, types.Version)
	}
	return nil
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCModule) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCModule) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// Disallow user-initiated channel closing for channels
	return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user cannot close channel")
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCModule) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return nil
}

// OnRecvPacket implements the IBCModule interface
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// var ack channeltypes.Acknowledgement
	var packet bandtypes.OracleResponsePacketData

	// Unmarshal the data from the module packet into the OracleResponsePacketData object.
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &packet); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Request has been resolved 
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeOracleResponsePacketDataPacket,
		sdk.NewAttribute(types.AttributeKeyRequestID, fmt.Sprintf("%d", packet.RequestID)),
	))

	if packet.ResolveStatus != bandtypes.RESOLVE_STATUS_SUCCESS {
		return channeltypes.NewErrorAcknowledgement(types.ErrOracleResolveStatusNotSuccess)
	}

	if err := im.keeper.StoreOracleResponsePacket(ctx, packet); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// TODO, store the OracleResponsePacket
	// -> what is the data type used to store the response packet?
	// -> needs to be general in order to match multiple oracle scripts data types response

	// if err := im.keeper.StoreOracleResponsePacket(ctx, packet); err != nil {
	// 	return channeltypes.NewErrorAcknowledgement(err)
	// }
	// this line is used by starport scaffolding # oracle/packet/module/recv

	// BEGIN SCAFFOLDING FROM IGNITE THAT WE MIGHT NOT NEED
	// var modulePacketData types.EscrowPacketData
	// if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
	// 	return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error()))
	// }

	// // Dispatch packet
	// switch packet := modulePacketData.Packet.(type) {
	// case *types.EscrowPacketData_OracleRequestPacketDataPacket:
	// 	packetAck, err := im.keeper.OnRecvOracleRequestPacketDataPacket(ctx, modulePacket, *packet.OracleRequestPacketDataPacket)
	// 	if err != nil {
	// 		ack = channeltypes.NewErrorAcknowledgement(err)
	// 	} else {
	// 		// Encode packet acknowledgment
	// 		packetAckBytes, err := types.ModuleCdc.MarshalJSON(&packetAck)
	// 		if err != nil {
	// 			return channeltypes.NewErrorAcknowledgement(sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error()))
	// 		}
	// 		ack = channeltypes.NewResultAcknowledgement(sdk.MustSortJSON(packetAckBytes))
	// 	}
	// 	ctx.EventManager().EmitEvent(
	// 		sdk.NewEvent(
	// 			types.EventTypeOracleRequestPacketDataPacket,
	// 			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	// 			sdk.NewAttribute(types.AttributeKeyAckSuccess, fmt.Sprintf("%t", err != nil)),
	// 		),
	// 	)
	// 	// this line is used by starport scaffolding # ibc/packet/module/recv
	// default:
	// 	err := fmt.Errorf("unrecognized %s packet type: %T", types.ModuleName, packet)
	// 	return channeltypes.NewErrorAcknowledgement(err)
	// }

	// END SCAFFOLDING FROM IGNITE THAT WE MIGHT NOT NEED

	// NOTE: acknowledgement will be written synchronously during IBC handler execution.
	return channeltypes.NewResultAcknowledgement(nil)
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet acknowledgement: %v", err)
	}

	// this line is used by starport scaffolding # oracle/packet/module/ack

	var modulePacketData types.EscrowPacketData
	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	var eventType string

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *types.EscrowPacketData_OracleRequestPacketDataPacket:
		err := im.keeper.OnAcknowledgementOracleRequestPacketDataPacket(ctx, modulePacket, *packet.OracleRequestPacketDataPacket, ack)
		if err != nil {
			return err
		}
		eventType = types.EventTypeOracleRequestPacketDataPacket
		// this line is used by starport scaffolding # ibc/packet/module/ack
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			eventType,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyAck, fmt.Sprintf("%v", ack)),
		),
	)

	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckSuccess, string(resp.Result)),
			),
		)
	case *channeltypes.Acknowledgement_Error:
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				eventType,
				sdk.NewAttribute(types.AttributeKeyAckError, resp.Error),
			),
		)
	}

	return nil
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCModule) OnTimeoutPacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	var modulePacketData types.EscrowPacketData
	if err := modulePacketData.Unmarshal(modulePacket.GetData()); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	// Dispatch packet
	switch packet := modulePacketData.Packet.(type) {
	case *types.EscrowPacketData_OracleRequestPacketDataPacket:
		err := im.keeper.OnTimeoutOracleRequestPacketDataPacket(ctx, modulePacket, *packet.OracleRequestPacketDataPacket)
		if err != nil {
			return err
		}
		// this line is used by starport scaffolding # ibc/packet/module/timeout
	default:
		errMsg := fmt.Sprintf("unrecognized %s packet type: %T", types.ModuleName, packet)
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}

	return nil
}
