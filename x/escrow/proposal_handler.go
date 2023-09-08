package escrow

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"
)

func NewUpdateChannelRequestProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateChannelRequestProposal:
			return handleUpdateChannelRequestProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized UpdateChannelRequest proposal content type: %T", c)
		}
	}
}

func handleUpdateChannelRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.UpdateChannelRequestProposal) error {
	k.HandleChannelRequest(ctx, p.ChannelRequest)
	return nil
}