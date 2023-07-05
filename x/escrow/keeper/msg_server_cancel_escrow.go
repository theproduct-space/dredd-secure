package keeper

import (
	"context"
	"dredd-secure/x/escrow/types"

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

	if escrow.Status != "open" {
		return nil, errors.Wrapf(types.ErrWrongEscrowStatus, "%v", escrow.Status)
	}

	initiator, _ := sdk.AccAddressFromBech32(escrow.Initiator)
	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, initiator, escrow.InitiatorCoins)
	if err != nil {
		return nil, err
	}

	escrow.Status = "cancelled"
	k.SetEscrow(ctx, escrow)
	return &types.MsgCancelEscrowResponse{}, nil
}
