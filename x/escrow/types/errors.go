package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/escrow module sentinel errors
var (
	ErrSample            = errors.Register(ModuleName, 1110, "sample error")
	ErrWrongEscrowStatus = errors.Register(ModuleName, 1111, "Wrong escrow status")
	ErrInitiatorCannotPay    = errors.Register(ModuleName, 1112, "Initiator cannot pay")
)
