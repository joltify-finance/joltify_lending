// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/assets/asset.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Asset defines a single exchangable asset.
type Asset struct {
	// Unique, sequentially-generated.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The human readable symbol of the `Asset` (e.g. `USDC`, `ATOM`).
	// Must be uppercase, unique and correspond to the canonical symbol of the
	// full coin.
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// The name of base denomination unit of the `Asset` (e.g. `uatom`,
	// 'ibc/xxxxx'). Must be unique and match the `denom` used in the `sdk.Coin`
	// type in the `x/bank` module.
	Denom string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	// The exponent of converting one unit of `denom` to a full coin.
	// For example, `name=USDC, denom=uusdc, denom_exponent=-6` defines that
	// `1 uusdc = 10^(-6) USDC`. Note that `uusdc` refers to a `Coin` type in
	// `x/bank`, where the prefix `u` means `micro` by convetion. `uusdc` is
	// a different concept from a "quantum" defined by `atomic_resolution` below.
	// To convert from an amount of `denom` to quantums:
	// `quantums = denom_amount * 10^(denom_exponent - atomic_resolution)`
	DenomExponent int32 `protobuf:"zigzag32,4,opt,name=denom_exponent,json=denomExponent,proto3" json:"denom_exponent,omitempty"`
	// `true` if this `Asset` has a valid `MarketId` value.
	HasMarket bool `protobuf:"varint,5,opt,name=has_market,json=hasMarket,proto3" json:"has_market,omitempty"`
	// The `Id` of the `Market` associated with this `Asset`. It acts as the
	// oracle price for the purposes of calculating collateral
	// and margin requirements.
	MarketId uint32 `protobuf:"varint,6,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	// The exponent for converting an atomic amount (1 'quantum')
	// to a full coin. For example, if `atomic_resolution = -8`
	// then an `asset_position` with `base_quantums = 1e8` is equivalent to
	// a position size of one full coin.
	AtomicResolution int32 `protobuf:"zigzag32,7,opt,name=atomic_resolution,json=atomicResolution,proto3" json:"atomic_resolution,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0b73b5c910a62b5, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}

func (m *Asset) XXX_Size() int {
	return m.Size()
}

func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Asset) GetDenomExponent() int32 {
	if m != nil {
		return m.DenomExponent
	}
	return 0
}

func (m *Asset) GetHasMarket() bool {
	if m != nil {
		return m.HasMarket
	}
	return false
}

func (m *Asset) GetMarketId() uint32 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

func (m *Asset) GetAtomicResolution() int32 {
	if m != nil {
		return m.AtomicResolution
	}
	return 0
}

func init() {
	proto.RegisterType((*Asset)(nil), "dydxprotocol.assets.Asset")
}

func init() { proto.RegisterFile("dydxprotocol/assets/asset.proto", fileDescriptor_d0b73b5c910a62b5) }

