package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateEscrow{}, "escrow/CreateEscrow", nil)
	cdc.RegisterConcrete(&MsgCancelEscrow{}, "escrow/CancelEscrow", nil)
	cdc.RegisterConcrete(&MsgFulfillEscrow{}, "escrow/FulfillEscrow", nil)
	cdc.RegisterConcrete(&MsgSendOracleRequestPacketData{}, "escrow/SendOracleRequestPacketData", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateEscrow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelEscrow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFulfillEscrow{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendOracleRequestPacketData{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
