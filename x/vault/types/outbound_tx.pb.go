// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/outbound_tx.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Entity struct {
	Address github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"address,omitempty"`
	Feecoin github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,2,rep,name=feecoin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"feecoin"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_76d8c68ba70de65f, []int{0}
}
func (m *Entity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return m.Size()
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Entity) GetFeecoin() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Feecoin
	}
	return nil
}

type Proposals struct {
	Entry []*Entity `protobuf:"bytes,1,rep,name=entry,proto3" json:"entry,omitempty"`
}

func (m *Proposals) Reset()         { *m = Proposals{} }
func (m *Proposals) String() string { return proto.CompactTextString(m) }
func (*Proposals) ProtoMessage()    {}
func (*Proposals) Descriptor() ([]byte, []int) {
	return fileDescriptor_76d8c68ba70de65f, []int{1}
}
func (m *Proposals) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposals.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposals.Merge(m, src)
}
func (m *Proposals) XXX_Size() int {
	return m.Size()
}
func (m *Proposals) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposals.DiscardUnknown(m)
}

var xxx_messageInfo_Proposals proto.InternalMessageInfo

func (m *Proposals) GetEntry() []*Entity {
	if m != nil {
		return m.Entry
	}
	return nil
}

type OutboundTx struct {
	Index           string                                        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Processed       bool                                          `protobuf:"varint,2,opt,name=processed,proto3" json:"processed,omitempty"`
	Items           map[string]Proposals                          `protobuf:"bytes,3,rep,name=items,proto3" json:"items" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ChainType       string                                        `protobuf:"bytes,4,opt,name=chainType,proto3" json:"chainType,omitempty"`
	InTxHash        string                                        `protobuf:"bytes,5,opt,name=inTxHash,proto3" json:"inTxHash,omitempty"`
	ReceiverAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,6,opt,name=receiverAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"receiverAddress,omitempty"`
	NeedMint        bool                                          `protobuf:"varint,7,opt,name=needMint,proto3" json:"needMint,omitempty"`
	Feecoin         github_com_cosmos_cosmos_sdk_types.Coins      `protobuf:"bytes,8,rep,name=feecoin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"feecoin"`
}

func (m *OutboundTx) Reset()         { *m = OutboundTx{} }
func (m *OutboundTx) String() string { return proto.CompactTextString(m) }
func (*OutboundTx) ProtoMessage()    {}
func (*OutboundTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_76d8c68ba70de65f, []int{2}
}
func (m *OutboundTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutboundTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutboundTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutboundTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutboundTx.Merge(m, src)
}
func (m *OutboundTx) XXX_Size() int {
	return m.Size()
}
func (m *OutboundTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutboundTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutboundTx proto.InternalMessageInfo

func (m *OutboundTx) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *OutboundTx) GetProcessed() bool {
	if m != nil {
		return m.Processed
	}
	return false
}

func (m *OutboundTx) GetItems() map[string]Proposals {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OutboundTx) GetChainType() string {
	if m != nil {
		return m.ChainType
	}
	return ""
}

func (m *OutboundTx) GetInTxHash() string {
	if m != nil {
		return m.InTxHash
	}
	return ""
}

func (m *OutboundTx) GetReceiverAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.ReceiverAddress
	}
	return nil
}

func (m *OutboundTx) GetNeedMint() bool {
	if m != nil {
		return m.NeedMint
	}
	return false
}

func (m *OutboundTx) GetFeecoin() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Feecoin
	}
	return nil
}

func init() {
	proto.RegisterType((*Entity)(nil), "joltify.vault.entity")
	proto.RegisterType((*Proposals)(nil), "joltify.vault.proposals")
	proto.RegisterType((*OutboundTx)(nil), "joltify.vault.OutboundTx")
	proto.RegisterMapType((map[string]Proposals)(nil), "joltify.vault.OutboundTx.ItemsEntry")
}

func init() { proto.RegisterFile("joltify/vault/outbound_tx.proto", fileDescriptor_76d8c68ba70de65f) }

