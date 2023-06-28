// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dreddsecure/escrow/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateEscrow struct {
	Creator        string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	InitiatorCoins []types.Coin `protobuf:"bytes,2,rep,name=initiatorCoins,proto3" json:"initiatorCoins"`
	FulfillerCoins []types.Coin `protobuf:"bytes,3,rep,name=fulfillerCoins,proto3" json:"fulfillerCoins"`
	StartDate      string       `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate        string       `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (m *MsgCreateEscrow) Reset()         { *m = MsgCreateEscrow{} }
func (m *MsgCreateEscrow) String() string { return proto.CompactTextString(m) }
func (*MsgCreateEscrow) ProtoMessage()    {}
func (*MsgCreateEscrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{0}
}
func (m *MsgCreateEscrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateEscrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateEscrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateEscrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateEscrow.Merge(m, src)
}
func (m *MsgCreateEscrow) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateEscrow) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateEscrow.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateEscrow proto.InternalMessageInfo

func (m *MsgCreateEscrow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateEscrow) GetInitiatorCoins() []types.Coin {
	if m != nil {
		return m.InitiatorCoins
	}
	return nil
}

func (m *MsgCreateEscrow) GetFulfillerCoins() []types.Coin {
	if m != nil {
		return m.FulfillerCoins
	}
	return nil
}

func (m *MsgCreateEscrow) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *MsgCreateEscrow) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

type MsgCreateEscrowResponse struct {
}

func (m *MsgCreateEscrowResponse) Reset()         { *m = MsgCreateEscrowResponse{} }
func (m *MsgCreateEscrowResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateEscrowResponse) ProtoMessage()    {}
func (*MsgCreateEscrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{1}
}
func (m *MsgCreateEscrowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateEscrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateEscrowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateEscrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateEscrowResponse.Merge(m, src)
}
func (m *MsgCreateEscrowResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateEscrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateEscrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateEscrowResponse proto.InternalMessageInfo

type MsgCancelEscrow struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCancelEscrow) Reset()         { *m = MsgCancelEscrow{} }
func (m *MsgCancelEscrow) String() string { return proto.CompactTextString(m) }
func (*MsgCancelEscrow) ProtoMessage()    {}
func (*MsgCancelEscrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{2}
}
func (m *MsgCancelEscrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelEscrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelEscrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelEscrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelEscrow.Merge(m, src)
}
func (m *MsgCancelEscrow) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelEscrow) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelEscrow.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelEscrow proto.InternalMessageInfo

func (m *MsgCancelEscrow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCancelEscrow) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgCancelEscrowResponse struct {
}

func (m *MsgCancelEscrowResponse) Reset()         { *m = MsgCancelEscrowResponse{} }
func (m *MsgCancelEscrowResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelEscrowResponse) ProtoMessage()    {}
func (*MsgCancelEscrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{3}
}
func (m *MsgCancelEscrowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelEscrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelEscrowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelEscrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelEscrowResponse.Merge(m, src)
}
func (m *MsgCancelEscrowResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelEscrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelEscrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelEscrowResponse proto.InternalMessageInfo

type MsgFulfillEscrow struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgFulfillEscrow) Reset()         { *m = MsgFulfillEscrow{} }
func (m *MsgFulfillEscrow) String() string { return proto.CompactTextString(m) }
func (*MsgFulfillEscrow) ProtoMessage()    {}
func (*MsgFulfillEscrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{4}
}
func (m *MsgFulfillEscrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFulfillEscrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFulfillEscrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFulfillEscrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFulfillEscrow.Merge(m, src)
}
func (m *MsgFulfillEscrow) XXX_Size() int {
	return m.Size()
}
func (m *MsgFulfillEscrow) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFulfillEscrow.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFulfillEscrow proto.InternalMessageInfo

func (m *MsgFulfillEscrow) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgFulfillEscrow) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgFulfillEscrowResponse struct {
}

func (m *MsgFulfillEscrowResponse) Reset()         { *m = MsgFulfillEscrowResponse{} }
func (m *MsgFulfillEscrowResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFulfillEscrowResponse) ProtoMessage()    {}
func (*MsgFulfillEscrowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92b765ed5c8100e9, []int{5}
}
func (m *MsgFulfillEscrowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFulfillEscrowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFulfillEscrowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFulfillEscrowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFulfillEscrowResponse.Merge(m, src)
}
func (m *MsgFulfillEscrowResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFulfillEscrowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFulfillEscrowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFulfillEscrowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateEscrow)(nil), "dreddsecure.escrow.MsgCreateEscrow")
	proto.RegisterType((*MsgCreateEscrowResponse)(nil), "dreddsecure.escrow.MsgCreateEscrowResponse")
	proto.RegisterType((*MsgCancelEscrow)(nil), "dreddsecure.escrow.MsgCancelEscrow")
	proto.RegisterType((*MsgCancelEscrowResponse)(nil), "dreddsecure.escrow.MsgCancelEscrowResponse")
	proto.RegisterType((*MsgFulfillEscrow)(nil), "dreddsecure.escrow.MsgFulfillEscrow")
	proto.RegisterType((*MsgFulfillEscrowResponse)(nil), "dreddsecure.escrow.MsgFulfillEscrowResponse")
}

