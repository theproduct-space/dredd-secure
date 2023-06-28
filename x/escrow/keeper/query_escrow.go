package keeper

import (
	"context"

	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EscrowAll(goCtx context.Context, req *types.QueryAllEscrowRequest) (*types.QueryAllEscrowResponse, error) {
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

		escrows = append(escrows, escrow)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEscrowResponse{Escrow: escrows, Pagination: pageRes}, nil
}

func (k Keeper) Escrow(goCtx context.Context, req *types.QueryGetEscrowRequest) (*types.QueryGetEscrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	escrow, found := k.GetEscrow(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetEscrowResponse{Escrow: escrow}, nil
}
