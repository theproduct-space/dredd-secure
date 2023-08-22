package types

// ValidateBasic is used for validating the packet
func (p OracleRequestPacketDataPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p OracleRequestPacketDataPacketData) GetBytes() ([]byte, error) {
	var modulePacket EscrowPacketData

	modulePacket.Packet = &EscrowPacketData_OracleRequestPacketDataPacket{&p}

	return modulePacket.Marshal()
}
