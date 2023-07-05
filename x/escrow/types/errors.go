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
	ErrFulfillerCannotPay    = errors.Register(ModuleName, 1113, "Fulfiller cannot pay")
	ErrCannotReleaseInitiatorAssets    = errors.Register(ModuleName, 1114, "Module cannot release Initiator assets")
	ErrCannotReleaseFulfillerAssets    = errors.Register(ModuleName, 1115, "Module cannot release Fulfiller assets")
)
