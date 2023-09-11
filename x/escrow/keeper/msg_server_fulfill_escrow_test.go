package keeper_test

import (
	"context"
	"dredd-secure/x/escrow"
	"dredd-secure/x/escrow/keeper"
	"dredd-secure/x/escrow/testutil"
	"dredd-secure/x/escrow/types"
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	keepertest "dredd-secure/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// setupMsgServerFulfillEscrow is a test helper function to setup the necessary dependencies for testing the FullfillEscrow message server function
func setupMsgServerFulfillEscrow(tb testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper, *testutil.MockTransferKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	ibcTransferMock := testutil.NewMockTransferKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock, ibcTransferMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	now := time.Now()

	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "BTC",
		ResolveTime: "120",
		Price: "25983000000000", // $25983 
	})
	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "ATOM",
		ResolveTime: "140",
		Price: "7000000000", // $7
	})

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
	// create an escrow that can be closed when the second party fulfills it.
	// with an array of two OracleConditions to validate, that will always be valid
	// ID : 0
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
		Tips: nil,
		StartDate: "1588148578",
		EndDate: "2788148978",
		OracleConditions: `[{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"BTC"}},{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"ATOM"}}]`,
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
	// ID : 1
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
		Tips: nil,
		StartDate: "4588148578",
		EndDate:   "4788148978",
		OracleConditions: "",
	})
	require.Nil(tb, errSecondCreate)

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(9),
	}})
	// create another escrow that can only be closed in the future
	// ID : 2
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(9),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(111),
		}},
		Tips: nil,
		StartDate: "2588148578",
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(88),
	}})
	// create an escrow that can only be closed in the near future
	// ID : 3
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(88),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(1100),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+5, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(8),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 4
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(8),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(110),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+6, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(7),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 5
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(7),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(99),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+7, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(6),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 6
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(6),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(88),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+160, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(5),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 7
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(5),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(77),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+180, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(4),
	}})
	// create another escrow that can only be closed in the near future
	// ID : 8
	server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{{
			Denom:  "token",
			Amount: sdk.NewInt(4),
		}},
		FulfillerCoins: []sdk.Coin{{
			Denom:  "stake",
			Amount: sdk.NewInt(66),
		}},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+200, 10),
		EndDate:   "2788148978",
		OracleConditions: "",
	})

	// The bank is expected to receive the CreatorCoins from the creator (to be escrowed)
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
	// create another escrow that can never be fulfill because of invalid OracleConditions
	// ID : 9
	_, errNinthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
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
		Tips: nil,
		StartDate: "1588148578",
		EndDate: "2788148978",
		OracleConditions: `[{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":0},{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"BTC"}},{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"ATOM"}}]`,
	})
	require.Nil(tb, errNinthCreate)

	// The bank is expected to receive the CreatorCoins from the creator (to be escrowed)
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
	// create another escrow that can never be fulfill because of invalid OracleConditions
	// ID : 10
	_, errTenthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
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
		Tips: nil,
		StartDate: "1588148578",
		EndDate: "2788148978",
		OracleConditions: `[{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"BTC"}},{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"ATOM"}}]`,
	})
	require.Nil(tb, errTenthCreate)

	// Return the necessary components for testing
	return server, *k, context, ctrl, bankMock, ibcTransferMock
}

func setupMsgServerFulfillEscrow02(tb testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper, *testutil.MockTransferKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	ibcTransferMock := testutil.NewMockTransferKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock, ibcTransferMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	now := time.Now()

	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "BTC",
		ResolveTime: "120",
		Price: "25983000000000", // $25983 
	})
	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "ATOM",
		ResolveTime: "140",
		Price: "7000000000", // $7
	})

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1000),
	}})
	// create an escrow that can be closed when the second party fulfills it.
	// with an array of two OracleConditions to validate, that will always be valid
	// ID : 0
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
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+200, 10),
		EndDate: "2788148978",
		OracleConditions: `[{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"BTC"}},{"label":"Oracle Token Price","name":"oracle-token-price","type":"oracleCondition","subConditions":[{"conditionType":"gt","dataType":"number","name":"price","label":"USD Price","value":1},{"conditionType":"lt","dataType":"number","name":"price","label":"USD Price","value":9999999999999}],"tokenOfInterest":{"symbol":"ATOM"}}]`,
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
	// ID : 1
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
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+200, 10),
		EndDate: "2788148978",
		OracleConditions: "",
	})
	require.Nil(tb, errSecondCreate)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(111),
		},
	})
	// Create an escrow that can only be closed in the future
	// ID : 2
	_, errThirdCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(111),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(2222),
			},
		},
		Tips: nil,
		StartDate: strconv.FormatInt(now.Unix()+200, 10),
		EndDate: "2788148979",
		OracleConditions: "",
	})
	require.Nil(tb, errThirdCreate)

	// Return the necessary components for testing
	return server, *k, context, ctrl, bankMock, ibcTransferMock
}

