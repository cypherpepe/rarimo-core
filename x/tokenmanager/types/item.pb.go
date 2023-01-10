// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenmanager/item.proto

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

type Type int32

const (
	Type_NATIVE       Type = 0
	Type_ERC20        Type = 1
	Type_ERC721       Type = 2
	Type_ERC1155      Type = 3
	Type_METAPLEX_NFT Type = 4
	Type_METAPLEX_FT  Type = 5
	Type_NEAR_FT      Type = 6
	Type_NEAR_NFT     Type = 7
)

var Type_name = map[int32]string{
	0: "NATIVE",
	1: "ERC20",
	2: "ERC721",
	3: "ERC1155",
	4: "METAPLEX_NFT",
	5: "METAPLEX_FT",
	6: "NEAR_FT",
	7: "NEAR_NFT",
}

var Type_value = map[string]int32{
	"NATIVE":       0,
	"ERC20":        1,
	"ERC721":       2,
	"ERC1155":      3,
	"METAPLEX_NFT": 4,
	"METAPLEX_FT":  5,
	"NEAR_FT":      6,
	"NEAR_NFT":     7,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{0}
}

type ItemMetadata struct {
	Uri      string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	ImageUri string `protobuf:"bytes,2,opt,name=imageUri,proto3" json:"imageUri,omitempty"`
	// Hash of the token image. Encoded into hex string. (optional)
	ImageHash string `protobuf:"bytes,3,opt,name=imageHash,proto3" json:"imageHash,omitempty"`
}

func (m *ItemMetadata) Reset()         { *m = ItemMetadata{} }
func (m *ItemMetadata) String() string { return proto.CompactTextString(m) }
func (*ItemMetadata) ProtoMessage()    {}
func (*ItemMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{0}
}
func (m *ItemMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ItemMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ItemMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemMetadata.Merge(m, src)
}
func (m *ItemMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ItemMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ItemMetadata proto.InternalMessageInfo

func (m *ItemMetadata) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *ItemMetadata) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func (m *ItemMetadata) GetImageHash() string {
	if m != nil {
		return m.ImageHash
	}
	return ""
}

