// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cscalist/proposal.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// EditCSCAListProposal allows to propose changes in CSCA list, which are required
// in case of a CSCA key compromise or adding new authorities. Use it when you want
// to add or remove a few CSCA keys.
type EditCSCAListProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Poseidon hashes of new CSCA public keys
	ToAdd []string `protobuf:"bytes,3,rep,name=toAdd,proto3" json:"toAdd,omitempty"`
	// Poseidon hashes of existing CSCA public keys
	ToRemove []string `protobuf:"bytes,4,rep,name=toRemove,proto3" json:"toRemove,omitempty"`
}

func (m *EditCSCAListProposal) Reset()         { *m = EditCSCAListProposal{} }
func (m *EditCSCAListProposal) String() string { return proto.CompactTextString(m) }
func (*EditCSCAListProposal) ProtoMessage()    {}
func (*EditCSCAListProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_557c5b3809954aeb, []int{0}
}
func (m *EditCSCAListProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditCSCAListProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditCSCAListProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditCSCAListProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditCSCAListProposal.Merge(m, src)
}
func (m *EditCSCAListProposal) XXX_Size() int {
	return m.Size()
}
func (m *EditCSCAListProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_EditCSCAListProposal.DiscardUnknown(m)
}

var xxx_messageInfo_EditCSCAListProposal proto.InternalMessageInfo

func (m *EditCSCAListProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditCSCAListProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EditCSCAListProposal) GetToAdd() []string {
	if m != nil {
		return m.ToAdd
	}
	return nil
}

func (m *EditCSCAListProposal) GetToRemove() []string {
	if m != nil {
		return m.ToRemove
	}
	return nil
}

// ReplaceCSCAListProposal allows to propose a different CSCA list. This drops the
// old tree and rebuilds the new one from scratch. Use EditCSCAListProposal instead,
// unless you know what you are doing.
type ReplaceCSCAListProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Poseidon hashes of all CSCA public keys
	Leaves []string `protobuf:"bytes,3,rep,name=leaves,proto3" json:"leaves,omitempty"`
}

func (m *ReplaceCSCAListProposal) Reset()         { *m = ReplaceCSCAListProposal{} }
func (m *ReplaceCSCAListProposal) String() string { return proto.CompactTextString(m) }
func (*ReplaceCSCAListProposal) ProtoMessage()    {}
func (*ReplaceCSCAListProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_557c5b3809954aeb, []int{1}
}
func (m *ReplaceCSCAListProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplaceCSCAListProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReplaceCSCAListProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReplaceCSCAListProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceCSCAListProposal.Merge(m, src)
}
func (m *ReplaceCSCAListProposal) XXX_Size() int {
	return m.Size()
}
func (m *ReplaceCSCAListProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceCSCAListProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceCSCAListProposal proto.InternalMessageInfo

func (m *ReplaceCSCAListProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ReplaceCSCAListProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ReplaceCSCAListProposal) GetLeaves() []string {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func init() {
	proto.RegisterType((*EditCSCAListProposal)(nil), "rarimo.rarimocore.cscalist.EditCSCAListProposal")
	proto.RegisterType((*ReplaceCSCAListProposal)(nil), "rarimo.rarimocore.cscalist.ReplaceCSCAListProposal")
}

func init() { proto.RegisterFile("cscalist/proposal.proto", fileDescriptor_557c5b3809954aeb) }

var fileDescriptor_557c5b3809954aeb = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2e, 0x4e, 0x4e,
	0xcc, 0xc9, 0x2c, 0x2e, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2a, 0x4a, 0x2c, 0xca, 0xcc, 0xcd, 0xd7, 0x83, 0x50, 0xc9, 0xf9,
	0x45, 0xa9, 0x7a, 0x30, 0xa5, 0x4a, 0x0d, 0x8c, 0x5c, 0x22, 0xae, 0x29, 0x99, 0x25, 0xce, 0xc1,
	0xce, 0x8e, 0x3e, 0x99, 0xc5, 0x25, 0x01, 0x50, 0xad, 0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25,
	0x39, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x02, 0x17, 0x77, 0x4a,
	0x6a, 0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0x13, 0x58, 0x0e, 0x59, 0x08,
	0xac, 0x2f, 0xdf, 0x31, 0x25, 0x45, 0x82, 0x59, 0x81, 0x19, 0xac, 0x0f, 0xc4, 0x11, 0x92, 0xe2,
	0xe2, 0x28, 0xc9, 0x0f, 0x4a, 0xcd, 0xcd, 0x2f, 0x4b, 0x95, 0x60, 0x01, 0x4b, 0xc0, 0xf9, 0x4a,
	0x99, 0x5c, 0xe2, 0x41, 0xa9, 0x05, 0x39, 0x89, 0xc9, 0xa9, 0x54, 0x73, 0x84, 0x18, 0x17, 0x5b,
	0x4e, 0x6a, 0x62, 0x59, 0x6a, 0x31, 0xd4, 0x15, 0x50, 0x9e, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x43, 0xc2, 0x09, 0x4a, 0xe9, 0x82, 0xc2, 0x4b, 0xbf, 0x42, 0x1f, 0x1e, 0xb8, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xa0, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xcf, 0x03, 0x20, 0x75, 0x01, 0x00, 0x00,
}

func (m *EditCSCAListProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditCSCAListProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditCSCAListProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ToRemove) > 0 {
		for iNdEx := len(m.ToRemove) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ToRemove[iNdEx])
			copy(dAtA[i:], m.ToRemove[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.ToRemove[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ToAdd) > 0 {
		for iNdEx := len(m.ToAdd) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ToAdd[iNdEx])
			copy(dAtA[i:], m.ToAdd[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.ToAdd[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ReplaceCSCAListProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplaceCSCAListProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplaceCSCAListProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leaves) > 0 {
		for iNdEx := len(m.Leaves) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Leaves[iNdEx])
			copy(dAtA[i:], m.Leaves[iNdEx])
			i = encodeVarintProposal(dAtA, i, uint64(len(m.Leaves[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EditCSCAListProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.ToAdd) > 0 {
		for _, s := range m.ToAdd {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	if len(m.ToRemove) > 0 {
		for _, s := range m.ToRemove {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *ReplaceCSCAListProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.Leaves) > 0 {
		for _, s := range m.Leaves {
			l = len(s)
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EditCSCAListProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: EditCSCAListProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditCSCAListProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAdd", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAdd = append(m.ToAdd, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToRemove", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToRemove = append(m.ToRemove, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *ReplaceCSCAListProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ReplaceCSCAListProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplaceCSCAListProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leaves = append(m.Leaves, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)