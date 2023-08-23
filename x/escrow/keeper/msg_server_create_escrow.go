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
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a new escrow object with the provided details
	escrow := types.Escrow{
		Status:         constants.StatusOpen,
		Initiator:      msg.Creator,
		Fulfiller:      "",
		InitiatorCoins: msg.InitiatorCoins,
		FulfillerCoins: msg.FulfillerCoins,
		Tips:			msg.Tips,
		StartDate:      msg.StartDate,
		EndDate:        msg.EndDate,
		ApiConditions:  msg.ApiConditions,
	}

	initiator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	// Transfer the initiator's coins from their account to the escrow module
	sendingCoins := escrow.InitiatorCoins
	if escrow.Tips != nil {
		sendingCoins = []sdk.Coin{escrow.InitiatorCoins[0], escrow.Tips[0]}
	}

	errSendCoins := k.bank.SendCoinsFromAccountToModule(ctx, initiator, types.ModuleName, sendingCoins)
	if errSendCoins != nil {
		return nil, errors.Wrapf(errSendCoins, types.ErrInitiatorCannotPay.Error())
	}

	// Append the newly created escrow to the store
	k.AppendEscrow(ctx, escrow)

	return &types.MsgCreateEscrowResponse{}, nil
}