type ItemChainParams struct {
	TokenID   string `protobuf:"bytes,1,opt,name=tokenID,proto3" json:"tokenID,omitempty"`
	TokenType Type   `protobuf:"varint,2,opt,name=tokenType,proto3,enum=rarimo.rarimocore.tokenmanager.Type" json:"tokenType,omitempty"`
	// Seed for deriving address for Solana networks. Encoded into hex string. (optional, present only for solana networks)
	Seed string `protobuf:"bytes,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (m *ItemChainParams) Reset()         { *m = ItemChainParams{} }
func (m *ItemChainParams) String() string { return proto.CompactTextString(m) }
func (*ItemChainParams) ProtoMessage()    {}
func (*ItemChainParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{1}
}
func (m *ItemChainParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ItemChainParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ItemChainParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemChainParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemChainParams.Merge(m, src)
}
func (m *ItemChainParams) XXX_Size() int {
	return m.Size()
}
func (m *ItemChainParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemChainParams.DiscardUnknown(m)
}

var xxx_messageInfo_ItemChainParams proto.InternalMessageInfo

func (m *ItemChainParams) GetTokenID() string {
	if m != nil {
		return m.TokenID
	}
	return ""
}

func (m *ItemChainParams) GetTokenType() Type {
	if m != nil {
		return m.TokenType
	}
	return Type_NATIVE
}

func (m *ItemChainParams) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type Item struct {
	// hex-encoded
	Index       string                      `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Collection  string                      `protobuf:"bytes,2,opt,name=collection,proto3" json:"collection,omitempty"`
	ChainParams map[string]*ItemChainParams `protobuf:"bytes,3,rep,name=chainParams,proto3" json:"chainParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Item metadata
	Metadata *ItemMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{2}
}
func (m *Item) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Item.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return m.Size()
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Item) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Item) GetChainParams() map[string]*ItemChainParams {
	if m != nil {
		return m.ChainParams
	}
	return nil
}

func (m *Item) GetMetadata() *ItemMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ItemList struct {
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (m *ItemList) Reset()         { *m = ItemList{} }
func (m *ItemList) String() string { return proto.CompactTextString(m) }
func (*ItemList) ProtoMessage()    {}
func (*ItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83fe9c71c21d068, []int{3}
}
func (m *ItemList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ItemList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemList.Merge(m, src)
}
func (m *ItemList) XXX_Size() int {
	return m.Size()
}
func (m *ItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemList.DiscardUnknown(m)
}

var xxx_messageInfo_ItemList proto.InternalMessageInfo

func (m *ItemList) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("rarimo.rarimocore.tokenmanager.Type", Type_name, Type_value)
	proto.RegisterType((*ItemMetadata)(nil), "rarimo.rarimocore.tokenmanager.ItemMetadata")
	proto.RegisterType((*ItemChainParams)(nil), "rarimo.rarimocore.tokenmanager.ItemChainParams")
	proto.RegisterType((*Item)(nil), "rarimo.rarimocore.tokenmanager.Item")
	proto.RegisterMapType((map[string]*ItemChainParams)(nil), "rarimo.rarimocore.tokenmanager.Item.ChainParamsEntry")
	proto.RegisterType((*ItemList)(nil), "rarimo.rarimocore.tokenmanager.ItemList")
}

func init() { proto.RegisterFile("tokenmanager/item.proto", fileDescriptor_e83fe9c71c21d068) }

var fileDescriptor_e83fe9c71c21d068 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8a, 0x1a, 0x41,
	0x10, 0xb6, 0xfd, 0x59, 0xb5, 0x46, 0xb2, 0x43, 0x11, 0xc8, 0xb0, 0x84, 0x41, 0x64, 0x0f, 0x12,
	0x92, 0x31, 0x4e, 0x90, 0x84, 0xbd, 0xb9, 0x66, 0x64, 0x05, 0x57, 0x96, 0xc6, 0xfc, 0xb0, 0x97,
	0xd0, 0xab, 0x8d, 0xdb, 0xe8, 0xcc, 0xc8, 0x4c, 0x6f, 0x58, 0xcf, 0x79, 0x81, 0x3c, 0x56, 0x8e,
	0x1e, 0x73, 0x0c, 0xfa, 0x22, 0xa1, 0x7b, 0x46, 0x9d, 0xe4, 0x90, 0x78, 0xea, 0xaa, 0xea, 0xaa,
	0xef, 0xfb, 0xaa, 0xaa, 0x1b, 0x9e, 0xc9, 0x70, 0xce, 0x03, 0x9f, 0x05, 0x6c, 0xc6, 0xa3, 0x96,
	0x90, 0xdc, 0x77, 0x96, 0x51, 0x28, 0x43, 0xb4, 0x23, 0x16, 0x09, 0x3f, 0x74, 0x92, 0x63, 0x12,
	0x46, 0xdc, 0xc9, 0xa6, 0x36, 0x6e, 0xa1, 0x36, 0x90, 0xdc, 0xbf, 0xe6, 0x92, 0x4d, 0x99, 0x64,
	0x68, 0x42, 0xe1, 0x21, 0x12, 0x16, 0xa9, 0x93, 0x66, 0x95, 0x2a, 0x13, 0xcf, 0xa0, 0x22, 0x7c,
	0x36, 0xe3, 0x1f, 0x22, 0x61, 0xe5, 0x75, 0x78, 0xef, 0xe3, 0x73, 0xa8, 0x6a, 0xfb, 0x8a, 0xc5,
	0xf7, 0x56, 0x41, 0x5f, 0x1e, 0x02, 0x8d, 0x6f, 0x04, 0x4e, 0x15, 0x78, 0xef, 0x9e, 0x89, 0xe0,
	0x86, 0x45, 0xcc, 0x8f, 0xd1, 0x82, 0xb2, 0xe6, 0x1f, 0xbc, 0x4f, 0x39, 0x76, 0x2e, 0x5e, 0x42,
	0x55, 0x9b, 0xe3, 0xd5, 0x92, 0x6b, 0xa2, 0x27, 0xee, 0xb9, 0xf3, 0x6f, 0xf5, 0x8e, 0x5c, 0x2d,
	0x39, 0x3d, 0x94, 0x21, 0x42, 0x31, 0xe6, 0x7c, 0x9a, 0x4a, 0xd1, 0x76, 0x63, 0x9d, 0x87, 0xa2,
	0x52, 0x81, 0x4f, 0xa1, 0x24, 0x82, 0x29, 0x7f, 0x4c, 0x89, 0x13, 0x07, 0x6d, 0x80, 0x49, 0xb8,
	0x58, 0xf0, 0x89, 0x14, 0x61, 0x90, 0x36, 0x98, 0x89, 0xe0, 0x27, 0x30, 0x26, 0x07, 0xfd, 0x56,
	0xa1, 0x5e, 0x68, 0x1a, 0x6e, 0xe7, 0x7f, 0xc2, 0x14, 0xa1, 0x93, 0xe9, 0xdb, 0x0b, 0x64, 0xb4,
	0xa2, 0x59, 0x24, 0xbc, 0x82, 0x8a, 0x9f, 0x4e, 0xdd, 0x2a, 0xd6, 0x49, 0xd3, 0x70, 0x5f, 0x1e,
	0x83, 0xba, 0xdb, 0x14, 0xdd, 0x57, 0x9f, 0x85, 0x60, 0xfe, 0x4d, 0xa5, 0xf6, 0x38, 0xe7, 0xab,
	0xdd, 0x1e, 0xe7, 0x7c, 0x85, 0x1e, 0x94, 0xbe, 0xb2, 0xc5, 0x43, 0x32, 0x5b, 0xc3, 0x6d, 0x1d,
	0x43, 0x96, 0x81, 0xa5, 0x49, 0xf5, 0x45, 0xfe, 0x1d, 0x69, 0xf4, 0xa1, 0xa2, 0x6e, 0x87, 0x22,
	0x96, 0x78, 0x01, 0x25, 0xf5, 0xdc, 0x62, 0x8b, 0xe8, 0xc9, 0x9c, 0x1f, 0x03, 0x4b, 0x93, 0x92,
	0x17, 0x12, 0x8a, 0x6a, 0x83, 0x08, 0x70, 0x32, 0xea, 0x8e, 0x07, 0x1f, 0x3d, 0x33, 0x87, 0x55,
	0x28, 0x79, 0xb4, 0xe7, 0xbe, 0x36, 0x89, 0x0a, 0x7b, 0xb4, 0xf7, 0xd6, 0x6d, 0x9b, 0x79, 0x34,
	0xa0, 0xec, 0xd1, 0x5e, 0xbb, 0xdd, 0xe9, 0x98, 0x05, 0x34, 0xa1, 0x76, 0xed, 0x8d, 0xbb, 0x37,
	0x43, 0xef, 0xf3, 0x97, 0x51, 0x7f, 0x6c, 0x16, 0xf1, 0x14, 0x8c, 0x7d, 0xa4, 0x3f, 0x36, 0x4b,
	0x2a, 0x7f, 0xe4, 0x75, 0xa9, 0x72, 0x4e, 0xb0, 0x06, 0x15, 0xed, 0xa8, 0xdc, 0xf2, 0xe5, 0xf0,
	0xc7, 0xc6, 0x26, 0xeb, 0x8d, 0x4d, 0x7e, 0x6d, 0x6c, 0xf2, 0x7d, 0x6b, 0xe7, 0xd6, 0x5b, 0x3b,
	0xf7, 0x73, 0x6b, 0xe7, 0x6e, 0xdd, 0x99, 0x90, 0x0b, 0x76, 0xe7, 0x4c, 0x42, 0xbf, 0x95, 0xe8,
	0x4f, 0x8f, 0x57, 0xaa, 0x8f, 0xd6, 0x63, 0xeb, 0x8f, 0x5f, 0xa6, 0xa4, 0xc7, 0x77, 0x27, 0xfa,
	0x9f, 0xbd, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x03, 0xa9, 0xb8, 0x21, 0x82, 0x03, 0x00, 0x00,
}

func (m *ItemMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ItemMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ImageHash) > 0 {
		i -= len(m.ImageHash)
		copy(dAtA[i:], m.ImageHash)
		i = encodeVarintItem(dAtA, i, uint64(len(m.ImageHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ImageUri) > 0 {
		i -= len(m.ImageUri)
		copy(dAtA[i:], m.ImageUri)
		i = encodeVarintItem(dAtA, i, uint64(len(m.ImageUri)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ItemChainParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemChainParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ItemChainParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TokenType != 0 {
		i = encodeVarintItem(dAtA, i, uint64(m.TokenType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TokenID) > 0 {
		i -= len(m.TokenID)
		copy(dAtA[i:], m.TokenID)
		i = encodeVarintItem(dAtA, i, uint64(len(m.TokenID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Item) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Item) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Item) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintItem(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainParams) > 0 {
		for k := range m.ChainParams {
			v := m.ChainParams[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintItem(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintItem(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintItem(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Collection) > 0 {
		i -= len(m.Collection)
		copy(dAtA[i:], m.Collection)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Collection)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintItem(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ItemList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ItemList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ItemList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintItem(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintItem(dAtA []byte, offset int, v uint64) int {
	offset -= sovItem(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ItemMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.ImageUri)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.ImageHash)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *ItemChainParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenID)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if m.TokenType != 0 {
		n += 1 + sovItem(uint64(m.TokenType))
	}
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *Item) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	l = len(m.Collection)
	if l > 0 {
		n += 1 + l + sovItem(uint64(l))
	}
	if len(m.ChainParams) > 0 {
		for k, v := range m.ChainParams {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovItem(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovItem(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovItem(uint64(mapEntrySize))
		}
	}
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovItem(uint64(l))
	}
	return n
}

func (m *ItemList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovItem(uint64(l))
		}
	}
	return n
}

func sovItem(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozItem(x uint64) (n int) {
	return sovItem(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ItemMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: ItemMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ItemMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
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
func (m *ItemChainParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: ItemChainParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ItemChainParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenType", wireType)
			}
			m.TokenType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TokenType |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
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
func (m *Item) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: Item: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Item: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainParams == nil {
				m.ChainParams = make(map[string]*ItemChainParams)
			}
			var mapkey string
			var mapvalue *ItemChainParams
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowItem
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
							return ErrIntOverflowItem
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
						return ErrInvalidLengthItem
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthItem
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowItem
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthItem
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthItem
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ItemChainParams{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipItem(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthItem
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ChainParams[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ItemMetadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
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
func (m *ItemList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItem
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
			return fmt.Errorf("proto: ItemList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ItemList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItem
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
				return ErrInvalidLengthItem
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthItem
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Item{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItem(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItem
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
func skipItem(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
					return 0, ErrIntOverflowItem
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
				return 0, ErrInvalidLengthItem
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupItem
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthItem
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthItem        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItem          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupItem = fmt.Errorf("proto: unexpected end of group")
)
