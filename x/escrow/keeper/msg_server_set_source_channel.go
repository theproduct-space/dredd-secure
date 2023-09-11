package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetSourceChannel(goCtx context.Context, msg *types.MsgSetSourceChannel) (*types.MsgSetSourceChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the set source request is from the gov module
	if k.govKeeper.GetAuthority() != msg.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "Cannot set the channel id: not from the gov module")
	}

	k.SetSrcChannel(ctx, msg.Channel)

	return &types.MsgSetSourceChannelResponse{}, nil
}
