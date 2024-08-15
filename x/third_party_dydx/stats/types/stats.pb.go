// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/stats/stats.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BlockStats is used to store stats transiently within the scope of a block.
type BlockStats struct {
	// The fills that occured on this block.
	Fills []*BlockStats_Fill `protobuf:"bytes,1,rep,name=fills,proto3" json:"fills,omitempty"`
}

func (m *BlockStats) Reset()         { *m = BlockStats{} }
func (m *BlockStats) String() string { return proto.CompactTextString(m) }
func (*BlockStats) ProtoMessage()    {}
func (*BlockStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{0}
}
func (m *BlockStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStats.Merge(m, src)
}
func (m *BlockStats) XXX_Size() int {
	return m.Size()
}
func (m *BlockStats) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStats.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStats proto.InternalMessageInfo

func (m *BlockStats) GetFills() []*BlockStats_Fill {
	if m != nil {
		return m.Fills
	}
	return nil
}

// Fill records data about a fill on this block.
type BlockStats_Fill struct {
	// Taker wallet address
	Taker string `protobuf:"bytes,1,opt,name=taker,proto3" json:"taker,omitempty"`
	// Maker wallet address
	Maker string `protobuf:"bytes,2,opt,name=maker,proto3" json:"maker,omitempty"`
	// Notional USDC filled in quantums
	Notional uint64 `protobuf:"varint,3,opt,name=notional,proto3" json:"notional,omitempty"`
}

func (m *BlockStats_Fill) Reset()         { *m = BlockStats_Fill{} }
func (m *BlockStats_Fill) String() string { return proto.CompactTextString(m) }
func (*BlockStats_Fill) ProtoMessage()    {}
func (*BlockStats_Fill) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{0, 0}
}
func (m *BlockStats_Fill) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockStats_Fill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockStats_Fill.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockStats_Fill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStats_Fill.Merge(m, src)
}
func (m *BlockStats_Fill) XXX_Size() int {
	return m.Size()
}
func (m *BlockStats_Fill) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStats_Fill.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStats_Fill proto.InternalMessageInfo

func (m *BlockStats_Fill) GetTaker() string {
	if m != nil {
		return m.Taker
	}
	return ""
}

func (m *BlockStats_Fill) GetMaker() string {
	if m != nil {
		return m.Maker
	}
	return ""
}

func (m *BlockStats_Fill) GetNotional() uint64 {
	if m != nil {
		return m.Notional
	}
	return 0
}

// StatsMetadata stores metadata for the x/stats module
type StatsMetadata struct {
	// The oldest epoch that is included in the stats. The next epoch to be
	// removed from the window.
	TrailingEpoch uint32 `protobuf:"varint,1,opt,name=trailing_epoch,json=trailingEpoch,proto3" json:"trailing_epoch,omitempty"`
}

func (m *StatsMetadata) Reset()         { *m = StatsMetadata{} }
func (m *StatsMetadata) String() string { return proto.CompactTextString(m) }
func (*StatsMetadata) ProtoMessage()    {}
func (*StatsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{1}
}
func (m *StatsMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatsMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsMetadata.Merge(m, src)
}
func (m *StatsMetadata) XXX_Size() int {
	return m.Size()
}
func (m *StatsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_StatsMetadata proto.InternalMessageInfo

func (m *StatsMetadata) GetTrailingEpoch() uint32 {
	if m != nil {
		return m.TrailingEpoch
	}
	return 0
}

// EpochStats stores stats for a particular epoch
type EpochStats struct {
	// Epoch end time
	EpochEndTime time.Time `protobuf:"bytes,1,opt,name=epoch_end_time,json=epochEndTime,proto3,stdtime" json:"epoch_end_time"`
	// Stats for each user in this epoch. Sorted by user.
	Stats []*EpochStats_UserWithStats `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (m *EpochStats) Reset()         { *m = EpochStats{} }
func (m *EpochStats) String() string { return proto.CompactTextString(m) }
func (*EpochStats) ProtoMessage()    {}
func (*EpochStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{2}
}
func (m *EpochStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochStats.Merge(m, src)
}
func (m *EpochStats) XXX_Size() int {
	return m.Size()
}
func (m *EpochStats) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochStats.DiscardUnknown(m)
}

var xxx_messageInfo_EpochStats proto.InternalMessageInfo

func (m *EpochStats) GetEpochEndTime() time.Time {
	if m != nil {
		return m.EpochEndTime
	}
	return time.Time{}
}

func (m *EpochStats) GetStats() []*EpochStats_UserWithStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// A user and its associated stats
type EpochStats_UserWithStats struct {
	User  string     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Stats *UserStats `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *EpochStats_UserWithStats) Reset()         { *m = EpochStats_UserWithStats{} }
func (m *EpochStats_UserWithStats) String() string { return proto.CompactTextString(m) }
func (*EpochStats_UserWithStats) ProtoMessage()    {}
func (*EpochStats_UserWithStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{2, 0}
}
func (m *EpochStats_UserWithStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochStats_UserWithStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochStats_UserWithStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochStats_UserWithStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochStats_UserWithStats.Merge(m, src)
}
func (m *EpochStats_UserWithStats) XXX_Size() int {
	return m.Size()
}
func (m *EpochStats_UserWithStats) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochStats_UserWithStats.DiscardUnknown(m)
}

