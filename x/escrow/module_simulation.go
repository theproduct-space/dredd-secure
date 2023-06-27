package escrow

import (
	"math/rand"

	"dredd-secure/testutil/sample"
	escrowsimulation "dredd-secure/x/escrow/simulation"
	"dredd-secure/x/escrow/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = escrowsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateEscrow = "op_weight_msg_create_escrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEscrow int = 100

	opWeightMsgFulfillEscrow = "op_weight_msg_fulfill_escrow"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFulfillEscrow int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	escrowGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&escrowGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEscrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEscrow, &weightMsgCreateEscrow, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEscrow = defaultWeightMsgCreateEscrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEscrow,
		escrowsimulation.SimulateMsgCreateEscrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFulfillEscrow int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFulfillEscrow, &weightMsgFulfillEscrow, nil,
		func(_ *rand.Rand) {
			weightMsgFulfillEscrow = defaultWeightMsgFulfillEscrow
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFulfillEscrow,
		escrowsimulation.SimulateMsgFulfillEscrow(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEscrow,
			defaultWeightMsgCreateEscrow,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				escrowsimulation.SimulateMsgCreateEscrow(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFulfillEscrow,
			defaultWeightMsgFulfillEscrow,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				escrowsimulation.SimulateMsgFulfillEscrow(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
