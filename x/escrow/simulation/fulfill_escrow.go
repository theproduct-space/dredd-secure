package simulation

import (
	"math/rand"

	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFulfillEscrow(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFulfillEscrow{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the FulfillEscrow simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "FulfillEscrow simulation not implemented"), nil, nil
	}
}
