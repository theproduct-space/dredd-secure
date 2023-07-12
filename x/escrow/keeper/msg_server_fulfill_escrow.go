package keeper

import (
	"context"
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"
	"fmt"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// FulfillEscrow fulfills an existing escrow of the given id
func (k msgServer) FulfillEscrow(goCtx context.Context, msg *types.MsgFulfillEscrow) (*types.MsgFulfillEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the escrow from the keeper
	escrow, found := k.GetEscrow(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
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
		errSendCoins := k.bank.SendCoins(ctx, fulfiller, initiator, escrow.FulfillerCoins)
		if errSendCoins != nil {
			return nil, errors.Wrapf(errSendCoins, types.ErrFulfillerCannotPay.Error())
		}

		// release the initiator assets and send them to the fulfiller
		errReleaseInitiatorCoins := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.InitiatorCoins)
		if errReleaseInitiatorCoins != nil {
			panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), errReleaseInitiatorCoins.Error()))
		}

		// change the escrow status to "closed"
		escrow.Status = constants.StatusClosed
	} else {
		// If not all conditions are met, escrow the fulfiller assets
		errEscrowInitiatorCoins := k.bank.SendCoinsFromAccountToModule(ctx, fulfiller, types.ModuleName, escrow.FulfillerCoins)
		if errEscrowInitiatorCoins != nil {
			return nil, errors.Wrapf(errEscrowInitiatorCoins, types.ErrFulfillerCannotPay.Error())
		}

		// change the escrow status to "pending"
		escrow.Status = constants.StatusPending
	}

	escrow.Fulfiller = msg.Creator
	k.SetEscrow(ctx, escrow)
	return &types.MsgFulfillEscrowResponse{}, nil
}
