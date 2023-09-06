package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSourceChannel(goCtx context.Context, req *types.QuerySourceChannelRequest) (*types.QuerySourceChannelResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)

	return &types.QuerySourceChannelResponse{Channel: params.GetSourceChannel()}, nil
}
