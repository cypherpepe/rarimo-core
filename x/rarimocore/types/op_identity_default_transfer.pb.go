// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/op_identity_default_transfer.proto

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

type IdentityDefaultTransfer struct {
	// Hex 0x
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Chain    string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	// Hex 0x
	GISTHash string `protobuf:"bytes,3,opt,name=GISTHash,proto3" json:"GISTHash,omitempty"`
	// Hex 0x
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// Hex 0x
	StateHash string `protobuf:"bytes,5,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	// Dec
	StateCreatedAtTimestamp string `protobuf:"bytes,6,opt,name=stateCreatedAtTimestamp,proto3" json:"stateCreatedAtTimestamp,omitempty"`
	StateCreatedAtBlock     string `protobuf:"bytes,7,opt,name=stateCreatedAtBlock,proto3" json:"stateCreatedAtBlock,omitempty"`
	// Hex 0x
	StateReplacedBy string `protobuf:"bytes,10,opt,name=stateReplacedBy,proto3" json:"stateReplacedBy,omitempty"`
	GISTReplacedBy  string `protobuf:"bytes,11,opt,name=GISTReplacedBy,proto3" json:"GISTReplacedBy,omitempty"`
	// Dec
	GISTCreatedAtTimestamp string `protobuf:"bytes,12,opt,name=GISTCreatedAtTimestamp,proto3" json:"GISTCreatedAtTimestamp,omitempty"`
	GISTCreatedAtBlock     string `protobuf:"bytes,13,opt,name=GISTCreatedAtBlock,proto3" json:"GISTCreatedAtBlock,omitempty"`
	// HEX 0x
	ReplacedStateHash string `protobuf:"bytes,16,opt,name=replacedStateHash,proto3" json:"replacedStateHash,omitempty"`
	ReplacedGISTHash  string `protobuf:"bytes,17,opt,name=replacedGISTHash,proto3" json:"replacedGISTHash,omitempty"`
}

func (m *IdentityDefaultTransfer) Reset()         { *m = IdentityDefaultTransfer{} }
func (m *IdentityDefaultTransfer) String() string { return proto.CompactTextString(m) }
func (*IdentityDefaultTransfer) ProtoMessage()    {}
func (*IdentityDefaultTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a957d0fc92995af, []int{0}
}
func (m *IdentityDefaultTransfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentityDefaultTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentityDefaultTransfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentityDefaultTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityDefaultTransfer.Merge(m, src)
}
func (m *IdentityDefaultTransfer) XXX_Size() int {
	return m.Size()
}
func (m *IdentityDefaultTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityDefaultTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityDefaultTransfer proto.InternalMessageInfo

func (m *IdentityDefaultTransfer) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetGISTHash() string {
	if m != nil {
		return m.GISTHash
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetStateHash() string {
	if m != nil {
		return m.StateHash
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetStateCreatedAtTimestamp() string {
	if m != nil {
		return m.StateCreatedAtTimestamp
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetStateCreatedAtBlock() string {
	if m != nil {
		return m.StateCreatedAtBlock
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetStateReplacedBy() string {
	if m != nil {
		return m.StateReplacedBy
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetGISTReplacedBy() string {
	if m != nil {
		return m.GISTReplacedBy
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetGISTCreatedAtTimestamp() string {
	if m != nil {
		return m.GISTCreatedAtTimestamp
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetGISTCreatedAtBlock() string {
	if m != nil {
		return m.GISTCreatedAtBlock
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetReplacedStateHash() string {
	if m != nil {
		return m.ReplacedStateHash
	}
	return ""
}

func (m *IdentityDefaultTransfer) GetReplacedGISTHash() string {
	if m != nil {
		return m.ReplacedGISTHash
	}
	return ""
}

func init() {
	proto.RegisterType((*IdentityDefaultTransfer)(nil), "rarimo.rarimocore.rarimocore.IdentityDefaultTransfer")
}

func init() {
	proto.RegisterFile("rarimocore/op_identity_default_transfer.proto", fileDescriptor_0a957d0fc92995af)
}

var fileDescriptor_0a957d0fc92995af = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x31, 0x05, 0x5a, 0xb6, 0x2d, 0x85, 0x2d, 0x2a, 0x16, 0x42, 0x56, 0xd5, 0x43, 0x85,
	0xaa, 0x62, 0x23, 0x55, 0xaa, 0x7a, 0x2d, 0xad, 0x94, 0x90, 0x23, 0x70, 0xca, 0x05, 0x2d, 0xf6,
	0x62, 0x56, 0xd8, 0x5e, 0x6b, 0x3d, 0x91, 0xc2, 0x5b, 0xe4, 0x92, 0x77, 0xca, 0x91, 0x63, 0x8e,
	0x11, 0xbc, 0x48, 0xe4, 0x59, 0xf3, 0x1f, 0x4e, 0x9e, 0xf9, 0x7d, 0xdf, 0xa7, 0x19, 0xdb, 0x43,
	0x3a, 0x8a, 0x29, 0x11, 0x4a, 0x57, 0x2a, 0xee, 0xc8, 0x78, 0x2c, 0x3c, 0x1e, 0x81, 0x80, 0xc5,
	0xd8, 0xe3, 0x53, 0x76, 0x17, 0xc0, 0x18, 0x14, 0x8b, 0x92, 0x29, 0x57, 0x76, 0xac, 0x24, 0x48,
	0xda, 0xd2, 0x76, 0x7b, 0x97, 0xda, 0x2b, 0x9b, 0x0d, 0x90, 0x73, 0x1e, 0x85, 0x2c, 0x62, 0x3e,
	0x57, 0x8e, 0x00, 0x1e, 0xea, 0x58, 0xb3, 0xee, 0x4b, 0x5f, 0x62, 0xe9, 0xa4, 0x95, 0xa6, 0xdf,
	0x1e, 0x0b, 0xa4, 0xd1, 0xcf, 0x06, 0xfe, 0xd7, 0xf3, 0x46, 0xd9, 0x38, 0xda, 0x24, 0xef, 0x5c,
	0x19, 0x81, 0x62, 0x2e, 0x98, 0xc6, 0x57, 0xa3, 0x5d, 0x1e, 0x6c, 0x7b, 0x5a, 0x27, 0x45, 0x77,
	0xc6, 0x44, 0x64, 0xe6, 0x51, 0xd0, 0x4d, 0x9a, 0xb8, 0xea, 0x0f, 0x47, 0xd7, 0x2c, 0x99, 0x99,
	0x6f, 0x74, 0x62, 0xd3, 0xd3, 0x0a, 0xc9, 0x0b, 0xcf, 0x2c, 0x20, 0xcd, 0x0b, 0x8f, 0xb6, 0x48,
	0x39, 0x01, 0x06, 0x1c, 0xcd, 0x45, 0xc4, 0x3b, 0x40, 0xff, 0x90, 0x06, 0x36, 0xff, 0x14, 0x67,
	0xc0, 0xbd, 0xbf, 0x30, 0x12, 0x21, 0x4f, 0x80, 0x85, 0xb1, 0x59, 0x42, 0xef, 0x25, 0x99, 0x76,
	0xc9, 0xe7, 0x43, 0xa9, 0x17, 0x48, 0x77, 0x6e, 0xbe, 0xc5, 0xd4, 0x39, 0x89, 0xb6, 0xc9, 0x27,
	0xc4, 0x03, 0x1e, 0x07, 0xcc, 0xe5, 0x5e, 0x6f, 0x61, 0x12, 0x74, 0x1f, 0x63, 0xfa, 0x9d, 0x54,
	0xd2, 0xf7, 0xd9, 0x33, 0xbe, 0x47, 0xe3, 0x11, 0xa5, 0xbf, 0xc9, 0x97, 0x94, 0x9c, 0x59, 0xfe,
	0x03, 0xfa, 0x2f, 0xa8, 0xd4, 0x26, 0xf4, 0x40, 0xd1, 0xab, 0x7f, 0xc4, 0xcc, 0x19, 0x85, 0xfe,
	0x24, 0x35, 0x95, 0x4d, 0x1d, 0x6e, 0xbf, 0x65, 0x15, 0xed, 0xa7, 0x02, 0xfd, 0x41, 0xaa, 0x1b,
	0xb8, 0xfd, 0x4b, 0x35, 0x34, 0x9f, 0xf0, 0xde, 0xcd, 0xd3, 0xca, 0x32, 0x96, 0x2b, 0xcb, 0x78,
	0x59, 0x59, 0xc6, 0xc3, 0xda, 0xca, 0x2d, 0xd7, 0x56, 0xee, 0x79, 0x6d, 0xe5, 0x6e, 0xbb, 0xbe,
	0x80, 0x80, 0x4d, 0x6c, 0x57, 0x86, 0x8e, 0xbe, 0xbb, 0xec, 0xd1, 0xc1, 0x03, 0xbe, 0x77, 0xf6,
	0xae, 0x19, 0x16, 0x31, 0x4f, 0x26, 0x25, 0x3c, 0xb5, 0x5f, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x65, 0xe1, 0x32, 0xe8, 0x02, 0x00, 0x00,
}

func (m *IdentityDefaultTransfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentityDefaultTransfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentityDefaultTransfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReplacedGISTHash) > 0 {
		i -= len(m.ReplacedGISTHash)
		copy(dAtA[i:], m.ReplacedGISTHash)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.ReplacedGISTHash)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if len(m.ReplacedStateHash) > 0 {
		i -= len(m.ReplacedStateHash)
		copy(dAtA[i:], m.ReplacedStateHash)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.ReplacedStateHash)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.GISTCreatedAtBlock) > 0 {
		i -= len(m.GISTCreatedAtBlock)
		copy(dAtA[i:], m.GISTCreatedAtBlock)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.GISTCreatedAtBlock)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.GISTCreatedAtTimestamp) > 0 {
		i -= len(m.GISTCreatedAtTimestamp)
		copy(dAtA[i:], m.GISTCreatedAtTimestamp)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.GISTCreatedAtTimestamp)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.GISTReplacedBy) > 0 {
		i -= len(m.GISTReplacedBy)
		copy(dAtA[i:], m.GISTReplacedBy)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.GISTReplacedBy)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.StateReplacedBy) > 0 {
		i -= len(m.StateReplacedBy)
		copy(dAtA[i:], m.StateReplacedBy)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.StateReplacedBy)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.StateCreatedAtBlock) > 0 {
		i -= len(m.StateCreatedAtBlock)
		copy(dAtA[i:], m.StateCreatedAtBlock)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.StateCreatedAtBlock)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.StateCreatedAtTimestamp) > 0 {
		i -= len(m.StateCreatedAtTimestamp)
		copy(dAtA[i:], m.StateCreatedAtTimestamp)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.StateCreatedAtTimestamp)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StateHash) > 0 {
		i -= len(m.StateHash)
		copy(dAtA[i:], m.StateHash)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.StateHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.GISTHash) > 0 {
		i -= len(m.GISTHash)
		copy(dAtA[i:], m.GISTHash)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.GISTHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintOpIdentityDefaultTransfer(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpIdentityDefaultTransfer(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpIdentityDefaultTransfer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdentityDefaultTransfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.GISTHash)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.StateHash)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.StateCreatedAtTimestamp)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.StateCreatedAtBlock)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.StateReplacedBy)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.GISTReplacedBy)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.GISTCreatedAtTimestamp)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.GISTCreatedAtBlock)
	if l > 0 {
		n += 1 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.ReplacedStateHash)
	if l > 0 {
		n += 2 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	l = len(m.ReplacedGISTHash)
	if l > 0 {
		n += 2 + l + sovOpIdentityDefaultTransfer(uint64(l))
	}
	return n
}

func sovOpIdentityDefaultTransfer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpIdentityDefaultTransfer(x uint64) (n int) {
	return sovOpIdentityDefaultTransfer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdentityDefaultTransfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpIdentityDefaultTransfer
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
			return fmt.Errorf("proto: IdentityDefaultTransfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentityDefaultTransfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
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
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
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
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GISTHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateCreatedAtTimestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateCreatedAtTimestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateCreatedAtBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateCreatedAtBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateReplacedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateReplacedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GISTReplacedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GISTReplacedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GISTCreatedAtTimestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GISTCreatedAtTimestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GISTCreatedAtBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GISTCreatedAtBlock = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplacedStateHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplacedStateHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplacedGISTHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpIdentityDefaultTransfer
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
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplacedGISTHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOpIdentityDefaultTransfer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOpIdentityDefaultTransfer
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
func skipOpIdentityDefaultTransfer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpIdentityDefaultTransfer
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
					return 0, ErrIntOverflowOpIdentityDefaultTransfer
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
					return 0, ErrIntOverflowOpIdentityDefaultTransfer
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
				return 0, ErrInvalidLengthOpIdentityDefaultTransfer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpIdentityDefaultTransfer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpIdentityDefaultTransfer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpIdentityDefaultTransfer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpIdentityDefaultTransfer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpIdentityDefaultTransfer = fmt.Errorf("proto: unexpected end of group")
)