var xxx_messageInfo_EpochStats_UserWithStats proto.InternalMessageInfo

func (m *EpochStats_UserWithStats) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *EpochStats_UserWithStats) GetStats() *UserStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// GlobalStats stores global stats
type GlobalStats struct {
	// Notional USDC traded in quantums
	NotionalTraded uint64 `protobuf:"varint,1,opt,name=notional_traded,json=notionalTraded,proto3" json:"notional_traded,omitempty"`
}

func (m *GlobalStats) Reset()         { *m = GlobalStats{} }
func (m *GlobalStats) String() string { return proto.CompactTextString(m) }
func (*GlobalStats) ProtoMessage()    {}
func (*GlobalStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{3}
}
func (m *GlobalStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GlobalStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GlobalStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GlobalStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalStats.Merge(m, src)
}
func (m *GlobalStats) XXX_Size() int {
	return m.Size()
}
func (m *GlobalStats) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalStats.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalStats proto.InternalMessageInfo

func (m *GlobalStats) GetNotionalTraded() uint64 {
	if m != nil {
		return m.NotionalTraded
	}
	return 0
}

// UserStats stores stats for a User
type UserStats struct {
	// Taker USDC in quantums
	TakerNotional uint64 `protobuf:"varint,1,opt,name=taker_notional,json=takerNotional,proto3" json:"taker_notional,omitempty"`
	// Maker USDC in quantums
	MakerNotional uint64 `protobuf:"varint,2,opt,name=maker_notional,json=makerNotional,proto3" json:"maker_notional,omitempty"`
}

func (m *UserStats) Reset()         { *m = UserStats{} }
func (m *UserStats) String() string { return proto.CompactTextString(m) }
func (*UserStats) ProtoMessage()    {}
func (*UserStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bd047160832bd95, []int{4}
}
func (m *UserStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStats.Merge(m, src)
}
func (m *UserStats) XXX_Size() int {
	return m.Size()
}
func (m *UserStats) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStats.DiscardUnknown(m)
}

var xxx_messageInfo_UserStats proto.InternalMessageInfo

func (m *UserStats) GetTakerNotional() uint64 {
	if m != nil {
		return m.TakerNotional
	}
	return 0
}

func (m *UserStats) GetMakerNotional() uint64 {
	if m != nil {
		return m.MakerNotional
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockStats)(nil), "joltify.third_party.dydxprotocol.stats.BlockStats")
	proto.RegisterType((*BlockStats_Fill)(nil), "joltify.third_party.dydxprotocol.stats.BlockStats.Fill")
	proto.RegisterType((*StatsMetadata)(nil), "joltify.third_party.dydxprotocol.stats.StatsMetadata")
	proto.RegisterType((*EpochStats)(nil), "joltify.third_party.dydxprotocol.stats.EpochStats")
	proto.RegisterType((*EpochStats_UserWithStats)(nil), "joltify.third_party.dydxprotocol.stats.EpochStats.UserWithStats")
	proto.RegisterType((*GlobalStats)(nil), "joltify.third_party.dydxprotocol.stats.GlobalStats")
	proto.RegisterType((*UserStats)(nil), "joltify.third_party.dydxprotocol.stats.UserStats")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/stats/stats.proto", fileDescriptor_5bd047160832bd95)
}

