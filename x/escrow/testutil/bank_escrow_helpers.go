package testutil

import (
	"context"

	"dredd-secure/x/escrow/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)

func (escrow *MockBankKeeper) ExpectAny(context context.Context) {
	escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
}


func (escrow *MockBankKeeper) ExpectPay(context context.Context, who string, coins []sdk.Coin) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromAccountToModule(sdk.UnwrapSDKContext(context), whoAddr, types.ModuleName, coins)
}

func (escrow *MockBankKeeper) ExpectRefund(context context.Context, who string, coins []sdk.Coin) *gomock.Call {
	whoAddr, err := sdk.AccAddressFromBech32(who)
	if err != nil {
		panic(err)
	}
	return escrow.EXPECT().SendCoinsFromModuleToAccount(sdk.UnwrapSDKContext(context), types.ModuleName, whoAddr, coins)
}

func (escrow *MockBankKeeper) ExpectSend(context context.Context, from string, to string, coins []sdk.Coin) *gomock.Call {
	fromAddr, errFrom := sdk.AccAddressFromBech32(from)
	toAddr, errTo := sdk.AccAddressFromBech32(to)
	if errFrom != nil {
		panic(errFrom)
	}
	if errTo != nil {
		panic(errTo)
	}
	return escrow.EXPECT().SendCoins(sdk.UnwrapSDKContext(context), fromAddr, toAddr, coins)
}