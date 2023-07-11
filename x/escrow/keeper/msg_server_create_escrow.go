package keeper

import (
	"context"

	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/types"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEscrow(goCtx context.Context, msg *types.MsgCreateEscrow) (*types.MsgCreateEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var escrow = types.Escrow{
		Status: constants.StatusOpen,
		Initiator: msg.Creator,
		Fulfiller: "",
		InitiatorCoins: msg.InitiatorCoins,
		FulfillerCoins: msg.FulfillerCoins,
		StartDate: msg.StartDate,
		EndDate: msg.EndDate,
    }
	
	initiator, err := sdk.AccAddressFromBech32(msg.Creator)
    if err != nil {
        panic(err)
    }

	errSendCoins := k.bank.SendCoinsFromAccountToModule(ctx, initiator, types.ModuleName, escrow.InitiatorCoins)
    if errSendCoins != nil {
		return nil, errors.Wrapf(errSendCoins, types.ErrInitiatorCannotPay.Error())
    }
    k.AppendEscrow(ctx, escrow)

	return &types.MsgCreateEscrowResponse{}, nil
}
