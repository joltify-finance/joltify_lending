// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/jolt/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// GenesisState defines the jolt module's genesis state.
type GenesisState struct {
	Params                    Params                                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PreviousAccumulationTimes GenesisAccumulationTimes                 `protobuf:"bytes,2,rep,name=previous_accumulation_times,json=previousAccumulationTimes,proto3,castrepeated=GenesisAccumulationTimes" json:"previous_accumulation_times"`
	Deposits                  Deposits                                 `protobuf:"bytes,3,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	Borrows                   Borrows                                  `protobuf:"bytes,4,rep,name=borrows,proto3,castrepeated=Borrows" json:"borrows"`
	TotalSupplied             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=total_supplied,json=totalSupplied,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_supplied"`
	TotalBorrowed             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=total_borrowed,json=totalBorrowed,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_borrowed"`
	TotalReserves             github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=total_reserves,json=totalReserves,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_reserves"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2bfb514156ff30f, []int{0}
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

func (m *GenesisState) GetPreviousAccumulationTimes() GenesisAccumulationTimes {
	if m != nil {
		return m.PreviousAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetDeposits() Deposits {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetBorrows() Borrows {
	if m != nil {
		return m.Borrows
	}
	return nil
}

func (m *GenesisState) GetTotalSupplied() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalSupplied
	}
	return nil
}

func (m *GenesisState) GetTotalBorrowed() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalBorrowed
	}
	return nil
}

func (m *GenesisState) GetTotalReserves() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalReserves
	}
	return nil
}

