package keeper_test

import (
	"context"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/testutil"
	"dredd-secure/x/escrow/types"
	"errors"
	"testing"

	keepertest "dredd-secure/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// setupMsgServerCancelEscrow is a test helper function to setup the necessary dependencies for testing the CancelEscrow message server function
func setupMsgServerCancelEscrow(tb testing.TB) (types.MsgServer, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})

	// Create an escrow using the message server and a valid MsgCreateEscrow
	_, err := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(1000),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		}},
		StartDate: "1588148578",
		EndDate:   "2788148978",
	})

	if err != nil {
		tb.Fatalf("Failed to create escrow: %s", err)
	}

	// Return the necessary components for testing
	return server, context, ctrl, bankMock
}


// TestCancelEscrow tests the CancelEscrow message server function
func TestCancelEscrow(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	// Expect the bank to refund the initiator's coins
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})

	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})

	require.Nil(t, err)
}

// TestCancelEscrowNotInitiator tests the scenario where a non-initiator tries to cancel an escrow
func TestCancelEscrowNotInitiator(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	// Attempt to cancel the escrow as a non-initiator
	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected ErrUnauthorized error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

// TestCancelEscrowDoesNotExist tests the scenario where the escrow to be canceled does not exist
func TestCancelEscrowDoesNotExist(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	// Attempt to cancel a non-existent escrow
	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      1,
	})

	// Ensure an error is returned and it matches the expected ErrKeyNotFound error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

// TestCancelEscrowWrongStatus tests the scenario where an escrow with an innapropriate status is canceled
// to accomplish this, we try cancelling the escrow two times.
func TestCancelEscrowWrongStatus(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	// Expect the bank to refund the initiator's coins
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})

	// Cancel the escrow for the first time
	_, errFirstCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})
	require.Nil(t, errFirstCancel)

	// Attempt to cancel the same escrow again
	_, errSecondCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected ErrWrongEscrowStatus error.
	require.NotNil(t, errSecondCancel)
	require.ErrorIs(t, errSecondCancel, types.ErrWrongEscrowStatus)
}

// TestCancelEscrowModuleCannotPay tests the scenario where the module cannot release the initiator's assets during the cancellation of an escrow.
func TestCancelEscrowModuleCannotPay(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)

	// Set up the expectation that the module will attempt to send coins from the module account to the initiator's account,
	// but an error "oops" will occur
	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(context, types.ModuleName, initiator, gomock.Any()).
		Return(errors.New("oops"))

	// Ensure that the code panics with the expected error message
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", r)
	}()

	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})

	if err != nil {
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", err.Error())
	}
}