var fileDescriptor_76d8c68ba70de65f = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x6e, 0xda, 0xed, 0xcf, 0x7a, 0x41, 0x20, 0x6b, 0x91, 0x4c, 0x85, 0xd2, 0xaa, 0xe2, 0x50,
	0x09, 0xd5, 0xa1, 0xcb, 0xa5, 0xe2, 0xc4, 0x16, 0x21, 0x81, 0x10, 0x42, 0x8a, 0x7a, 0x82, 0xc3,
	0x2a, 0x4d, 0xa6, 0xad, 0xd9, 0xd4, 0x8e, 0x62, 0xa7, 0x4a, 0xde, 0x82, 0xe7, 0xe0, 0x19, 0x78,
	0x80, 0x3d, 0xee, 0x91, 0xd3, 0x82, 0xda, 0x17, 0xe0, 0xcc, 0x09, 0xd9, 0x4e, 0xdb, 0xa5, 0x27,
	0x84, 0xc4, 0x29, 0x1e, 0xcf, 0x37, 0xdf, 0x37, 0x33, 0x5f, 0x8c, 0x3a, 0x9f, 0x44, 0xac, 0xd8,
	0xac, 0xf0, 0x56, 0x41, 0x16, 0x2b, 0x4f, 0x64, 0x6a, 0x2a, 0x32, 0x1e, 0x5d, 0xa8, 0x9c, 0x26,
	0xa9, 0x50, 0x02, 0xdf, 0x2d, 0x01, 0xd4, 0x00, 0xda, 0xa7, 0x73, 0x31, 0x17, 0x26, 0xe3, 0xe9,
	0x93, 0x05, 0xb5, 0xdd, 0x50, 0xc8, 0xa5, 0x90, 0xde, 0x34, 0x90, 0xe0, 0xad, 0x86, 0x53, 0x50,
	0xc1, 0xd0, 0x0b, 0x05, 0xe3, 0x36, 0xdf, 0xfb, 0xea, 0xa0, 0x06, 0x70, 0xc5, 0x54, 0x81, 0xdf,
	0xa2, 0x66, 0x10, 0x45, 0x29, 0x48, 0x49, 0x9c, 0xae, 0xd3, 0xbf, 0x33, 0x1e, 0xfe, 0xba, 0xe9,
	0x0c, 0xe6, 0x4c, 0x2d, 0xb2, 0x29, 0x0d, 0xc5, 0xd2, 0x2b, 0xa9, 0xec, 0x67, 0x20, 0xa3, 0x4b,
	0x4f, 0x15, 0x09, 0x48, 0x7a, 0x1e, 0x86, 0xe7, 0xb6, 0xd0, 0xdf, 0x32, 0x60, 0x40, 0xcd, 0x19,
	0x80, 0x16, 0x22, 0xd5, 0x6e, 0xad, 0x7f, 0x72, 0xf6, 0x90, 0xda, 0x3a, 0xaa, 0x3b, 0xa1, 0x65,
	0x27, 0xf4, 0xa5, 0x60, 0x7c, 0xfc, 0xf4, 0xea, 0xa6, 0x53, 0xf9, 0xf2, 0xbd, 0xd3, 0xff, 0x0b,
	0x2d, 0x5d, 0x20, 0xfd, 0x2d, 0x77, 0x6f, 0x84, 0x8e, 0x93, 0x54, 0x24, 0x42, 0x06, 0xb1, 0xc4,
	0x4f, 0x50, 0x1d, 0xb8, 0x4a, 0x0b, 0xe2, 0x18, 0xc5, 0x07, 0xf4, 0x8f, 0x05, 0x51, 0x3b, 0xa6,
	0x6f, 0x31, 0xbd, 0x9f, 0x35, 0x84, 0xde, 0x97, 0x3b, 0x9d, 0xe4, 0xf8, 0x14, 0xd5, 0x19, 0x8f,
	0x20, 0x37, 0xa3, 0x1f, 0xfb, 0x36, 0xc0, 0x8f, 0x0c, 0x7d, 0x08, 0x52, 0x42, 0x44, 0xaa, 0x5d,
	0xa7, 0xdf, 0xf2, 0xf7, 0x17, 0xf8, 0x05, 0xaa, 0x33, 0x05, 0x4b, 0x49, 0x6a, 0x46, 0xef, 0xf1,
	0x81, 0xde, 0x9e, 0x9d, 0xbe, 0xd1, 0xb0, 0x57, 0x5a, 0x77, 0x7c, 0xa4, 0x87, 0xf5, 0x6d, 0xa1,
	0xe6, 0x0f, 0x17, 0x01, 0xe3, 0x93, 0x22, 0x01, 0x72, 0x64, 0x94, 0xf7, 0x17, 0xb8, 0x8d, 0x5a,
	0x8c, 0x4f, 0xf2, 0xd7, 0x81, 0x5c, 0x90, 0xba, 0x49, 0xee, 0x62, 0xfc, 0x11, 0xdd, 0x4b, 0x21,
	0x04, 0xb6, 0x82, 0xb4, 0xdc, 0x3d, 0x69, 0xfc, 0xab, 0x69, 0x87, 0x4c, 0x5a, 0x98, 0x03, 0x44,
	0xef, 0x18, 0x57, 0xa4, 0x69, 0xa6, 0xde, 0xc5, 0xb7, 0x8d, 0x6d, 0xfd, 0x3f, 0x63, 0xdb, 0x3e,
	0x42, 0xfb, 0xa5, 0xe1, 0xfb, 0xa8, 0x76, 0x09, 0x45, 0xe9, 0x8d, 0x3e, 0x62, 0x8a, 0xea, 0xab,
	0x20, 0xce, 0xc0, 0xb8, 0x72, 0x72, 0x46, 0x0e, 0x76, 0xbf, 0xfb, 0x29, 0x7c, 0x0b, 0x7b, 0x5e,
	0x1d, 0x39, 0x63, 0xff, 0x6a, 0xed, 0x3a, 0xd7, 0x6b, 0xd7, 0xf9, 0xb1, 0x76, 0x9d, 0xcf, 0x1b,
	0xb7, 0x72, 0xbd, 0x71, 0x2b, 0xdf, 0x36, 0x6e, 0xe5, 0xc3, 0xe8, 0x56, 0x83, 0x25, 0xd1, 0x60,
	0xc6, 0x78, 0xc0, 0x43, 0xd8, 0xc6, 0x17, 0x31, 0xf0, 0x88, 0xf1, 0xb9, 0x97, 0x97, 0x0f, 0xd2,
	0xb4, 0x3d, 0x6d, 0x98, 0x67, 0xf4, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x25, 0xef,
	0xf4, 0xae, 0x03, 0x00, 0x00,
}

