// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dreddsecure/escrow/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	AskCount       uint64                                   `protobuf:"varint,1,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty"`
	MinCount       uint64                                   `protobuf:"varint,2,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty"`
	MinDsCount     uint64                                   `protobuf:"varint,3,opt,name=min_ds_count,json=minDsCount,proto3" json:"min_ds_count,omitempty"`
	PrepareGasBase uint64                                   `protobuf:"varint,4,opt,name=prepare_gas_base,json=prepareGasBase,proto3" json:"prepare_gas_base,omitempty"`
	PrepareGasEach uint64                                   `protobuf:"varint,5,opt,name=prepare_gas_each,json=prepareGasEach,proto3" json:"prepare_gas_each,omitempty"`
	ExecuteGasBase uint64                                   `protobuf:"varint,6,opt,name=execute_gas_base,json=executeGasBase,proto3" json:"execute_gas_base,omitempty"`
	ExecuteGasEach uint64                                   `protobuf:"varint,7,opt,name=execute_gas_each,json=executeGasEach,proto3" json:"execute_gas_each,omitempty"`
	SourceChannel  string                                   `protobuf:"bytes,8,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
	FeeLimit       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,9,rep,name=fee_limit,json=feeLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_limit"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_84997c7643bed4b5, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetAskCount() uint64 {
	if m != nil {
		return m.AskCount
	}
	return 0
}

func (m *Params) GetMinCount() uint64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *Params) GetMinDsCount() uint64 {
	if m != nil {
		return m.MinDsCount
	}
	return 0
}

func (m *Params) GetPrepareGasBase() uint64 {
	if m != nil {
		return m.PrepareGasBase
	}
	return 0
}

func (m *Params) GetPrepareGasEach() uint64 {
	if m != nil {
		return m.PrepareGasEach
	}
	return 0
}

func (m *Params) GetExecuteGasBase() uint64 {
	if m != nil {
		return m.ExecuteGasBase
	}
	return 0
}

func (m *Params) GetExecuteGasEach() uint64 {
	if m != nil {
		return m.ExecuteGasEach
	}
	return 0
}

func (m *Params) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *Params) GetFeeLimit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "dreddsecure.escrow.Params")
}

func init() { proto.RegisterFile("dreddsecure/escrow/params.proto", fileDescriptor_84997c7643bed4b5) }

var fileDescriptor_84997c7643bed4b5 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0x5a, 0x4a, 0x63, 0xe0, 0x84, 0x22, 0x86, 0x70, 0x27, 0x25, 0x11, 0x12, 0x52,
	0x96, 0x8b, 0x39, 0x10, 0x0b, 0x63, 0x0b, 0x62, 0x61, 0x40, 0x1d, 0x59, 0x22, 0xc7, 0x79, 0x97,
	0x58, 0xbd, 0xd8, 0x91, 0x5f, 0x02, 0xc7, 0xb7, 0x60, 0x64, 0x44, 0x62, 0xe3, 0x93, 0x74, 0xec,
	0xc8, 0x04, 0xa8, 0xfd, 0x22, 0xc8, 0x76, 0x90, 0xa2, 0x4e, 0x49, 0xfe, 0xbf, 0x5f, 0xfe, 0xb2,
	0xfc, 0x1e, 0x49, 0x2a, 0x0d, 0x55, 0x85, 0xc0, 0x07, 0x0d, 0x14, 0x90, 0x6b, 0xf5, 0x99, 0x76,
	0x4c, 0xb3, 0x16, 0xf3, 0x4e, 0xab, 0x5e, 0x85, 0xe1, 0x44, 0xc8, 0x9d, 0x70, 0xfe, 0xb8, 0x56,
	0xb5, 0xb2, 0x98, 0x9a, 0x37, 0x67, 0x9e, 0xc7, 0x5c, 0x61, 0xab, 0x90, 0x96, 0x0c, 0x81, 0x7e,
	0xba, 0x2a, 0xa1, 0x67, 0x57, 0x94, 0x2b, 0x21, 0x1d, 0x7f, 0xfa, 0x63, 0x46, 0x16, 0x1f, 0x6c,
	0x75, 0x78, 0x41, 0x02, 0x86, 0xdb, 0x82, 0xab, 0x41, 0xf6, 0x91, 0x9f, 0xfa, 0xd9, 0x7c, 0xb3,
	0x64, 0xb8, 0x5d, 0x9b, 0x6f, 0x03, 0x5b, 0x21, 0x47, 0x78, 0xc7, 0xc1, 0x56, 0x48, 0x07, 0x53,
	0xf2, 0xc0, 0xc0, 0x0a, 0x47, 0x3e, 0xb3, 0x9c, 0xb4, 0x42, 0xbe, 0x41, 0x67, 0x64, 0xe4, 0x51,
	0xa7, 0xa1, 0x63, 0x1a, 0x8a, 0x9a, 0x61, 0x61, 0x4e, 0x13, 0xcd, 0xad, 0x75, 0x36, 0xe6, 0xef,
	0x18, 0xae, 0x18, 0xc2, 0xa9, 0x09, 0x8c, 0x37, 0xd1, 0xdd, 0x53, 0xf3, 0x2d, 0xe3, 0x8d, 0x31,
	0xe1, 0x16, 0xf8, 0xd0, 0x4f, 0x3a, 0x17, 0xce, 0x1c, 0xf3, 0x49, 0xe7, 0xd4, 0xb4, 0x9d, 0xf7,
	0x4e, 0x4d, 0xdb, 0xf9, 0x8c, 0x9c, 0xa1, 0x1a, 0x34, 0x87, 0x82, 0x37, 0x4c, 0x4a, 0xb8, 0x89,
	0x96, 0xa9, 0x9f, 0x05, 0x9b, 0x87, 0x2e, 0x5d, 0xbb, 0x30, 0x6c, 0x48, 0x70, 0x0d, 0x50, 0xdc,
	0x88, 0x56, 0xf4, 0x51, 0x90, 0xce, 0xb2, 0xfb, 0x2f, 0x9e, 0xe4, 0xee, 0xa6, 0x73, 0x73, 0x8e,
	0x7c, 0xbc, 0xe9, 0x7c, 0xad, 0x84, 0x5c, 0x3d, 0xdf, 0xfd, 0x4e, 0xbc, 0x9f, 0x7f, 0x92, 0xac,
	0x16, 0x7d, 0x33, 0x94, 0x39, 0x57, 0x2d, 0x1d, 0xc7, 0xe2, 0x1e, 0x97, 0x58, 0x6d, 0x69, 0xff,
	0xa5, 0x03, 0xb4, 0x3f, 0xe0, 0x66, 0x79, 0x0d, 0xf0, 0xde, 0x94, 0xbf, 0x9e, 0x7f, 0xfb, 0x9e,
	0x78, 0xab, 0x57, 0xbb, 0x43, 0xec, 0xef, 0x0f, 0xb1, 0xff, 0xf7, 0x10, 0xfb, 0x5f, 0x8f, 0xb1,
	0xb7, 0x3f, 0xc6, 0xde, 0xaf, 0x63, 0xec, 0x7d, 0xbc, 0xb0, 0x9b, 0x70, 0x39, 0xee, 0xca, 0xed,
	0xff, 0x6d, 0xb1, 0x65, 0xe5, 0xc2, 0xce, 0xf8, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8,
	0x82, 0xa0, 0x0e, 0x50, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeLimit) > 0 {
		for iNdEx := len(m.FeeLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x42
	}
	if m.ExecuteGasEach != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExecuteGasEach))
		i--
		dAtA[i] = 0x38
	}
	if m.ExecuteGasBase != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExecuteGasBase))
		i--
		dAtA[i] = 0x30
	}
	if m.PrepareGasEach != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PrepareGasEach))
		i--
		dAtA[i] = 0x28
	}
	if m.PrepareGasBase != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PrepareGasBase))
		i--
		dAtA[i] = 0x20
	}
	if m.MinDsCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinDsCount))
		i--
		dAtA[i] = 0x18
	}
	if m.MinCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x10
	}
	if m.AskCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AskCount != 0 {
		n += 1 + sovParams(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovParams(uint64(m.MinCount))
	}
	if m.MinDsCount != 0 {
		n += 1 + sovParams(uint64(m.MinDsCount))
	}
	if m.PrepareGasBase != 0 {
		n += 1 + sovParams(uint64(m.PrepareGasBase))
	}
	if m.PrepareGasEach != 0 {
		n += 1 + sovParams(uint64(m.PrepareGasEach))
	}
	if m.ExecuteGasBase != 0 {
		n += 1 + sovParams(uint64(m.ExecuteGasBase))
	}
	if m.ExecuteGasEach != 0 {
		n += 1 + sovParams(uint64(m.ExecuteGasEach))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.FeeLimit) > 0 {
		for _, e := range m.FeeLimit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDsCount", wireType)
			}
			m.MinDsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinDsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGasBase", wireType)
			}
			m.PrepareGasBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGasBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGasEach", wireType)
			}
			m.PrepareGasEach = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGasEach |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGasBase", wireType)
			}
			m.ExecuteGasBase = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGasBase |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGasEach", wireType)
			}
			m.ExecuteGasEach = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGasEach |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeLimit = append(m.FeeLimit, types.Coin{})
			if err := m.FeeLimit[len(m.FeeLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
