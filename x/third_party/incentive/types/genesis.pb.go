// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/incentive/v1beta1/genesis.proto

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

// AccumulationTime stores the previous reward distribution time and its corresponding collateral type
type AccumulationTime struct {
	CollateralType           string    `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	PreviousAccumulationTime time.Time `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time"`
}

func (m *AccumulationTime) Reset()         { *m = AccumulationTime{} }
func (m *AccumulationTime) String() string { return proto.CompactTextString(m) }
func (*AccumulationTime) ProtoMessage()    {}
func (*AccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9c5f5bdd0c28d6, []int{0}
}
func (m *AccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccumulationTime.Merge(m, src)
}
func (m *AccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *AccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_AccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_AccumulationTime proto.InternalMessageInfo

// GenesisRewardState groups together the global state for a particular reward so it can be exported in genesis.
type GenesisRewardState struct {
	AccumulationTimes  AccumulationTimes  `protobuf:"bytes,1,rep,name=accumulation_times,json=accumulationTimes,proto3,castrepeated=AccumulationTimes" json:"accumulation_times"`
	MultiRewardIndexes MultiRewardIndexes `protobuf:"bytes,2,rep,name=multi_reward_indexes,json=multiRewardIndexes,proto3,castrepeated=MultiRewardIndexes" json:"multi_reward_indexes"`
}

func (m *GenesisRewardState) Reset()         { *m = GenesisRewardState{} }
func (m *GenesisRewardState) String() string { return proto.CompactTextString(m) }
func (*GenesisRewardState) ProtoMessage()    {}
func (*GenesisRewardState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9c5f5bdd0c28d6, []int{1}
}
func (m *GenesisRewardState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisRewardState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisRewardState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisRewardState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisRewardState.Merge(m, src)
}
func (m *GenesisRewardState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisRewardState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisRewardState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisRewardState proto.InternalMessageInfo

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Params                      Params                      `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	JoltSupplyRewardState       GenesisRewardState          `protobuf:"bytes,2,opt,name=jolt_supply_reward_state,json=joltSupplyRewardState,proto3" json:"jolt_supply_reward_state"`
	JoltBorrowRewardState       GenesisRewardState          `protobuf:"bytes,3,opt,name=jolt_borrow_reward_state,json=joltBorrowRewardState,proto3" json:"jolt_borrow_reward_state"`
	SwapRewardState             GenesisRewardState          `protobuf:"bytes,4,opt,name=swap_reward_state,json=swapRewardState,proto3" json:"swap_reward_state"`
	JoltLiquidityProviderClaims JoltLiquidityProviderClaims `protobuf:"bytes,5,rep,name=jolt_liquidity_provider_claims,json=joltLiquidityProviderClaims,proto3,castrepeated=JoltLiquidityProviderClaims" json:"jolt_liquidity_provider_claims"`
	SwapClaims                  SwapClaims                  `protobuf:"bytes,6,rep,name=swap_claims,json=swapClaims,proto3,castrepeated=SwapClaims" json:"swap_claims"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e9c5f5bdd0c28d6, []int{2}
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

func init() {
	proto.RegisterType((*AccumulationTime)(nil), "joltify.third_party.incentive.v1beta1.AccumulationTime")
	proto.RegisterType((*GenesisRewardState)(nil), "joltify.third_party.incentive.v1beta1.GenesisRewardState")
	proto.RegisterType((*GenesisState)(nil), "joltify.third_party.incentive.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("joltify/third_party/incentive/v1beta1/genesis.proto", fileDescriptor_9e9c5f5bdd0c28d6)
}

var fileDescriptor_9e9c5f5bdd0c28d6 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xc7, 0x9b, 0xfd, 0xd3, 0xef, 0xe7, 0x22, 0xc6, 0xac, 0x21, 0x85, 0x4e, 0x4a, 0xa7, 0x21,
	0xc4, 0x2e, 0x4b, 0x58, 0x77, 0x40, 0x1c, 0x17, 0x0e, 0xfc, 0x97, 0xa6, 0x6c, 0x27, 0x2e, 0x91,
	0x9b, 0xba, 0x99, 0x87, 0x13, 0x1b, 0xdb, 0x69, 0xd7, 0x23, 0x07, 0xee, 0x7b, 0x07, 0x70, 0x86,
	0x37, 0xb2, 0xe3, 0xb8, 0x71, 0x62, 0xb0, 0xbd, 0x11, 0x64, 0xc7, 0x29, 0x5d, 0xa7, 0x4a, 0x99,
	0xb4, 0x5b, 0x62, 0xfb, 0xfb, 0xfd, 0x7c, 0xfd, 0x3c, 0xd6, 0x03, 0x76, 0x8e, 0x18, 0x55, 0xa4,
	0x3f, 0x0a, 0xd4, 0x21, 0x11, 0xbd, 0x98, 0x23, 0xa1, 0x46, 0x01, 0xc9, 0x13, 0x9c, 0x2b, 0x32,
	0xc0, 0xc1, 0x60, 0xbb, 0x8b, 0x15, 0xda, 0x0e, 0x52, 0x9c, 0x63, 0x49, 0xa4, 0xcf, 0x05, 0x53,
	0x0c, 0x3e, 0xb2, 0x22, 0x7f, 0x42, 0xe4, 0x8f, 0x45, 0xbe, 0x15, 0xb5, 0x56, 0x53, 0x96, 0x32,
	0xa3, 0x08, 0xf4, 0x57, 0x29, 0x6e, 0x75, 0xea, 0x11, 0x13, 0x8a, 0x48, 0x26, 0x6f, 0xa6, 0xe1,
	0x48, 0xa0, 0xb1, 0xa6, 0x9d, 0x32, 0x96, 0x52, 0x1c, 0x98, 0xbf, 0x6e, 0xd1, 0x0f, 0x14, 0xc9,
	0xb0, 0x54, 0x28, 0xe3, 0xe5, 0x81, 0x8d, 0x2f, 0x0e, 0xb8, 0xb7, 0x9b, 0x24, 0x45, 0x56, 0x50,
	0xa4, 0x08, 0xcb, 0x0f, 0x48, 0x86, 0xe1, 0x63, 0xb0, 0x9c, 0x30, 0x4a, 0x91, 0xc2, 0x02, 0xd1,
	0x58, 0x8d, 0x38, 0x76, 0x9d, 0x75, 0x67, 0xf3, 0xff, 0xe8, 0xee, 0xbf, 0xe5, 0x83, 0x11, 0xc7,
	0xb0, 0x0b, 0x5a, 0x5c, 0xe0, 0x01, 0x61, 0x85, 0x8c, 0xd1, 0x84, 0x4b, 0xac, 0x31, 0xee, 0xdc,
	0xba, 0xb3, 0xd9, 0xec, 0xb4, 0xfc, 0x32, 0x83, 0x5f, 0x65, 0xf0, 0x0f, 0xaa, 0x0c, 0xe1, 0x7f,
	0xa7, 0xbf, 0xda, 0x8d, 0x93, 0xf3, 0xb6, 0x13, 0xb9, 0x95, 0xcf, 0x74, 0x98, 0x8d, 0xaf, 0x73,
	0x00, 0xbe, 0x28, 0x2b, 0x1f, 0xe1, 0x21, 0x12, 0xbd, 0x7d, 0x85, 0x14, 0x86, 0x9f, 0x1c, 0x00,
	0xaf, 0x21, 0xa5, 0xeb, 0xac, 0xcf, 0x6f, 0x36, 0x3b, 0x4f, 0xfd, 0x5a, 0xcd, 0xf1, 0xa7, 0x61,
	0xe1, 0x03, 0x1d, 0xe8, 0xdb, 0x79, 0x7b, 0x65, 0x7a, 0x47, 0x46, 0x2b, 0x68, 0x7a, 0x09, 0x7e,
	0x76, 0xc0, 0x6a, 0x56, 0x50, 0x45, 0x62, 0x61, 0x92, 0xc5, 0x24, 0xef, 0xe1, 0x63, 0x2c, 0xdd,
	0xb9, 0x1b, 0xa5, 0x78, 0xa7, 0x2d, 0xca, 0xbb, 0xbd, 0xd2, 0x06, 0x61, 0xcb, 0xa6, 0x80, 0xd3,
	0x3b, 0x58, 0x46, 0x30, 0xbb, 0xb6, 0xb6, 0xf1, 0x63, 0x11, 0xdc, 0xb1, 0x25, 0x2a, 0x8b, 0xf3,
	0x06, 0x2c, 0x95, 0xcf, 0xc0, 0xf4, 0xad, 0xd9, 0xd9, 0xaa, 0x99, 0x64, 0xcf, 0x88, 0xc2, 0x05,
	0xcd, 0x8f, 0xac, 0x05, 0x3c, 0x06, 0xae, 0x56, 0xc7, 0xb2, 0xe0, 0x9c, 0x8e, 0xaa, 0xab, 0x4a,
	0x0d, 0xb2, 0x2d, 0x7e, 0x56, 0xd3, 0xfe, 0x7a, 0x1b, 0x2d, 0xea, 0xbe, 0xd6, 0xef, 0x1b, 0xff,
	0xc9, 0x1e, 0x57, 0xe4, 0x2e, 0x13, 0x82, 0x0d, 0xaf, 0x92, 0xe7, 0x6f, 0x91, 0x1c, 0x1a, 0xff,
	0x49, 0xf2, 0x07, 0xb0, 0x22, 0x87, 0x88, 0x5f, 0x45, 0x2e, 0xdc, 0x0e, 0x72, 0x59, 0x3b, 0x4f,
	0xc2, 0xbe, 0x3b, 0xc0, 0x33, 0xf7, 0xa4, 0xe4, 0x63, 0x41, 0x7a, 0x44, 0x8d, 0x62, 0x2e, 0xd8,
	0x80, 0xf4, 0xb0, 0x88, 0xcb, 0x09, 0xe0, 0x2e, 0x9a, 0x07, 0xb5, 0x5b, 0x13, 0xfd, 0x9a, 0x51,
	0xf5, 0xb6, 0xf2, 0xda, 0xb3, 0x56, 0xcf, 0xb5, 0x53, 0xf8, 0xd0, 0x3e, 0xad, 0xb5, 0xd9, 0x67,
	0x64, 0xb4, 0x76, 0x34, 0x7b, 0x13, 0x62, 0xd0, 0x34, 0xa5, 0xb1, 0xc9, 0x96, 0x4c, 0xb2, 0x27,
	0x35, 0x93, 0xed, 0x0f, 0x11, 0x2f, 0x83, 0x40, 0x1b, 0x04, 0x8c, 0x97, 0x64, 0x04, 0xe4, 0xf8,
	0x3b, 0xec, 0x9f, 0xfe, 0xf1, 0x1a, 0xa7, 0x17, 0x9e, 0x73, 0x76, 0xe1, 0x39, 0xbf, 0x2f, 0x3c,
	0xe7, 0xe4, 0xd2, 0x6b, 0x9c, 0x5d, 0x7a, 0x8d, 0x9f, 0x97, 0x5e, 0xe3, 0xfd, 0xcb, 0x94, 0xa8,
	0xc3, 0xa2, 0xeb, 0x27, 0x2c, 0x0b, 0x2c, 0x79, 0xab, 0x4f, 0x72, 0x94, 0x27, 0xb8, 0xfa, 0x8f,
	0x29, 0xce, 0x7b, 0x24, 0x4f, 0x83, 0xe3, 0x19, 0x03, 0x53, 0xcf, 0x35, 0xd9, 0x5d, 0x32, 0x63,
	0x69, 0xe7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xcf, 0x36, 0x70, 0x04, 0x06, 0x00, 0x00,
}

func (m *AccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisRewardState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisRewardState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisRewardState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MultiRewardIndexes) > 0 {
		for iNdEx := len(m.MultiRewardIndexes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MultiRewardIndexes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.AccumulationTimes) > 0 {
		for iNdEx := len(m.AccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
	if len(m.SwapClaims) > 0 {
		for iNdEx := len(m.SwapClaims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SwapClaims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.JoltLiquidityProviderClaims) > 0 {
		for iNdEx := len(m.JoltLiquidityProviderClaims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.JoltLiquidityProviderClaims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.SwapRewardState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.JoltBorrowRewardState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.JoltSupplyRewardState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
func (m *AccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisRewardState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccumulationTimes) > 0 {
		for _, e := range m.AccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MultiRewardIndexes) > 0 {
		for _, e := range m.MultiRewardIndexes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.JoltSupplyRewardState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.JoltBorrowRewardState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SwapRewardState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.JoltLiquidityProviderClaims) > 0 {
		for _, e := range m.JoltLiquidityProviderClaims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SwapClaims) > 0 {
		for _, e := range m.SwapClaims {
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
func (m *AccumulationTime) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisRewardState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisRewardState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisRewardState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccumulationTimes", wireType)
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
			m.AccumulationTimes = append(m.AccumulationTimes, AccumulationTime{})
			if err := m.AccumulationTimes[len(m.AccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiRewardIndexes", wireType)
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
			m.MultiRewardIndexes = append(m.MultiRewardIndexes, MultiRewardIndex{})
			if err := m.MultiRewardIndexes[len(m.MultiRewardIndexes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field JoltSupplyRewardState", wireType)
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
			if err := m.JoltSupplyRewardState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoltBorrowRewardState", wireType)
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
			if err := m.JoltBorrowRewardState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapRewardState", wireType)
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
			if err := m.SwapRewardState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoltLiquidityProviderClaims", wireType)
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
			m.JoltLiquidityProviderClaims = append(m.JoltLiquidityProviderClaims, JoltLiquidityProviderClaim{})
			if err := m.JoltLiquidityProviderClaims[len(m.JoltLiquidityProviderClaims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapClaims", wireType)
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
			m.SwapClaims = append(m.SwapClaims, SwapClaim{})
			if err := m.SwapClaims[len(m.SwapClaims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
