package keeper_test

import (
	"context"
	"errors"
	"testing"
	"time"
	"strconv"

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

func setupMsgServerFulfillEscrow(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(t, bankMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	now := time.Now()

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	// create an escrow that can be closed when the second party fulfills it.
	// ID : 0
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

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(99),
	}})
	// create an escrow that can only be closed in the future
	// ID : 1
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(99),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(1111),
		}},
		StartDate:      "2588148578",
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(9),
	}})
	// create another escrow that can only be closed in the future
	// ID : 2
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(9),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(111),
		}},
		StartDate:      "2588148578",
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(88),
	}})
	// create an escrow that can only be closed in the near future
	// ID : 3
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(88),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(1100),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 5, 10),
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(8),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 4
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(8),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(110),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 6, 10),
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(7),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 5
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(7),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(99),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 7, 10),
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(6),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 6
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(6),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(88),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 160, 10),
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(5),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 7
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(5),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(77),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 180, 10),
		EndDate:        "2788148978",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(4),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 8
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom: "token",
			Amount: sdk.NewInt(4),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(66),
		}},
		StartDate:      strconv.FormatInt(now.Unix() + 200, 10),
		EndDate:        "2788148978",
	})
	
	return server, *k, context, ctrl, bankMock
}

// Testing to fulfill the escrow that can be closed when the second party fulfills it.
func TestFulfillEscrow(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	// the bank is expected to "refund" the fulfiller (send escrowed InitiatorCoins to the fulfiller)
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	// the bank is expected to send the FulfillerCoins to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(9000),
	}})
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})

	require.Nil(t, err)
}

// Testing to fulfill the escrow that can only be closed in the future
func TestFulfillEscrowFuture(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(1111),
	}})
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 1,
	})

	require.Nil(t, err)
}

// Testing to fulfill multiple escrows that can only be closed in the future
func TestFulfillEscrowsFuture(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(1111),
	}})
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 1,
	})

	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(111),
	}})
	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 2,
	})

	require.Nil(t, err)
	require.Nil(t, err2)
}

// Testing to fulfill multiple escrows that can only be closed in the future
func TestFulfillEscrowsNearFuture(t *testing.T) {
	msgServer, k, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	// We fulfill the escrows in random order
	
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(1100),
	}})
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 3,
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(99),
	}})
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 5,
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(66),
	}})
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 8,
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(110),
	}})
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 4,
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(77),
	}})
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 7,
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(88),
	}})
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 6,
	})

	// Pending escrows list needs to be in order of start date
	// They have been added in the order 3, 5, 8, 4, 7, 6 and should be ordered as 3, 4, 5, 6, 7, 8
	// Which coincidently corresponds to the ID order
	pendingEscrowsIdList := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdList := []uint64{3, 4, 5, 6, 7, 8}

	require.EqualValues(t, pendingEscrowsIdList, controlPendingEscrowsIdList)

	// Wait 20 seconds
    time.Sleep(10 * time.Second)

	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(88),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(1100),
	}})
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(8),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(110),
	}})
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(7),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(99),
	}})

	// Fulfill pending escrows 3, 4 and 5
	k.FulfillPendingEscrows(sdk.UnwrapSDKContext(context))
	
	// Confirm they have been removed from the pending list
	pendingEscrowsIdListAfter := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdListAfter := []uint64{6, 7, 8}

	require.EqualValues(t, pendingEscrowsIdListAfter, controlPendingEscrowsIdListAfter)

	require.Nil(t, err)
}

func TestFulfillEscrowAsInitiator(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id: 0,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

func TestFulfillEscrowDoesNotExist(t *testing.T) {
	msgServer, _, context, ctrl, _ := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id: 55,
	})

	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

func TestFulfillEscrowWrongStatus (t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	
	// fullfill the escrow once
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom: "token",
		Amount: sdk.NewInt(1000),
	}})
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(9000),
	}})
	_, errFirstFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})
	require.Nil(t, errFirstFulfill)

	// then try to fulfill it again
	_, errSecondFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})
	require.NotNil(t, errSecondFulfill)
	require.ErrorIs(t, errSecondFulfill, types.ErrWrongEscrowStatus)
}

// Testing to fulfill the escrow that can be closed when the second party fulfills it, but the module cannot refund
func TestFulfillEscrowModuleCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// the bank is expected to send the FulfillerCoins from the fulfiller to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice ,[]sdk.Coin{{
		Denom: "stake",
		Amount: sdk.NewInt(9000),
	}})
	// the bank is expected to fail to unescrow the InitiatorCoins to send them to the fulfiller
    bankMock.EXPECT().
		SendCoinsFromModuleToAccount(context,types.ModuleName, fulfiller, gomock.Any()).
		Return(errors.New("oops"))
	defer func() {
		r := recover()
		require.NotNil(t, r, "The code did not panic")
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", r )
	}()
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})
}

// Testing to fulfill the escrow that can be closed when the second party fulfills it, but the fulfiller cannot pay the initiator
func TestFulfillEscrowFulfillerCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	initiator, _ := sdk.AccAddressFromBech32(testutil.Alice)
	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// the bank is expected to fail to send the FulfillerCoins from the fulfiller to the initiator
	bankMock.EXPECT().
		SendCoins(context,fulfiller, initiator, []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(9000),
		}}).
		Return(errors.New("oops"))
	
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 0,
	})

	require.NotNil(t, err)
    require.EqualError(t, err, "Fulfiller cannot pay: oops")
}

// Testing to fulfill the escrow that can only be closed in the future, but the fulfiller cannot pay the module
func TestFulfillEscrowFulfillerCannotPayModule(t *testing.T) {
	msgServer, _, context, ctrl, bankMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	fulfiller, _ := sdk.AccAddressFromBech32(testutil.Bob)

	// the bank is expected to fail to send the FulfillerCoins from the fulfiller to the initiator
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(context, fulfiller, types.ModuleName, []sdk.Coin{{
			Denom: "stake",
			Amount: sdk.NewInt(1111),
		}}).
		Return(errors.New("oops"))
	
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id: 1,
	})

	require.NotNil(t, err)
    require.EqualError(t, err, "Fulfiller cannot pay: oops")
}