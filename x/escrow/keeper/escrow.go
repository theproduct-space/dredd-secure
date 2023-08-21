package keeper

import (
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"io/ioutil"
	"net/http"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tidwall/gjson"
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

	k.AddExpiringEscrow(ctx, escrow)

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

// ValidateConditions validates the escrow conditions
func (k Keeper) ValidateConditions(ctx sdk.Context, escrow types.Escrow) bool {
	// Validate the StartDate, EndDate and ApiConditions
	if !k.ValidateStartDate(ctx, escrow) || !k.ValidateEndDate(ctx, escrow) || !k.ValidateApiConditions(ctx, escrow) {
		return false
	}

	return true
}

// ValidateStartDate validates that the startDate is not in the future
func (k Keeper) ValidateStartDate(ctx sdk.Context, escrow types.Escrow) bool {
	now := time.Now()
	unixTimeNow := now.Unix()

	startDateInt, errParseIntStartDate := strconv.ParseInt(escrow.StartDate, 10, 64)

	if (errParseIntStartDate != nil) {
		panic(errParseIntStartDate.Error())
	}

	if (unixTimeNow < startDateInt) {
		return false
	}
	return true
}

// ValidateEndDate validates that the endDate is not in the past
func (k Keeper) ValidateEndDate(ctx sdk.Context, escrow types.Escrow) bool {
	now := time.Now()
	unixTimeNow := now.Unix()

	endDateInt, errParseIntEndDate := strconv.ParseInt(escrow.EndDate, 10, 64)

	if errParseIntEndDate != nil {
		panic(errParseIntEndDate.Error())
	}

	if unixTimeNow > endDateInt {
		return false
	}
	return true
}

// ValidateApiConditions validates the ApiConditions by making the api calls and comparing the relevant fields with their expected values
func (k Keeper) ValidateApiConditions(ctx sdk.Context, escrow types.Escrow) bool {
	apiConditionsString := escrow.ApiConditions;

	var apiConditions []types.ApiCondition
	err := json.Unmarshal([]byte(apiConditionsString), &apiConditions)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	for _, condition := range apiConditions {
		apiRes, err := makeAPIRequest(condition.Name, strconv.Itoa(condition.TokenOfInterest.ID))

		if (err != nil) {
			return false
		}

		for _, subCondition := range condition.SubConditions {
			result := ValidateSubCondition(subCondition, apiRes);
			// If the result is false, return false immediately; otherwise, continue validating
			if !result {
				return false
			}
		}
	}

	return true
}

// Validates a SubCondition by comparing the subCondition value with the one fetched from the API
func ValidateSubCondition(subCondition types.SubCondition, apiRes string) (bool) {
	// access the data of interest (navigate apiRes using subcondition.path) in the appropriate type (using subCondition.dataType)
	pathString := strings.Join(subCondition.Path, ".")
	var dataToCompare interface{};
	if (subCondition.DataType == "number") {
		dataToCompare = gjson.Get(apiRes, pathString).Float()
	} else if (subCondition.DataType == "text") {
		dataToCompare = gjson.Get(apiRes, pathString).String()
	} else {
		fmt.Println("Invalid data type")
		return false
	}

	// Depending on the dataToCompare type,
	switch value := dataToCompare.(type) {
		case float64:
			// retrieve the expected value as a float64
			expectedValue := subCondition.Value.(float64);
			// make the appropriate comparison as described in subCondition.ConditionType
			switch (subCondition.ConditionType) {
				case "eq":
					return value == expectedValue
				case "lt":
					return value < expectedValue
				case "gt":
					return value > expectedValue
				default: 
					fmt.Println("Unknown data type")
					return false
			}
		case string:
			// retrieve the expected value as a string
			expectedValue := subCondition.Value.(string);
			// make a strict string comparison
			return  value == expectedValue
		default:
			fmt.Println("Unknown data type")
			return false
	}
}

// Make an API Request to the configured endpoints. Uses the "name" identifier to know which endpoint to use.
func makeAPIRequest(name string, tokenId string) (string, error) {
	var url string;
	var headers []types.Header;

	switch name {
		case "coinmarketcap-token-info": 
			url = "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?id=" + tokenId; // TODO have a more flexible "additionalParams" instead of tokenId which is only for token api calls
			headers = append(headers, types.Header{Key: "X-CMC_PRO_API_KEY", Value: "8857ddb8-2b98-4334-8278-eadc75e5dda3"}) // TODO do not hardcode the API KEY...
		default:
			return "", errors.Wrapf(types.ErrInvalidApiConditionName, "%v", name)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add headers to the request
	for _, header := range headers {
		req.Header.Add(header.Key, header.Value)
	}

	// Perform the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

    return string(body), nil
}

// ReleaseAssets releases the escrowed assets to the respective parties. The Initiator receives the FulfillerCoins, vice-versa
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

// RefundAssets releases the escrowed assets to the original parties. The Initiator receives the InitiatorCoins and fulfiller the FulfillerCoins
func (k Keeper) RefundAssets(ctx sdk.Context, escrow types.Escrow) {
	if escrow.Initiator != "" {
		// Release initiator assets
		initiator, err := sdk.AccAddressFromBech32(escrow.Initiator)
		if err != nil {
			panic(err)
		}
		errSendCoinsInitiator := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, initiator, escrow.InitiatorCoins)
		if errSendCoinsInitiator != nil {
			panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), errSendCoinsInitiator.Error()))
		}
	}

	if escrow.Fulfiller != "" {
		// Release fulfiller assets
		fulfiller, err := sdk.AccAddressFromBech32(escrow.Fulfiller)
		if err != nil {
			panic(err)
		}
		errSendCoinsFulfiller := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.FulfillerCoins)
		if errSendCoinsFulfiller != nil {
			panic(fmt.Sprintf(types.ErrCannotReleaseFulfillerAssets.Error(), errSendCoinsFulfiller.Error()))
		}
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
	if (len(bz) == 0) {
		return
	} 
	err := json.Unmarshal(bz, &list)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return 
}

