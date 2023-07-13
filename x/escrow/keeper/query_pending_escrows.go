package keeper

import (
	"context"
	"encoding/binary"

	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PendingEscrows(goCtx context.Context, req *types.QueryPendingEscrowsRequest) (*types.QueryPendingEscrowsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ids []uint64

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingEscrowKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val uint64 = binary.BigEndian.Uint64(iterator.Value())
		ids = append(ids, val)
	}

	return &types.QueryPendingEscrowsResponse{Ids: ids}, nil
}
