package keeper

import (
	"context"

	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEscrow(goCtx context.Context, msg *types.MsgCreateEscrow) (*types.MsgCreateEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var escrow = types.Escrow{
		Status: "open",
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

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, initiator, types.ModuleName, escrow.InitiatorCoins)
    if sdkError != nil {
        return nil, sdkError
    }
    k.AppendEscrow(ctx, escrow)

	return &types.MsgCreateEscrowResponse{}, nil
}
