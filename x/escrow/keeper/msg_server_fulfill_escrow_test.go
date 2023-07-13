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

// setupMsgServerFulfillEscrow is a test helper function to setup the necessary dependencies for testing the FullfillEscrow message server function
func setupMsgServerFulfillEscrow(tb testing.TB) (types.MsgServer, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(1000),
		},
	})

	// Create an escrow that can be closed when the second party fulfills it
	_, errFirstCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(1000),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(9000),
			},
		},
		StartDate: "1588148578",
		EndDate:   "2788148978",
	})
	require.Nil(tb, errFirstCreate)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(99),
		},
	})

	// Create an escrow that can only be closed in the future
	_, errSecondCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(99),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(1111),
			},
		},
		StartDate: "4588148578",
		EndDate:   "4788148978",
	})
	require.Nil(tb, errSecondCreate)

	// Return the necessary components for testing
	return server, context, ctrl, bankMock
}

// TestFulfillEscrow tests the fulfillment of an escrow that can be closed when the second party fulfills it.
func TestFulfillEscrow(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// The bank is expected to "refund" the fulfiller (send escrowed InitiatorCoins to the fulfiller)
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(1000),
		},
	})

	// The bank is expected to send the FulfillerCoins to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		},
	})

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	require.Nil(t, err)
}

// TestFulfillEscrowFuture tests the fulfillment of an escrow that can only be closed in the future.
func TestFulfillEscrowFuture(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// The bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(1111),
		},
	})

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
	})

	require.Nil(t, err)
}

// TestFulfillEscrowAsInitiator tests the case where the initiator tries to fulfill the escrow.
func TestFulfillEscrowAsInitiator(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// Attempt to fulfill the escrow as the initiator
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected ErrUnauthorized error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

// TestFulfillEscrowDoesNotExist tests the case where the escrow to be fulfilled does not exist.
func TestFulfillEscrowDoesNotExist(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// Attempt to fulfill a non-existent escrow
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id:      55,
	})

	// Ensure an error is returned and it matches the expected ErrKeyNotFound error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

// TestFulfillEscrowWrongStatus tests the case where the escrow has already been fulfilled.
// to accomplish this, we try fulfilling the escrow two times.
func TestFulfillEscrowWrongStatus(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	
	// The bank is expected to "refund" the fulfiller (send escrowed InitiatorCoins to the fulfiller)
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{ 
		{
			Denom:  "token",
			Amount: sdk.NewInt(1000),
		},
	})
	// The bank is expected to send the FulfillerCoins to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		},
	})
	// Fulfill the escrow once
	_, errFirstFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})
	require.Nil(t, errFirstFulfill)

	// Attempt to fulfill the escrow again
	_, errSecondFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected ErrWrongEscrowStatus error.
	require.NotNil(t, errSecondFulfill)
	require.ErrorIs(t, errSecondFulfill, types.ErrWrongEscrowStatus)
}

// TestFulfillEscrowModuleCannotPay tests the case where the module cannot refund the initiator's assets.
func TestFulfillEscrowModuleCannotPay(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// The bank is expected to send the FulfillerCoins from the fulfiller to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		},
	})

	// The bank is expected to fail to unescrow the InitiatorCoins to send them to the fulfiller
	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(context, types.ModuleName, fulfiller, gomock.Any()).
		Return(errors.New("oops"))

	// Expect a panic to occur with the specified error message
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", r)
	}()

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	if err != nil {
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", err.Error())
	}
}

// TestFulfillEscrowFulfillerCannotPay tests the case where the fulfiller cannot pay the initiator.
func TestFulfillEscrowFulfillerCannotPay(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)
	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// The bank is expected to fail to send the FulfillerCoins from the fulfiller to the initiator
	bankMock.EXPECT().
		SendCoins(context, fulfiller, initiator, []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(9000),
			},
		}).
		Return(errors.New("oops"))

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected error.
	require.NotNil(t, err)
	require.EqualError(t, err, "Fulfiller cannot pay: oops")
}

// TestFulfillEscrowFulfillerCannotPayModule tests the case where the fulfiller cannot pay the module.
func TestFulfillEscrowFulfillerCannotPayModule(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// The bank is expected to fail to send the FulfillerCoins from the fulfiller to the module
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(context, fulfiller, types.ModuleName, []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(1111),
			},
		}).
		Return(errors.New("oops"))

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
	})

	// Ensure an error is returned and it matches the expected error.
	require.NotNil(t, err)
	require.EqualError(t, err, "Fulfiller cannot pay: oops")
}
