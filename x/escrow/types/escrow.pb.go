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
	Id                uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status            string     `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Initiator         string     `protobuf:"bytes,3,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Fulfiller         string     `protobuf:"bytes,4,opt,name=fulfiller,proto3" json:"fulfiller,omitempty"`
	InitiatorCoin     string     `protobuf:"bytes,3,opt,name=initiatorCoin,proto3" json:"initiatorCoin,omitempty"`
	FulfillerCoin     string     `protobuf:"bytes,4,opt,name=fulfillerCoin,proto3" json:"fulfillerCoin,omitempty"`
	StartDate         string     `protobuf:"bytes,7,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate           string     `protobuf:"bytes,8,opt,name=endDate,proto3" json:"endDate,omitempty"`
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

func (m *Escrow) GetInitiatorCoin() string {
	if m != nil {
		return m.InitiatorCoin
	}
	return ""
}

func (m *Escrow) GetFulfillerCoin() string {
	if m != nil {
		return m.FulfillerCoin
	}
	return ""
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

func init() {
	proto.RegisterType((*Escrow)(nil), "dreddsecure.escrow.Escrow")
}

func init() { proto.RegisterFile("dreddsecure/escrow/escrow.proto", fileDescriptor_1399909c3997950a) }

var fileDescriptor_1399909c3997950a = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xe3, 0x7c, 0xfd, 0x52, 0x6a, 0x04, 0x83, 0x85, 0x90, 0x29, 0xc8, 0xad, 0x98, 0xba,
	0xe0, 0xa8, 0x20, 0x5e, 0xa0, 0xd0, 0x17, 0xe8, 0xc8, 0xe6, 0xc4, 0x6e, 0x65, 0x29, 0xc4, 0x95,
	0xed, 0xf0, 0xe7, 0x2d, 0x78, 0x16, 0x9e, 0xa2, 0x63, 0x47, 0x26, 0x84, 0x92, 0x17, 0x41, 0xb6,
	0xd3, 0x40, 0x37, 0xa6, 0xe4, 0xde, 0xdf, 0xb9, 0x47, 0x47, 0x3e, 0x70, 0xc4, 0xb5, 0xe0, 0xdc,
	0x88, 0xbc, 0xd2, 0x22, 0x15, 0x26, 0xd7, 0xea, 0xb9, 0xfd, 0xd0, 0xb5, 0x56, 0x56, 0x21, 0xf4,
	0x4b, 0x40, 0x03, 0x19, 0x9e, 0xac, 0xd4, 0x4a, 0x79, 0x9c, 0xba, 0xbf, 0xa0, 0x1c, 0x92, 0x5c,
	0x99, 0x47, 0x65, 0xd2, 0x8c, 0x19, 0x91, 0x3e, 0x4d, 0x33, 0x61, 0xd9, 0x34, 0xcd, 0x95, 0x2c,
	0x03, 0xbf, 0x7c, 0x8f, 0x61, 0x32, 0xf7, 0x06, 0xe8, 0x18, 0xc6, 0x92, 0x63, 0x30, 0x06, 0x93,
	0xde, 0x22, 0x96, 0x1c, 0x9d, 0xc2, 0xc4, 0x58, 0x66, 0x2b, 0x83, 0xe3, 0x31, 0x98, 0x0c, 0x16,
	0xed, 0x84, 0x2e, 0xe0, 0x40, 0x96, 0xd2, 0x4a, 0x66, 0x95, 0xc6, 0xff, 0x3c, 0xfa, 0x59, 0x38,
	0xba, 0xac, 0x8a, 0xa5, 0x2c, 0x0a, 0xa1, 0x71, 0x2f, 0xd0, 0x6e, 0x81, 0xe6, 0xf0, 0xa8, 0x93,
	0xde, 0x29, 0x59, 0xe2, 0xff, 0x63, 0x30, 0x39, 0xbc, 0x3e, 0xa3, 0x21, 0x26, 0x75, 0x31, 0x69,
	0x1b, 0x93, 0x3a, 0xc1, 0xac, 0xb7, 0xf9, 0x1c, 0x45, 0x8b, 0xfd, 0x2b, 0x67, 0xd3, 0x79, 0x7a,
	0x9b, 0xe4, 0x8f, 0x36, 0x7b, 0x57, 0x2e, 0xab, 0xb1, 0x4c, 0xdb, 0x7b, 0x66, 0x05, 0xee, 0x87,
	0xac, 0xdd, 0x02, 0x61, 0xd8, 0x17, 0x25, 0xf7, 0xec, 0xc0, 0xb3, 0xdd, 0x38, 0xbb, 0xdd, 0xd4,
	0x04, 0x6c, 0x6b, 0x02, 0xbe, 0x6a, 0x02, 0xde, 0x1a, 0x12, 0x6d, 0x1b, 0x12, 0x7d, 0x34, 0x24,
	0x7a, 0x38, 0xf7, 0xc5, 0x5c, 0xb5, 0xd5, 0xbd, 0xec, 0xca, 0xb3, 0xaf, 0x6b, 0x61, 0xb2, 0xc4,
	0x3f, 0xf9, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x4f, 0x61, 0x77, 0xdf, 0x01, 0x00,
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
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.FulfillerCoin) > 0 {
		i -= len(m.FulfillerCoin)
		copy(dAtA[i:], m.FulfillerCoin)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.FulfillerCoin)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.InitiatorCoin) > 0 {
		i -= len(m.InitiatorCoin)
		copy(dAtA[i:], m.InitiatorCoin)
		i = encodeVarintEscrow(dAtA, i, uint64(len(m.InitiatorCoin)))
		i--
		dAtA[i] = 0x3a
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
	l = len(m.InitiatorCoin)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.FulfillerCoin)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovEscrow(uint64(l))
	}
	l = len(m.EndDate)
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
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillerCoin", wireType)
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
			m.FulfillerCoin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitiatorCoin", wireType)
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
			m.InitiatorCoin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
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
		case 8:
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
