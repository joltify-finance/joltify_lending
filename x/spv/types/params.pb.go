// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/spv/params.proto

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

type Moneymarket struct {
	Denom            string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ConversionFactor int32  `protobuf:"varint,2,opt,name=conversion_factor,json=conversionFactor,proto3" json:"conversion_factor,omitempty"`
}

func (m *Moneymarket) Reset()         { *m = Moneymarket{} }
func (m *Moneymarket) String() string { return proto.CompactTextString(m) }
func (*Moneymarket) ProtoMessage()    {}
func (*Moneymarket) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e27a54d08a0a3a, []int{0}
}
func (m *Moneymarket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Moneymarket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Moneymarket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Moneymarket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Moneymarket.Merge(m, src)
}
func (m *Moneymarket) XXX_Size() int {
	return m.Size()
}
func (m *Moneymarket) XXX_DiscardUnknown() {
	xxx_messageInfo_Moneymarket.DiscardUnknown(m)
}

var xxx_messageInfo_Moneymarket proto.InternalMessageInfo

func (m *Moneymarket) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Moneymarket) GetConversionFactor() int32 {
	if m != nil {
		return m.ConversionFactor
	}
	return 0
}

type Incentive struct {
	Poolid string `protobuf:"bytes,1,opt,name=poolid,proto3" json:"poolid,omitempty"`
	Spy    string `protobuf:"bytes,2,opt,name=spy,proto3" json:"spy,omitempty"`
}

func (m *Incentive) Reset()         { *m = Incentive{} }
func (m *Incentive) String() string { return proto.CompactTextString(m) }
func (*Incentive) ProtoMessage()    {}
func (*Incentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e27a54d08a0a3a, []int{1}
}
func (m *Incentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Incentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Incentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Incentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Incentive.Merge(m, src)
}
func (m *Incentive) XXX_Size() int {
	return m.Size()
}
func (m *Incentive) XXX_DiscardUnknown() {
	xxx_messageInfo_Incentive.DiscardUnknown(m)
}

var xxx_messageInfo_Incentive proto.InternalMessageInfo

func (m *Incentive) GetPoolid() string {
	if m != nil {
		return m.Poolid
	}
	return ""
}

func (m *Incentive) GetSpy() string {
	if m != nil {
		return m.Spy
	}
	return ""
}

