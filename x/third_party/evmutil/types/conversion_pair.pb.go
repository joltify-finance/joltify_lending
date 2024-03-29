// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/evmutil/v1beta1/conversion_pair.proto

package types

import (
	bytes "bytes"
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

// ConversionPair defines a Jolt ERC20 address and corresponding denom that is
// allowed to be converted between ERC20 and sdk.Coin
type ConversionPair struct {
	// ERC20 address of the token on the Jolt EVM
	JoltERC20Address HexBytes `protobuf:"bytes,1,opt,name=jolt_erc20_address,json=joltErc20Address,proto3,casttype=HexBytes" json:"jolt_erc20_address,omitempty"`
	// Denom of the corresponding sdk.Coin
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *ConversionPair) Reset()         { *m = ConversionPair{} }
func (m *ConversionPair) String() string { return proto.CompactTextString(m) }
func (*ConversionPair) ProtoMessage()    {}
func (*ConversionPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c4d4bd101e3de9, []int{0}
}
func (m *ConversionPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConversionPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConversionPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConversionPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionPair.Merge(m, src)
}
func (m *ConversionPair) XXX_Size() int {
	return m.Size()
}
func (m *ConversionPair) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionPair.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionPair proto.InternalMessageInfo

// AllowedCosmosCoinERC20Token defines allowed cosmos-sdk denom & metadata
// for evm token representations of sdk assets.
// NOTE: once evm token contracts are deployed, changes to metadata for a given
// cosmos_denom will not change metadata of deployed contract.
type AllowedCosmosCoinERC20Token struct {
	// Denom of the sdk.Coin
	CosmosDenom string `protobuf:"bytes,1,opt,name=cosmos_denom,json=cosmosDenom,proto3" json:"cosmos_denom,omitempty"`
	// Name of ERC20 contract
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of ERC20 contract
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Number of decimals ERC20 contract is deployed with.
	Decimals uint32 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *AllowedCosmosCoinERC20Token) Reset()         { *m = AllowedCosmosCoinERC20Token{} }
func (m *AllowedCosmosCoinERC20Token) String() string { return proto.CompactTextString(m) }
func (*AllowedCosmosCoinERC20Token) ProtoMessage()    {}
func (*AllowedCosmosCoinERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_60c4d4bd101e3de9, []int{1}
}
func (m *AllowedCosmosCoinERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedCosmosCoinERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedCosmosCoinERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedCosmosCoinERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedCosmosCoinERC20Token.Merge(m, src)
}
func (m *AllowedCosmosCoinERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *AllowedCosmosCoinERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedCosmosCoinERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedCosmosCoinERC20Token proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionPair)(nil), "joltify.third_party.evmutil.v1beta1.ConversionPair")
	proto.RegisterType((*AllowedCosmosCoinERC20Token)(nil), "joltify.third_party.evmutil.v1beta1.AllowedCosmosCoinERC20Token")
}

func init() {
	proto.RegisterFile("joltify/third_party/evmutil/v1beta1/conversion_pair.proto", fileDescriptor_60c4d4bd101e3de9)
}

