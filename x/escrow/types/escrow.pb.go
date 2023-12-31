// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dreddsecure/escrow/escrow.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Escrow struct {
	Id               uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status           string       `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Initiator        string       `protobuf:"bytes,3,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Fulfiller        string       `protobuf:"bytes,4,opt,name=fulfiller,proto3" json:"fulfiller,omitempty"`
	InitiatorCoins   []types.Coin `protobuf:"bytes,5,rep,name=initiatorCoins,proto3" json:"initiatorCoins"`
	FulfillerCoins   []types.Coin `protobuf:"bytes,6,rep,name=fulfillerCoins,proto3" json:"fulfillerCoins"`
	Tips             []types.Coin `protobuf:"bytes,7,rep,name=tips,proto3" json:"tips"`
	StartDate        string       `protobuf:"bytes,8,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate          string       `protobuf:"bytes,9,opt,name=endDate,proto3" json:"endDate,omitempty"`
	OracleConditions string       `protobuf:"bytes,10,opt,name=oracleConditions,proto3" json:"oracleConditions,omitempty"`
}

func (m *Escrow) Reset()         { *m = Escrow{} }
func (m *Escrow) String() string { return proto.CompactTextString(m) }
func (*Escrow) ProtoMessage()    {}
func (*Escrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_1399909c3997950a, []int{0}
}
func (m *Escrow) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Escrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Escrow.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Escrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Escrow.Merge(m, src)
}
func (m *Escrow) XXX_Size() int {
	return m.Size()
}
func (m *Escrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Escrow.DiscardUnknown(m)
}

var xxx_messageInfo_Escrow proto.InternalMessageInfo

func (m *Escrow) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Escrow) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Escrow) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Escrow) GetFulfiller() string {
	if m != nil {
		return m.Fulfiller
	}
	return ""
}

func (m *Escrow) GetInitiatorCoins() []types.Coin {
	if m != nil {
		return m.InitiatorCoins
	}
	return nil
}

func (m *Escrow) GetFulfillerCoins() []types.Coin {
	if m != nil {
		return m.FulfillerCoins
	}
	return nil
}

func (m *Escrow) GetTips() []types.Coin {
	if m != nil {
		return m.Tips
	}
	return nil
}

func (m *Escrow) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Escrow) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Escrow) GetOracleConditions() string {
	if m != nil {
		return m.OracleConditions
	}
	return ""
}

func init() {
	proto.RegisterType((*Escrow)(nil), "dreddsecure.escrow.Escrow")
}

func init() { proto.RegisterFile("dreddsecure/escrow/escrow.proto", fileDescriptor_1399909c3997950a) }

var fileDescriptor_1399909c3997950a = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x77, 0x61, 0x5d, 0xa4, 0x26, 0xc4, 0x34, 0xc6, 0x54, 0x34, 0x85, 0x78, 0x22, 0x26,
	0xee, 0x06, 0x89, 0x2f, 0x00, 0x1a, 0xef, 0x1c, 0xbd, 0x95, 0x6d, 0x21, 0x4d, 0xd6, 0x1d, 0xd2,
	0x16, 0xff, 0xbc, 0x85, 0x37, 0x5f, 0x89, 0x23, 0x47, 0x4f, 0xc6, 0xc0, 0x8b, 0x98, 0xb6, 0xcb,
	0xfa, 0xef, 0xc2, 0x69, 0xb7, 0xdf, 0xef, 0x9b, 0x2f, 0x33, 0x99, 0x41, 0x1d, 0xae, 0x04, 0xe7,
	0x5a, 0x64, 0x0b, 0x25, 0x52, 0xa1, 0x33, 0x05, 0x4f, 0xe5, 0x27, 0x99, 0x2b, 0x30, 0x80, 0xf1,
	0x0f, 0x43, 0xe2, 0x49, 0xfb, 0x68, 0x06, 0x33, 0x70, 0x38, 0xb5, 0x7f, 0xde, 0xd9, 0xa6, 0x19,
	0xe8, 0x07, 0xd0, 0xe9, 0x84, 0x69, 0x91, 0x3e, 0xf6, 0x27, 0xc2, 0xb0, 0x7e, 0x9a, 0x81, 0x2c,
	0x3c, 0x3f, 0x7f, 0xab, 0xa3, 0xf8, 0xd6, 0x05, 0xe0, 0x16, 0xaa, 0x49, 0x4e, 0xc2, 0x6e, 0xd8,
	0x8b, 0xc6, 0x35, 0xc9, 0xf1, 0x31, 0x8a, 0xb5, 0x61, 0x66, 0xa1, 0x49, 0xad, 0x1b, 0xf6, 0x9a,
	0xe3, 0xf2, 0x85, 0xcf, 0x50, 0x53, 0x16, 0xd2, 0x48, 0x66, 0x40, 0x91, 0xba, 0x43, 0xdf, 0x82,
	0xa5, 0xd3, 0x45, 0x3e, 0x95, 0x79, 0x2e, 0x14, 0x89, 0x3c, 0xad, 0x04, 0x7c, 0x87, 0x5a, 0x95,
	0x75, 0x04, 0xb2, 0xd0, 0x64, 0xaf, 0x5b, 0xef, 0x1d, 0x5c, 0x9d, 0x24, 0xbe, 0xcf, 0xc4, 0xf6,
	0x99, 0x94, 0x7d, 0x26, 0xd6, 0x31, 0x8c, 0x96, 0x1f, 0x9d, 0x60, 0xfc, 0xa7, 0xcc, 0x06, 0x55,
	0xa9, 0x3e, 0x28, 0xde, 0x31, 0xe8, 0x77, 0x19, 0x1e, 0xa0, 0xc8, 0xc8, 0xb9, 0x26, 0x8d, 0xdd,
	0xca, 0x9d, 0xd9, 0x0e, 0xa9, 0x0d, 0x53, 0xe6, 0x86, 0x19, 0x41, 0xf6, 0xfd, 0x90, 0x95, 0x80,
	0x09, 0x6a, 0x88, 0x82, 0x3b, 0xd6, 0x74, 0x6c, 0xfb, 0xc4, 0x17, 0xe8, 0x10, 0x14, 0xcb, 0x72,
	0x31, 0x82, 0x82, 0x4b, 0x23, 0xa1, 0xd0, 0x04, 0x39, 0xcb, 0x3f, 0x7d, 0x78, 0xbd, 0x5c, 0xd3,
	0x70, 0xb5, 0xa6, 0xe1, 0xe7, 0x9a, 0x86, 0xaf, 0x1b, 0x1a, 0xac, 0x36, 0x34, 0x78, 0xdf, 0xd0,
	0xe0, 0xfe, 0xd4, 0x6d, 0xff, 0xb2, 0xbc, 0x8f, 0xe7, 0xed, 0x85, 0x98, 0x97, 0xb9, 0xd0, 0x93,
	0xd8, 0xed, 0x75, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0x64, 0x1f, 0xc7, 0xa9, 0x44, 0x02, 0x00,
	0x00,
}

func (m *Escrow) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Escrow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Escrow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OracleConditions) > 0 {
		i -= len(m.OracleConditions)
		copy(dAtA[i:], m.OracleConditions)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.OracleConditions)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Tips) > 0 {
		for iNdEx := len(m.Tips) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tips[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEscrow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.FulfillerCoins) > 0 {
		for iNdEx := len(m.FulfillerCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FulfillerCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEscrow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
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
				i = encodeVarintEscrow(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Fulfiller) > 0 {
		i -= len(m.Fulfiller)
		copy(dAtA[i:], m.Fulfiller)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.Fulfiller)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Initiator) > 0 {
		i -= len(m.Initiator)
		copy(dAtA[i:], m.Initiator)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.Initiator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintEscrow(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEscrow(dAtA []byte, offset int, v uint64) int {
	offset -= sovEscrow(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Escrow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovEscrow(uint64(m.Id))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.Initiator)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.Fulfiller)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	if len(m.InitiatorCoins) > 0 {
		for _, e := range m.InitiatorCoins {
			l = e.Size()
			n += 1 + l + sovEscrow(uint64(l))
		}
	}
	if len(m.FulfillerCoins) > 0 {
		for _, e := range m.FulfillerCoins {
			l = e.Size()
			n += 1 + l + sovEscrow(uint64(l))
		}
	}
	if len(m.Tips) > 0 {
		for _, e := range m.Tips {
			l = e.Size()
			n += 1 + l + sovEscrow(uint64(l))
		}
	}
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.EndDate)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.OracleConditions)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	return n
}

func sovEscrow(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEscrow(x uint64) (n int) {
	return sovEscrow(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Escrow) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEscrow
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
			return fmt.Errorf("proto: Escrow: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Escrow: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Initiator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Initiator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fulfiller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fulfiller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiatorCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitiatorCoins = append(m.InitiatorCoins, types.Coin{})
			if err := m.InitiatorCoins[len(m.InitiatorCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillerCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FulfillerCoins = append(m.FulfillerCoins, types.Coin{})
			if err := m.FulfillerCoins[len(m.FulfillerCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tips", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tips = append(m.Tips, types.Coin{})
			if err := m.Tips[len(m.Tips)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleConditions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEscrow
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
				return ErrInvalidLengthEscrow
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEscrow
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OracleConditions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEscrow(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEscrow
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
func skipEscrow(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEscrow
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
					return 0, ErrIntOverflowEscrow
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
					return 0, ErrIntOverflowEscrow
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
				return 0, ErrInvalidLengthEscrow
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEscrow
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEscrow
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEscrow        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEscrow          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEscrow = fmt.Errorf("proto: unexpected end of group")
)
