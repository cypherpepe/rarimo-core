// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the rarimocore module's genesis state.
type GenesisState struct {
	Params             Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DepositList        []Deposit        `protobuf:"bytes,2,rep,name=depositList,proto3" json:"depositList"`
	ConfirmationList   []Confirmation   `protobuf:"bytes,3,rep,name=confirmationList,proto3" json:"confirmationList"`
	ChangeKeyECDSAList []ChangeKeyECDSA `protobuf:"bytes,4,rep,name=changeKeyECDSAList,proto3" json:"changeKeyECDSAList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c823eb3ed15d1a7, []int{0}
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

func (m *GenesisState) GetDepositList() []Deposit {
	if m != nil {
		return m.DepositList
	}
	return nil
}

func (m *GenesisState) GetConfirmationList() []Confirmation {
	if m != nil {
		return m.ConfirmationList
	}
	return nil
}

func (m *GenesisState) GetChangeKeyECDSAList() []ChangeKeyECDSA {
	if m != nil {
		return m.ChangeKeyECDSAList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "rarifyprotocol.rarimocore.rarimocore.GenesisState")
}

func init() { proto.RegisterFile("rarimocore/genesis.proto", fileDescriptor_7c823eb3ed15d1a7) }

var fileDescriptor_7c823eb3ed15d1a7 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0xdb, 0xaa, 0x83, 0x7b, 0x07, 0x64, 0x21, 0x51, 0x45, 0xc2, 0x14, 0xc4, 0xd0,
	0x81, 0x24, 0x52, 0x60, 0x61, 0xa4, 0x2d, 0x42, 0x02, 0x06, 0x44, 0xc4, 0xc2, 0x52, 0xb9, 0x89,
	0x1b, 0x0c, 0x4d, 0x1c, 0xd9, 0x1e, 0xc8, 0xcc, 0x0b, 0xf0, 0x58, 0x1d, 0x3b, 0x32, 0x21, 0x94,
	0xbc, 0x08, 0xaa, 0x6d, 0x84, 0x11, 0x0c, 0xd9, 0x1c, 0xfd, 0xe7, 0xff, 0xbe, 0x1c, 0x1b, 0x0c,
	0x38, 0xe6, 0x34, 0x67, 0x09, 0xe3, 0x24, 0xcc, 0x48, 0x41, 0x04, 0x15, 0x41, 0xc9, 0x99, 0x64,
	0xf0, 0x70, 0x93, 0x2c, 0x2a, 0xf5, 0x91, 0xb0, 0x65, 0xf0, 0x3d, 0x68, 0x1d, 0xbd, 0xed, 0x8c,
	0x65, 0x4c, 0xcd, 0x84, 0x9b, 0x93, 0xee, 0x7a, 0x3b, 0x16, 0xb5, 0xc4, 0x1c, 0xe7, 0x06, 0xea,
	0xd9, 0xba, 0x94, 0x94, 0x4c, 0x50, 0x69, 0x92, 0x5d, 0x2b, 0x49, 0x58, 0xb1, 0xa0, 0x3c, 0xc7,
	0x92, 0xb2, 0xc2, 0xc4, 0xfb, 0x76, 0xfc, 0x80, 0x8b, 0x8c, 0xcc, 0x9e, 0x48, 0x35, 0x23, 0x49,
	0x2a, 0xb0, 0x1e, 0x39, 0x78, 0xe9, 0x80, 0xff, 0x17, 0x7a, 0x85, 0x58, 0x62, 0x49, 0xe0, 0x25,
	0xe8, 0x69, 0xf9, 0xc0, 0x1d, 0xba, 0xa3, 0x7e, 0x74, 0x14, 0xb4, 0x59, 0x29, 0xb8, 0x51, 0x9d,
	0x71, 0x77, 0xf5, 0xbe, 0xe7, 0xdc, 0x1a, 0x02, 0xbc, 0x03, 0x7d, 0xf3, 0xbf, 0xd7, 0x54, 0xc8,
	0xc1, 0xbf, 0x61, 0x67, 0xd4, 0x8f, 0xfc, 0x76, 0xc0, 0xa9, 0x2e, 0x1a, 0xa2, 0xcd, 0x81, 0x29,
	0xd8, 0xb2, 0x97, 0x55, 0xec, 0x8e, 0x62, 0x47, 0xed, 0xd8, 0x13, 0xab, 0x6d, 0x04, 0xbf, 0x88,
	0xf0, 0x11, 0x40, 0x7d, 0x67, 0x57, 0xa4, 0x3a, 0x9f, 0x4c, 0xe3, 0x33, 0xe5, 0xe9, 0x2a, 0xcf,
	0x49, 0x4b, 0xcf, 0x8f, 0xbe, 0x31, 0xfd, 0x41, 0x1d, 0xc7, 0xab, 0x1a, 0xb9, 0xeb, 0x1a, 0xb9,
	0x1f, 0x35, 0x72, 0x5f, 0x1b, 0xe4, 0xac, 0x1b, 0xe4, 0xbc, 0x35, 0xc8, 0xb9, 0x3f, 0xcd, 0xa8,
	0x5c, 0xe2, 0x79, 0x90, 0xb0, 0x3c, 0xd4, 0x4e, 0xff, 0x4b, 0x1a, 0x6a, 0x93, 0xaf, 0x9e, 0xf7,
	0x39, 0xb4, 0xde, 0x5a, 0x56, 0x25, 0x11, 0xf3, 0x9e, 0x1a, 0x3c, 0xfe, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x21, 0x30, 0x46, 0x86, 0xae, 0x02, 0x00, 0x00,
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
	if len(m.ChangeKeyECDSAList) > 0 {
		for iNdEx := len(m.ChangeKeyECDSAList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChangeKeyECDSAList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ConfirmationList) > 0 {
		for iNdEx := len(m.ConfirmationList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConfirmationList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DepositList) > 0 {
		for iNdEx := len(m.DepositList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DepositList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DepositList) > 0 {
		for _, e := range m.DepositList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ConfirmationList) > 0 {
		for _, e := range m.ConfirmationList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ChangeKeyECDSAList) > 0 {
		for _, e := range m.ChangeKeyECDSAList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field DepositList", wireType)
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
			m.DepositList = append(m.DepositList, Deposit{})
			if err := m.DepositList[len(m.DepositList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationList", wireType)
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
			m.ConfirmationList = append(m.ConfirmationList, Confirmation{})
			if err := m.ConfirmationList[len(m.ConfirmationList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChangeKeyECDSAList", wireType)
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
			m.ChangeKeyECDSAList = append(m.ChangeKeyECDSAList, ChangeKeyECDSA{})
			if err := m.ChangeKeyECDSAList[len(m.ChangeKeyECDSAList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
