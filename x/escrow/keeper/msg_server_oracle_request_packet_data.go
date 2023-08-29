package keeper

import (
	"fmt"
	"context"

	"strconv"

	"github.com/google/uuid"

	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
)

func (k msgServer) SendOracleRequestPacketData(goCtx context.Context, msg *types.MsgSendOracleRequestPacketData) (*types.MsgSendOracleRequestPacketDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Construct the packet
	var packet types.OracleRequestPacketDataPacketData

	// Generate a new UUID
	uid := uuid.New()

	// using the oracleScriptId in the clientId for data treater upon OracleResponsePacketData reception
	packet.ClientID = strconv.FormatUint(msg.OracleScriptID, 10) + "_" + uid.String()
	packet.OracleScriptID = msg.OracleScriptID
	packet.Calldata = msg.Calldata
	packet.AskCount = msg.AskCount
	packet.MinCount = msg.MinCount
	packet.FeeLimit = msg.FeeLimit
	packet.PrepareGas = msg.PrepareGas
	packet.ExecuteGas = msg.ExecuteGas

	fmt.Println("ClientID : ", packet.ClientID)
	fmt.Println("OracleScriptID : ", packet.OracleScriptID)
	fmt.Println("Calldata : ", packet.Calldata)
	fmt.Println("AskCount : ", packet.AskCount)
	fmt.Println("MinCount : ", packet.MinCount)
	fmt.Println("FeeLimit : ", packet.FeeLimit)
	fmt.Println("PrepareGas : ", packet.PrepareGas)
	fmt.Println("ExecuteGas : ", packet.ExecuteGas)

	oracleRequestPacket := bandtypes.NewOracleRequestPacketData(
		types.ModuleName,
		msg.OracleScriptID,
		msg.Calldata,
		msg.AskCount,
		msg.MinCount,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	err := k.RequestBandChainData(ctx, msg.ChannelID, oracleRequestPacket, clienttypes.ZeroHeight(), msg.TimeoutTimestamp)

	// Transmit the packet
	/*_, err := k.TransmitOracleRequestPacketDataPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)*/
	if err != nil {
		return nil, err
	}

	return &types.MsgSendOracleRequestPacketDataResponse{}, nil
}
