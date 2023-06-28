package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/escrow module sentinel errors
var (
	ErrSample            = errors.Register(ModuleName, 1100, "sample error")
	ErrWrongEscrowStatus = errors.Register(ModuleName, 2, "Wrong escrow status")
)
