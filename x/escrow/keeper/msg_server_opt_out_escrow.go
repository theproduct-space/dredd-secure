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

// OptOutEscrow opts out a fulfiller of an escrow of the given id
func (k msgServer) OptOutEscrow(goCtx context.Context, msg *types.MsgOptOutEscrow) (*types.MsgOptOutEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Retrieve the escrow from the keeper
	escrow, found := k.GetEscrow(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
	}

	// Check if the opt out request is from the fulfiller of the escrow
	if escrow.Fulfiller != msg.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "Cannot opt out: not from the fulfiller")
	}

	// Make sure the escrow status is "pending"
	if escrow.Status != constants.StatusPending {
		return nil, errors.Wrapf(types.ErrWrongEscrowStatus, "%v", escrow.Status)
	}

	fulfiller, _ := sdk.AccAddressFromBech32(escrow.Fulfiller)

	// Send back the fulfiller's coins from the escrow module
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, fulfiller, escrow.FulfillerCoins)
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotReleaseFulfillerAssets.Error(), err.Error()))
	}

	// Update the escrow's status back to "open"
	k.SetStatus(ctx, &escrow, constants.StatusOpen)
	k.SetEscrow(ctx, escrow)

	return &types.MsgOptOutEscrowResponse{}, nil
}
