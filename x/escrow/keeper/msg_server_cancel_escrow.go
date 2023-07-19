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

// CancelEscrow cancels an existing escrow of the given id
func (k msgServer) CancelEscrow(goCtx context.Context, msg *types.MsgCancelEscrow) (*types.MsgCancelEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the escrow from the keeper
	escrow, found := k.GetEscrow(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
	}

	// Check if the cancel request is from the initiator of the escrow
	if escrow.Initiator != msg.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not from the initiator")
	}

	// Make sure the escrow status is "open"
	if escrow.Status != constants.StatusOpen {
		return nil, errors.Wrapf(types.ErrWrongEscrowStatus, "%v", escrow.Status)
	}

	initiator, _ := sdk.AccAddressFromBech32(escrow.Initiator)

	// Send back the initiator's coins from the escrow module
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, initiator, escrow.InitiatorCoins)
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), err.Error()))
	}

	// Update the escrow's status to "cancelled"
	escrow.Status = constants.StatusCancelled
	k.SetEscrow(ctx, escrow)

	return &types.MsgCancelEscrowResponse{}, nil
}