// GenesisAccumulationTime stores the previous distribution time and its corresponding denom.
type GenesisAccumulationTime struct {
	CollateralType           string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty"`
	PreviousAccumulationTime time.Time                              `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time"`
	SupplyInterestFactor     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=supply_interest_factor,json=supplyInterestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"supply_interest_factor"`
	BorrowInterestFactor     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=borrow_interest_factor,json=borrowInterestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"borrow_interest_factor"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2bfb514156ff30f, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "joltify.third_party.jolt.v1beta1.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "joltify.third_party.jolt.v1beta1.GenesisAccumulationTime")
}

func init() {
	proto.RegisterFile("joltify/third_party/jolt/v1beta1/genesis.proto", fileDescriptor_f2bfb514156ff30f)
}

var fileDescriptor_f2bfb514156ff30f = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0x9b, 0x75, 0xff, 0xf0, 0x60, 0x43, 0xd1, 0x04, 0x59, 0x91, 0xd2, 0x6a, 0x07, 0x28,
	0x42, 0x4b, 0xd8, 0x38, 0x21, 0x71, 0x21, 0x9b, 0x86, 0xb8, 0xa1, 0xac, 0x12, 0x12, 0x12, 0x8a,
	0x9c, 0xe4, 0x35, 0x33, 0x24, 0x71, 0x64, 0x3b, 0x85, 0x7e, 0x09, 0xb4, 0x3b, 0xdf, 0x80, 0x33,
	0x1f, 0x62, 0xc7, 0x89, 0x13, 0xe2, 0xb0, 0xa1, 0xf6, 0x7b, 0x20, 0x14, 0xdb, 0xfd, 0x83, 0xaa,
	0xaa, 0x3b, 0xb0, 0xd3, 0xf6, 0xec, 0xe7, 0xdf, 0xef, 0xd9, 0x79, 0xaf, 0xc8, 0xf9, 0x40, 0x53,
	0x41, 0xba, 0x7d, 0x57, 0x9c, 0x12, 0x16, 0x07, 0x05, 0x66, 0xa2, 0xef, 0x56, 0x6b, 0x6e, 0x6f,
	0x3f, 0x04, 0x81, 0xf7, 0xdd, 0x04, 0x72, 0xe0, 0x84, 0x3b, 0x05, 0xa3, 0x82, 0x9a, 0x2d, 0x9d,
	0xef, 0x4c, 0xe5, 0x4b, 0x86, 0xa3, 0xf3, 0x1b, 0x76, 0x44, 0x79, 0x46, 0xb9, 0x1b, 0x62, 0x0e,
	0x63, 0x48, 0x44, 0x49, 0xae, 0x08, 0x8d, 0x1d, 0xb5, 0x1f, 0xc8, 0xc8, 0x55, 0x81, 0xde, 0x7a,
	0xb2, 0xb0, 0x18, 0x69, 0x52, 0xc9, 0xcd, 0x84, 0xd2, 0x24, 0x05, 0x57, 0x46, 0x61, 0xd9, 0x75,
	0x05, 0xc9, 0x80, 0x0b, 0x9c, 0x15, 0x3a, 0x61, 0x3b, 0xa1, 0x09, 0x55, 0x96, 0xea, 0x3f, 0xb5,
	0xba, 0xfb, 0x67, 0x05, 0xdd, 0x7e, 0xa5, 0xae, 0x74, 0x22, 0xb0, 0x00, 0xf3, 0x18, 0xad, 0x16,
	0x98, 0xe1, 0x8c, 0x5b, 0x46, 0xcb, 0x68, 0x6f, 0x1c, 0xb4, 0x9d, 0x45, 0x57, 0x74, 0xde, 0xc8,
	0x7c, 0x6f, 0xf9, 0xfc, 0xb2, 0x59, 0xf3, 0xf5, 0x69, 0xf3, 0xab, 0x81, 0x1e, 0x14, 0x0c, 0x7a,
	0x84, 0x96, 0x3c, 0xc0, 0x51, 0x54, 0x66, 0x65, 0x8a, 0x05, 0xa1, 0x79, 0x20, 0x0b, 0xb3, 0x96,
	0x5a, 0xf5, 0xf6, 0xc6, 0xc1, 0xf3, 0xc5, 0x74, 0x5d, 0xdd, 0xcb, 0x29, 0x44, 0x87, 0x64, 0xe0,
	0xb5, 0x2a, 0xdd, 0xb7, 0xab, 0xa6, 0x35, 0x27, 0x81, 0xfb, 0x3b, 0x23, 0xff, 0xcc, 0x96, 0xf9,
	0x16, 0xad, 0xc7, 0x50, 0x50, 0x4e, 0x04, 0xb7, 0xea, 0xb2, 0x92, 0xc7, 0x8b, 0x2b, 0x39, 0x52,
	0x27, 0xbc, 0xbb, 0xda, 0xbc, 0xae, 0x17, 0xb8, 0x3f, 0x86, 0x99, 0x27, 0x68, 0x2d, 0xa4, 0x8c,
	0xd1, 0x4f, 0xdc, 0x5a, 0x96, 0xdc, 0x6b, 0xbc, 0x9f, 0x27, 0x0f, 0x78, 0x5b, 0x1a, 0xbb, 0xa6,
	0x62, 0xee, 0x8f, 0x48, 0x26, 0x43, 0x9b, 0x82, 0x0a, 0x9c, 0x06, 0xbc, 0x2c, 0x8a, 0x94, 0x40,
	0x6c, 0xad, 0x48, 0xf6, 0x8e, 0xa3, 0xfb, 0xa5, 0x6a, 0xae, 0x31, 0xee, 0x90, 0x92, 0xdc, 0x7b,
	0xaa, 0x61, 0xed, 0x84, 0x88, 0xd3, 0x32, 0x74, 0x22, 0x9a, 0xe9, 0xe6, 0xd2, 0x7f, 0xf6, 0x78,
	0xfc, 0xd1, 0x15, 0xfd, 0x02, 0xb8, 0x3c, 0xc0, 0xfd, 0x3b, 0x52, 0x71, 0xa2, 0x0d, 0x13, 0xa7,
	0x2a, 0x02, 0x62, 0x6b, 0xf5, 0xa6, 0x9c, 0x9e, 0x36, 0x4c, 0x9c, 0x0c, 0x38, 0xb0, 0x1e, 0x70,
	0x6b, 0xed, 0xa6, 0x9c, 0xbe, 0x36, 0xec, 0x7e, 0xa9, 0xa3, 0xfb, 0x73, 0x3a, 0xc8, 0x7c, 0x84,
	0xb6, 0x22, 0x9a, 0xa6, 0x58, 0x00, 0xc3, 0x69, 0x50, 0x41, 0xe4, 0x50, 0xdc, 0xf2, 0x37, 0x27,
	0xcb, 0x9d, 0x7e, 0x01, 0x66, 0x88, 0x1a, 0xf3, 0x7b, 0xdd, 0x5a, 0x92, 0x83, 0xd4, 0x70, 0xd4,
	0x84, 0x3a, 0xa3, 0x09, 0x75, 0x3a, 0xa3, 0x09, 0xf5, 0xd6, 0xab, 0x5b, 0x9c, 0x5d, 0x35, 0x0d,
	0xdf, 0x9a, 0xd7, 0xb3, 0x26, 0x43, 0xf7, 0xe4, 0xe7, 0xef, 0x07, 0x24, 0x17, 0xc0, 0x80, 0x8b,
	0xa0, 0x8b, 0x23, 0x41, 0x99, 0x55, 0xaf, 0x6a, 0xf2, 0x5e, 0x54, 0x8c, 0x5f, 0x97, 0xcd, 0x87,
	0xd7, 0x78, 0x89, 0x23, 0x88, 0x7e, 0x7c, 0xdf, 0x43, 0xfa, 0x55, 0x8f, 0x20, 0xf2, 0xb7, 0x15,
	0xfb, 0xb5, 0x46, 0x1f, 0x4b, 0x72, 0xe5, 0x54, 0x9f, 0x7f, 0xc6, 0xb9, 0xfc, 0x3f, 0x9c, 0x8a,
	0xfd, 0xaf, 0xd3, 0x7b, 0x7f, 0x3e, 0xb0, 0x8d, 0x8b, 0x81, 0x6d, 0xfc, 0x1e, 0xd8, 0xc6, 0xd9,
	0xd0, 0xae, 0x5d, 0x0c, 0xed, 0xda, 0xcf, 0xa1, 0x5d, 0x7b, 0x77, 0x38, 0x65, 0xd1, 0x43, 0xb5,
	0xd7, 0x25, 0x39, 0xce, 0x23, 0x18, 0xc5, 0x41, 0x0a, 0x79, 0x4c, 0xf2, 0xc4, 0xfd, 0x3c, 0xfb,
	0xa3, 0x29, 0xcb, 0x08, 0x57, 0xe5, 0xf3, 0x3f, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x22,
	0xbd, 0x02, 0xea, 0x05, 0x00, 0x00,
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
	if len(m.TotalReserves) > 0 {
		for iNdEx := len(m.TotalReserves) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalReserves[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.TotalBorrowed) > 0 {
		for iNdEx := len(m.TotalBorrowed) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalBorrowed[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.TotalSupplied) > 0 {
		for iNdEx := len(m.TotalSupplied) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalSupplied[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Borrows) > 0 {
		for iNdEx := len(m.Borrows) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Borrows[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PreviousAccumulationTimes) > 0 {
		for iNdEx := len(m.PreviousAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreviousAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BorrowInterestFactor.Size()
		i -= size
		if _, err := m.BorrowInterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SupplyInterestFactor.Size()
		i -= size
		if _, err := m.SupplyInterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
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
	if len(m.PreviousAccumulationTimes) > 0 {
		for _, e := range m.PreviousAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Borrows) > 0 {
		for _, e := range m.Borrows {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalSupplied) > 0 {
		for _, e := range m.TotalSupplied {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalBorrowed) > 0 {
		for _, e := range m.TotalBorrowed {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalReserves) > 0 {
		for _, e := range m.TotalReserves {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SupplyInterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.BorrowInterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTimes", wireType)
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
			m.PreviousAccumulationTimes = append(m.PreviousAccumulationTimes, GenesisAccumulationTime{})
			if err := m.PreviousAccumulationTimes[len(m.PreviousAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrows", wireType)
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
			m.Borrows = append(m.Borrows, Borrow{})
			if err := m.Borrows[len(m.Borrows)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupplied", wireType)
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
			m.TotalSupplied = append(m.TotalSupplied, types.Coin{})
			if err := m.TotalSupplied[len(m.TotalSupplied)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBorrowed", wireType)
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
			m.TotalBorrowed = append(m.TotalBorrowed, types.Coin{})
			if err := m.TotalBorrowed[len(m.TotalBorrowed)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalReserves", wireType)
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
			m.TotalReserves = append(m.TotalReserves, types.Coin{})
			if err := m.TotalReserves[len(m.TotalReserves)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplyInterestFactor", wireType)
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
			if err := m.SupplyInterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowInterestFactor", wireType)
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
			if err := m.BorrowInterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
