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

func setupMsgServerCreateEscrow(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(t, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, bankMock
}

func TestCreateEscrow(t *testing.T) {
	msgServer, keeper, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})

	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Bob,
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

	// After the CreateEscrow method call, we expect to find an escrow with Id=0
	escrow, found := keeper.GetEscrow(sdk.UnwrapSDKContext(context), 0)
	require.True(t, found)
	require.EqualValues(t, types.Escrow{
		Id:        0,
		Status:    constants.StatusOpen,
		Initiator: testutil.Bob,
		Fulfiller: "",
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
	}, escrow)

	require.Nil(t, err)
}

func TestCreateEscrowInitiatorCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerCreateEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(context, initiator, types.ModuleName, gomock.Any()).
		Return(errors.New("oops"))
	_, err := msgServer.CreateEscrow(context, &types.MsgCreateEscrow{
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
	require.NotNil(t, err)
	require.EqualError(t, err, "Initiator cannot pay: oops")
}
