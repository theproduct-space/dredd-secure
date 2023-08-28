package keeper

import (
	"context"

	"strconv"

	"github.com/google/uuid"

	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

func (k msgServer) SendOracleRequestPacketData(goCtx context.Context, msg *types.MsgSendOracleRequestPacketData) (*types.MsgSendOracleRequestPacketDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Generate a new UUID
	uid := uuid.New()

	// Construct the packet
	var packet types.OracleRequestPacketDataPacketData

	// using the oracleScriptId in the clientId for data treater upon OracleResponsePacketData reception
	packet.ClientId = strconv.FormatUint(msg.OracleScriptId, 10) + "_" + uid.String()
	packet.OracleScriptId = msg.OracleScriptId
	packet.Calldata = msg.Calldata
	packet.AskCount = msg.AskCount
	packet.MinCount = msg.MinCount
	packet.FeeLimit = msg.FeeLimit
	packet.PrepareGas = msg.PrepareGas
	packet.ExecuteGas = msg.ExecuteGas

	// Transmit the packet
	_, err := k.TransmitOracleRequestPacketDataPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendOracleRequestPacketDataResponse{}, nil
}