func (m *Entity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feecoin) > 0 {
		for iNdEx := len(m.Feecoin) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Feecoin[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proposals) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposals) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposals) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entry) > 0 {
		for iNdEx := len(m.Entry) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entry[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OutboundTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutboundTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutboundTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Feecoin) > 0 {
		for iNdEx := len(m.Feecoin) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Feecoin[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.NeedMint {
		i--
		if m.NeedMint {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.ReceiverAddress) > 0 {
		i -= len(m.ReceiverAddress)
		copy(dAtA[i:], m.ReceiverAddress)
		i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.ReceiverAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.InTxHash) > 0 {
		i -= len(m.InTxHash)
		copy(dAtA[i:], m.InTxHash)
		i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.InTxHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ChainType) > 0 {
		i -= len(m.ChainType)
		copy(dAtA[i:], m.ChainType)
		i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.ChainType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Items) > 0 {
		for k := range m.Items {
			v := m.Items[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOutboundTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintOutboundTx(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintOutboundTx(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Processed {
		i--
		if m.Processed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutboundTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutboundTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Entity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovOutboundTx(uint64(l))
	}
	if len(m.Feecoin) > 0 {
		for _, e := range m.Feecoin {
			l = e.Size()
			n += 1 + l + sovOutboundTx(uint64(l))
		}
	}
	return n
}

func (m *Proposals) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Entry) > 0 {
		for _, e := range m.Entry {
			l = e.Size()
			n += 1 + l + sovOutboundTx(uint64(l))
		}
	}
	return n
}

func (m *OutboundTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOutboundTx(uint64(l))
	}
	if m.Processed {
		n += 2
	}
	if len(m.Items) > 0 {
		for k, v := range m.Items {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovOutboundTx(uint64(len(k))) + 1 + l + sovOutboundTx(uint64(l))
			n += mapEntrySize + 1 + sovOutboundTx(uint64(mapEntrySize))
		}
	}
	l = len(m.ChainType)
	if l > 0 {
		n += 1 + l + sovOutboundTx(uint64(l))
	}
	l = len(m.InTxHash)
	if l > 0 {
		n += 1 + l + sovOutboundTx(uint64(l))
	}
	l = len(m.ReceiverAddress)
	if l > 0 {
		n += 1 + l + sovOutboundTx(uint64(l))
	}
	if m.NeedMint {
		n += 2
	}
	if len(m.Feecoin) > 0 {
		for _, e := range m.Feecoin {
			l = e.Size()
			n += 1 + l + sovOutboundTx(uint64(l))
		}
	}
	return n
}

func sovOutboundTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutboundTx(x uint64) (n int) {
	return sovOutboundTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundTx
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
			return fmt.Errorf("proto: entity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: entity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feecoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feecoin = append(m.Feecoin, types.Coin{})
			if err := m.Feecoin[len(m.Feecoin)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundTx
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
func (m *Proposals) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundTx
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
			return fmt.Errorf("proto: proposals: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: proposals: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entry = append(m.Entry, &Entity{})
			if err := m.Entry[len(m.Entry)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundTx
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
func (m *OutboundTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutboundTx
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
			return fmt.Errorf("proto: OutboundTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutboundTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Processed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
			m.Processed = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Items == nil {
				m.Items = make(map[string]Proposals)
			}
			var mapkey string
			mapvalue := &Proposals{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowOutboundTx
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
							return ErrIntOverflowOutboundTx
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
						return ErrInvalidLengthOutboundTx
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthOutboundTx
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
							return ErrIntOverflowOutboundTx
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
						return ErrInvalidLengthOutboundTx
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthOutboundTx
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Proposals{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipOutboundTx(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthOutboundTx
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Items[mapkey] = *mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InTxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InTxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverAddress = append(m.ReceiverAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ReceiverAddress == nil {
				m.ReceiverAddress = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NeedMint", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
			m.NeedMint = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feecoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutboundTx
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
				return ErrInvalidLengthOutboundTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOutboundTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Feecoin = append(m.Feecoin, types.Coin{})
			if err := m.Feecoin[len(m.Feecoin)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutboundTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutboundTx
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
func skipOutboundTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutboundTx
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
					return 0, ErrIntOverflowOutboundTx
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
					return 0, ErrIntOverflowOutboundTx
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
				return 0, ErrInvalidLengthOutboundTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOutboundTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOutboundTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOutboundTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutboundTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOutboundTx = fmt.Errorf("proto: unexpected end of group")
)
