// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oraclemanager/oracle.proto

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

type OracleStatus int32

const (
	OracleStatus_Inactive OracleStatus = 0
	OracleStatus_Active   OracleStatus = 1
	OracleStatus_Jailed   OracleStatus = 2
	OracleStatus_Freezed  OracleStatus = 3
	OracleStatus_Slashed  OracleStatus = 4
)

var OracleStatus_name = map[int32]string{
	0: "Inactive",
	1: "Active",
	2: "Jailed",
	3: "Freezed",
	4: "Slashed",
}

var OracleStatus_value = map[string]int32{
	"Inactive": 0,
	"Active":   1,
	"Jailed":   2,
	"Freezed":  3,
	"Slashed":  4,
}

func (x OracleStatus) String() string {
	return proto.EnumName(OracleStatus_name, int32(x))
}

func (OracleStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{0}
}

type OracleIndex struct {
	Chain   string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *OracleIndex) Reset()         { *m = OracleIndex{} }
func (m *OracleIndex) String() string { return proto.CompactTextString(m) }
func (*OracleIndex) ProtoMessage()    {}
func (*OracleIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{0}
}
func (m *OracleIndex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleIndex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleIndex.Merge(m, src)
}
func (m *OracleIndex) XXX_Size() int {
	return m.Size()
}
func (m *OracleIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleIndex.DiscardUnknown(m)
}

var xxx_messageInfo_OracleIndex proto.InternalMessageInfo

