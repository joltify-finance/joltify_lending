// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/clob/liquidations_config.proto

package types

import (
	fmt "fmt"
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

// LiquidationsConfig stores all configurable fields related to liquidations.
type LiquidationsConfig struct {
	// The maximum liquidation fee (in parts-per-million). This fee goes
	// 100% to the insurance fund.
	MaxLiquidationFeePpm uint32 `protobuf:"varint,1,opt,name=max_liquidation_fee_ppm,json=maxLiquidationFeePpm,proto3" json:"max_liquidation_fee_ppm,omitempty"`
	// Limits around how much of a single position can be liquidated
	// within a single block.
	PositionBlockLimits PositionBlockLimits `protobuf:"bytes,2,opt,name=position_block_limits,json=positionBlockLimits,proto3" json:"position_block_limits"`
	// Limits around how many quote quantums from a single subaccount can
	// be liquidated within a single block.
	SubaccountBlockLimits SubaccountBlockLimits `protobuf:"bytes,3,opt,name=subaccount_block_limits,json=subaccountBlockLimits,proto3" json:"subaccount_block_limits"`
	// Config about how the fillable-price spread from the oracle price
	// increases based on the adjusted bankruptcy rating of the subaccount.
	FillablePriceConfig FillablePriceConfig `protobuf:"bytes,4,opt,name=fillable_price_config,json=fillablePriceConfig,proto3" json:"fillable_price_config"`
}

func (m *LiquidationsConfig) Reset()         { *m = LiquidationsConfig{} }
func (m *LiquidationsConfig) String() string { return proto.CompactTextString(m) }
func (*LiquidationsConfig) ProtoMessage()    {}
func (*LiquidationsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1bf5be993d07d99, []int{0}
}
func (m *LiquidationsConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidationsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidationsConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidationsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidationsConfig.Merge(m, src)
}
func (m *LiquidationsConfig) XXX_Size() int {
	return m.Size()
}
func (m *LiquidationsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidationsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidationsConfig proto.InternalMessageInfo

func (m *LiquidationsConfig) GetMaxLiquidationFeePpm() uint32 {
	if m != nil {
		return m.MaxLiquidationFeePpm
	}
	return 0
}

func (m *LiquidationsConfig) GetPositionBlockLimits() PositionBlockLimits {
	if m != nil {
		return m.PositionBlockLimits
	}
	return PositionBlockLimits{}
}

func (m *LiquidationsConfig) GetSubaccountBlockLimits() SubaccountBlockLimits {
	if m != nil {
		return m.SubaccountBlockLimits
	}
	return SubaccountBlockLimits{}
}

func (m *LiquidationsConfig) GetFillablePriceConfig() FillablePriceConfig {
	if m != nil {
		return m.FillablePriceConfig
	}
	return FillablePriceConfig{}
}

// PositionBlockLimits stores all configurable fields related to limits
// around how much of a single position can be liquidated within a single block.
type PositionBlockLimits struct {
	// The minimum amount of quantums to liquidate for each message (in
	// quote quantums).
	// Overridden by the maximum size of the position.
	MinPositionNotionalLiquidated uint64 `protobuf:"varint,1,opt,name=min_position_notional_liquidated,json=minPositionNotionalLiquidated,proto3" json:"min_position_notional_liquidated,omitempty"`
	// The maximum portion of the position liquidated (in parts-per-
	// million). Overridden by min_position_notional_liquidated.
	MaxPositionPortionLiquidatedPpm uint32 `protobuf:"varint,2,opt,name=max_position_portion_liquidated_ppm,json=maxPositionPortionLiquidatedPpm,proto3" json:"max_position_portion_liquidated_ppm,omitempty"`
}

func (m *PositionBlockLimits) Reset()         { *m = PositionBlockLimits{} }
func (m *PositionBlockLimits) String() string { return proto.CompactTextString(m) }
func (*PositionBlockLimits) ProtoMessage()    {}
func (*PositionBlockLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1bf5be993d07d99, []int{1}
}
func (m *PositionBlockLimits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PositionBlockLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PositionBlockLimits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PositionBlockLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionBlockLimits.Merge(m, src)
}
func (m *PositionBlockLimits) XXX_Size() int {
	return m.Size()
}
func (m *PositionBlockLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionBlockLimits.DiscardUnknown(m)
}

var xxx_messageInfo_PositionBlockLimits proto.InternalMessageInfo

func (m *PositionBlockLimits) GetMinPositionNotionalLiquidated() uint64 {
	if m != nil {
		return m.MinPositionNotionalLiquidated
	}
	return 0
}

func (m *PositionBlockLimits) GetMaxPositionPortionLiquidatedPpm() uint32 {
	if m != nil {
		return m.MaxPositionPortionLiquidatedPpm
	}
	return 0
}

// SubaccountBlockLimits stores all configurable fields related to limits
// around how many quote quantums from a single subaccount can
// be liquidated within a single block.
type SubaccountBlockLimits struct {
	// The maximum notional amount that a single subaccount can have
	// liquidated (in quote quantums) per block.
	MaxNotionalLiquidated uint64 `protobuf:"varint,1,opt,name=max_notional_liquidated,json=maxNotionalLiquidated,proto3" json:"max_notional_liquidated,omitempty"`
	// The maximum insurance-fund payout amount for a given subaccount
	// per block. I.e. how much it can cover for that subaccount.
	MaxQuantumsInsuranceLost uint64 `protobuf:"varint,2,opt,name=max_quantums_insurance_lost,json=maxQuantumsInsuranceLost,proto3" json:"max_quantums_insurance_lost,omitempty"`
}

func (m *SubaccountBlockLimits) Reset()         { *m = SubaccountBlockLimits{} }
func (m *SubaccountBlockLimits) String() string { return proto.CompactTextString(m) }
func (*SubaccountBlockLimits) ProtoMessage()    {}
func (*SubaccountBlockLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1bf5be993d07d99, []int{2}
}
func (m *SubaccountBlockLimits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountBlockLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountBlockLimits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountBlockLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountBlockLimits.Merge(m, src)
}
func (m *SubaccountBlockLimits) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountBlockLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountBlockLimits.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountBlockLimits proto.InternalMessageInfo

func (m *SubaccountBlockLimits) GetMaxNotionalLiquidated() uint64 {
	if m != nil {
		return m.MaxNotionalLiquidated
	}
	return 0
}

func (m *SubaccountBlockLimits) GetMaxQuantumsInsuranceLost() uint64 {
	if m != nil {
		return m.MaxQuantumsInsuranceLost
	}
	return 0
}

// FillablePriceConfig stores all configurable fields related to calculating
// the fillable price for liquidating a position.
type FillablePriceConfig struct {
	// The rate at which the Adjusted Bankruptcy Rating increases.
	BankruptcyAdjustmentPpm uint32 `protobuf:"varint,1,opt,name=bankruptcy_adjustment_ppm,json=bankruptcyAdjustmentPpm,proto3" json:"bankruptcy_adjustment_ppm,omitempty"`
	// The maximum value that the liquidation spread can take, as
	// a ratio against the position's maintenance margin.
	SpreadToMaintenanceMarginRatioPpm uint32 `protobuf:"varint,2,opt,name=spread_to_maintenance_margin_ratio_ppm,json=spreadToMaintenanceMarginRatioPpm,proto3" json:"spread_to_maintenance_margin_ratio_ppm,omitempty"`
}

func (m *FillablePriceConfig) Reset()         { *m = FillablePriceConfig{} }
func (m *FillablePriceConfig) String() string { return proto.CompactTextString(m) }
func (*FillablePriceConfig) ProtoMessage()    {}
func (*FillablePriceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1bf5be993d07d99, []int{3}
}
func (m *FillablePriceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FillablePriceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FillablePriceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FillablePriceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FillablePriceConfig.Merge(m, src)
}
func (m *FillablePriceConfig) XXX_Size() int {
	return m.Size()
}
func (m *FillablePriceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FillablePriceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FillablePriceConfig proto.InternalMessageInfo

func (m *FillablePriceConfig) GetBankruptcyAdjustmentPpm() uint32 {
	if m != nil {
		return m.BankruptcyAdjustmentPpm
	}
	return 0
}

func (m *FillablePriceConfig) GetSpreadToMaintenanceMarginRatioPpm() uint32 {
	if m != nil {
		return m.SpreadToMaintenanceMarginRatioPpm
	}
	return 0
}

func init() {
	proto.RegisterType((*LiquidationsConfig)(nil), "joltify.third_party.dydxprotocol.clob.LiquidationsConfig")
	proto.RegisterType((*PositionBlockLimits)(nil), "joltify.third_party.dydxprotocol.clob.PositionBlockLimits")
	proto.RegisterType((*SubaccountBlockLimits)(nil), "joltify.third_party.dydxprotocol.clob.SubaccountBlockLimits")
	proto.RegisterType((*FillablePriceConfig)(nil), "joltify.third_party.dydxprotocol.clob.FillablePriceConfig")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/clob/liquidations_config.proto", fileDescriptor_f1bf5be993d07d99)
}

var fileDescriptor_f1bf5be993d07d99 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0xe3, 0x36, 0xba, 0x8b, 0xb9, 0x62, 0xe3, 0x34, 0x4a, 0x00, 0xe1, 0x86, 0x20, 0x50,
	0x37, 0xd8, 0x12, 0x08, 0x16, 0x15, 0x08, 0x11, 0xa4, 0x02, 0x52, 0x8a, 0xdc, 0xc0, 0x8a, 0xcd,
	0x68, 0x6c, 0x8f, 0xdd, 0x69, 0xe7, 0x5f, 0x3d, 0x63, 0xc9, 0xe1, 0x21, 0x10, 0x0f, 0xc1, 0x92,
	0x1d, 0x2f, 0xd1, 0x65, 0x97, 0xac, 0x10, 0x4a, 0x5e, 0x04, 0xcd, 0xd8, 0x71, 0x8c, 0x12, 0xa4,
	0xb2, 0x1a, 0xdb, 0xe7, 0x9b, 0xdf, 0x39, 0x73, 0xce, 0xe7, 0x01, 0x2f, 0xce, 0x04, 0xd5, 0x24,
	0x9d, 0x07, 0xfa, 0x94, 0xe4, 0x09, 0x94, 0x28, 0xd7, 0xf3, 0x20, 0x99, 0x27, 0xa5, 0xcc, 0x85,
	0x16, 0xb1, 0xa0, 0x41, 0x4c, 0x45, 0x14, 0x50, 0x72, 0x51, 0x90, 0x04, 0x69, 0x22, 0xb8, 0x82,
	0xb1, 0xe0, 0x29, 0xc9, 0x7c, 0xab, 0x70, 0xef, 0xd7, 0x00, 0xbf, 0x05, 0xf0, 0xdb, 0x00, 0xdf,
	0x00, 0x6e, 0xed, 0x65, 0x22, 0x13, 0xf6, 0x53, 0x60, 0x9e, 0xaa, 0xcd, 0xe3, 0xef, 0xbb, 0xc0,
	0x9d, 0xb6, 0xd0, 0xaf, 0x2c, 0xd9, 0x7d, 0x02, 0x06, 0x0c, 0x95, 0xb0, 0x95, 0x14, 0xa6, 0x18,
	0x43, 0x29, 0xd9, 0xd0, 0x19, 0x39, 0x07, 0x37, 0x66, 0x7b, 0x0c, 0x95, 0xad, 0x7d, 0x47, 0x18,
	0x87, 0x92, 0xb9, 0x1a, 0xf4, 0xa5, 0x50, 0xc4, 0xea, 0x23, 0x2a, 0xe2, 0x73, 0x48, 0x09, 0x23,
	0x5a, 0x0d, 0x77, 0x46, 0xce, 0xc1, 0xff, 0x8f, 0x0e, 0xfd, 0x6b, 0x95, 0xea, 0x87, 0x35, 0x63,
	0x62, 0x10, 0x53, 0x4b, 0x98, 0x74, 0x2f, 0x7f, 0xee, 0x77, 0x66, 0x3d, 0xb9, 0x19, 0x72, 0x3f,
	0x81, 0x81, 0x2a, 0x22, 0x14, 0xc7, 0xa2, 0xe0, 0xfa, 0xcf, 0xbc, 0xbb, 0x36, 0xef, 0xb3, 0x6b,
	0xe6, 0x7d, 0xdf, 0x50, 0x36, 0x33, 0xf7, 0xd5, 0xb6, 0xa0, 0x39, 0x71, 0x4a, 0x28, 0x45, 0x11,
	0xc5, 0x50, 0xe6, 0x24, 0xc6, 0xf5, 0x6c, 0x86, 0xdd, 0x7f, 0x3a, 0xf1, 0x51, 0xcd, 0x08, 0x0d,
	0xa2, 0x9a, 0xc1, 0xea, 0xc4, 0xe9, 0x66, 0x68, 0xfc, 0xcd, 0x01, 0xbd, 0x2d, 0x4d, 0x72, 0x5f,
	0x83, 0x11, 0x23, 0x1c, 0x36, 0x33, 0xe0, 0xc2, 0x2c, 0x88, 0x36, 0x83, 0xc4, 0x89, 0x9d, 0x5f,
	0x77, 0x76, 0x87, 0x11, 0xbe, 0x22, 0xbc, 0xab, 0x55, 0xd3, 0x46, 0xe4, 0x4e, 0xc1, 0x3d, 0x33,
	0xff, 0x06, 0x24, 0x45, 0x6e, 0xd7, 0x35, 0xc7, 0x7a, 0x61, 0xc7, 0x7a, 0x61, 0x9f, 0xa1, 0x72,
	0xc5, 0x0a, 0x2b, 0xe1, 0x1a, 0x15, 0x4a, 0x36, 0xfe, 0xec, 0x80, 0xfe, 0xd6, 0xde, 0xba, 0x4f,
	0x2b, 0x9f, 0xfd, 0xbd, 0xce, 0x3e, 0x43, 0xe5, 0x96, 0xfa, 0x9e, 0x83, 0xdb, 0x66, 0xdf, 0x45,
	0x81, 0xb8, 0x2e, 0x98, 0x82, 0x84, 0xab, 0x22, 0x47, 0x3c, 0xc6, 0x90, 0x0a, 0xa5, 0x6d, 0x5d,
	0xdd, 0xd9, 0x90, 0xa1, 0xf2, 0xa4, 0x56, 0xbc, 0x5d, 0x09, 0xa6, 0x42, 0xe9, 0xf1, 0x57, 0x07,
	0xf4, 0xb6, 0xb4, 0xdc, 0x3d, 0x04, 0x37, 0x23, 0xc4, 0xcf, 0xf3, 0x42, 0xea, 0x78, 0x0e, 0x51,
	0x72, 0x56, 0x28, 0xcd, 0x30, 0xd7, 0x2d, 0xe3, 0x0f, 0xd6, 0x82, 0x97, 0x4d, 0xdc, 0x78, 0xff,
	0x04, 0x3c, 0x50, 0x32, 0xc7, 0x28, 0x81, 0x5a, 0x40, 0x86, 0x08, 0xd7, 0x98, 0xdb, 0x8a, 0x18,
	0xca, 0x33, 0xc2, 0x61, 0x6e, 0x7e, 0x94, 0x56, 0xd7, 0xee, 0x56, 0xea, 0x0f, 0xe2, 0x78, 0xad,
	0x3d, 0xb6, 0xd2, 0x99, 0x51, 0x86, 0x92, 0x4d, 0xa2, 0xcb, 0x85, 0xe7, 0x5c, 0x2d, 0x3c, 0xe7,
	0xd7, 0xc2, 0x73, 0xbe, 0x2c, 0xbd, 0xce, 0xd5, 0xd2, 0xeb, 0xfc, 0x58, 0x7a, 0x9d, 0x8f, 0x6f,
	0x32, 0xa2, 0x4f, 0x8b, 0xc8, 0x8f, 0x05, 0x0b, 0x6a, 0x87, 0x3d, 0x4c, 0x89, 0x65, 0xac, 0xde,
	0x21, 0xc5, 0x3c, 0x21, 0x3c, 0x0b, 0xca, 0xf6, 0xcd, 0x02, 0x8d, 0xf7, 0xaa, 0x1b, 0x45, 0xcf,
	0x25, 0x56, 0xd1, 0x7f, 0xd6, 0x88, 0x8f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x5b, 0x80,
	0xfa, 0x87, 0x04, 0x00, 0x00,
}

func (m *LiquidationsConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidationsConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidationsConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FillablePriceConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.SubaccountBlockLimits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PositionBlockLimits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.MaxLiquidationFeePpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxLiquidationFeePpm))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PositionBlockLimits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PositionBlockLimits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PositionBlockLimits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxPositionPortionLiquidatedPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxPositionPortionLiquidatedPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.MinPositionNotionalLiquidated != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MinPositionNotionalLiquidated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubaccountBlockLimits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountBlockLimits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountBlockLimits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxQuantumsInsuranceLost != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxQuantumsInsuranceLost))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxNotionalLiquidated != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.MaxNotionalLiquidated))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FillablePriceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FillablePriceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FillablePriceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SpreadToMaintenanceMarginRatioPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.SpreadToMaintenanceMarginRatioPpm))
		i--
		dAtA[i] = 0x10
	}
	if m.BankruptcyAdjustmentPpm != 0 {
		i = encodeVarintLiquidationsConfig(dAtA, i, uint64(m.BankruptcyAdjustmentPpm))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidationsConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidationsConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LiquidationsConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxLiquidationFeePpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxLiquidationFeePpm))
	}
	l = m.PositionBlockLimits.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	l = m.SubaccountBlockLimits.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	l = m.FillablePriceConfig.Size()
	n += 1 + l + sovLiquidationsConfig(uint64(l))
	return n
}