func setupMsgServerFulfillEscrow03(tb testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockBankKeeper, *testutil.MockTransferKeeper) {
	tb.Helper()

	// Setup the necessary dependencies
	ctrl := gomock.NewController(tb)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	ibcTransferMock := testutil.NewMockTransferKeeper(ctrl)
	k, ctx := keepertest.EscrowKeeperWithMocks(tb, bankMock, ibcTransferMock)
	escrow.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "BTC",
		ResolveTime: "120",
		Price: "25983000000000", // $25983 
	})
	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "ATOM",
		ResolveTime: "130",
		Price: "7000000000", // $7
	})
	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "AVAX",
		ResolveTime: "140",
		Price: "7000000000", // $7
	})
	k.SetOraclePrice(ctx, types.OraclePrice{
		Symbol: "AKRO",
		ResolveTime: "150",
		Price: "7000000000", // $7
	})

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(1),
	}})
	// create an escrow that can be closed when the second party fulfills it.
	// with an array of two OracleConditions to validate, that will always be valid
	// ID : 0
	_, errFirstCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(1),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(2),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694036136",
		EndDate: "253402214400",
		OracleConditions: "[{\"label\":\"Token Price\",\"name\":\"oracle-token-price\",\"type\":\"oracleCondition\",\"subConditions\":[{\"conditionType\":\"gt\",\"dataType\":\"number\",\"name\":\"price\",\"label\":\"USD Price\",\"value\":999999999999999}],\"tokenOfInterest\":{\"symbol\":\"AVAX\",\"name\":\"Avalanche\"}}]",
	})
	require.Nil(tb, errFirstCreate)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(12),
		},
	})
	// Create an escrow that can only be closed in the future
	// ID : 1
	_, errSecondCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(12),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(12),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1693972800",
		EndDate: "1695182400",
		OracleConditions: "[{\"label\":\"Token Price\",\"name\":\"oracle-token-price\",\"type\":\"oracleCondition\",\"subConditions\":[{\"conditionType\":\"eq\",\"dataType\":\"number\",\"name\":\"price\",\"label\":\"USD Price\"}]}]",
	})
	require.Nil(tb, errSecondCreate)

	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(2),
		},
	})
	// Create an escrow that can only be closed in the future
	// ID : 2
	_, errThirdCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(2),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(4),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694036728",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errThirdCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(123),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 3
	_, errFourthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(123),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(123),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694037779",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errFourthCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(123),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 4
	_, errFifthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(123),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(1),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694037937",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errFifthCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(12),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 5
	_, errSixthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(12),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(12),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694038181",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errSixthCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(12),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 6
	_, errSeventhCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(12),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(12),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694038342",
		EndDate: "253402214400",
		OracleConditions: "[{\"label\":\"Token Price\",\"name\":\"oracle-token-price\",\"type\":\"oracleCondition\",\"subConditions\":[{\"conditionType\":\"eq\",\"dataType\":\"number\",\"name\":\"price\",\"label\":\"USD Price\"}]}]",
	})
	require.Nil(tb, errSeventhCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(4),
	}})
	bankMock.ExpectSend(context, testutil.Alice, testutil.TipAccount, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(8),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 7
	_, errEighthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(4),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(5),
			},
		},
		Tips: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(8),
			},
		},
		StartDate: "1694099831",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errEighthCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(97),
	}})
	bankMock.ExpectSend(context, testutil.Alice, testutil.TipAccount, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(59985),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 8
	_, errNinethCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(97),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(6),
			},
		},
		Tips: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(59985),
			},
		},
		StartDate: "1694100013",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errNinethCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(12),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 9
	_, errTenthCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(12),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(12),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694059200",
		EndDate: "253402214400",
		OracleConditions: "[]",
	})
	require.Nil(tb, errTenthCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(3),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 10
	_, errEleventhCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(3),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(5),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694116982",
		EndDate: "253402214400",
		OracleConditions: "[{\"label\":\"Token Price\",\"name\":\"oracle-token-price\",\"type\":\"oracleCondition\",\"subConditions\":[{\"conditionType\":\"gt\",\"dataType\":\"number\",\"name\":\"price\",\"label\":\"USD Price\",\"value\":25860}],\"tokenOfInterest\":{\"symbol\":\"BTC\",\"name\":\"Bitcoin\"}}]",
	})
	require.Nil(tb, errEleventhCreate)
	// Expect the bank to receive payment from the creator
	bankMock.ExpectPay(context, testutil.Alice, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(5),
	}})
	// Create an escrow that can only be closed in the future
	// ID : 11
	_, errTwelvethCreate := server.CreateEscrow(context, &types.MsgCreateEscrow{
		Creator: testutil.Alice,
		InitiatorCoins: []sdk.Coin{
			{
				Denom:  "token",
				Amount: sdk.NewInt(5),
			},
		},
		FulfillerCoins: []sdk.Coin{
			{
				Denom:  "stake",
				Amount: sdk.NewInt(6),
			},
		},
		Tips: []sdk.Coin{},
		StartDate: "1694189541",
		EndDate: "253402214400",
		OracleConditions: "[{\"label\":\"Token Price\",\"name\":\"oracle-token-price\",\"type\":\"oracleCondition\",\"subConditions\":[{\"conditionType\":\"gt\",\"dataType\":\"number\",\"name\":\"price\",\"label\":\"USD Price\",\"value\":44444444444444440}],\"tokenOfInterest\":{\"symbol\":\"AKRO\",\"name\":\"Akropolis\"}}]",
	})
	require.Nil(tb, errTwelvethCreate)

	// Return the necessary components for testing
	return server, *k, context, ctrl, bankMock, ibcTransferMock
}

