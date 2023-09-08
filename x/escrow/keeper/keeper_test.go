package keeper_test

import (
	"testing"
	
	"github.com/stretchr/testify/require"

	testkeeper "dredd-secure/testutil/keeper"
)

func TestGetSetSrcChannel(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.EscrowKeeper(t)

	channelRequest := "channel-1"

	k.SetSrcChannel(ctx, channelRequest)

	storedSrcChannel:= k.GetSrcChannel(ctx)

	require.EqualValues(t, channelRequest, storedSrcChannel)
}

func TestHandleChannelRequest(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.EscrowKeeper(t)

	channelRequest := "channel-1"

	k.HandleChannelRequest(ctx, channelRequest)

	storedChannelRequest := k.GetSrcChannel(ctx)

	require.EqualValues(t, channelRequest, storedChannelRequest)
}