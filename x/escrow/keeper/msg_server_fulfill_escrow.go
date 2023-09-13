package keeper

import (
	"context"
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"
	"fmt"
	"strings"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ibcTransferTypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

// FulfillEscrow fulfills an existing escrow of the given id
func (k msgServer) FulfillEscrow(goCtx context.Context, msg *types.MsgFulfillEscrow) (*types.MsgFulfillEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the escrow from the keeper
	escrow, found := k.GetEscrow(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
	}

	// Create a new slice to hold the updated FulfillerCoins
	var updatedFulfillerCoins sdk.Coins
	for _, storedCoin := range escrow.FulfillerCoins {
		// For every fulfillercoins stored in the escrow contract, we want to verify that the denom passed as argument matches the denom stored on the contract
		// The IBC denom needs to be analyzed in order to retrieve their base_denom.
		// Then we can compare the base_denom foundfrom the ibc hash to the base_denom that was stored in the escrow contract
		storedDenom := storedCoin.Denom
		var denom string
		for _, keyVal := range msg.DenomMap {
			if keyVal.Key == storedDenom {
				denom = keyVal.Value
				break
			}
		}
		ibcDenomCheck := strings.Split(denom, "/")[0]

		if ibcDenomCheck == "ibc" {
			hash := strings.Split(denom, "/")[1]

			hashIbc, err := ibcTransferTypes.ParseHexHash(hash)
			if err != nil {
				return nil, err
			}

			denomTrace, found := k.ibcTransfer.GetDenomTrace(ctx, hashIbc)
			if !found {
				return nil, errors.Wrapf(types.ErrDenomTraceNotExist, "%v", hash)
			}

			baseDenom := denomTrace.BaseDenom

			// the stored denom should be equal to the base_denom found in GetDenomTrace
			if storedDenom != baseDenom {
				// TODO return proper type instead of ErrKeyNotFound
				return nil, errors.Wrapf(types.ErrIncompatibleDenom, "%v", baseDenom)
			}

		} else if storedDenom != denom {
			// token is not IBC, compare the denom and return error if they do not match
			return nil, errors.Wrapf(types.ErrIncompatibleDenom, "%v", denom)
		}

		// if no error is found, append a new coin with the amount found in the contract, but with the denom provided in the msg.
		updatedCoin := sdk.Coin{Denom: denom, Amount: storedCoin.Amount}
		updatedFulfillerCoins = append(updatedFulfillerCoins, updatedCoin)
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
		// If all the conditions are met, send fulfiller coins to initator
		errSendCoins := k.bank.SendCoins(ctx, fulfiller, initiator, updatedFulfillerCoins)
		if errSendCoins != nil {
			return nil, errors.Wrapf(errSendCoins, types.ErrFulfillerCannotPay.Error())
		}

		// release the initiator assets and send them to the fulfiller
		errReleaseInitiatorCoins := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.InitiatorCoins)
		if errReleaseInitiatorCoins != nil {
			panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), errReleaseInitiatorCoins.Error()))
		}

		// change the escrow status to "closed"
		k.SetStatus(ctx, &escrow, constants.StatusClosed)
	} else {
		// If not all conditions are met, escrow the fulfiller assets
		errEscrowInitiatorCoins := k.bank.SendCoinsFromAccountToModule(ctx, fulfiller, types.ModuleName, updatedFulfillerCoins)
		if errEscrowInitiatorCoins != nil {
			return nil, errors.Wrapf(errEscrowInitiatorCoins, types.ErrFulfillerCannotPay.Error())
		}

		// Add the escrow to the list of pending escrows
		k.AddPendingEscrow(ctx, escrow)

		// change the escrow status to "pending"
		k.SetStatus(ctx, &escrow, constants.StatusPending)
	}

	escrow.Fulfiller = msg.Creator
	k.SetEscrow(ctx, escrow)
	return &types.MsgFulfillEscrowResponse{}, nil
}