var fileDescriptor_5bd047160832bd95 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x6d, 0xba, 0x16, 0x6d, 0x5f, 0x49, 0x91, 0xa2, 0x1d, 0xaa, 0x1c, 0xd2, 0x2a, 0x12, 0xd0,
	0x0b, 0x8e, 0x28, 0xd2, 0xb8, 0xa2, 0x4a, 0x63, 0x02, 0x69, 0x3b, 0x84, 0x01, 0x82, 0x4b, 0xe4,
	0x26, 0x6e, 0x6a, 0xe6, 0xc4, 0x51, 0xe2, 0x4a, 0xeb, 0xbf, 0xd8, 0x2f, 0xe0, 0x0f, 0xf0, 0x47,
	0x76, 0xdc, 0x91, 0x13, 0xa0, 0xf6, 0x8f, 0x20, 0x7f, 0x6e, 0xda, 0xec, 0xd6, 0x4b, 0xe4, 0xf7,
	0xf2, 0xde, 0xf7, 0xec, 0x67, 0x19, 0x26, 0x3f, 0xa4, 0x50, 0x7c, 0xbe, 0x0a, 0xd4, 0x82, 0x97,
	0x49, 0x54, 0xd0, 0x52, 0xad, 0x82, 0x64, 0x95, 0xdc, 0x16, 0xa5, 0x54, 0x32, 0x96, 0x22, 0xa8,
	0x14, 0x55, 0x95, 0xf9, 0x12, 0x24, 0x9d, 0x17, 0x5b, 0x0f, 0x69, 0x78, 0x48, 0xd3, 0x43, 0x50,
	0xed, 0x9e, 0xa6, 0x32, 0x95, 0xc8, 0x05, 0x7a, 0x65, 0xdc, 0xee, 0x30, 0x95, 0x32, 0x15, 0x2c,
	0x40, 0x34, 0x5b, 0xce, 0x03, 0xc5, 0x33, 0x56, 0x29, 0x9a, 0x15, 0x46, 0xe0, 0xff, 0xb2, 0x00,
	0xa6, 0x42, 0xc6, 0x37, 0x9f, 0xf4, 0x14, 0xe7, 0x12, 0xba, 0x73, 0x2e, 0x44, 0x35, 0xb0, 0x46,
	0x47, 0xe3, 0xde, 0xe4, 0x2d, 0x39, 0x2c, 0x9d, 0xec, 0x47, 0x90, 0xf7, 0x5c, 0x88, 0xd0, 0x4c,
	0x71, 0xaf, 0xa0, 0xa3, 0xa1, 0x73, 0x0a, 0x5d, 0x45, 0x6f, 0x58, 0x39, 0xb0, 0x46, 0xd6, 0xf8,
	0x24, 0x34, 0x40, 0xb3, 0x19, 0xb2, 0x6d, 0xc3, 0x22, 0x70, 0x5c, 0x38, 0xce, 0xa5, 0xe2, 0x32,
	0xa7, 0x62, 0x70, 0x34, 0xb2, 0xc6, 0x9d, 0x70, 0x87, 0xfd, 0x33, 0xb0, 0x31, 0xe4, 0x92, 0x29,
	0x9a, 0x50, 0x45, 0x9d, 0xe7, 0xd0, 0x57, 0x25, 0xe5, 0x82, 0xe7, 0x69, 0xc4, 0x0a, 0x19, 0x2f,
	0x30, 0xc1, 0x0e, 0xed, 0x9a, 0x3d, 0xd7, 0xa4, 0xff, 0xb3, 0x0d, 0x80, 0x2b, 0x73, 0xca, 0x8f,
	0xd0, 0x47, 0x71, 0xc4, 0xf2, 0x24, 0xd2, 0x8d, 0xa0, 0xab, 0x37, 0x71, 0x89, 0xa9, 0x8b, 0xd4,
	0x75, 0x91, 0xeb, 0xba, 0xae, 0xe9, 0xf1, 0xfd, 0x9f, 0x61, 0xeb, 0xee, 0xef, 0xd0, 0x0a, 0x9f,
	0xa2, 0xf7, 0x3c, 0x4f, 0xf4, 0x4f, 0xe7, 0x0b, 0x74, 0xb1, 0x82, 0x41, 0x1b, 0x1b, 0x7b, 0x77,
	0x68, 0x63, 0xfb, 0xed, 0x90, 0xcf, 0x15, 0x2b, 0xbf, 0x72, 0x65, 0x50, 0x68, 0xc6, 0xb9, 0x02,
	0xec, 0x47, 0xbc, 0xe3, 0x40, 0x67, 0x59, 0xed, 0x2a, 0xc4, 0xb5, 0x73, 0xb1, 0x0f, 0xd7, 0xfb,
	0x7f, 0x7d, 0x68, 0xb8, 0x9e, 0xdc, 0x4c, 0xf3, 0xcf, 0xa0, 0x77, 0x21, 0xe4, 0x8c, 0x0a, 0x93,
	0xf5, 0x12, 0x9e, 0xd5, 0x9d, 0x47, 0xaa, 0xa4, 0x09, 0x4b, 0x30, 0xb6, 0x13, 0xf6, 0x6b, 0xfa,
	0x1a, 0x59, 0xff, 0x1b, 0x9c, 0xec, 0x66, 0xe1, 0x65, 0xe8, 0x2b, 0x8c, 0x76, 0xf7, 0x67, 0x4c,
	0x36, 0xb2, 0x57, 0x5b, 0x52, 0xcb, 0xb2, 0xc7, 0xb2, 0xb6, 0x91, 0x65, 0x4d, 0xd9, 0x34, 0xbe,
	0x5f, 0x7b, 0xd6, 0xc3, 0xda, 0xb3, 0xfe, 0xad, 0x3d, 0xeb, 0x6e, 0xe3, 0xb5, 0x1e, 0x36, 0x5e,
	0xeb, 0xf7, 0xc6, 0x6b, 0x7d, 0xff, 0x90, 0x72, 0xb5, 0x58, 0xce, 0x48, 0x2c, 0xb3, 0x60, 0x7b,
	0xe0, 0x57, 0x73, 0x9e, 0xd3, 0x3c, 0x66, 0x35, 0x8e, 0x04, 0xcb, 0x13, 0x9e, 0xa7, 0xc1, 0x6d,
	0xf3, 0xad, 0x45, 0xba, 0x8a, 0xed, 0x1b, 0x53, 0xab, 0x82, 0x55, 0xb3, 0x27, 0x58, 0xcc, 0x9b,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x09, 0x99, 0x2c, 0x9a, 0x03, 0x00, 0x00,
}

