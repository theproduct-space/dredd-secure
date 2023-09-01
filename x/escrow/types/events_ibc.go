package types

// IBC events
const (
	EventTypeTimeout                        = "timeout"
	EventTypeOracleRequestPacketDataPacket  = "oracleRequestPacketData_packet"
	EventTypeOracleResponsePacketDataPacket = "oracleResponsePacketData_packet"
	EventTypePriceUpdate    = "price_update"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
	AttributeKeyRequestID  = "request_id"
	AttributeKeySymbol     = "symbol"
	AttributeKeyPrice      = "price"
	AttributeKeyTimestamp  = "timestamp"
)
