package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetSourceChannel(goCtx context.Context, msg *types.MsgSetSourceChannel) (*types.MsgSetSourceChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	params.SourceChannel = msg.Channel
	k.SetParams(ctx, params)

	return &types.MsgSetSourceChannelResponse{}, nil
}
