// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rarimocore/params.proto

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

type ParamsUpdateType int32

const (
	ParamsUpdateType_CHANGE_SET ParamsUpdateType = 0
	ParamsUpdateType_OTHER      ParamsUpdateType = 1
)

var ParamsUpdateType_name = map[int32]string{
	0: "CHANGE_SET",
	1: "OTHER",
}

var ParamsUpdateType_value = map[string]int32{
	"CHANGE_SET": 0,
	"OTHER":      1,
}

func (x ParamsUpdateType) String() string {
	return proto.EnumName(ParamsUpdateType_name, int32(x))
}

func (ParamsUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}

type PartyStatus int32

const (
	PartyStatus_Active   PartyStatus = 0
	PartyStatus_Frozen   PartyStatus = 1
	PartyStatus_Slashed  PartyStatus = 2
	PartyStatus_Inactive PartyStatus = 3
)

var PartyStatus_name = map[int32]string{
	0: "Active",
	1: "Frozen",
	2: "Slashed",
	3: "Inactive",
}

var PartyStatus_value = map[string]int32{
	"Active":   0,
	"Frozen":   1,
	"Slashed":  2,
	"Inactive": 3,
}

func (x PartyStatus) String() string {
	return proto.EnumName(PartyStatus_name, int32(x))
}

func (PartyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{1}
}

type Party struct {
	// PublicKey in hex format
	PubKey string `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	// Server address connect to
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Rarimo core account
	Account         string      `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Status          PartyStatus `protobuf:"varint,4,opt,name=status,proto3,enum=rarimo.rarimocore.rarimocore.PartyStatus" json:"status,omitempty"`
	ViolationsCount uint64      `protobuf:"varint,5,opt,name=violationsCount,proto3" json:"violationsCount,omitempty"`
	FreezeEndBlock  uint64      `protobuf:"varint,6,opt,name=freezeEndBlock,proto3" json:"freezeEndBlock,omitempty"`
	Delegator       string      `protobuf:"bytes,7,opt,name=delegator,proto3" json:"delegator,omitempty"`
	// service information used in initial setup
	CommittedGlobalPublicKey string `protobuf:"bytes,8,opt,name=committedGlobalPublicKey,proto3" json:"committedGlobalPublicKey,omitempty"`
	// session where party was reported and violations count was incremented
	ReportedSessions []string `protobuf:"bytes,9,rep,name=reportedSessions,proto3" json:"reportedSessions,omitempty"`
}

func (m *Party) Reset()         { *m = Party{} }
func (m *Party) String() string { return proto.CompactTextString(m) }
func (*Party) ProtoMessage()    {}
func (*Party) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{0}
}
func (m *Party) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Party) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Party.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Party) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Party.Merge(m, src)
}
func (m *Party) XXX_Size() int {
	return m.Size()
}
func (m *Party) XXX_DiscardUnknown() {
	xxx_messageInfo_Party.DiscardUnknown(m)
}

var xxx_messageInfo_Party proto.InternalMessageInfo

func (m *Party) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *Party) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Party) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Party) GetStatus() PartyStatus {
	if m != nil {
		return m.Status
	}
	return PartyStatus_Active
}

func (m *Party) GetViolationsCount() uint64 {
	if m != nil {
		return m.ViolationsCount
	}
	return 0
}

func (m *Party) GetFreezeEndBlock() uint64 {
	if m != nil {
		return m.FreezeEndBlock
	}
	return 0
}

func (m *Party) GetDelegator() string {
	if m != nil {
		return m.Delegator
	}
	return ""
}

func (m *Party) GetCommittedGlobalPublicKey() string {
	if m != nil {
		return m.CommittedGlobalPublicKey
	}
	return ""
}

func (m *Party) GetReportedSessions() []string {
	if m != nil {
		return m.ReportedSessions
	}
	return nil
}

// Params defines the parameters for the module.
type Params struct {
	KeyECDSA           string   `protobuf:"bytes,1,opt,name=keyECDSA,proto3" json:"keyECDSA,omitempty"`
	Threshold          uint64   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Parties            []*Party `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	IsUpdateRequired   bool     `protobuf:"varint,5,opt,name=isUpdateRequired,proto3" json:"isUpdateRequired,omitempty"`
	LastSignature      string   `protobuf:"bytes,6,opt,name=lastSignature,proto3" json:"lastSignature,omitempty"`
	StakeAmount        string   `protobuf:"bytes,7,opt,name=stakeAmount,proto3" json:"stakeAmount,omitempty"`
	StakeDenom         string   `protobuf:"bytes,8,opt,name=stakeDenom,proto3" json:"stakeDenom,omitempty"`
	MaxViolationsCount uint64   `protobuf:"varint,9,opt,name=maxViolationsCount,proto3" json:"maxViolationsCount,omitempty"`
	FreezeBlocksPeriod uint64   `protobuf:"varint,10,opt,name=freezeBlocksPeriod,proto3" json:"freezeBlocksPeriod,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_997c11d1d0285e72, []int{1}
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

func (m *Params) GetKeyECDSA() string {
	if m != nil {
		return m.KeyECDSA
	}
	return ""
}

func (m *Params) GetThreshold() uint64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *Params) GetParties() []*Party {
	if m != nil {
		return m.Parties
	}
	return nil
}

