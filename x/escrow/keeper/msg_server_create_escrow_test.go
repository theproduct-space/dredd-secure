package keeper_test

import (
	"context"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/constants"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/testutil"
	"dredd-secure/x/escrow/types"
	"errors"
	"testing"

	keepertest "dredd-secure/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// setupMsgServerCreateEscrow is a test helper function to setup the necessary dependencies for testing the CreateEscrow message server function
func setupMsgServerCreateEscrow(tb testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	ibcTransferMock := testutil.NewMockTransferKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock, ibcTransferMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// Return the necessary components for testing
	return server, *k, context, ctrl, bankMock
}

// TestCreateEscrow tests the CreateEscrow function of the message server.
func TestCreateEscrowTips(t *testing.T) {
	msgServer, keeper, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{Denom: "token", Amount: sdk.NewInt(999)},
	})

	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1)}})

	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator:          testutil.Bob,
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(999)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1)}},
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	})

	// Verify that the escrow was created successfully
	escrow, found := keeper.GetEscrow(sdk.UnwrapSDKContext(context), 0)
	require.True(t, found)
	require.EqualValues(t, types.Escrow{
		Id:               0,
		Status:           constants.StatusOpen,
		Initiator:        testutil.Bob,
		Fulfiller:        "",
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(999)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1)}},
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	}, escrow)

	require.Nil(t, err)
}

// TestCreateEscrow tests the CreateEscrow function of the message server.
func TestCreateEscrow(t *testing.T) {
	msgServer, keeper, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{Denom: "token", Amount: sdk.NewInt(1000)},
	})

	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator:          testutil.Bob,
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1000)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             nil,
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	})

	// Verify that the escrow was created successfully
	escrow, found := keeper.GetEscrow(sdk.UnwrapSDKContext(context), 0)
	require.True(t, found)
	require.EqualValues(t, types.Escrow{
		Id:               0,
		Status:           constants.StatusOpen,
		Initiator:        testutil.Bob,
		Fulfiller:        "",
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1000)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             nil,
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	}, escrow)

	require.Nil(t, err)
}

// TestCreateEscrowInitiatorCannotPay tests the scenario where the initiator cannot pay the required coins for escrow creation.
func TestCreateEscrowInitiatorCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)

	// Set up the expectation that the bank will attempt to send coins from the initiator's account to the module account
	// and an error "oops" will occur
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(context, initiator, types.ModuleName, gomock.Any()).
		Return(errors.New("oops"))

	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator:          testutil.Alice,
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1000)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             nil,
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	})

	// Verify that the expected error is returned
	require.NotNil(t, err)
	require.EqualError(t, err, "Initiator cannot pay: oops")
}

// TestCreateEscrowInitiatorCannotPayTips tests the scenario where the initiator cannot pay the required tip coins for escrow creation.
func TestCreateEscrowInitiatorCannotPayTips(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)

	// Set up the expectation that the bank will attempt to send coins from the initiator's account to the module account
	// and an error "oops" will occur
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(context, initiator, types.ModuleName, gomock.Any()).
		Return(errors.New("oops"))

	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator:          testutil.Alice,
		InitiatorCoins:   []sdk.Coin{{Denom: "token", Amount: sdk.NewInt(1000)}},
		FulfillerCoins:   []sdk.Coin{{Denom: "stake", Amount: sdk.NewInt(9000)}},
		Tips:             []sdk.Coin{{Denom: "uax", Amount: sdk.NewInt(9)}},
		StartDate:        "1588148578",
		EndDate:          "2788148978",
		OracleConditions: "",
	})

	// Verify that the expected error is returned
	require.NotNil(t, err)
	require.EqualError(t, err, "Initiator cannot pay: oops")
}