var fileDescriptor_60c4d4bd101e3de9 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x28, 0xa7, 0xc3, 0x1c, 0xe8, 0x64, 0x9d, 0x50, 0x74, 0x48, 0x6e, 0x38, 0x96,
	0x2e, 0x24, 0xd7, 0x32, 0xc1, 0xd6, 0x86, 0x22, 0xc4, 0x80, 0x50, 0xc4, 0xc4, 0x12, 0x39, 0x89,
	0x9b, 0x1a, 0x1c, 0xbf, 0x91, 0xed, 0x96, 0x46, 0xe2, 0x03, 0x30, 0x21, 0x3e, 0x02, 0x23, 0x1f,
	0x85, 0xb1, 0x23, 0x53, 0x55, 0xd2, 0x6f, 0xc1, 0x84, 0xf2, 0x87, 0x8a, 0x81, 0xed, 0x7d, 0x9e,
	0x37, 0xbf, 0x9f, 0xa2, 0xd7, 0xf8, 0xe9, 0x7b, 0x90, 0x56, 0x2c, 0xaa, 0xc0, 0x2e, 0x85, 0xce,
	0xe2, 0x92, 0x69, 0x5b, 0x05, 0x7c, 0x5d, 0xac, 0xac, 0x90, 0xc1, 0x7a, 0x9c, 0x70, 0xcb, 0xc6,
	0x41, 0x0a, 0x6a, 0xcd, 0xb5, 0x11, 0xa0, 0xe2, 0x92, 0x09, 0xed, 0x97, 0x1a, 0x2c, 0x90, 0x47,
	0x3d, 0xea, 0xff, 0x83, 0xfa, 0x3d, 0xea, 0xf7, 0xe8, 0xe5, 0x45, 0x0e, 0x39, 0xb4, 0xdf, 0x07,
	0xcd, 0xd4, 0xa1, 0x57, 0x9f, 0xf0, 0xbd, 0xf0, 0xe8, 0x7c, 0xc3, 0x84, 0x26, 0xaf, 0x31, 0x69,
	0x74, 0x31, 0xd7, 0xe9, 0xe4, 0x3a, 0x66, 0x59, 0xa6, 0xb9, 0x31, 0x2e, 0xf2, 0xd0, 0xe8, 0x6c,
	0xe6, 0xd5, 0xbb, 0xe1, 0xf9, 0x2b, 0x90, 0x76, 0x1e, 0x85, 0x93, 0xeb, 0x69, 0xb7, 0xfb, 0xbd,
	0x1b, 0x9e, 0xbe, 0xe4, 0x9b, 0x59, 0x65, 0xb9, 0x89, 0xce, 0x1b, 0x76, 0xde, 0xa0, 0xfd, 0x96,
	0x5c, 0xe0, 0x5b, 0x19, 0x57, 0x50, 0xb8, 0x37, 0x3c, 0x34, 0xba, 0x1d, 0x75, 0xe1, 0xd9, 0xe0,
	0xf3, 0xb7, 0xa1, 0x73, 0xf5, 0x05, 0xe1, 0x07, 0x53, 0x29, 0xe1, 0x23, 0xcf, 0x42, 0x30, 0x05,
	0x98, 0x10, 0x84, 0x6a, 0xdd, 0x6f, 0xe1, 0x03, 0x57, 0xe4, 0x21, 0x3e, 0x4b, 0xdb, 0x3e, 0xee,
	0x14, 0xa8, 0x55, 0xdc, 0xe9, 0xba, 0xe7, 0x4d, 0x45, 0x08, 0x1e, 0x28, 0x56, 0xf0, 0xde, 0xde,
	0xce, 0xe4, 0x3e, 0x3e, 0x31, 0x55, 0x91, 0x80, 0x74, 0x6f, 0xb6, 0x6d, 0x9f, 0xc8, 0x25, 0x3e,
	0xcd, 0x78, 0x2a, 0x0a, 0x26, 0x8d, 0x3b, 0xf0, 0xd0, 0xe8, 0x6e, 0x74, 0xcc, 0xdd, 0x0f, 0xcd,
	0x96, 0xfb, 0x5f, 0x14, 0x7d, 0xaf, 0x29, 0xfa, 0x51, 0x53, 0xb4, 0xad, 0x29, 0xda, 0xd7, 0x14,
	0x7d, 0x3d, 0x50, 0x67, 0x7b, 0xa0, 0xce, 0xcf, 0x03, 0x75, 0xde, 0xbd, 0xc8, 0x85, 0x5d, 0xae,
	0x12, 0x3f, 0x85, 0x22, 0xe8, 0xcf, 0xfe, 0x78, 0x21, 0x14, 0x53, 0x29, 0xff, 0x9b, 0x63, 0xc9,
	0x55, 0x26, 0x54, 0x1e, 0x6c, 0xfe, 0xfb, 0x96, 0xb6, 0x2a, 0xb9, 0x49, 0x4e, 0xda, 0xfb, 0x3f,
	0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x11, 0x27, 0xc9, 0xf7, 0x01, 0x00, 0x00,
}

