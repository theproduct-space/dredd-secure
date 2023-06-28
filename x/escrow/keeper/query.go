package keeper

import (
	"dredd-secure/x/escrow/types"
)

var _ types.QueryServer = Keeper{}