var fileDescriptor_d0b73b5c910a62b5 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0xf0, 0xcb, 0x69, 0xeb, 0x35, 0x70, 0x87, 0x17, 0x45, 0x02, 0x62, 0x2c, 0x82, 0x50,
	0x10, 0xdb, 0x41, 0x07, 0x57, 0x05, 0x07, 0x07, 0x97, 0x8c, 0x2e, 0x25, 0x6d, 0x82, 0x0d, 0xb6,
	0x4d, 0x69, 0x72, 0xd2, 0x7e, 0x0b, 0x3f, 0x96, 0xe3, 0x2d, 0x82, 0xa3, 0xb4, 0x5f, 0x44, 0x4c,
	0xce, 0x43, 0xa7, 0xbc, 0xf7, 0xfb, 0x3f, 0xc8, 0xe3, 0xc1, 0x53, 0xde, 0xf3, 0xae, 0x69, 0x95,
	0x51, 0xb9, 0x2a, 0x13, 0xa6, 0xb5, 0x30, 0xda, 0x3d, 0xb1, 0x55, 0x74, 0xf0, 0x77, 0x20, 0x76,
	0x03, 0x67, 0x1f, 0x00, 0x7a, 0xb7, 0x3f, 0x25, 0x5a, 0xc0, 0xa9, 0xe4, 0x18, 0x84, 0x20, 0x9a,
	0xd3, 0xa9, 0xe4, 0xe8, 0x08, 0xfa, 0xba, 0xaf, 0x32, 0x55, 0xe2, 0x69, 0x08, 0xa2, 0x80, 0x6e,
	0x3a, 0x74, 0x08, 0x3d, 0x2e, 0x6a, 0x55, 0xe1, 0x1d, 0xcb, 0xae, 0x41, 0xe7, 0x70, 0x61, 0x8b,
	0x54, 0x74, 0x8d, 0xaa, 0x45, 0x6d, 0xf0, 0x6e, 0x08, 0xa2, 0x25, 0x9d, 0x5b, 0xbd, 0xdf, 0x20,
	0x3a, 0x81, 0xb0, 0x60, 0x3a, 0xad, 0x58, 0xfb, 0x22, 0x0c, 0xf6, 0x42, 0x10, 0xcd, 0x68, 0x50,
	0x30, 0xfd, 0x68, 0x01, 0x1d, 0xc3, 0xc0, 0x45, 0xa9, 0xe4, 0xd8, 0xb7, 0xab, 0xcc, 0x1c, 0x3c,
	0x70, 0x74, 0x01, 0x97, 0xcc, 0xa8, 0x4a, 0xe6, 0x69, 0x2b, 0xb4, 0x2a, 0x57, 0x46, 0xaa, 0x1a,
	0xef, 0xd9, 0x5f, 0xf6, 0x5d, 0x40, 0xb7, 0x7e, 0x47, 0xdf, 0x07, 0x02, 0xd6, 0x03, 0x01, 0x5f,
	0x03, 0x01, 0x6f, 0x23, 0x99, 0xac, 0x47, 0x32, 0xf9, 0x1c, 0xc9, 0xe4, 0xe9, 0xe6, 0x59, 0x9a,
	0x62, 0x95, 0xc5, 0xb9, 0xaa, 0x92, 0x7f, 0x27, 0x7b, 0xbd, 0xbe, 0xcc, 0x0b, 0x26, 0xeb, 0x64,
	0x2b, 0xdd, 0xef, 0x19, 0x4d, 0xdf, 0x08, 0x9d, 0xf9, 0x36, 0xb8, 0xfa, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x03, 0x1b, 0xd1, 0xd6, 0x6a, 0x01, 0x00, 0x00,
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AtomicResolution != 0 {
		i = encodeVarintAsset(dAtA, i, uint64((uint32(m.AtomicResolution)<<1)^uint32((m.AtomicResolution>>31))))
		i--
		dAtA[i] = 0x38
	}
	if m.MarketId != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x30
	}
	if m.HasMarket {
		i--
		if m.HasMarket {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.DenomExponent != 0 {
		i = encodeVarintAsset(dAtA, i, uint64((uint32(m.DenomExponent)<<1)^uint32((m.DenomExponent>>31))))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintAsset(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintAsset(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAsset(uint64(m.Id))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovAsset(uint64(l))
	}
	if m.DenomExponent != 0 {
		n += 1 + sozAsset(uint64(m.DenomExponent))
	}
	if m.HasMarket {
		n += 2
	}
	if m.MarketId != 0 {
		n += 1 + sovAsset(uint64(m.MarketId))
	}
	if m.AtomicResolution != 0 {
		n += 1 + sozAsset(uint64(m.AtomicResolution))
	}
	return n
}

func sovAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozAsset(x uint64) (n int) {
	return sovAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAsset
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
				return ErrInvalidLengthAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomExponent", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.DenomExponent = v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasMarket", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
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
			m.HasMarket = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AtomicResolution", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.AtomicResolution = v
		default:
			iNdEx = preIndex
			skippy, err := skipAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAsset
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

func skipAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
					return 0, ErrIntOverflowAsset
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
				return 0, ErrInvalidLengthAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAsset = fmt.Errorf("proto: unexpected end of group")
)
