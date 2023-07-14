package keeper

import (
	"dredd-secure/x/escrow/constants"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
	"bytes"
	"sort"

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

// Validate the escrow conditions
func (k Keeper) ValidateConditions(ctx sdk.Context, escrow types.Escrow) bool {
	// Validate time conditions
	now := time.Now()
	unixTimeNow := now.Unix()

	endDateInt, errParseIntEndDate := strconv.ParseInt(escrow.EndDate, 10, 64)
	startDateInt, errParseIntStartDate := strconv.ParseInt(escrow.StartDate, 10, 64)
	if (errParseIntEndDate != nil) {
		panic(errParseIntEndDate.Error())
	}
	if (errParseIntStartDate != nil) {
		panic(errParseIntStartDate.Error())
	}
	
	// If the current date is before start date or after end date, time conditions are not met
	if (unixTimeNow < startDateInt || unixTimeNow > endDateInt) {
		return false
	}

	return true
}

// Release assets to the respective parties. The Initiator receives the FulfillerCoins, vice-versa
func (k Keeper) ReleaseAssets(ctx sdk.Context, escrow types.Escrow) {
	// Release initiator assets
	initiator, err := sdk.AccAddressFromBech32(escrow.Initiator)
    if err != nil {
        panic(err)
    }
	errSendCoinsInitiator := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, initiator, escrow.FulfillerCoins)
	if errSendCoinsInitiator != nil {
		panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), errSendCoinsInitiator.Error()))
	}

	// Release fulfiller assets
	fulfiller, err := sdk.AccAddressFromBech32(escrow.Fulfiller)
	if err != nil {
		panic(err)
	}
	errSendCoinsFulfiller := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.InitiatorCoins)
	if errSendCoinsFulfiller != nil {
		panic(fmt.Sprintf(types.ErrCannotReleaseFulfillerAssets.Error(), errSendCoinsFulfiller.Error()))
	}
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

// GetAllPendingEscrows returns all pending escrows ID
func (k Keeper) GetAllPendingEscrows(ctx sdk.Context) (list []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PendingEscrowKey)
	bz := store.Get(byteKey)

	if bz == nil {
		return
	}

	for i := 0; i <= len(bz); i += 8 {
		barr := bz[i:]
		if (len(barr) >= 8) {
			var val uint64 = binary.BigEndian.Uint64(barr)
			list = append(list, val)
		}
	}

	return
}

// Fulfills escrows ordered in start date as ascending, removes fulfilled escrows from the array
func (k Keeper) FulfillPendingEscrows(ctx sdk.Context) {
	var pendingEscrows []uint64 = k.GetAllPendingEscrows(ctx)
	var i int = -1
	for index, v := range pendingEscrows {
		escrow, found := k.GetEscrow(ctx, v)
		if (found && k.ValidateConditions(ctx, escrow)) {
			k.ReleaseAssets(ctx, escrow)
			escrow.Status = constants.StatusClosed
			k.SetEscrow(ctx, escrow)
			i = index
		} else if (found && !k.ValidateConditions(ctx, escrow)) {
			break
		}
	}

	if (len(pendingEscrows) > i + 1) {
		pendingEscrows = pendingEscrows[i + 1:]
	} else {
		pendingEscrows = []uint64{}
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	
	var buf *bytes.Buffer = new(bytes.Buffer)
	byteKey := types.KeyPrefix(types.PendingEscrowKey)
	binary.Write(buf, binary.BigEndian, pendingEscrows)
	if (buf.Bytes() == nil) {
		store.Set(byteKey, []byte{})
	} else {
		store.Set(byteKey, buf.Bytes())
	}
}

// Add escrow id to pending escrows id array in order
func (k Keeper) AddPendingEscrow(ctx sdk.Context, escrow types.Escrow) {
    pendingEscrows := k.GetAllPendingEscrows(ctx)

	// Either add in order or add to the list if its the first element
	if (len(pendingEscrows) > 0) {
		index := sort.Search(len(pendingEscrows), func(i int) bool {
			return escrow.GetId() == pendingEscrows[i]
		})
	
		// Escrow already in the list
		if (index < len(pendingEscrows)) {
			return
		}

		i := sort.Search(len(pendingEscrows), func(i int) bool { 
			escr, found := k.GetEscrow(ctx, pendingEscrows[i])
			if (found) {
				return escr.GetStartDate() >= escrow.GetStartDate() 
			}
			return false
		})
		pendingEscrows = append(pendingEscrows, escrow.GetId())
		copy(pendingEscrows[i+1:], pendingEscrows[i:])
		pendingEscrows[i] = escrow.GetId()
	} else {
		pendingEscrows = append(pendingEscrows, escrow.GetId())
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	
	var buf *bytes.Buffer = new(bytes.Buffer)
	byteKey := types.KeyPrefix(types.PendingEscrowKey)
	binary.Write(buf, binary.BigEndian, pendingEscrows)
	store.Set(byteKey, buf.Bytes())
}