func (m *BlockStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fills) > 0 {
		for iNdEx := len(m.Fills) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fills[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BlockStats_Fill) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockStats_Fill) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockStats_Fill) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Notional != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.Notional))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Maker) > 0 {
		i -= len(m.Maker)
		copy(dAtA[i:], m.Maker)
		i = encodeVarintStats(dAtA, i, uint64(len(m.Maker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Taker) > 0 {
		i -= len(m.Taker)
		copy(dAtA[i:], m.Taker)
		i = encodeVarintStats(dAtA, i, uint64(len(m.Taker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatsMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatsMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatsMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TrailingEpoch != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.TrailingEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EpochStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Stats) > 0 {
		for iNdEx := len(m.Stats) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Stats[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.EpochEndTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EpochEndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintStats(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EpochStats_UserWithStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochStats_UserWithStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochStats_UserWithStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStats(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintStats(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GlobalStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GlobalStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GlobalStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NotionalTraded != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.NotionalTraded))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UserStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MakerNotional != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.MakerNotional))
		i--
		dAtA[i] = 0x10
	}
	if m.TakerNotional != 0 {
		i = encodeVarintStats(dAtA, i, uint64(m.TakerNotional))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Fills) > 0 {
		for _, e := range m.Fills {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func (m *BlockStats_Fill) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Taker)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	l = len(m.Maker)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.Notional != 0 {
		n += 1 + sovStats(uint64(m.Notional))
	}
	return n
}

func (m *StatsMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TrailingEpoch != 0 {
		n += 1 + sovStats(uint64(m.TrailingEpoch))
	}
	return n
}

func (m *EpochStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.EpochEndTime)
	n += 1 + l + sovStats(uint64(l))
	if len(m.Stats) > 0 {
		for _, e := range m.Stats {
			l = e.Size()
			n += 1 + l + sovStats(uint64(l))
		}
	}
	return n
}

func (m *EpochStats_UserWithStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovStats(uint64(l))
	}
	return n
}

func (m *GlobalStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NotionalTraded != 0 {
		n += 1 + sovStats(uint64(m.NotionalTraded))
	}
	return n
}

func (m *UserStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TakerNotional != 0 {
		n += 1 + sovStats(uint64(m.TakerNotional))
	}
	if m.MakerNotional != 0 {
		n += 1 + sovStats(uint64(m.MakerNotional))
	}
	return n
}

func sovStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: BlockStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fills", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fills = append(m.Fills, &BlockStats_Fill{})
			if err := m.Fills[len(m.Fills)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *BlockStats_Fill) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: Fill: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fill: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Taker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Taker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Maker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Notional", wireType)
			}
			m.Notional = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Notional |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *StatsMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: StatsMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatsMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrailingEpoch", wireType)
			}
			m.TrailingEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrailingEpoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *EpochStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: EpochStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.EpochEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stats = append(m.Stats, &EpochStats_UserWithStats{})
			if err := m.Stats[len(m.Stats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *EpochStats_UserWithStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: UserWithStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserWithStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &UserStats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *GlobalStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: GlobalStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GlobalStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotionalTraded", wireType)
			}
			m.NotionalTraded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotionalTraded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func (m *UserStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: UserStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TakerNotional", wireType)
			}
			m.TakerNotional = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TakerNotional |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MakerNotional", wireType)
			}
			m.MakerNotional = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MakerNotional |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStats = fmt.Errorf("proto: unexpected end of group")
)
