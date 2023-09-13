package escrow

import (
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the escrow
	for _, elem := range genState.EscrowList {
		k.SetEscrow(ctx, elem)
	}

	k.SetPendingEscrows(ctx, genState.PendingEscrows)
	k.SetExpiringEscrows(ctx, genState.ExpiringEscrows)
	k.SetLastExecs(ctx, genState.LastExecs)
	k.SetSrcChannel(ctx, genState.SourceChannel)

	// Set escrow count
	k.SetEscrowCount(ctx, genState.EscrowCount)
	// Set all the oraclePrice
	for _, elem := range genState.OraclePriceList {
		k.SetOraclePrice(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EscrowList = k.GetAllEscrow(ctx)
	genesis.EscrowCount = k.GetEscrowCount(ctx)
	genesis.PendingEscrows = k.GetAllPendingEscrows(ctx)
	genesis.ExpiringEscrows = k.GetAllExpiringEscrows(ctx)
	genesis.LastExecs = k.GetLastExecs(ctx)
	genesis.SourceChannel = k.GetSrcChannel(ctx)
	genesis.PortId = k.GetPort(ctx)
	genesis.OraclePriceList = k.GetAllOraclePrice(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
