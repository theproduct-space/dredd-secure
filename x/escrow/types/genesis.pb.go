// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dreddsecure/escrow/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the escrow module's genesis state.
type GenesisState struct {
	Params          Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId          string            `protobuf:"bytes,7,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	EscrowList      []Escrow          `protobuf:"bytes,2,rep,name=escrowList,proto3" json:"escrowList"`
	EscrowCount     uint64            `protobuf:"varint,3,opt,name=escrowCount,proto3" json:"escrowCount,omitempty"`
	PendingEscrows  []uint64          `protobuf:"varint,4,rep,packed,name=pendingEscrows,proto3" json:"pendingEscrows,omitempty"`
	ExpiringEscrows []uint64          `protobuf:"varint,5,rep,packed,name=expiringEscrows,proto3" json:"expiringEscrows,omitempty"`
	LastExecs       map[string]string `protobuf:"bytes,6,rep,name=lastExecs,proto3" json:"lastExecs" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OraclePriceList []OraclePrice     `protobuf:"bytes,8,rep,name=oraclePriceList,proto3" json:"oraclePriceList"`
	SourceChannel   string            `protobuf:"bytes,9,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3772c078e9194019, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetEscrowList() []Escrow {
	if m != nil {
		return m.EscrowList
	}
	return nil
}

func (m *GenesisState) GetEscrowCount() uint64 {
	if m != nil {
		return m.EscrowCount
	}
	return 0
}

func (m *GenesisState) GetPendingEscrows() []uint64 {
	if m != nil {
		return m.PendingEscrows
	}
	return nil
}

func (m *GenesisState) GetExpiringEscrows() []uint64 {
	if m != nil {
		return m.ExpiringEscrows
	}
	return nil
}

func (m *GenesisState) GetLastExecs() map[string]string {
	if m != nil {
		return m.LastExecs
	}
	return nil
}

func (m *GenesisState) GetOraclePriceList() []OraclePrice {
	if m != nil {
		return m.OraclePriceList
	}
	return nil
}

func (m *GenesisState) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "dreddsecure.escrow.GenesisState")
	proto.RegisterMapType((map[string]string)(nil), "dreddsecure.escrow.GenesisState.LastExecsEntry")
}

func init() { proto.RegisterFile("dreddsecure/escrow/genesis.proto", fileDescriptor_3772c078e9194019) }

var fileDescriptor_3772c078e9194019 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x4d, 0x36, 0x6b, 0x66, 0xb5, 0x2b, 0xc3, 0x82, 0x43, 0x84, 0x34, 0x08, 0x0b,
	0x39, 0x68, 0x02, 0x8a, 0xb0, 0x88, 0x07, 0xd9, 0xa5, 0x88, 0xb0, 0xb0, 0x25, 0xbd, 0x79, 0x29,
	0x31, 0x79, 0xc4, 0x60, 0xcc, 0x84, 0x99, 0x89, 0xb6, 0xdf, 0xc2, 0x4f, 0xe2, 0xe7, 0xe8, 0xb1,
	0x47, 0x4f, 0x22, 0xed, 0x17, 0x91, 0xcc, 0xa4, 0x34, 0xd6, 0xe8, 0x29, 0xf3, 0xde, 0xfb, 0xbd,
	0x7f, 0xfe, 0x33, 0xef, 0x61, 0x3f, 0xe3, 0x90, 0x65, 0x02, 0xd2, 0x86, 0x43, 0x04, 0x22, 0xe5,
	0xec, 0x6b, 0x94, 0x43, 0x05, 0xa2, 0x10, 0x61, 0xcd, 0x99, 0x64, 0x84, 0xf4, 0x88, 0x50, 0x13,
	0xee, 0x45, 0xce, 0x72, 0xa6, 0xca, 0x51, 0x7b, 0xd2, 0xa4, 0x3b, 0x19, 0xd0, 0xaa, 0x13, 0x9e,
	0x7c, 0x16, 0xff, 0x01, 0xf4, 0xa7, 0x03, 0x2e, 0x07, 0x00, 0xc6, 0x93, 0xb4, 0x84, 0x45, 0xcd,
	0x8b, 0x14, 0x34, 0xf6, 0xe4, 0xbb, 0x85, 0xef, 0xbf, 0xd5, 0x26, 0xe7, 0x32, 0x91, 0x40, 0xae,
	0xb0, 0xad, 0x7f, 0x44, 0x91, 0x8f, 0x82, 0xb3, 0xe7, 0x6e, 0xf8, 0xb7, 0xe9, 0x70, 0xa6, 0x88,
	0x6b, 0x6b, 0xfd, 0x73, 0x62, 0xc4, 0x1d, 0x4f, 0x1e, 0xe1, 0xd3, 0x9a, 0x71, 0xb9, 0x28, 0x32,
	0x7a, 0xea, 0xa3, 0xc0, 0x89, 0xed, 0x36, 0x7c, 0x97, 0x91, 0x37, 0x18, 0xeb, 0xbe, 0xdb, 0x42,
	0x48, 0x3a, 0xf2, 0xcd, 0x7f, 0xc9, 0x4e, 0xd5, 0xa7, 0x93, 0xed, 0xf5, 0x10, 0x1f, 0x9f, 0xe9,
	0xe8, 0x86, 0x35, 0x95, 0xa4, 0xa6, 0x8f, 0x02, 0x2b, 0xee, 0xa7, 0xc8, 0x53, 0x3c, 0xae, 0xa1,
	0xca, 0x8a, 0x2a, 0xd7, 0x22, 0x82, 0x5a, 0xbe, 0x19, 0x58, 0x9d, 0xd6, 0x51, 0x8d, 0x84, 0xf8,
	0x1c, 0x96, 0x75, 0xc1, 0x7b, 0xf8, 0x49, 0x0f, 0x3f, 0x2e, 0x92, 0x39, 0x76, 0xca, 0x44, 0xc8,
	0xe9, 0x12, 0x52, 0x41, 0x6d, 0x75, 0x81, 0x68, 0xe8, 0x02, 0xfd, 0x97, 0x0c, 0x6f, 0xf7, 0x1d,
	0xd3, 0x4a, 0xf2, 0x55, 0x27, 0x7d, 0xd0, 0x21, 0x77, 0xf8, 0x5c, 0x0f, 0x64, 0xd6, 0xce, 0x43,
	0xbd, 0xcd, 0x3d, 0x25, 0x3d, 0x19, 0x92, 0xbe, 0x3b, 0xa0, 0x7b, 0x97, 0x47, 0xdd, 0xe4, 0x12,
	0x8f, 0x05, 0x6b, 0x78, 0x0a, 0x8b, 0xf4, 0x63, 0x52, 0x55, 0x50, 0x52, 0x47, 0xcd, 0xe1, 0x81,
	0xce, 0xde, 0xe8, 0xa4, 0xfb, 0x1a, 0x8f, 0xff, 0xb4, 0x46, 0x1e, 0x62, 0xf3, 0x13, 0xac, 0xd4,
	0xc0, 0x9d, 0xb8, 0x3d, 0x92, 0x0b, 0x7c, 0xf2, 0x25, 0x29, 0x1b, 0xa0, 0x23, 0x95, 0xd3, 0xc1,
	0xab, 0xd1, 0x15, 0xba, 0x7e, 0xb9, 0xde, 0x7a, 0x68, 0xb3, 0xf5, 0xd0, 0xaf, 0xad, 0x87, 0xbe,
	0xed, 0x3c, 0x63, 0xb3, 0xf3, 0x8c, 0x1f, 0x3b, 0xcf, 0x78, 0xff, 0x58, 0xb9, 0x7e, 0xd6, 0xad,
	0xdc, 0x72, 0xbf, 0x74, 0x72, 0x55, 0x83, 0xf8, 0x60, 0xab, 0x75, 0x7b, 0xf1, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0xfb, 0xe9, 0x74, 0x25, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.OraclePriceList) > 0 {
		for iNdEx := len(m.OraclePriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OraclePriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastExecs) > 0 {
		for k := range m.LastExecs {
			v := m.LastExecs[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGenesis(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ExpiringEscrows) > 0 {
		dAtA2 := make([]byte, len(m.ExpiringEscrows)*10)
		var j1 int
		for _, num := range m.ExpiringEscrows {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PendingEscrows) > 0 {
		dAtA4 := make([]byte, len(m.PendingEscrows)*10)
		var j3 int
		for _, num := range m.PendingEscrows {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGenesis(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x22
	}
	if m.EscrowCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EscrowCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EscrowList) > 0 {
		for iNdEx := len(m.EscrowList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EscrowList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.EscrowList) > 0 {
		for _, e := range m.EscrowList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EscrowCount != 0 {
		n += 1 + sovGenesis(uint64(m.EscrowCount))
	}
	if len(m.PendingEscrows) > 0 {
		l = 0
		for _, e := range m.PendingEscrows {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	if len(m.ExpiringEscrows) > 0 {
		l = 0
		for _, e := range m.ExpiringEscrows {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	if len(m.LastExecs) > 0 {
		for k, v := range m.LastExecs {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + len(v) + sovGenesis(uint64(len(v)))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.OraclePriceList) > 0 {
		for _, e := range m.OraclePriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowList = append(m.EscrowList, Escrow{})
			if err := m.EscrowList[len(m.EscrowList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowCount", wireType)
			}
			m.EscrowCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EscrowCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PendingEscrows = append(m.PendingEscrows, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PendingEscrows) == 0 {
					m.PendingEscrows = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PendingEscrows = append(m.PendingEscrows, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingEscrows", wireType)
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExpiringEscrows = append(m.ExpiringEscrows, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ExpiringEscrows) == 0 {
					m.ExpiringEscrows = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExpiringEscrows = append(m.ExpiringEscrows, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiringEscrows", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastExecs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastExecs == nil {
				m.LastExecs = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.LastExecs[mapkey] = mapvalue
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OraclePriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OraclePriceList = append(m.OraclePriceList, OraclePrice{})
			if err := m.OraclePriceList[len(m.OraclePriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