// TestFulfillEscrow tests the fulfillment of an escrow that can be closed when the second party fulfills it.
func TestFulfillEscrow(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
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

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	require.Nil(t, err)
}

// TestFulfillEscrowFuture tests the fulfillment of an escrow that can only be closed in the future.
func TestFulfillEscrowFuture(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// The bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(1111),
		},
	})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator:  testutil.Bob,
		Id:       1,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	require.Nil(t, err)
}

// Testing to fulfill multiple escrows that can only be closed in the future
func TestFulfillEscrowsFuture(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(1111),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(111),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      2,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	require.Nil(t, err)
	require.Nil(t, err2)
}

// Testing to fulfill multiple escrows that can only be closed in the future
func TestFulfillEscrowsNearFuture(t *testing.T) {
	msgServer, k, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()
	// We fulfill the escrows in random order

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(1100),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      3,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(99),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      5,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(66),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      8,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(110),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      4,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(77),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")
	
	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      7,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(88),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      6,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
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
		Denom:  "token",
		Amount: sdk.NewInt(88),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(1100),
	}})
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(8),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(110),
	}})
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{{
		Denom:  "token",
		Amount: sdk.NewInt(7),
	}})
	bankMock.ExpectRefund(context, testutil.Alice, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(99),
	}})

	// Fulfill pending escrows 3, 4 and 5
	k.CancelExpiredEscrows(sdk.UnwrapSDKContext(context))
	k.FulfillPendingEscrows(sdk.UnwrapSDKContext(context))

	// Confirm they have been removed from the pending list
	pendingEscrowsIdListAfter := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdListAfter := []uint64{6, 7, 8}

	require.EqualValues(t, pendingEscrowsIdListAfter, controlPendingEscrowsIdListAfter)

	require.Nil(t, err)
}

// Testing the pending list tracking
func TestFulfillEscrowsPendingList(t *testing.T) {
	msgServer, k, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(88),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err1 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      6,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(77),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      7,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(66),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err3 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      8,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Pending escrows list needs to be in order of start date
	// They have been added in the order 6, 7, 8 and should be ordered as 6, 7, 8
	pendingEscrowsIdList := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdList := []uint64{6, 7, 8}

	require.EqualValues(t, pendingEscrowsIdList, controlPendingEscrowsIdList)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
}

