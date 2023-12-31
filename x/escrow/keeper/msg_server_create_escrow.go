package keeper

import (
	"context"
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateEscrow creates a new escrow with with the provided msg details
func (k msgServer) CreateEscrow(goCtx context.Context, msg *types.MsgCreateEscrow) (*types.MsgCreateEscrowResponse, error) {
	tipAddress, errTipsAddr := sdk.AccAddressFromBech32(constants.Tips)
	if errTipsAddr != nil {
		panic(errTipsAddr.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a new escrow object with the provided details
	escrow := types.Escrow{
		Status:           constants.StatusOpen,
		Initiator:        msg.Creator,
		Fulfiller:        "",
		InitiatorCoins:   msg.InitiatorCoins,
		FulfillerCoins:   msg.FulfillerCoins,
		Tips:             msg.Tips,
		StartDate:        msg.StartDate,
		EndDate:          msg.EndDate,
		OracleConditions: msg.OracleConditions,
	}

	initiator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err.Error())
	}

	// Transfer the initiator's coins from their account to the escrow module
	errSendCoins := k.bank.SendCoinsFromAccountToModule(ctx, initiator, types.ModuleName, escrow.InitiatorCoins)
	if errSendCoins != nil {
		return nil, errors.Wrapf(errSendCoins, types.ErrInitiatorCannotPay.Error())
	}

	if escrow.Tips != nil && len(escrow.Tips) != 0 {
		errSendCoinsTips := k.bank.SendCoins(ctx, initiator, tipAddress, escrow.Tips)
		if errSendCoins != nil {
			return nil, errors.Wrapf(errSendCoinsTips, types.ErrInitiatorCannotPay.Error())
		}
	}

	// Append the newly created escrow to the store
	k.AppendEscrow(ctx, escrow)

	return &types.MsgCreateEscrowResponse{}, nil
}