func (m *PositionBlockLimits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinPositionNotionalLiquidated != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MinPositionNotionalLiquidated))
	}
	if m.MaxPositionPortionLiquidatedPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxPositionPortionLiquidatedPpm))
	}
	return n
}

func (m *SubaccountBlockLimits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxNotionalLiquidated != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxNotionalLiquidated))
	}
	if m.MaxQuantumsInsuranceLost != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.MaxQuantumsInsuranceLost))
	}
	return n
}

func (m *FillablePriceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BankruptcyAdjustmentPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.BankruptcyAdjustmentPpm))
	}
	if m.SpreadToMaintenanceMarginRatioPpm != 0 {
		n += 1 + sovLiquidationsConfig(uint64(m.SpreadToMaintenanceMarginRatioPpm))
	}
	return n
}

func sovLiquidationsConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidationsConfig(x uint64) (n int) {
	return sovLiquidationsConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LiquidationsConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: LiquidationsConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidationsConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLiquidationFeePpm", wireType)
			}
			m.MaxLiquidationFeePpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxLiquidationFeePpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionBlockLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PositionBlockLimits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountBlockLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubaccountBlockLimits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FillablePriceConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
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
				return ErrInvalidLengthLiquidationsConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidationsConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FillablePriceConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *PositionBlockLimits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: PositionBlockLimits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PositionBlockLimits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPositionNotionalLiquidated", wireType)
			}
			m.MinPositionNotionalLiquidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPositionNotionalLiquidated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPositionPortionLiquidatedPpm", wireType)
			}
			m.MaxPositionPortionLiquidatedPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPositionPortionLiquidatedPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *SubaccountBlockLimits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: SubaccountBlockLimits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountBlockLimits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNotionalLiquidated", wireType)
			}
			m.MaxNotionalLiquidated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxNotionalLiquidated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxQuantumsInsuranceLost", wireType)
			}
			m.MaxQuantumsInsuranceLost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxQuantumsInsuranceLost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func (m *FillablePriceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidationsConfig
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
			return fmt.Errorf("proto: FillablePriceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FillablePriceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BankruptcyAdjustmentPpm", wireType)
			}
			m.BankruptcyAdjustmentPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BankruptcyAdjustmentPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadToMaintenanceMarginRatioPpm", wireType)
			}
			m.SpreadToMaintenanceMarginRatioPpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidationsConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpreadToMaintenanceMarginRatioPpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidationsConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidationsConfig
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
func skipLiquidationsConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidationsConfig
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
					return 0, ErrIntOverflowLiquidationsConfig
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
					return 0, ErrIntOverflowLiquidationsConfig
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
				return 0, ErrInvalidLengthLiquidationsConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidationsConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidationsConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidationsConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidationsConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidationsConfig = fmt.Errorf("proto: unexpected end of group")
)