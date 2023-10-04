// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_identity_aggregated_transfer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
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

type IdentityAggregatedTransfer struct {
	// Hex 0x
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	// Rarimo
	Chain string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	// Hex 0x
	GISTHash      string `protobuf:"bytes,3,opt,name=GISTHash,proto3" json:"GISTHash,omitempty"`
	StateRootHash string `protobuf:"bytes,5,opt,name=stateRootHash,proto3" json:"stateRootHash,omitempty"`
	// Dec
	Timestamp uint64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *IdentityAggregatedTransfer) Reset()         { *m = IdentityAggregatedTransfer{} }
func (m *IdentityAggregatedTransfer) String() string { return proto.CompactTextString(m) }
func (*IdentityAggregatedTransfer) ProtoMessage()    {}
func (*IdentityAggregatedTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e61feb9c5fbd8544, []int{0}
}
func (m *IdentityAggregatedTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentityAggregatedTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentityAggregatedTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentityAggregatedTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityAggregatedTransfer.Merge(m, src)
}
func (m *IdentityAggregatedTransfer) XXX_Size() int {
	return m.Size()
}
func (m *IdentityAggregatedTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityAggregatedTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityAggregatedTransfer proto.InternalMessageInfo

func (m *IdentityAggregatedTransfer) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *IdentityAggregatedTransfer) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *IdentityAggregatedTransfer) GetGISTHash() string {
	if m != nil {
		return m.GISTHash
	}
	return ""
}

func (m *IdentityAggregatedTransfer) GetStateRootHash() string {
	if m != nil {
		return m.StateRootHash
	}
	return ""
}

func (m *IdentityAggregatedTransfer) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentityAggregatedTransfer)(nil), "rarimo.rarimocore.rarimocore.IdentityAggregatedTransfer")
}

func init() {
	proto.RegisterFile("rarimocore/op_identity_aggregated_transfer.proto", fileDescriptor_e61feb9c5fbd8544)
}

var fileDescriptor_e61feb9c5fbd8544 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xff, 0x6f, 0x11, 0xb5, 0xc4, 0x12, 0x55, 0x22, 0x8a, 0x2a, 0xab, 0x42, 0x0c,
	0x5d, 0x48, 0x2a, 0xf1, 0x04, 0xb0, 0x40, 0x19, 0x43, 0x27, 0x96, 0xe8, 0x36, 0x35, 0xae, 0x05,
	0xf6, 0x8d, 0x9c, 0x3b, 0xd0, 0xb7, 0xe0, 0x49, 0x78, 0x0e, 0xc6, 0x8e, 0x8c, 0x28, 0x79, 0x11,
	0x54, 0x3b, 0xb4, 0x65, 0xf2, 0x39, 0xc7, 0xf7, 0xd3, 0x91, 0x0e, 0x9f, 0x39, 0x70, 0xda, 0x60,
	0x89, 0x4e, 0x66, 0x58, 0x15, 0x7a, 0x25, 0x2d, 0x69, 0xda, 0x14, 0xa0, 0x94, 0x93, 0x0a, 0x48,
	0xae, 0x0a, 0x72, 0x60, 0xeb, 0x67, 0xe9, 0xd2, 0xca, 0x21, 0x61, 0x34, 0x0e, 0x44, 0x7a, 0x00,
	0x8f, 0x64, 0x72, 0x4e, 0xf8, 0x22, 0xad, 0x01, 0x0b, 0x4a, 0xba, 0x4c, 0x93, 0x34, 0x01, 0x4b,
	0x46, 0x0a, 0x15, 0x7a, 0x99, 0xed, 0x54, 0x48, 0x2f, 0x3e, 0x18, 0x4f, 0xe6, 0x5d, 0xe7, 0xcd,
	0xbe, 0x72, 0xd1, 0x35, 0x46, 0x09, 0x3f, 0x2d, 0xd1, 0x92, 0x83, 0x92, 0x62, 0x36, 0x61, 0xd3,
	0x61, 0xbe, 0xf7, 0xd1, 0x88, 0x0f, 0xca, 0x35, 0x68, 0x1b, 0xff, 0xf3, 0x1f, 0xc1, 0xec, 0x88,
	0xbb, 0xf9, 0xe3, 0xe2, 0x1e, 0xea, 0x75, 0xfc, 0x3f, 0x10, 0xbf, 0x3e, 0xba, 0xe4, 0x67, 0x35,
	0x01, 0xc9, 0x1c, 0x91, 0xfc, 0xc1, 0xc0, 0x1f, 0xfc, 0x0d, 0xa3, 0x31, 0x1f, 0x92, 0x36, 0xb2,
	0x26, 0x30, 0x55, 0xdc, 0x9f, 0xb0, 0x69, 0x3f, 0x3f, 0x04, 0xb7, 0x0f, 0x9f, 0x8d, 0x60, 0xdb,
	0x46, 0xb0, 0xef, 0x46, 0xb0, 0xf7, 0x56, 0xf4, 0xb6, 0xad, 0xe8, 0x7d, 0xb5, 0xa2, 0xf7, 0x34,
	0x53, 0x9a, 0x5e, 0x61, 0x99, 0x96, 0x68, 0xb2, 0x30, 0x48, 0xf7, 0x5c, 0xf9, 0x71, 0xdf, 0xb2,
	0xa3, 0xa5, 0x69, 0x53, 0xc9, 0x7a, 0x79, 0xe2, 0x37, 0xb8, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x3e, 0x33, 0x25, 0x36, 0x84, 0x01, 0x00, 0x00,
}

