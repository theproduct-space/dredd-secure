package testutil

import (
	"context"

	bytes "github.com/cometbft/cometbft/libs/bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
)


func (escrow *MockTransferKeeper) ExpectGetDenomTrace(context context.Context, denomTraceHash bytes.HexBytes) *gomock.Call {

	if denomTraceHash == nil {
		panic("No denom hash was provided")
	}
	return escrow.EXPECT().GetDenomTrace(sdk.UnwrapSDKContext(context), denomTraceHash)
}
