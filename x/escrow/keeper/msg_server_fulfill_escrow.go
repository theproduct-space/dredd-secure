package keeper

import (
	"context"
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"
	"fmt"
	"strings"

	"cosmossdk.io/errors"

	// tmbytes "github.com/cometbft/cometbft/libs/bytes"

	"github.com/cometbft/cometbft/libs/bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// FulfillEscrow fulfills an existing escrow of the given id
func (k msgServer) FulfillEscrow(goCtx context.Context, msg *types.MsgFulfillEscrow) (*types.MsgFulfillEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// req := &ibcTransferTypes.QueryDenomTraceRequest{
	// 	Hash: "EC2E067E9E24BD23A2425F115D35F5AC148F05F51EEBD7EE7814E04A59DE482D",
	// }

	// Retrieve the escrow from the keeper
	escrow, found := k.GetEscrow(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
	}


	// Create a new slice to hold the updated FulfillerCoins
	var updatedFulfillerCoins sdk.Coins
	for _, storedCoin := range escrow.FulfillerCoins {
		// For every fulfillercoins stored in the escrow contract, we want to verify that the denom passed as argument matches the denom stored on the contract
		// The native denom are trivial
		// But the IBC denom needs to be analyzed in order to retrieve their base_denom.
		// Then we can compare the found base_denom from the ibc hash to the base_denom that was stored in the escrow contract
		storedDenom := storedCoin.Denom
		fmt.Println("storedDenom", storedDenom)
		fmt.Println("msg.DenomMap", msg.DenomMap)
		denom := msg.DenomMap[storedDenom];
		fmt.Println("denom", denom)
		ibcDenomCheck := strings.Split(denom, "/")[0]

		if (ibcDenomCheck == "ibc") {
			hash := strings.Split(denom, "/")[1]

			fmt.Println("hash", hash)
			fmt.Println("hexBytes hash", bytes.HexBytes(hash))

			denomTrace, found := k.ibcTransfer.GetDenomTrace(ctx, bytes.HexBytes(hash))
			if (!found) {
				// TODO return proper type instead of ErrKeyNotFound
				return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The denomTrace with hash %v doesn't exist", hash)
			}

			fmt.Println("denomTrace", denomTrace)
			baseDenom := denomTrace.BaseDenom;

			fmt.Println("baseDenom from denomTrace", baseDenom)

			// the stored denom should be equal to the base_denom found in GetDenomTrace
			if (storedDenom != baseDenom) {
				// TODO return proper type instead of ErrKeyNotFound
				return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The base denom %v found from hash %v does not match the denom stored in the contract %v", baseDenom, hash , storedDenom)
			}
			
			
		} else {
			// token is not IBC, compare the denom still
			if (storedDenom != denom) {
				return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The provided denom %v does not match the denom stored in the contract %v", denom, storedDenom)
				// TODO return error: "The base denom found from the ibc hashed denom does not fit the denom stored in the contract"
			}
		}

		// if no error is found, append a new coin with the amount found in the contract, but with the denom provided in the msg.
		updatedCoin := sdk.Coin{Denom: denom, Amount: storedCoin.Amount}
		updatedFulfillerCoins = append(updatedFulfillerCoins, updatedCoin)

		fmt.Println("updatedFulfillerCoins", updatedFulfillerCoins)
	}

	// Make sure the fulfill request is not from the initiator of the escrow contract
	if escrow.Initiator == msg.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "Initator of the escrow can not fulfill it")
	}

	// Make sure the escrow status is "open"
	if escrow.Status != constants.StatusOpen {
		return nil, errors.Wrapf(types.ErrWrongEscrowStatus, "%v", escrow.Status)
	}

	initiator, _ := sdk.AccAddressFromBech32(escrow.Initiator)
	fulfiller, _ := sdk.AccAddressFromBech32(msg.Creator)

	conditionsValidity := k.ValidateConditions(ctx, escrow)

	if conditionsValidity {
		fmt.Println("TRYING TO SEND updatedFulfillerCoins TO THE INITIATOR")
		// If all the conditions are met, send fulfiller coins to initator
		errSendCoins := k.bank.SendCoins(ctx, fulfiller, initiator, updatedFulfillerCoins)
		if errSendCoins != nil {
			fmt.Println("ERROR WHILE SENDING updatedFulfillerCoins TO THE INITIATOR")
			return nil, errors.Wrapf(errSendCoins, types.ErrFulfillerCannotPay.Error())
		}

		// release the initiator assets and send them to the fulfiller
		errReleaseInitiatorCoins := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.InitiatorCoins)
		if errReleaseInitiatorCoins != nil {
			panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), errReleaseInitiatorCoins.Error()))
		}

		// change the escrow status to "closed"
		k.SetStatus(ctx, &escrow, constants.StatusClosed)
		// escrow.Status = constants.StatusClosed
	} else {
		fmt.Println("TRYING TO SEND updatedFulfillerCoins TO THE MODULE")
		// If not all conditions are met, escrow the fulfiller assets
		errEscrowInitiatorCoins := k.bank.SendCoinsFromAccountToModule(ctx, fulfiller, types.ModuleName, updatedFulfillerCoins)
		if errEscrowInitiatorCoins != nil {
			fmt.Println("ERROR WHILE SENDING updatedFulfillerCoins TO THE MODULE")
			return nil, errors.Wrapf(errEscrowInitiatorCoins, types.ErrFulfillerCannotPay.Error())
		}

		// Add the escrow to the list of pending escrows
		k.AddPendingEscrow(ctx, escrow)

		// change the escrow status to "pending"
		k.SetStatus(ctx, &escrow, constants.StatusPending)
		// escrow.Status = constants.StatusPending
	}

	escrow.Fulfiller = msg.Creator
	k.SetEscrow(ctx, escrow)
	return &types.MsgFulfillEscrowResponse{}, nil
}