func (m *IdentityAggregatedTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentityAggregatedTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentityAggregatedTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateRootHash) > 0 {
		i -= len(m.StateRootHash)
		copy(dAtA[i:], m.StateRootHash)
		i = encodeVarintOpIdentityAggregatedTransfer(dAtA, i, uint64(len(m.StateRootHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Timestamp != 0 {
		i = encodeVarintOpIdentityAggregatedTransfer(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.GISTHash) > 0 {
		i -= len(m.GISTHash)
		copy(dAtA[i:], m.GISTHash)
		i = encodeVarintOpIdentityAggregatedTransfer(dAtA, i, uint64(len(m.GISTHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOpIdentityAggregatedTransfer(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintOpIdentityAggregatedTransfer(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpIdentityAggregatedTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpIdentityAggregatedTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdentityAggregatedTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovOpIdentityAggregatedTransfer(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOpIdentityAggregatedTransfer(uint64(l))
	}
	l = len(m.GISTHash)
	if l > 0 {
		n += 1 + l + sovOpIdentityAggregatedTransfer(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovOpIdentityAggregatedTransfer(uint64(m.Timestamp))
	}
	l = len(m.StateRootHash)
	if l > 0 {
		n += 1 + l + sovOpIdentityAggregatedTransfer(uint64(l))
	}
	return n
}

func sovOpIdentityAggregatedTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpIdentityAggregatedTransfer(x uint64) (n int) {
	return sovOpIdentityAggregatedTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdentityAggregatedTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpIdentityAggregatedTransfer
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
			return fmt.Errorf("proto: IdentityAggregatedTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentityAggregatedTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityAggregatedTransfer
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
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityAggregatedTransfer
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
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GISTHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityAggregatedTransfer
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
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GISTHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityAggregatedTransfer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRootHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityAggregatedTransfer
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
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRootHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpIdentityAggregatedTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpIdentityAggregatedTransfer
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
func skipOpIdentityAggregatedTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpIdentityAggregatedTransfer
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
					return 0, ErrIntOverflowOpIdentityAggregatedTransfer
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
					return 0, ErrIntOverflowOpIdentityAggregatedTransfer
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
				return 0, ErrInvalidLengthOpIdentityAggregatedTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpIdentityAggregatedTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpIdentityAggregatedTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpIdentityAggregatedTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpIdentityAggregatedTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpIdentityAggregatedTransfer = fmt.Errorf("proto: unexpected end of group")
)