// Testing the pending list tracking
func TestFulfillEscrowsPendingList02(t *testing.T) {
	msgServer, k, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow02(t)
	defer ctrl.Finish()

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(9000),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err1 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(1111),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Pending escrows list needs to be in order of start date
	// They have been added in the order 1, 0 and should be ordered as 1, 0
	pendingEscrowsIdList := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdList := []uint64{1, 0}

	require.EqualValues(t, pendingEscrowsIdList, controlPendingEscrowsIdList)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

// Testing the pending list tracking
func TestFulfillEscrowsPendingList03(t *testing.T) {
	msgServer, k, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow02(t)
	defer ctrl.Finish()

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(1111),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(9000),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err1 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Pending escrows list needs to be in order of start date
	// They have been added in the order 0, 1 and should be ordered as 0, 1
	pendingEscrowsIdList := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdList := []uint64{0, 1}

	require.EqualValues(t, pendingEscrowsIdList, controlPendingEscrowsIdList)
	require.Nil(t, err1)
	require.Nil(t, err2)
}

// Testing the pending list tracking
func TestFulfillEscrowsPendingList04(t *testing.T) {
	msgServer, k, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow03(t)
	defer ctrl.Finish()

	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(2),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err1 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// The bank is expected to "refund" the fulfiller (send escrowed InitiatorCoins to the fulfiller)
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(2),
		},
	})
	// The bank is expected to send the FulfillerCoins to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(4),
		},
	})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err2 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      2,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// The bank is expected to "refund" the fulfiller (send escrowed InitiatorCoins to the fulfiller)
	bankMock.ExpectRefund(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "token",
			Amount: sdk.NewInt(3),
		},
	})
	// The bank is expected to send the FulfillerCoins to the initiator
	bankMock.ExpectSend(context, testutil.Bob, testutil.Alice, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(5),
		},
	})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err3 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      10,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	// the bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{{
		Denom:  "stake",
		Amount: sdk.NewInt(6),
	}})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err4 := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      11,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Pending escrows list needs to be in order of start date
	// They have been added in the order 0, 1 and should be ordered as 0, 1
	pendingEscrowsIdList := k.GetAllPendingEscrows(sdk.UnwrapSDKContext(context))
	controlPendingEscrowsIdList := []uint64{0, 11}

	require.EqualValues(t, controlPendingEscrowsIdList, pendingEscrowsIdList)
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
	require.Nil(t, err4)
}

// Testing the expiring list tracking
func TestFulfillEscrowsExpiringList02(t *testing.T) {
	_, k, context, ctrl, _, _ := setupMsgServerFulfillEscrow02(t)
	defer ctrl.Finish()

	// Expiring escrows list needs to be in order of end date
	// They have been added in the order 0, 1, 2 and should be ordered as 1, 0, 2
	expiringEscrowsIdList := k.GetAllExpiringEscrows(sdk.UnwrapSDKContext(context))
	controlExpiringEscrowsIdList := []uint64{1, 0, 2}

	require.EqualValues(t, expiringEscrowsIdList, controlExpiringEscrowsIdList)
}

func test1Function(args ...interface{}) interface{} {
	return "Test1"
}

func test2Function(args ...interface{}) interface{} {
	if len(args) > 0 {
		return args[0]
	}
	return nil
}