func (m *Params) GetIsUpdateRequired() bool {
	if m != nil {
		return m.IsUpdateRequired
	}
	return false
}

func (m *Params) GetLastSignature() string {
	if m != nil {
		return m.LastSignature
	}
	return ""
}

func (m *Params) GetStakeAmount() string {
	if m != nil {
		return m.StakeAmount
	}
	return ""
}

func (m *Params) GetStakeDenom() string {
	if m != nil {
		return m.StakeDenom
	}
	return ""
}

func (m *Params) GetMaxViolationsCount() uint64 {
	if m != nil {
		return m.MaxViolationsCount
	}
	return 0
}

func (m *Params) GetFreezeBlocksPeriod() uint64 {
	if m != nil {
		return m.FreezeBlocksPeriod
	}
	return 0
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.ParamsUpdateType", ParamsUpdateType_name, ParamsUpdateType_value)
	proto.RegisterEnum("rarimo.rarimocore.rarimocore.PartyStatus", PartyStatus_name, PartyStatus_value)
	proto.RegisterType((*Party)(nil), "rarimo.rarimocore.rarimocore.Party")
	proto.RegisterType((*Params)(nil), "rarimo.rarimocore.rarimocore.Params")
}

func init() { proto.RegisterFile("rarimocore/params.proto", fileDescriptor_997c11d1d0285e72) }

