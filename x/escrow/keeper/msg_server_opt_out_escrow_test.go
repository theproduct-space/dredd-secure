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
	"time"

	keepertest "dredd-secure/testutil/keeper"

	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// setupMsgServerOptOutEscrow is a test helper function to setup the necessary dependencies for testing the OptOutEscrow message server function
func setupMsgServerOptOutEscrow(tb testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	now := time.Now()

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
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+60, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	if err != nil {
		tb.Fatalf("Failed to create escrow: %s", err)
	}

	// The bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		},
	})

	// Fulfill the escrow once
	_, errFulfill := server.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})
	require.Nil(tb, errFulfill)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(100),
	}})

	// Create an escrow using the message server and a valid MsgCreateEscrow
	_, err2 := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(100),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(900),
		}},
		Tips: nil,
		StartDate: "1588148578",
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	if err2 != nil {
		tb.Fatalf("Failed to create escrow: %s", err2)
	}

	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(900),
		},
	})

	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(100),
	}})

	// Fulfill the escrow once
	_, errFulfill2 := server.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
	})
	require.Nil(tb, errFulfill2)

	// Return the necessary components for testing
	return server, *k, context, ctrl, bankMock
}

// TestOptOutEscrow tests the OptOutEscrow message server function
func TestOptOutEscrow(t *testing.T) {
	msgServer, k, context, ctrl, bankMock := setupMsgServerOptOutEscrow(t)
	defer ctrl.Finish()

	// Expect the bank to refund the initiator's coins
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(9000),
	}})

	_, err := msgServer.OptOutEscrow(context, &types.MsgOptOutEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	escrow, _ := k.GetEscrow(sdk.UnwrapSDKContext(context), 0)

	require.Equal(t, constants.StatusOpen, escrow.GetStatus())

	require.Nil(t, err)
}

// TestOptOutEscrowNotFulfiller tests the scenario where a non-fulfiller tries to opt out of an escrow
func TestOptOutEscrowNotFulfiller(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerOptOutEscrow(t)
	defer ctrl.Finish()

	// Attempt to opt out of the escrow as a non-fulfiller
	_, err := msgServer.OptOutEscrow(context, &types.MsgOptOutEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})

	// Ensure an error is returned and it matches the expected ErrUnauthorized error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

// TestOptOutEscrowDoesNotExist tests the scenario where the escrow to be opt-ed out from does not exist
func TestOptOutEscrowDoesNotExist(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerOptOutEscrow(t)
	defer ctrl.Finish()

	// Attempt to opt out of a non-existent escrow
	_, err := msgServer.OptOutEscrow(context, &types.MsgOptOutEscrow{
		Creator: testutil.Bob,
		Id:      10,
	})

	// Ensure an error is returned and it matches the expected ErrKeyNotFound error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

// TestOptOutEscrowWrongStatus tests the scenario where an escrow with an innapropriate status is opt-ed out from
// to accomplish this, we try to opt out of a closed escrow.
func TestOptOutEscrowWrongStatus(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerOptOutEscrow(t)
	defer ctrl.Finish()

	// Attempt to opt out from a closed escrow
	_, errOptOut := msgServer.OptOutEscrow(context, &types.MsgOptOutEscrow{
		Creator: testutil.Bob,
		Id:      1,
	})

	// Ensure an error is returned and it matches the expected ErrWrongEscrowStatus error.
	require.NotNil(t, errOptOut)
	require.ErrorIs(t, errOptOut, types.ErrWrongEscrowStatus)
}

// TestOptOutEscrowModuleCannotPay tests the scenario where the module cannot release the fulfiller's assets during the opt out of an escrow.
func TestOptOutEscrowModuleCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerOptOutEscrow(t)
	defer ctrl.Finish()

	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// Set up the expectation that the module will attempt to send coins from the module account to the fulfiller's account,
	// but an error "oops" will occur
	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(context, types.ModuleName, fulfiller, gomock.Any()).
		Return(errors.New("oops"))

	// Ensure that the code panics with the expected error message
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "Module cannot release Fulfiller assets%!(EXTRA string=oops)", r)
	}()

	_, err := msgServer.OptOutEscrow(context, &types.MsgOptOutEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	if err != nil {
		require.Equal(t, "Module cannot release Fulfiller assets%!(EXTRA string=oops)", err.Error())
	}
}
