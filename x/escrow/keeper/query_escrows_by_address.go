package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EscrowsByAddress(goCtx context.Context, req *types.QueryEscrowsByAddressRequest) (*types.QueryEscrowsByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var escrows []types.Escrow

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	escrowStore := prefix.NewStore(store, types.KeyPrefix(types.EscrowKey))

	pageRes, err := query.Paginate(escrowStore, req.Pagination, func(key []byte, value []byte) error {
		var escrow types.Escrow
		if err := k.cdc.Unmarshal(value, &escrow); err != nil {
			return err
		}

		if escrow.Fulfiller == req.Address || escrow.Initiator == req.Address {
			escrows = append(escrows, escrow)
		}

		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEscrowsByAddressResponse{Escrow: escrows, Pagination: pageRes}, nil
}
