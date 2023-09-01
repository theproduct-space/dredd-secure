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

func (k Keeper) OraclePriceAll(goCtx context.Context, req *types.QueryAllOraclePriceRequest) (*types.QueryAllOraclePriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var oraclePrices []types.OraclePrice
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	oraclePriceStore := prefix.NewStore(store, types.KeyPrefix(types.OraclePriceKeyPrefix))

	pageRes, err := query.Paginate(oraclePriceStore, req.Pagination, func(key []byte, value []byte) error {
		var oraclePrice types.OraclePrice
		if err := k.cdc.Unmarshal(value, &oraclePrice); err != nil {
			return err
		}

		oraclePrices = append(oraclePrices, oraclePrice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOraclePriceResponse{OraclePrice: oraclePrices, Pagination: pageRes}, nil
}

func (k Keeper) OraclePrice(goCtx context.Context, req *types.QueryGetOraclePriceRequest) (*types.QueryGetOraclePriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetOraclePrice(
		ctx,
		req.Symbol,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetOraclePriceResponse{OraclePrice: val}, nil
}
