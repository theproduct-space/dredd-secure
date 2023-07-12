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

func setupMsgServerCancelEscrow(tb testing.TB) (types.MsgServer, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	tb.Helper()
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
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

	return server, context, ctrl, bankMock
}

func TestCancelEscrow(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

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

func TestCancelEscrowNotInitiator(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Bob,
		Id:      0,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestCancelEscrowDoesNotExist(t *testing.T) {
	msgServer, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      1,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

func TestCancelEscrowWrongStatus(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
	_, errFirstCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})
	require.Nil(t, errFirstCancel)

	_, errSecondCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id:      0,
	})
	require.NotNil(t, errSecondCancel)
	require.ErrorIs(t, errSecondCancel, types.ErrWrongEscrowStatus)
}

func TestCancelEscrowModuleCannotPay(t *testing.T) {
	msgServer, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)
	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(context, types.ModuleName, initiator, gomock.Any()).
		Return(errors.New("oops"))
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