// Tests if the function used to execute a function every X seconds works correctly
func TestExecTimerUtilFunc(t *testing.T) {
	_, k, context, _, _, _ := setupMsgServerFulfillEscrow(t)

	execs := []keeper.Exec{
		{
			ID:       "test1",
			Function: test1Function,
			Args:     nil,
			DelayS:   -1,
		},
		{
			ID:       "test2",
			Function: test2Function,
			Args:     []interface{}{"Test2"},
			DelayS:   1,
		},
	}

	exec1 := []string{"Test1", "Test2"}
	exec2 := []string{"Test1"}
	exec3 := []string{"Test1", "Test2"}
	results := k.ExecuteAfterNSeconds(sdk.UnwrapSDKContext(context), execs)
	castResults := make([]string, 0)
	for _, result := range results {
		castResults = append(castResults, fmt.Sprintf("%v", result))
	}
	require.EqualValues(t, exec1, castResults)
	results = k.ExecuteAfterNSeconds(sdk.UnwrapSDKContext(context), execs)
	castResults = make([]string, 0)
	for _, result := range results {
		castResults = append(castResults, fmt.Sprintf("%v", result))
	}
	require.EqualValues(t, exec2, castResults)
	time.Sleep(2 * time.Second)
	results = k.ExecuteAfterNSeconds(sdk.UnwrapSDKContext(context), execs)
	castResults = make([]string, 0)
	for _, result := range results {
		castResults = append(castResults, fmt.Sprintf("%v", result))
	}
	require.EqualValues(t, exec3, castResults)
}

// TestFulfillEscrowAsInitiator tests the case where the initiator tries to fulfill the escrow.
func TestFulfillEscrowAsInitiator(t *testing.T) {
	msgServer, _, context, ctrl, _, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")
	// Attempt to fulfill the escrow as the initiator
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Ensure an error is returned and it matches the expected ErrUnauthorized error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrUnauthorized)
}

// TestFulfillEscrowDoesNotExist tests the case where the escrow to be fulfilled does not exist.
func TestFulfillEscrowDoesNotExist(t *testing.T) {
	msgServer, _, context, ctrl, _, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	// Attempt to fulfill a non-existent escrow
	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Alice,
		Id:      55,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Ensure an error is returned and it matches the expected ErrKeyNotFound error.
	require.NotNil(t, err)
	require.ErrorIs(t, err, sdkerrors.ErrKeyNotFound)
}

// TestFulfillEscrowWrongStatus tests the case where the escrow has already been fulfilled.
// to accomplish this, we try fulfilling the escrow two times.
func TestFulfillEscrowWrongStatus(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
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

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	// Fulfill the escrow once
	_, errFirstFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})
	require.Nil(t, errFirstFulfill)

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	// Attempt to fulfill the escrow again
	_, errSecondFulfill := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Ensure an error is returned and it matches the expected ErrWrongEscrowStatus error.
	require.NotNil(t, errSecondFulfill)
	require.ErrorIs(t, errSecondFulfill, types.ErrWrongEscrowStatus)
}

// TestFulfillEscrowModuleCannotPay tests the case where the module cannot refund the initiator's assets.
func TestFulfillEscrowModuleCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
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

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	if err != nil {
		require.Equal(t, "Module cannot release Initiator assets%!(EXTRA string=oops)", err.Error())
	}
}

// TestFulfillEscrowFulfillerCannotPay tests the case where the fulfiller cannot pay the initiator.
func TestFulfillEscrowFulfillerCannotPay(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
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

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      0,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Ensure an error is returned and it matches the expected error.
	require.NotNil(t, err)
	require.EqualError(t, err, "Fulfiller cannot pay: oops")
}

// TestFulfillEscrowFulfillerCannotPayModule tests the case where the fulfiller cannot pay the module.
func TestFulfillEscrowFulfillerCannotPayModule(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
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

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      1,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	// Ensure an error is returned and it matches the expected error.
	require.NotNil(t, err)
	require.EqualError(t, err, "Fulfiller cannot pay: oops")
}

// TestFulfillEscrow tests the fulfillment of an escrow that can never be closed due to invalid OracleConditions
func TestFulfillEscrowInvalidOracleConditions(t *testing.T) {
	msgServer, _, context, ctrl, bankMock, ibcTransferMock := setupMsgServerFulfillEscrow(t)
	defer ctrl.Finish()

	// The bank is expected to receive the FulfillerCoins from the fulfiller (to be escrowed)
	bankMock.ExpectPay(context, testutil.Bob, []sdk.Coin{
		{
			Denom:  "stake",
			Amount: sdk.NewInt(9000),
		},
	})

	ibcTransferMock.ExpectGetDenomTrace(context, "stake")

	_, err := msgServer.FulfillEscrow(context, &types.MsgFulfillEscrow{
		Creator: testutil.Bob,
		Id:      9,
		DenomMap: []*types.KeyVal{
			{
				Key:   "stake",
				Value: "stake",
			},
		},
	})

	require.Nil(t, err)
}