// GetAllExpiringEscrows returns all expiring escrows ID
func (k Keeper) GetAllExpiringEscrows(ctx sdk.Context) (list []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ExpiringEscrowKey)
	bz := store.Get(byteKey)
	if (len(bz) == 0) {
		return
	} 
	err := json.Unmarshal(bz, &list)
	if err != nil {
		fmt.Println("Error:", err)
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
			k.RemoveFromExpiringList(ctx, escrow)
			k.SetEscrow(ctx, escrow)
			i = index
		} else if (found && !k.ValidateStartDate(ctx, escrow)) {
			break
		}
	}

	if (len(pendingEscrows) > i + 1) {
		pendingEscrows = pendingEscrows[i + 1:]
	} else {
		pendingEscrows = []uint64{}
	}
	
	k.SetPendingEscrows(ctx, pendingEscrows)
}

func (k Keeper) RemoveFromPendingList(ctx sdk.Context, escrow types.Escrow) {
	pendingEscrows := k.GetAllPendingEscrows(ctx) 
	index := sort.Search(len(pendingEscrows), func(i int) bool {
		return escrow.GetId() == pendingEscrows[i]
	})

	// Escrow found in the list
	if index < len(pendingEscrows) {
		if len(pendingEscrows) == 1 {
			pendingEscrows = []uint64{}
		} else {
			pendingEscrows = append(pendingEscrows[:index], pendingEscrows[index+1:]...)
		}
	}

	k.SetPendingEscrows(ctx, pendingEscrows)
}

func (k Keeper) RemoveFromExpiringList(ctx sdk.Context, escrow types.Escrow) {
	expiringEscrows := k.GetAllExpiringEscrows(ctx) 
	index := sort.Search(len(expiringEscrows), func(i int) bool {
		return escrow.GetId() == expiringEscrows[i]
	})

	// Escrow found in the list
	if index < len(expiringEscrows) {
		if len(expiringEscrows) == 1 {
			expiringEscrows = []uint64{}
		} else {
			expiringEscrows = append(expiringEscrows[:index], expiringEscrows[index+1:]...)
		}
	}

	k.SetExpiringEscrows(ctx, expiringEscrows)
}

func (k Keeper) SetExpiringEscrows(ctx sdk.Context, expiringEscrows []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ExpiringEscrowKey)
	
	jsonData, err := json.Marshal(expiringEscrows)
	if err != nil {
		fmt.Println("Error:", err)
	}

	store.Set(byteKey, jsonData)
}

func (k Keeper) SetPendingEscrows(ctx sdk.Context, pendingEscrows []uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PendingEscrowKey)

	jsonData, err := json.Marshal(pendingEscrows)
	if err != nil {
		fmt.Println("Error:", err)
	}

	store.Set(byteKey, jsonData)
}

