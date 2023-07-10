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

func (k msgServer) CancelEscrow(goCtx context.Context, msg *types.MsgCancelEscrow) (*types.MsgCancelEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	escrow, found := k.GetEscrow(ctx, msg.Id)

	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrKeyNotFound, "The escrow with key %d doesn't exist", msg.Id)
	}

	if escrow.Initiator != msg.Creator {
		return nil, errors.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not from the initiator")
	}

	if escrow.Status != constants.StatusOpen {
		return nil, errors.Wrapf(types.ErrWrongEscrowStatus, "%v", escrow.Status)
	}

	initiator, _ := sdk.AccAddressFromBech32(escrow.Initiator)
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, initiator, escrow.InitiatorCoins)
	if err != nil {
		panic(fmt.Sprintf(types.ErrCannotReleaseInitiatorAssets.Error(), err.Error()))
	}

	escrow.Status = constants.StatusCancelled
	k.SetEscrow(ctx, escrow)
	return &types.MsgCancelEscrowResponse{}, nil
}
