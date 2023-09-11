package testutil

import (
	"context"

	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibcTransferTypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	"github.com/golang/mock/gomock"
)


func (escrow *MockTransferKeeper) ExpectGetDenomTrace(context context.Context, denomTraceHash string) (*gomock.Call, bool) {
	if denomTraceHash == "" {
		panic("No denom hash was provided")
	}
	ibcDenomCheck := strings.Split(denomTraceHash, "/")[0]
	if (ibcDenomCheck == "ibc") {
		hash := strings.Split(denomTraceHash, "/")[1]
		hashIbc, err := ibcTransferTypes.ParseHexHash(hash)
		if err != nil {
			// TODO proper error
			panic("Could not parse hex hash")
		}
	return escrow.EXPECT().GetDenomTrace(sdk.UnwrapSDKContext(context), hashIbc), true
	}
	return nil, false
}
