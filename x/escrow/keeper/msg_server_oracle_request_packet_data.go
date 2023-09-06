package keeper

import (
	"context"

	"strconv"

	"github.com/google/uuid"

	"dredd-secure/x/escrow/types"

	bandtypes "github.com/bandprotocol/oracle-consumer/types/band"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

func (k msgServer) SendOracleRequestPacketData(goCtx context.Context, msg *types.MsgSendOracleRequestPacketData) (*types.MsgSendOracleRequestPacketDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Generate a new UUID
	uid := uuid.New()

	oracleRequestPacket := bandtypes.NewOracleRequestPacketData(
		strconv.FormatUint(msg.OracleScriptID, 10) + "_" + uid.String(),
		msg.OracleScriptID,
		msg.Calldata,
		msg.AskCount,
		msg.MinCount,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	err := k.RequestBandChainData(ctx, msg.ChannelID, oracleRequestPacket, clienttypes.ZeroHeight(), msg.TimeoutTimestamp)

	if err != nil {
		return nil, err
	}

	return &types.MsgSendOracleRequestPacketDataResponse{}, nil
}