var fileDescriptor_997c11d1d0285e72 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0xd6, 0x89, 0x27, 0x50, 0xac, 0x3d, 0x80, 0x85, 0x2a, 0xcb, 0x2a, 0x08, 0x99,
	0x4a, 0x75, 0x51, 0xb9, 0x21, 0x21, 0x91, 0xb6, 0xa1, 0x05, 0x24, 0x88, 0xec, 0xc2, 0x81, 0x0b,
	0xda, 0xd8, 0x43, 0xbb, 0xaa, 0xed, 0x35, 0xbb, 0xeb, 0xaa, 0xe9, 0x57, 0xf0, 0x59, 0x1c, 0x7b,
	0xe4, 0x88, 0x9a, 0x3b, 0x37, 0xee, 0xc8, 0x6b, 0x97, 0x84, 0x04, 0x10, 0x27, 0xef, 0xbc, 0x37,
	0x23, 0xbf, 0x7d, 0x6f, 0x07, 0xee, 0x08, 0x2a, 0x58, 0xc6, 0x63, 0x2e, 0x70, 0xbb, 0xa0, 0x82,
	0x66, 0x32, 0x28, 0x04, 0x57, 0x9c, 0xac, 0xd7, 0x44, 0x30, 0xe3, 0xe7, 0x8e, 0x1b, 0xdf, 0xdb,
	0xb0, 0x3a, 0xa2, 0x42, 0x4d, 0xc8, 0x6d, 0x30, 0x8b, 0x72, 0xfc, 0x0a, 0x27, 0x8e, 0xe1, 0x19,
	0xbe, 0x15, 0x36, 0x15, 0x71, 0xa0, 0x4b, 0x93, 0x44, 0xa0, 0x94, 0x4e, 0x5b, 0x13, 0xd7, 0xa5,
	0x66, 0xe2, 0x98, 0x97, 0xb9, 0x72, 0x3a, 0x0d, 0x53, 0x97, 0x64, 0x00, 0xa6, 0x54, 0x54, 0x95,
	0xd2, 0x59, 0xf1, 0x0c, 0x7f, 0x6d, 0xe7, 0x61, 0xf0, 0x2f, 0x11, 0x81, 0x16, 0x10, 0xe9, 0x81,
	0xb0, 0x19, 0x24, 0x3e, 0xdc, 0x3a, 0x63, 0x3c, 0xa5, 0x8a, 0xf1, 0x5c, 0xee, 0xe9, 0x9f, 0xac,
	0x7a, 0x86, 0xbf, 0x12, 0x2e, 0xc2, 0xe4, 0x01, 0xac, 0x7d, 0x14, 0x88, 0x17, 0x38, 0xcc, 0x93,
	0xdd, 0x94, 0xc7, 0xa7, 0x8e, 0xa9, 0x1b, 0x17, 0x50, 0xb2, 0x0e, 0x56, 0x82, 0x29, 0x1e, 0x53,
	0xc5, 0x85, 0xd3, 0xd5, 0x82, 0x67, 0x00, 0x79, 0x02, 0x4e, 0xcc, 0xb3, 0x8c, 0x29, 0x85, 0xc9,
	0x41, 0xca, 0xc7, 0x34, 0x1d, 0x95, 0xe3, 0x94, 0xc5, 0x95, 0x21, 0x3d, 0xdd, 0xfc, 0x57, 0x9e,
	0x6c, 0x82, 0x2d, 0xb0, 0xe0, 0x42, 0x61, 0x12, 0xa1, 0x94, 0x95, 0x34, 0xc7, 0xf2, 0x3a, 0xbe,
	0x15, 0x2e, 0xe1, 0x1b, 0x3f, 0xda, 0x60, 0x8e, 0x74, 0x3e, 0xe4, 0x2e, 0xf4, 0x4e, 0x71, 0x32,
	0xdc, 0xdb, 0x8f, 0x06, 0x8d, 0xe7, 0xbf, 0xea, 0x4a, 0xac, 0x3a, 0x11, 0x28, 0x4f, 0x78, 0x9a,
	0x68, 0xdf, 0x57, 0xc2, 0x19, 0x40, 0x9e, 0x42, 0xb7, 0xa0, 0x42, 0x31, 0x94, 0x4e, 0xc7, 0xeb,
	0xf8, 0xfd, 0x9d, 0x7b, 0xff, 0x61, 0x70, 0x78, 0x3d, 0x53, 0xe9, 0x65, 0xf2, 0x6d, 0x91, 0x50,
	0x85, 0x21, 0x7e, 0x2a, 0x99, 0xc0, 0x44, 0x9b, 0xdb, 0x0b, 0x97, 0x70, 0x72, 0x1f, 0x6e, 0xa6,
	0x54, 0xaa, 0x88, 0x1d, 0xe7, 0x54, 0x95, 0x02, 0xb5, 0xb9, 0x56, 0xf8, 0x3b, 0x48, 0x3c, 0xe8,
	0x4b, 0x45, 0x4f, 0x71, 0x90, 0xe9, 0xa4, 0x6a, 0x77, 0xe7, 0x21, 0xe2, 0x02, 0xe8, 0x72, 0x1f,
	0x73, 0x9e, 0x35, 0x8e, 0xce, 0x21, 0x24, 0x00, 0x92, 0xd1, 0xf3, 0x77, 0x0b, 0x91, 0x5b, 0xfa,
	0xe6, 0x7f, 0x60, 0xaa, 0xfe, 0x3a, 0x5f, 0x1d, 0xae, 0x1c, 0xa1, 0x60, 0x3c, 0x71, 0xa0, 0xee,
	0x5f, 0x66, 0x36, 0xb7, 0xc0, 0xae, 0x6d, 0xaf, 0xef, 0x77, 0x34, 0x29, 0x90, 0xac, 0x01, 0xec,
	0x1d, 0x0e, 0x5e, 0x1f, 0x0c, 0x3f, 0x44, 0xc3, 0x23, 0xbb, 0x45, 0x2c, 0x58, 0x7d, 0x73, 0x74,
	0x38, 0x0c, 0x6d, 0x63, 0xf3, 0x19, 0xf4, 0xe7, 0x5e, 0x25, 0x01, 0x30, 0x07, 0xb1, 0x62, 0x67,
	0x68, 0xb7, 0xaa, 0xf3, 0x73, 0xc1, 0x2f, 0x30, 0xb7, 0x0d, 0xd2, 0x87, 0x6e, 0x94, 0x52, 0x79,
	0x82, 0x89, 0xdd, 0x26, 0x37, 0xa0, 0xf7, 0x22, 0xa7, 0x75, 0x5b, 0x67, 0xf7, 0xe5, 0x97, 0x2b,
	0xd7, 0xb8, 0xbc, 0x72, 0x8d, 0x6f, 0x57, 0xae, 0xf1, 0x79, 0xea, 0xb6, 0x2e, 0xa7, 0x6e, 0xeb,
	0xeb, 0xd4, 0x6d, 0xbd, 0x7f, 0x74, 0xcc, 0x54, 0x4a, 0xc7, 0x41, 0xcc, 0xb3, 0xed, 0x3a, 0xa4,
	0xe6, 0xb3, 0xa5, 0xb7, 0xf7, 0x7c, 0x7b, 0x6e, 0x95, 0xd5, 0xa4, 0x40, 0x39, 0x36, 0xf5, 0x2a,
	0x3f, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xab, 0xe0, 0xed, 0xe5, 0x03, 0x00, 0x00,
}

