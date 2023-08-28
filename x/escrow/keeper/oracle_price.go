package keeper

import (
	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetOraclePrice set a specific oraclePrice in the store from its index
func (k Keeper) SetOraclePrice(ctx sdk.Context, oraclePrice types.OraclePrice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OraclePriceKeyPrefix))
	b := k.cdc.MustMarshal(&oraclePrice)
	store.Set(types.OraclePriceKey(
		oraclePrice.Symbol,
	), b)
}

// GetOraclePrice returns a oraclePrice from its index
func (k Keeper) GetOraclePrice(
	ctx sdk.Context,
	symbol string,

) (val types.OraclePrice, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OraclePriceKeyPrefix))

	b := store.Get(types.OraclePriceKey(
		symbol,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOraclePrice removes a oraclePrice from the store
func (k Keeper) RemoveOraclePrice(
	ctx sdk.Context,
	symbol string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OraclePriceKeyPrefix))
	store.Delete(types.OraclePriceKey(
		symbol,
	))
}

// GetAllOraclePrice returns all oraclePrice
func (k Keeper) GetAllOraclePrice(ctx sdk.Context) (list []types.OraclePrice) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OraclePriceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OraclePrice
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