func (this *ConversionPair) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *ConversionPair")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *ConversionPair but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *ConversionPair but is not nil && this == nil")
	}
	if !bytes.Equal(this.JoltERC20Address, that1.JoltERC20Address) {
		return fmt.Errorf("JoltERC20Address this(%v) Not Equal that(%v)", this.JoltERC20Address, that1.JoltERC20Address)
	}
	if this.Denom != that1.Denom {
		return fmt.Errorf("Denom this(%v) Not Equal that(%v)", this.Denom, that1.Denom)
	}
	return nil
}
func (this *ConversionPair) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConversionPair)
	if !ok {
		that2, ok := that.(ConversionPair)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.JoltERC20Address, that1.JoltERC20Address) {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	return true
}
func (this *AllowedCosmosCoinERC20Token) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*AllowedCosmosCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedCosmosCoinERC20Token)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *AllowedCosmosCoinERC20Token")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *AllowedCosmosCoinERC20Token but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *AllowedCosmosCoinERC20Token but is not nil && this == nil")
	}
	if this.CosmosDenom != that1.CosmosDenom {
		return fmt.Errorf("CosmosDenom this(%v) Not Equal that(%v)", this.CosmosDenom, that1.CosmosDenom)
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Symbol != that1.Symbol {
		return fmt.Errorf("Symbol this(%v) Not Equal that(%v)", this.Symbol, that1.Symbol)
	}
	if this.Decimals != that1.Decimals {
		return fmt.Errorf("Decimals this(%v) Not Equal that(%v)", this.Decimals, that1.Decimals)
	}
	return nil
}
func (this *AllowedCosmosCoinERC20Token) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AllowedCosmosCoinERC20Token)
	if !ok {
		that2, ok := that.(AllowedCosmosCoinERC20Token)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CosmosDenom != that1.CosmosDenom {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Decimals != that1.Decimals {
		return false
	}
	return true
}
func (m *ConversionPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConversionPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConversionPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.JoltERC20Address) > 0 {
		i -= len(m.JoltERC20Address)
		copy(dAtA[i:], m.JoltERC20Address)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.JoltERC20Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllowedCosmosCoinERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedCosmosCoinERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedCosmosCoinERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintConversionPair(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CosmosDenom) > 0 {
		i -= len(m.CosmosDenom)
		copy(dAtA[i:], m.CosmosDenom)
		i = encodeVarintConversionPair(dAtA, i, uint64(len(m.CosmosDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConversionPair(dAtA []byte, offset int, v uint64) int {
	offset -= sovConversionPair(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConversionPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.JoltERC20Address)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	return n
}

func (m *AllowedCosmosCoinERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CosmosDenom)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovConversionPair(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovConversionPair(uint64(m.Decimals))
	}
	return n
}

func sovConversionPair(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConversionPair(x uint64) (n int) {
	return sovConversionPair(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConversionPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: ConversionPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConversionPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoltERC20Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JoltERC20Address = append(m.JoltERC20Address[:0], dAtA[iNdEx:postIndex]...)
			if m.JoltERC20Address == nil {
				m.JoltERC20Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func (m *AllowedCosmosCoinERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConversionPair
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
			return fmt.Errorf("proto: AllowedCosmosCoinERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedCosmosCoinERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CosmosDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
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
				return ErrInvalidLengthConversionPair
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConversionPair
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConversionPair
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConversionPair(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConversionPair
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
func skipConversionPair(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
					return 0, ErrIntOverflowConversionPair
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
				return 0, ErrInvalidLengthConversionPair
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConversionPair
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConversionPair
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConversionPair        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConversionPair          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConversionPair = fmt.Errorf("proto: unexpected end of group")
)
