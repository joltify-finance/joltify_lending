// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/outbound_tx.proto

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
	OutboundTxs     []string                                      `protobuf:"bytes,3,rep,name=outboundTxs,proto3" json:"outboundTxs,omitempty"`
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

func (m *OutboundTx) GetOutboundTxs() []string {
	if m != nil {
		return m.OutboundTxs
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
	proto.RegisterType((*Entity)(nil), "joltify.vault.Entity")
	proto.RegisterType((*Proposals)(nil), "joltify.vault.Proposals")
	proto.RegisterType((*OutboundTx)(nil), "joltify.vault.OutboundTx")
}

func init() { proto.RegisterFile("joltify/vault/outbound_tx.proto", fileDescriptor_76d8c68ba70de65f) }

var fileDescriptor_76d8c68ba70de65f = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0x6e, 0x5a, 0xfb, 0x35, 0xab, 0x08, 0xc3, 0x0a, 0x63, 0x91, 0x34, 0xf4, 0x14, 0x90, 0xce,
	0x58, 0xbd, 0xec, 0x75, 0x2b, 0x82, 0x20, 0xa2, 0x84, 0x9e, 0xf4, 0xb0, 0x24, 0x93, 0xb7, 0xed,
	0x68, 0x77, 0x26, 0x64, 0xa6, 0x25, 0xf9, 0x17, 0xfe, 0x0e, 0x7f, 0x83, 0x3f, 0x60, 0x8f, 0x7b,
	0xf4, 0xb4, 0x4a, 0xfb, 0x2f, 0xc4, 0x83, 0x64, 0x26, 0xed, 0x56, 0x4f, 0x22, 0x78, 0x4a, 0xde,
	0x8f, 0xe7, 0x79, 0xde, 0xf7, 0x7d, 0x12, 0x34, 0xfc, 0xa0, 0x56, 0x46, 0xcc, 0x4b, 0xb6, 0x89,
	0xd7, 0x2b, 0xc3, 0xd4, 0xda, 0x24, 0x6a, 0x2d, 0xd3, 0x0b, 0x53, 0xd0, 0x2c, 0x57, 0x46, 0xe1,
	0x7b, 0x75, 0x03, 0xb5, 0x0d, 0x83, 0xd3, 0x85, 0x5a, 0x28, 0x5b, 0x61, 0xd5, 0x9b, 0x6b, 0x1a,
	0xf8, 0x5c, 0xe9, 0x4b, 0xa5, 0x59, 0x12, 0x6b, 0x60, 0x9b, 0x49, 0x02, 0x26, 0x9e, 0x30, 0xae,
	0x84, 0x74, 0xf5, 0xd1, 0x17, 0x0f, 0x75, 0x5e, 0x48, 0x23, 0x4c, 0x89, 0x5f, 0xa1, 0x6e, 0x9c,
	0xa6, 0x39, 0x68, 0x4d, 0xbc, 0xc0, 0x0b, 0xef, 0x4e, 0x27, 0x3f, 0x6e, 0x86, 0xe3, 0x85, 0x30,
	0xcb, 0x75, 0x42, 0xb9, 0xba, 0x64, 0x35, 0x95, 0x7b, 0x8c, 0x75, 0xfa, 0x91, 0x99, 0x32, 0x03,
	0x4d, 0xcf, 0x39, 0x3f, 0x77, 0xc0, 0x68, 0xcf, 0x80, 0x01, 0x75, 0xe7, 0x00, 0x95, 0x10, 0x69,
	0x06, 0xad, 0xf0, 0xe4, 0xe9, 0x43, 0xea, 0x70, 0xb4, 0x9a, 0x84, 0xd6, 0x93, 0xd0, 0xe7, 0x4a,
	0xc8, 0xe9, 0x93, 0xab, 0x9b, 0x61, 0xe3, 0xf3, 0xb7, 0x61, 0xf8, 0x17, 0x5a, 0x15, 0x40, 0x47,
	0x7b, 0xee, 0xd1, 0x19, 0xea, 0xbf, 0xcd, 0x55, 0xa6, 0x74, 0xbc, 0xd2, 0xf8, 0x31, 0x6a, 0x83,
	0x34, 0x79, 0x49, 0x3c, 0xab, 0xf8, 0x80, 0xfe, 0x76, 0x20, 0xea, 0xd6, 0x8c, 0x5c, 0xcf, 0xe8,
	0x67, 0x13, 0xa1, 0x37, 0xf5, 0x4d, 0x67, 0x05, 0x3e, 0x45, 0x6d, 0x21, 0x53, 0x28, 0xec, 0xea,
	0xfd, 0xc8, 0x05, 0xf8, 0x11, 0xea, 0x67, 0xb9, 0xe2, 0xa0, 0x35, 0xa4, 0xa4, 0x19, 0x78, 0x61,
	0x2f, 0xba, 0x4d, 0xe0, 0x00, 0x9d, 0xa8, 0x03, 0x83, 0x26, 0xad, 0xa0, 0x15, 0xf6, 0xa3, 0xe3,
	0x54, 0x85, 0xe7, 0xcb, 0x58, 0xc8, 0x59, 0x99, 0x01, 0xb9, 0x63, 0x99, 0x6f, 0x13, 0x78, 0x80,
	0x7a, 0x42, 0xce, 0x8a, 0x97, 0xb1, 0x5e, 0x92, 0xb6, 0x2d, 0x1e, 0x62, 0xfc, 0x1e, 0xdd, 0xcf,
	0x81, 0x83, 0xd8, 0x40, 0x5e, 0xdf, 0x96, 0x74, 0xfe, 0xd5, 0x94, 0x3f, 0x99, 0x2a, 0x61, 0x09,
	0x90, 0xbe, 0x16, 0xd2, 0x90, 0xae, 0xdd, 0xea, 0x10, 0x1f, 0x1b, 0xd7, 0xfb, 0x7f, 0xc6, 0x4d,
	0xa3, 0xab, 0xad, 0xef, 0x5d, 0x6f, 0x7d, 0xef, 0xfb, 0xd6, 0xf7, 0x3e, 0xed, 0xfc, 0xc6, 0xf5,
	0xce, 0x6f, 0x7c, 0xdd, 0xf9, 0x8d, 0x77, 0x67, 0x47, 0x64, 0xb5, 0x81, 0xe3, 0xb9, 0x90, 0xb1,
	0xe4, 0xb0, 0x8f, 0x2f, 0x56, 0x20, 0x53, 0x21, 0x17, 0xac, 0xa8, 0x7f, 0x0e, 0x2b, 0x91, 0x74,
	0xec, 0x27, 0xfd, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x94, 0x9d, 0xc4, 0xfb, 0x3a, 0x03,
	0x00, 0x00,
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
	if len(m.OutboundTxs) > 0 {
		for iNdEx := len(m.OutboundTxs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OutboundTxs[iNdEx])
			copy(dAtA[i:], m.OutboundTxs[iNdEx])
			i = encodeVarintOutboundTx(dAtA, i, uint64(len(m.OutboundTxs[iNdEx])))
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
	if len(m.OutboundTxs) > 0 {
		for _, s := range m.OutboundTxs {
			l = len(s)
			n += 1 + l + sovOutboundTx(uint64(l))
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
			return fmt.Errorf("proto: Entity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entity: illegal tag %d (wire type %d)", fieldNum, wire)
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
			return fmt.Errorf("proto: Proposals: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposals: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTxs", wireType)
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
			m.OutboundTxs = append(m.OutboundTxs, string(dAtA[iNdEx:postIndex]))
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
