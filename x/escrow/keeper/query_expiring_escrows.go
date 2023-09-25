package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExpiringEscrows(goCtx context.Context, req *types.QueryExpiringEscrowsRequest) (*types.QueryExpiringEscrowsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ids := k.GetAllExpiringEscrows(ctx)

	return &types.QueryExpiringEscrowsResponse{Ids: ids}, nil
}