func (m *Party) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Party) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Party) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReportedSessions) > 0 {
		for iNdEx := len(m.ReportedSessions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ReportedSessions[iNdEx])
			copy(dAtA[i:], m.ReportedSessions[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.ReportedSessions[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.CommittedGlobalPublicKey) > 0 {
		i -= len(m.CommittedGlobalPublicKey)
		copy(dAtA[i:], m.CommittedGlobalPublicKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.CommittedGlobalPublicKey)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Delegator) > 0 {
		i -= len(m.Delegator)
		copy(dAtA[i:], m.Delegator)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Delegator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.FreezeEndBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FreezeEndBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.ViolationsCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ViolationsCount))
		i--
		dAtA[i] = 0x28
	}
	if m.Status != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintParams(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if m.FreezeBlocksPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.FreezeBlocksPeriod))
		i--
		dAtA[i] = 0x50
	}
	if m.MaxViolationsCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxViolationsCount))
		i--
		dAtA[i] = 0x48
	}
	if len(m.StakeDenom) > 0 {
		i -= len(m.StakeDenom)
		copy(dAtA[i:], m.StakeDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.StakeDenom)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.StakeAmount) > 0 {
		i -= len(m.StakeAmount)
		copy(dAtA[i:], m.StakeAmount)
		i = encodeVarintParams(dAtA, i, uint64(len(m.StakeAmount)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastSignature) > 0 {
		i -= len(m.LastSignature)
		copy(dAtA[i:], m.LastSignature)
		i = encodeVarintParams(dAtA, i, uint64(len(m.LastSignature)))
		i--
		dAtA[i] = 0x32
	}
	if m.IsUpdateRequired {
		i--
		if m.IsUpdateRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Parties) > 0 {
		for iNdEx := len(m.Parties) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Parties[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x10
	}
	if len(m.KeyECDSA) > 0 {
		i -= len(m.KeyECDSA)
		copy(dAtA[i:], m.KeyECDSA)
		i = encodeVarintParams(dAtA, i, uint64(len(m.KeyECDSA)))
		i--
		dAtA[i] = 0xa
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
func (m *Party) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovParams(uint64(m.Status))
	}
	if m.ViolationsCount != 0 {
		n += 1 + sovParams(uint64(m.ViolationsCount))
	}
	if m.FreezeEndBlock != 0 {
		n += 1 + sovParams(uint64(m.FreezeEndBlock))
	}
	l = len(m.Delegator)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.CommittedGlobalPublicKey)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if len(m.ReportedSessions) > 0 {
		for _, s := range m.ReportedSessions {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.KeyECDSA)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Threshold != 0 {
		n += 1 + sovParams(uint64(m.Threshold))
	}
	if len(m.Parties) > 0 {
		for _, e := range m.Parties {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.IsUpdateRequired {
		n += 2
	}
	l = len(m.LastSignature)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.StakeAmount)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.StakeDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MaxViolationsCount != 0 {
		n += 1 + sovParams(uint64(m.MaxViolationsCount))
	}
	if m.FreezeBlocksPeriod != 0 {
		n += 1 + sovParams(uint64(m.FreezeBlocksPeriod))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Party) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Party: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Party: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= PartyStatus(b&0x7F) << shift
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
					return ErrIntOverflowParams
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
					return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delegator", wireType)
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
			m.Delegator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommittedGlobalPublicKey", wireType)
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
			m.CommittedGlobalPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReportedSessions", wireType)
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
			m.ReportedSessions = append(m.ReportedSessions, string(dAtA[iNdEx:postIndex]))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyECDSA", wireType)
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
			m.KeyECDSA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parties", wireType)
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
			m.Parties = append(m.Parties, &Party{})
			if err := m.Parties[len(m.Parties)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsUpdateRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsUpdateRequired = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSignature", wireType)
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
			m.LastSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAmount", wireType)
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
			m.StakeAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeDenom", wireType)
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
			m.StakeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxViolationsCount", wireType)
			}
			m.MaxViolationsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxViolationsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreezeBlocksPeriod", wireType)
			}
			m.FreezeBlocksPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FreezeBlocksPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
