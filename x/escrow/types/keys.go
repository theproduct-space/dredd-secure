package types

const (
	// ModuleName defines the module name
	ModuleName = "escrow"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_escrow"

    // Version defines the current version the IBC module supports
	Version = "bandchain-1"

	// PortID is the default port id that module binds to
	PortID = "escrow"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("escrow-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	EscrowKey         = "Escrow/value/"
	EscrowCountKey    = "Escrow/count/"
	PendingEscrowKey  = "Escrow/pending/"
	ExpiringEscrowKey = "Escrow/expiring/"
	LastExecsKey      = "Escrow/lastExec/"
)
