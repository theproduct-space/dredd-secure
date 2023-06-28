package keeper

import (
	"context"

	"dredd-secure/x/escrow/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FulfillEscrow(goCtx context.Context, msg *types.MsgFulfillEscrow) (*types.MsgFulfillEscrowResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFulfillEscrowResponse{}, nil
}