func (m *OracleIndex) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *OracleIndex) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type Oracle struct {
	Index                 *OracleIndex `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Status                OracleStatus `protobuf:"varint,2,opt,name=status,proto3,enum=rarimo.rarimocore.oraclemanager.OracleStatus" json:"status,omitempty"`
	Stake                 string       `protobuf:"bytes,3,opt,name=stake,proto3" json:"stake,omitempty"`
	MissedCount           uint64       `protobuf:"varint,4,opt,name=missedCount,proto3" json:"missedCount,omitempty"`
	ViolationsCount       uint64       `protobuf:"varint,5,opt,name=violationsCount,proto3" json:"violationsCount,omitempty"`
	FreezeEndBlock        uint64       `protobuf:"varint,6,opt,name=freezeEndBlock,proto3" json:"freezeEndBlock,omitempty"`
	VotesCount            uint64       `protobuf:"varint,7,opt,name=votesCount,proto3" json:"votesCount,omitempty"`
	CreateOperationsCount uint64       `protobuf:"varint,8,opt,name=createOperationsCount,proto3" json:"createOperationsCount,omitempty"`
}

func (m *Oracle) Reset()         { *m = Oracle{} }
func (m *Oracle) String() string { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()    {}
func (*Oracle) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c1a9fae0dab29d9, []int{1}
}
func (m *Oracle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Oracle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Oracle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Oracle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oracle.Merge(m, src)
}
func (m *Oracle) XXX_Size() int {
	return m.Size()
}
func (m *Oracle) XXX_DiscardUnknown() {
	xxx_messageInfo_Oracle.DiscardUnknown(m)
}

var xxx_messageInfo_Oracle proto.InternalMessageInfo

func (m *Oracle) GetIndex() *OracleIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *Oracle) GetStatus() OracleStatus {
	if m != nil {
		return m.Status
	}
	return OracleStatus_Inactive
}

func (m *Oracle) GetStake() string {
	if m != nil {
		return m.Stake
	}
	return ""
}

func (m *Oracle) GetMissedCount() uint64 {
	if m != nil {
		return m.MissedCount
	}
	return 0
}

func (m *Oracle) GetViolationsCount() uint64 {
	if m != nil {
		return m.ViolationsCount
	}
	return 0
}

func (m *Oracle) GetFreezeEndBlock() uint64 {
	if m != nil {
		return m.FreezeEndBlock
	}
	return 0
}

func (m *Oracle) GetVotesCount() uint64 {
	if m != nil {
		return m.VotesCount
	}
	return 0
}

func (m *Oracle) GetCreateOperationsCount() uint64 {
	if m != nil {
		return m.CreateOperationsCount
	}
	return 0
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.oraclemanager.OracleStatus", OracleStatus_name, OracleStatus_value)
	proto.RegisterType((*OracleIndex)(nil), "rarimo.rarimocore.oraclemanager.OracleIndex")
	proto.RegisterType((*Oracle)(nil), "rarimo.rarimocore.oraclemanager.Oracle")
}

func init() { proto.RegisterFile("oraclemanager/oracle.proto", fileDescriptor_6c1a9fae0dab29d9) }

var fileDescriptor_6c1a9fae0dab29d9 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x4d, 0xfa, 0xda, 0xf4, 0x79, 0xf3, 0x78, 0x86, 0xe1, 0x09, 0xe1, 0x2d, 0x62, 0x79, 0x0b,
	0x29, 0x62, 0x13, 0x68, 0xdd, 0xba, 0xb0, 0x52, 0xa1, 0x82, 0x16, 0xd2, 0x9d, 0xbb, 0xe9, 0xe4,
	0x9a, 0x0e, 0x4d, 0x32, 0x65, 0x66, 0x5a, 0xaa, 0x5f, 0xe1, 0x67, 0xb9, 0xec, 0xd2, 0xa5, 0xb4,
	0x9f, 0xe0, 0x0f, 0x48, 0x66, 0x2a, 0xa4, 0x45, 0xf0, 0xad, 0xe6, 0x9e, 0x33, 0xe7, 0xdc, 0x3b,
	0x87, 0xb9, 0x70, 0x2f, 0x24, 0x65, 0x05, 0x96, 0xb4, 0xa2, 0x39, 0xca, 0xc4, 0xa2, 0x78, 0x2d,
	0x85, 0x16, 0xe4, 0xb9, 0xa4, 0x92, 0x97, 0x22, 0xb6, 0x07, 0x13, 0x12, 0xe3, 0x33, 0xf5, 0xfd,
	0x5d, 0x2e, 0x72, 0x61, 0xb4, 0x49, 0x5d, 0x59, 0xdb, 0xc3, 0x1b, 0xf0, 0x67, 0x46, 0x36, 0xad,
	0x32, 0xdc, 0x91, 0x3b, 0xe8, 0xb0, 0x25, 0xe5, 0x55, 0xe8, 0xf6, 0xdc, 0xfe, 0x93, 0xd4, 0x02,
	0x12, 0x42, 0x97, 0x32, 0x26, 0x36, 0x95, 0x0e, 0x5b, 0x86, 0xff, 0x0b, 0x1f, 0x7e, 0xb7, 0xc0,
	0xb3, 0x7e, 0x32, 0x86, 0x0e, 0xaf, 0x7b, 0x18, 0xab, 0x3f, 0x7c, 0x15, 0xff, 0xe7, 0x41, 0x71,
	0x63, 0x6e, 0x6a, 0xad, 0x64, 0x02, 0x9e, 0xd2, 0x54, 0x6f, 0x94, 0x99, 0x73, 0x3b, 0x1c, 0x3c,
	0xb2, 0xc9, 0xdc, 0x98, 0xd2, 0x93, 0xb9, 0x4e, 0xa1, 0x34, 0x5d, 0x61, 0x78, 0x65, 0x53, 0x18,
	0x40, 0x7a, 0xe0, 0x97, 0x5c, 0x29, 0xcc, 0xde, 0x99, 0x24, 0xed, 0x9e, 0xdb, 0x6f, 0xa7, 0x4d,
	0x8a, 0xf4, 0xe1, 0xe9, 0x96, 0x8b, 0x82, 0x6a, 0x2e, 0x2a, 0x65, 0x55, 0x1d, 0xa3, 0xba, 0xa4,
	0xc9, 0x0b, 0xb8, 0xfd, 0x22, 0x11, 0xbf, 0xe1, 0xa4, 0xca, 0xc6, 0x85, 0x60, 0xab, 0xd0, 0x33,
	0xc2, 0x0b, 0x96, 0x44, 0x00, 0x5b, 0xa1, 0xf1, 0xd4, 0xac, 0x6b, 0x34, 0x0d, 0x86, 0xbc, 0x86,
	0x67, 0x4c, 0x22, 0xd5, 0x38, 0x5b, 0xa3, 0x6c, 0xce, 0xbd, 0x36, 0xd2, 0x7f, 0x5f, 0xbe, 0xfc,
	0x04, 0x37, 0xcd, 0xdc, 0xe4, 0x06, 0xae, 0xa7, 0x15, 0x65, 0x9a, 0x6f, 0x31, 0x70, 0x08, 0x80,
	0xf7, 0xd6, 0xd6, 0x6e, 0x5d, 0x7f, 0xa0, 0xbc, 0xc0, 0x2c, 0x68, 0x11, 0x1f, 0xba, 0xef, 0xcd,
	0xeb, 0xb2, 0xe0, 0xaa, 0x06, 0xf3, 0x82, 0xaa, 0x25, 0x66, 0x41, 0x7b, 0xfc, 0xf1, 0xc7, 0x21,
	0x72, 0xf7, 0x87, 0xc8, 0xfd, 0x75, 0x88, 0xdc, 0xef, 0xc7, 0xc8, 0xd9, 0x1f, 0x23, 0xe7, 0xe7,
	0x31, 0x72, 0x3e, 0x8f, 0x72, 0xae, 0x0b, 0xba, 0x88, 0x99, 0x28, 0x13, 0xfb, 0x07, 0xa7, 0x63,
	0x50, 0xff, 0x45, 0xb2, 0x4b, 0xce, 0x37, 0x52, 0x7f, 0x5d, 0xa3, 0x5a, 0x78, 0x66, 0xb5, 0x46,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0x0c, 0x18, 0x67, 0xaf, 0x02, 0x00, 0x00,
}

func (m *OracleIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleIndex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleIndex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Oracle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Oracle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Oracle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreateOperationsCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.CreateOperationsCount))
		i--
		dAtA[i] = 0x40
	}
	if m.VotesCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.VotesCount))
		i--
		dAtA[i] = 0x38
	}
	if m.FreezeEndBlock != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.FreezeEndBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.ViolationsCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ViolationsCount))
		i--
		dAtA[i] = 0x28
	}
	if m.MissedCount != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.MissedCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Stake) > 0 {
		i -= len(m.Stake)
		copy(dAtA[i:], m.Stake)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Stake)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Index != nil {
		{
			size, err := m.Index.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOracle(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleIndex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *Oracle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != nil {
		l = m.Index.Size()
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovOracle(uint64(m.Status))
	}
	l = len(m.Stake)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.MissedCount != 0 {
		n += 1 + sovOracle(uint64(m.MissedCount))
	}
	if m.ViolationsCount != 0 {
		n += 1 + sovOracle(uint64(m.ViolationsCount))
	}
	if m.FreezeEndBlock != 0 {
		n += 1 + sovOracle(uint64(m.FreezeEndBlock))
	}
	if m.VotesCount != 0 {
		n += 1 + sovOracle(uint64(m.VotesCount))
	}
	if m.CreateOperationsCount != 0 {
		n += 1 + sovOracle(uint64(m.CreateOperationsCount))
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: OracleIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Oracle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Oracle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Oracle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Index == nil {
				m.Index = &OracleIndex{}
			}
			if err := m.Index.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= OracleStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stake = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedCount", wireType)
			}
			m.MissedCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViolationsCount", wireType)
			}
			m.ViolationsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ViolationsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreezeEndBlock", wireType)
			}
			m.FreezeEndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FreezeEndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesCount", wireType)
			}
			m.VotesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotesCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateOperationsCount", wireType)
			}
			m.CreateOperationsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateOperationsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
