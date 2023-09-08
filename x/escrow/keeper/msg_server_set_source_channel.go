package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetSourceChannel(goCtx context.Context, msg *types.MsgSetSourceChannel) (*types.MsgSetSourceChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetSrcChannel(ctx, msg.Channel)

	return &types.MsgSetSourceChannelResponse{}, nil
}