// Params defines the parameters for the module.
type Params struct {
	BurnThreshold github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=burn_threshold,json=burnThreshold,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burn_threshold"`
	Markets       []Moneymarket                            `protobuf:"bytes,2,rep,name=markets,proto3" json:"markets"`
	Incentives    []Incentive                              `protobuf:"bytes,3,rep,name=incentives,proto3" json:"incentives"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e27a54d08a0a3a, []int{2}
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

func (m *Params) GetBurnThreshold() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BurnThreshold
	}
	return nil
}

func (m *Params) GetMarkets() []Moneymarket {
	if m != nil {
		return m.Markets
	}
	return nil
}

func (m *Params) GetIncentives() []Incentive {
	if m != nil {
		return m.Incentives
	}
	return nil
}

func init() {
	proto.RegisterType((*Moneymarket)(nil), "joltify.spv.Moneymarket")
	proto.RegisterType((*Incentive)(nil), "joltify.spv.Incentive")
	proto.RegisterType((*Params)(nil), "joltify.spv.Params")
}

func init() { proto.RegisterFile("joltify/spv/params.proto", fileDescriptor_b1e27a54d08a0a3a) }

var fileDescriptor_b1e27a54d08a0a3a = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x93, 0xdb, 0x7b, 0x2b, 0x9d, 0xa2, 0x5c, 0x87, 0x52, 0x62, 0x17, 0x69, 0xe9, 0x2a,
	0x20, 0x9d, 0xb1, 0x8a, 0x22, 0xe2, 0xaa, 0x82, 0xe0, 0x42, 0x28, 0xc1, 0x95, 0x9b, 0x92, 0x3f,
	0xd3, 0x74, 0x6c, 0x32, 0x27, 0xcc, 0x4c, 0x83, 0x79, 0x0b, 0x97, 0x2e, 0x5d, 0xfb, 0x24, 0x5d,
	0x76, 0xe9, 0x4a, 0xa5, 0x7d, 0x0b, 0x57, 0x92, 0x49, 0xa2, 0xb9, 0xab, 0xcc, 0x99, 0x73, 0xbe,
	0xf3, 0xf1, 0xfb, 0x32, 0xc8, 0xf9, 0x04, 0xa9, 0xe6, 0xdb, 0x92, 0xaa, 0xbc, 0xa0, 0x79, 0x20,
	0x83, 0x4c, 0x91, 0x5c, 0x82, 0x06, 0x3c, 0x6c, 0x3a, 0x44, 0xe5, 0xc5, 0x64, 0x94, 0x40, 0x02,
	0xe6, 0x9e, 0x56, 0xa7, 0x7a, 0x64, 0xe2, 0x46, 0xa0, 0x32, 0x50, 0x34, 0x0c, 0x14, 0xa3, 0xc5,
	0x32, 0x64, 0x3a, 0x58, 0xd2, 0x08, 0xb8, 0xa8, 0xfb, 0xf3, 0x35, 0x1a, 0xbe, 0x07, 0xc1, 0xca,
	0x2c, 0x90, 0x7b, 0xa6, 0xf1, 0x08, 0xdd, 0xc4, 0x4c, 0x40, 0xe6, 0xd8, 0x33, 0xdb, 0x1b, 0xf8,
	0x75, 0x81, 0x1f, 0xa3, 0x87, 0x11, 0x88, 0x82, 0x49, 0xc5, 0x41, 0x6c, 0xb6, 0x41, 0xa4, 0x41,
	0x3a, 0x57, 0x33, 0xdb, 0xbb, 0xf1, 0x6f, 0xff, 0x37, 0xde, 0x9a, 0xfb, 0xf9, 0x73, 0x34, 0x78,
	0x27, 0x22, 0x26, 0x34, 0x2f, 0x18, 0x1e, 0xa3, 0x7e, 0x0e, 0x90, 0xf2, 0xb8, 0x59, 0xd8, 0x54,
	0xf8, 0x16, 0xf5, 0x54, 0x5e, 0x9a, 0x1d, 0x03, 0xbf, 0x3a, 0xce, 0xff, 0xd8, 0xa8, 0xbf, 0x36,
	0x70, 0x58, 0xa2, 0x07, 0xe1, 0x41, 0x8a, 0x8d, 0xde, 0x49, 0xa6, 0x76, 0x90, 0x56, 0xe2, 0x9e,
	0x37, 0x7c, 0xfa, 0x88, 0xd4, 0x30, 0xa4, 0x82, 0x21, 0x0d, 0x0c, 0x79, 0x03, 0x5c, 0xac, 0x9e,
	0x1c, 0x7f, 0x4e, 0xad, 0xef, 0xbf, 0xa6, 0x5e, 0xc2, 0xf5, 0xee, 0x10, 0x92, 0x08, 0x32, 0xda,
	0x90, 0xd7, 0x9f, 0x85, 0x8a, 0xf7, 0x54, 0x97, 0x39, 0x53, 0x46, 0xa0, 0xfc, 0xfb, 0x95, 0xc5,
	0x87, 0xd6, 0x01, 0xbf, 0x44, 0xf7, 0xea, 0x08, 0x94, 0x73, 0x65, 0xcc, 0x1c, 0xd2, 0x09, 0x97,
	0x74, 0x32, 0x5a, 0x5d, 0x57, 0x5e, 0x7e, 0x3b, 0x8e, 0x5f, 0x23, 0xc4, 0x5b, 0x5e, 0xe5, 0xf4,
	0x8c, 0x78, 0x7c, 0x47, 0xfc, 0x2f, 0x8e, 0x46, 0xda, 0x99, 0x7f, 0x75, 0xfd, 0xf5, 0xdb, 0xd4,
	0x5a, 0xad, 0x8f, 0x67, 0xd7, 0x3e, 0x9d, 0x5d, 0xfb, 0xf7, 0xd9, 0xb5, 0xbf, 0x5c, 0x5c, 0xeb,
	0x74, 0x71, 0xad, 0x1f, 0x17, 0xd7, 0xfa, 0xf8, 0xa2, 0x03, 0xd4, 0xec, 0x5c, 0x6c, 0xb9, 0x08,
	0x44, 0xc4, 0xda, 0x7a, 0x93, 0x32, 0x11, 0x73, 0x91, 0xd0, 0xcf, 0xe6, 0x85, 0x18, 0xc8, 0xb0,
	0x6f, 0x7e, 0xef, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x2c, 0xf4, 0xb2, 0x3d, 0x02,
	0x00, 0x00,
}

func (m *Moneymarket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Moneymarket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Moneymarket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConversionFactor != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConversionFactor))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Incentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Incentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Incentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Spy) > 0 {
		i -= len(m.Spy)
		copy(dAtA[i:], m.Spy)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Spy)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Poolid) > 0 {
		i -= len(m.Poolid)
		copy(dAtA[i:], m.Poolid)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Poolid)))
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
	if len(m.Incentives) > 0 {
		for iNdEx := len(m.Incentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Incentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Markets) > 0 {
		for iNdEx := len(m.Markets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BurnThreshold) > 0 {
		for iNdEx := len(m.BurnThreshold) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnThreshold[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *Moneymarket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ConversionFactor != 0 {
		n += 1 + sovParams(uint64(m.ConversionFactor))
	}
	return n
}

func (m *Incentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Poolid)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Spy)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BurnThreshold) > 0 {
		for _, e := range m.BurnThreshold {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.Markets) > 0 {
		for _, e := range m.Markets {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.Incentives) > 0 {
		for _, e := range m.Incentives {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Moneymarket) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Moneymarket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Moneymarket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversionFactor", wireType)
			}
			m.ConversionFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConversionFactor |= int32(b&0x7F) << shift
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
func (m *Incentive) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Incentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Incentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Poolid", wireType)
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
			m.Poolid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spy", wireType)
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
			m.Spy = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field BurnThreshold", wireType)
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
			m.BurnThreshold = append(m.BurnThreshold, types.Coin{})
			if err := m.BurnThreshold[len(m.BurnThreshold)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
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
			m.Markets = append(m.Markets, Moneymarket{})
			if err := m.Markets[len(m.Markets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Incentives", wireType)
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
			m.Incentives = append(m.Incentives, Incentive{})
			if err := m.Incentives[len(m.Incentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
