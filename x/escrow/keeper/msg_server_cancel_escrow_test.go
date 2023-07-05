package keeper_test

import (
	"context"
	"testing"

	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/testutil"
	"dredd-secure/x/escrow/types"

	keepertest "dredd-secure/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMsgServerCancelEscrow(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(t, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(1000),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(9000),
		}},
		StartDate:      "1588148578",
		EndDate:        "2788148978",
	})
	
	return server, *k, context, ctrl, bankMock
}

func TestCancelEscrow(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()
	
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id: 0,
	})

	require.Nil(t, err)
}

func TestCancelEscrowNotInitiator(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()
	
	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestCancelEscrowDoesNotExist(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()
	
	_, err := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id: 1,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

func TestCancelEscrowWrongStatus (t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerCancelEscrow(t)
	defer ctrl.Finish()
	
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	_, errFirstCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id: 0,
	})
	require.Nil(t, errFirstCancel)

	_, errSecondCancel := msgServer.CancelEscrow(context, &types.MsgCancelEscrow{
		Creator: testutil.Alice,
		Id: 0,
	})
	require.NotNil(t, errSecondCancel)
	require.ErrorIs(t, errSecondCancel, types.ErrWrongEscrowStatus)
}

// func TestCancelEscrowModuleCannotPay(t *testing.T) {
// 	msgServer, _, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
// 	defer ctrl.Finish()

// 	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)
//     bankMock.EXPECT().
//         SendCoinsFromAccountToModule(context, initiator, types.ModuleName, gomock.Any()).
//         Return(errors.New("oops"))
// 	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
// 		Creator: testutil.Alice,
// 		InitiatorCoins: []sdk.Coin{{
// 			Denom: "token",
// 			Amount: sdk.NewInt(1000),
// 		}},
// 		FulfillerCoins: []sdk.Coin{{
// 			Denom: "stake",
// 			Amount: sdk.NewInt(9000),
// 		}},
// 		StartDate:      "1588148578",
// 		EndDate:        "2788148978",
// 	})
//     require.NotNil(t, err)
//     require.EqualError(t, err, "Initiator cannot pay: oops")
// }