// Cancels escrows ordered in end date as ascending, removes cancelled escrows from the array
func (k Keeper) CancelExpiredEscrows(ctx sdk.Context) {
	expiringEscrows := k.GetAllExpiringEscrows(ctx)
	i := -1
	for index, v := range expiringEscrows {
		escrow, found := k.GetEscrow(ctx, v)
		if found && !k.ValidateEndDate(ctx, escrow) {
			k.RefundAssets(ctx, escrow)
			escrow.Status = constants.StatusCancelled
			k.RemoveFromPendingList(ctx, escrow)
			k.SetEscrow(ctx, escrow)
			i = index
		} else if found && k.ValidateEndDate(ctx, escrow) {
			break
		}
	}

	if len(expiringEscrows) > i+1 {
		expiringEscrows = expiringEscrows[i+1:]
	} else {
		expiringEscrows = []uint64{}
	}
	
	k.SetExpiringEscrows(ctx, expiringEscrows)
}

// Add escrow id to pending escrows id array in order
func (k Keeper) AddPendingEscrow(ctx sdk.Context, escrow types.Escrow) {
    pendingEscrows := k.GetAllPendingEscrows(ctx)

	// Either add in order or add to the list if its the first element
	if len(pendingEscrows) > 0 {
		_, f := sort.Find(len(pendingEscrows), func(i int) int {
			if escrow.GetId() == pendingEscrows[i] {
				return 0
			} else if (escrow.GetId() > pendingEscrows[i]) {
				return 1
			} 
			return -1
		})
	
		// Escrow already in the list
		if f {
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

	k.SetPendingEscrows(ctx, pendingEscrows)
}

// Add escrow id to expiring escrows id array in order
func (k Keeper) AddExpiringEscrow(ctx sdk.Context, escrow types.Escrow) {
	expiringEscrows := k.GetAllExpiringEscrows(ctx)

	// Either add in order or add to the list if its the first element
	if len(expiringEscrows) > 0 {
		_, f := sort.Find(len(expiringEscrows), func(i int) int {
			if escrow.GetId() == expiringEscrows[i] {
				return 0
			} else if (escrow.GetId() > expiringEscrows[i]) {
				return 1
			} 
			return -1
		})

		// Escrow already in the list
		if f {
			return
		}

		i := sort.Search(len(expiringEscrows), func(i int) bool {
			escr, found := k.GetEscrow(ctx, expiringEscrows[i])
			if found {
				return escr.GetEndDate() >= escrow.GetEndDate()
			}
			return false
		})
		expiringEscrows = append(expiringEscrows, escrow.GetId())
		copy(expiringEscrows[i+1:], expiringEscrows[i:])
		expiringEscrows[i] = escrow.GetId()
	} else {
		expiringEscrows = append(expiringEscrows, escrow.GetId())
	}

	k.SetExpiringEscrows(ctx, expiringEscrows)
}

// Setter for the status
func (k Keeper) SetStatus(ctx sdk.Context, escrow *types.Escrow, newStatus string) {
	oldStatus := escrow.Status

	if (newStatus == constants.StatusOpen) {
		k.AddExpiringEscrow(ctx, *escrow)
	}

	if (newStatus == constants.StatusClosed || newStatus == constants.StatusCancelled) {
		k.RemoveFromExpiringList(ctx, *escrow)
	}

	if (oldStatus == constants.StatusPending && newStatus != constants.StatusPending) {
		k.RemoveFromPendingList(ctx, *escrow)
	}

	if (newStatus == constants.StatusPending) {
		k.AddPendingEscrow(ctx, *escrow)
	}

	escrow.Status = newStatus
}

// Utility function used for executing functions after a certain amount of time
func (k Keeper) ExecuteAfterNSeconds(
	ctx sdk.Context, 
	execs []Exec) []interface{} {
	results := make([]interface{}, 0)
	var lastExecs map[string] string
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LastExecsKey))
	byteKey := types.KeyPrefix(types.LastExecsKey)
	bz := store.Get(byteKey)
	if (len(bz) == 0) {
		lastExecs = make(map[string]string)
	} else {
		err := json.Unmarshal(bz, &lastExecs)
		if err != nil {
			fmt.Println("Error:", err)
			return results
		}
	}

	currentTime := time.Now()
	epoch := currentTime.Unix()
	
	for _, exec := range execs {
		epochString, found := lastExecs[exec.ID]
		if !found {
			epochString = "0"
			lastExecs[exec.ID] = epochString
		}

		epochInt, err := strconv.ParseInt(epochString, 10, 64)
		if err != nil {
			fmt.Println("Error converting epoch string to int:", err)
		} else {
			if (epochInt + exec.DelayS < epoch) {
				result := exec.Function(exec.Args...)
				results = append(results, result)
				lastExecs[exec.ID] = strconv.FormatInt(epoch, 10)
			}
		}
	}

	jsonData, err2 := json.Marshal(lastExecs)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return results
	}

	store.Set(byteKey, jsonData)

	return results
}