func init() { proto.RegisterFile("dreddsecure/escrow/tx.proto", fileDescriptor_92b765ed5c8100e9) }

var fileDescriptor_92b765ed5c8100e9 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0xab, 0x40,
	0x14, 0x06, 0xda, 0x7b, 0x6f, 0x3a, 0x57, 0xab, 0x21, 0x26, 0x52, 0x34, 0xd8, 0x54, 0x17, 0x4d,
	0xd4, 0x21, 0xad, 0x71, 0xa5, 0xab, 0xd6, 0x9f, 0x55, 0x37, 0x2c, 0x5d, 0x49, 0xe1, 0x94, 0x4c,
	0x82, 0x4c, 0xc3, 0x4c, 0xb5, 0xbe, 0x85, 0x4f, 0xe2, 0x73, 0x74, 0xd9, 0xa5, 0x2b, 0x63, 0xda,
	0x67, 0x70, 0x6f, 0x60, 0xa0, 0x02, 0xda, 0x58, 0x77, 0x1c, 0xce, 0xf7, 0x73, 0xce, 0x07, 0x07,
	0xed, 0xb8, 0x21, 0xb8, 0x2e, 0x03, 0x67, 0x14, 0x82, 0x09, 0xcc, 0x09, 0xe9, 0x83, 0xc9, 0xc7,
	0x78, 0x18, 0x52, 0x4e, 0x55, 0x35, 0xd3, 0xc4, 0xa2, 0xa9, 0x6f, 0x79, 0xd4, 0xa3, 0x71, 0xdb,
	0x8c, 0x9e, 0x04, 0x52, 0x37, 0x1c, 0xca, 0xee, 0x28, 0x33, 0xfb, 0x36, 0x03, 0xf3, 0xbe, 0xd5,
	0x07, 0x6e, 0xb7, 0x4c, 0x87, 0x92, 0x40, 0xf4, 0x1b, 0xef, 0x32, 0xda, 0xe8, 0x31, 0xaf, 0x1b,
	0x82, 0xcd, 0xe1, 0x32, 0x56, 0x52, 0x35, 0xf4, 0xcf, 0x89, 0x6a, 0x1a, 0x6a, 0x72, 0x5d, 0x6e,
	0x56, 0xac, 0xb4, 0x54, 0xaf, 0x51, 0x95, 0x04, 0x84, 0x93, 0xa8, 0xe8, 0x52, 0x12, 0x30, 0x4d,
	0xa9, 0x97, 0x9a, 0xff, 0xdb, 0x35, 0x2c, 0x6c, 0x70, 0x64, 0x83, 0x13, 0x1b, 0x1c, 0x21, 0x3a,
	0xe5, 0xc9, 0xeb, 0x9e, 0x64, 0x15, 0x68, 0x91, 0xd0, 0x60, 0xe4, 0x0f, 0x88, 0xef, 0x43, 0x22,
	0x54, 0x5a, 0x51, 0x28, 0x4f, 0x53, 0x77, 0x51, 0x85, 0x71, 0x3b, 0xe4, 0x17, 0x36, 0x07, 0xad,
	0x1c, 0x4f, 0xfb, 0xf9, 0x22, 0xda, 0x04, 0x02, 0x37, 0xee, 0xfd, 0x11, 0x9b, 0x24, 0x65, 0xa3,
	0x86, 0xb6, 0x0b, 0x6b, 0x5b, 0xc0, 0x86, 0x34, 0x60, 0xd0, 0x38, 0x13, 0x89, 0xd8, 0x81, 0x03,
	0xfe, 0x8f, 0x89, 0x54, 0x91, 0x42, 0x5c, 0x4d, 0xa9, 0xcb, 0xcd, 0xb2, 0xa5, 0x10, 0x37, 0xd5,
	0xcd, 0x90, 0x17, 0xba, 0xe7, 0x68, 0xb3, 0xc7, 0xbc, 0x2b, 0x31, 0xff, 0xaf, 0x85, 0x75, 0xa4,
	0x15, 0xd9, 0xa9, 0x72, 0xfb, 0x59, 0x41, 0xa5, 0x1e, 0xf3, 0xd4, 0x5b, 0xb4, 0x96, 0xfb, 0x90,
	0xfb, 0xf8, 0xeb, 0x7f, 0x82, 0x0b, 0x6b, 0xeb, 0x87, 0x2b, 0x80, 0x52, 0xa7, 0xd8, 0x21, 0x1b,
	0xcc, 0x52, 0x87, 0x0c, 0x68, 0xb9, 0xc3, 0x37, 0x29, 0xa9, 0x0e, 0x5a, 0xcf, 0x47, 0x74, 0xb0,
	0x84, 0x9d, 0x43, 0xe9, 0x47, 0xab, 0xa0, 0x52, 0x93, 0xce, 0xe9, 0x64, 0x66, 0xc8, 0xd3, 0x99,
	0x21, 0xbf, 0xcd, 0x0c, 0xf9, 0x69, 0x6e, 0x48, 0xd3, 0xb9, 0x21, 0xbd, 0xcc, 0x0d, 0xe9, 0x46,
	0x9c, 0xdd, 0x71, 0x72, 0x77, 0xe3, 0xc5, 0xe5, 0x3d, 0x0e, 0x81, 0xf5, 0xff, 0xc6, 0x37, 0x73,
	0xf2, 0x11, 0x00, 0x00, 0xff, 0xff, 0x84, 0xec, 0x99, 0x35, 0x9c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateEscrow(ctx context.Context, in *MsgCreateEscrow, opts ...grpc.CallOption) (*MsgCreateEscrowResponse, error)
	CancelEscrow(ctx context.Context, in *MsgCancelEscrow, opts ...grpc.CallOption) (*MsgCancelEscrowResponse, error)
	FulfillEscrow(ctx context.Context, in *MsgFulfillEscrow, opts ...grpc.CallOption) (*MsgFulfillEscrowResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateEscrow(ctx context.Context, in *MsgCreateEscrow, opts ...grpc.CallOption) (*MsgCreateEscrowResponse, error) {
	out := new(MsgCreateEscrowResponse)
	err := c.cc.Invoke(ctx, "/dreddsecure.escrow.Msg/CreateEscrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelEscrow(ctx context.Context, in *MsgCancelEscrow, opts ...grpc.CallOption) (*MsgCancelEscrowResponse, error) {
	out := new(MsgCancelEscrowResponse)
	err := c.cc.Invoke(ctx, "/dreddsecure.escrow.Msg/CancelEscrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FulfillEscrow(ctx context.Context, in *MsgFulfillEscrow, opts ...grpc.CallOption) (*MsgFulfillEscrowResponse, error) {
	out := new(MsgFulfillEscrowResponse)
	err := c.cc.Invoke(ctx, "/dreddsecure.escrow.Msg/FulfillEscrow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateEscrow(context.Context, *MsgCreateEscrow) (*MsgCreateEscrowResponse, error)
	CancelEscrow(context.Context, *MsgCancelEscrow) (*MsgCancelEscrowResponse, error)
	FulfillEscrow(context.Context, *MsgFulfillEscrow) (*MsgFulfillEscrowResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateEscrow(ctx context.Context, req *MsgCreateEscrow) (*MsgCreateEscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEscrow not implemented")
}
func (*UnimplementedMsgServer) CancelEscrow(ctx context.Context, req *MsgCancelEscrow) (*MsgCancelEscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelEscrow not implemented")
}
func (*UnimplementedMsgServer) FulfillEscrow(ctx context.Context, req *MsgFulfillEscrow) (*MsgFulfillEscrowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FulfillEscrow not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateEscrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateEscrow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateEscrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dreddsecure.escrow.Msg/CreateEscrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateEscrow(ctx, req.(*MsgCreateEscrow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelEscrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelEscrow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelEscrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dreddsecure.escrow.Msg/CancelEscrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelEscrow(ctx, req.(*MsgCancelEscrow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FulfillEscrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFulfillEscrow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FulfillEscrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dreddsecure.escrow.Msg/FulfillEscrow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FulfillEscrow(ctx, req.(*MsgFulfillEscrow))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dreddsecure.escrow.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEscrow",
			Handler:    _Msg_CreateEscrow_Handler,
		},
		{
			MethodName: "CancelEscrow",
			Handler:    _Msg_CancelEscrow_Handler,
		},
		{
			MethodName: "FulfillEscrow",
			Handler:    _Msg_FulfillEscrow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dreddsecure/escrow/tx.proto",
}

func (m *MsgCreateEscrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateEscrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateEscrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FulfillerCoins) > 0 {
		for iNdEx := len(m.FulfillerCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FulfillerCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.InitiatorCoins) > 0 {
		for iNdEx := len(m.InitiatorCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitiatorCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateEscrowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateEscrowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateEscrowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCancelEscrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelEscrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelEscrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelEscrowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelEscrowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelEscrowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgFulfillEscrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFulfillEscrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFulfillEscrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFulfillEscrowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFulfillEscrowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFulfillEscrowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateEscrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.InitiatorCoins) > 0 {
		for _, e := range m.InitiatorCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.FulfillerCoins) > 0 {
		for _, e := range m.FulfillerCoins {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EndDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateEscrowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCancelEscrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgCancelEscrowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgFulfillEscrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgFulfillEscrowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateEscrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateEscrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateEscrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiatorCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitiatorCoins = append(m.InitiatorCoins, types.Coin{})
			if err := m.InitiatorCoins[len(m.InitiatorCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillerCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FulfillerCoins = append(m.FulfillerCoins, types.Coin{})
			if err := m.FulfillerCoins[len(m.FulfillerCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateEscrowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateEscrowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateEscrowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCancelEscrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCancelEscrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelEscrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCancelEscrowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCancelEscrowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelEscrowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgFulfillEscrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgFulfillEscrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFulfillEscrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgFulfillEscrowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgFulfillEscrowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFulfillEscrowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
