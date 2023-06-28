package keeper

import (
	"encoding/binary"

	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetEscrowCount get the total number of escrow
func (k Keeper) GetEscrowCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EscrowCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetEscrowCount set the total number of escrow
func (k Keeper) SetEscrowCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.EscrowCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendEscrow appends a escrow in the store with a new id and update the count
func (k Keeper) AppendEscrow(
	ctx sdk.Context,
	escrow types.Escrow,
) uint64 {
	// Create the escrow
	count := k.GetEscrowCount(ctx)

	// Set the ID of the appended value
	escrow.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKey))
	appendedValue := k.cdc.MustMarshal(&escrow)
	store.Set(GetEscrowIDBytes(escrow.Id), appendedValue)

	// Update escrow count
	k.SetEscrowCount(ctx, count+1)

	return count
}

// SetEscrow set a specific escrow in the store
func (k Keeper) SetEscrow(ctx sdk.Context, escrow types.Escrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKey))
	b := k.cdc.MustMarshal(&escrow)
	store.Set(GetEscrowIDBytes(escrow.Id), b)
}

// GetEscrow returns a escrow from its id
func (k Keeper) GetEscrow(ctx sdk.Context, id uint64) (val types.Escrow, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKey))
	b := store.Get(GetEscrowIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEscrow removes a escrow from the store
func (k Keeper) RemoveEscrow(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKey))
	store.Delete(GetEscrowIDBytes(id))
}

// GetAllEscrow returns all escrow
func (k Keeper) GetAllEscrow(ctx sdk.Context) (list []types.Escrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EscrowKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Escrow
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetEscrowIDBytes returns the byte representation of the ID
func GetEscrowIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetEscrowIDFromBytes returns ID in uint64 format from a byte array
func GetEscrowIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
