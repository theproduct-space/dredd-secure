package keeper

import (
	"context"

	"dredd-secure/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EscrowsByAddress(goCtx context.Context, req *types.QueryEscrowsByAddressRequest) (*types.QueryEscrowsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryEscrowsByAddressResponse{}, nil
}
