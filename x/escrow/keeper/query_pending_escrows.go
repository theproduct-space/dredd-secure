package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingEscrows(goCtx context.Context, req *types.QueryPendingEscrowsRequest) (*types.QueryPendingEscrowsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var ids []uint64 = k.GetAllPendingEscrows(ctx)

	return &types.QueryPendingEscrowsResponse{Ids: ids}, nil
}
