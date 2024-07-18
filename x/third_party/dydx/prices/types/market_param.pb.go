// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/prices/market_param.proto

package types

import (
	fmt "fmt"
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

// MarketParam represents the x/prices configuration for markets, including
// representing price values, resolving markets on individual exchanges, and
// generating price updates. This configuration is specific to the quote
// currency.
type MarketParam struct {
	// Unique, sequentially-generated value.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The human-readable name of the market pair (e.g. `BTC-USD`).
	Pair string `protobuf:"bytes,2,opt,name=pair,proto3" json:"pair,omitempty"`
	// Static value. The exponent of the price.
	// For example if `Exponent == -5` then a `Value` of `1,000,000,000`
	// represents “$10,000`. Therefore `10 ^ Exponent` represents the smallest
	// price step (in dollars) that can be recorded.
	Exponent int32 `protobuf:"zigzag32,3,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// The minimum number of exchanges that should be reporting a live price for
	// a price update to be considered valid.
	MinExchanges uint32 `protobuf:"varint,4,opt,name=min_exchanges,json=minExchanges,proto3" json:"min_exchanges,omitempty"`
	// The minimum allowable change in `price` value that would cause a price
	// update on the network. Measured as `1e-6` (parts per million).
	MinPriceChangePpm uint32 `protobuf:"varint,5,opt,name=min_price_change_ppm,json=minPriceChangePpm,proto3" json:"min_price_change_ppm,omitempty"`
	// A string of json that encodes the configuration for resolving the price
	// of this market on various exchanges.
	ExchangeConfigJson string `protobuf:"bytes,6,opt,name=exchange_config_json,json=exchangeConfigJson,proto3" json:"exchange_config_json,omitempty"`
}

func (m *MarketParam) Reset()         { *m = MarketParam{} }
func (m *MarketParam) String() string { return proto.CompactTextString(m) }
func (*MarketParam) ProtoMessage()    {}
func (*MarketParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea95d07ac52885cf, []int{0}
}
func (m *MarketParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketParam.Merge(m, src)
}
func (m *MarketParam) XXX_Size() int {
	return m.Size()
}
func (m *MarketParam) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketParam.DiscardUnknown(m)
}

var xxx_messageInfo_MarketParam proto.InternalMessageInfo

func (m *MarketParam) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MarketParam) GetPair() string {
	if m != nil {
		return m.Pair
	}
	return ""
}

func (m *MarketParam) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *MarketParam) GetMinExchanges() uint32 {
	if m != nil {
		return m.MinExchanges
	}
	return 0
}

func (m *MarketParam) GetMinPriceChangePpm() uint32 {
	if m != nil {
		return m.MinPriceChangePpm
	}
	return 0
}

func (m *MarketParam) GetExchangeConfigJson() string {
	if m != nil {
		return m.ExchangeConfigJson
	}
	return ""
}

func init() {
	proto.RegisterType((*MarketParam)(nil), "joltify.third_party.dydxprotocol.prices.MarketParam")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/prices/market_param.proto", fileDescriptor_ea95d07ac52885cf)
}

var fileDescriptor_ea95d07ac52885cf = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0x9b, 0x7e, 0xfd, 0x8a, 0x46, 0x2b, 0x34, 0xf4, 0xb0, 0x78, 0x58, 0x8a, 0x1e, 0xec,
	0xc5, 0x8d, 0xe0, 0xcd, 0xa3, 0xc5, 0x4b, 0x41, 0x28, 0x3d, 0x7a, 0x09, 0xdb, 0x4d, 0xba, 0x9d,
	0xda, 0x4c, 0x42, 0x36, 0xc2, 0xee, 0xbf, 0xf0, 0x67, 0x79, 0xec, 0x45, 0xf0, 0x28, 0xed, 0x1f,
	0x91, 0x4d, 0xbb, 0x52, 0xf0, 0x96, 0xcc, 0xf3, 0x3e, 0xc3, 0xf0, 0xd2, 0x87, 0x95, 0x59, 0x7b,
	0x58, 0x54, 0xdc, 0x2f, 0xc1, 0x49, 0x61, 0x53, 0xe7, 0x2b, 0x2e, 0x2b, 0x59, 0x5a, 0x67, 0xbc,
	0xc9, 0xcc, 0x9a, 0x5b, 0x07, 0x99, 0x2a, 0xb8, 0x4e, 0xdd, 0xab, 0xf2, 0x75, 0x20, 0xd5, 0x49,
	0x80, 0xec, 0xe6, 0xe0, 0x26, 0x47, 0x6e, 0x72, 0xec, 0x26, 0x7b, 0xf7, 0xea, 0x93, 0xd0, 0xb3,
	0xe7, 0xe0, 0x4f, 0x6b, 0x9d, 0x5d, 0xd0, 0x36, 0xc8, 0x88, 0x0c, 0xc9, 0xa8, 0x37, 0x6b, 0x83,
	0x64, 0x8c, 0x76, 0x6c, 0x0a, 0x2e, 0x6a, 0x0f, 0xc9, 0xe8, 0x74, 0x16, 0xde, 0xec, 0x92, 0x9e,
	0xa8, 0xd2, 0x1a, 0x54, 0xe8, 0xa3, 0x7f, 0x43, 0x32, 0xea, 0xcf, 0x7e, 0xff, 0xec, 0x9a, 0xf6,
	0x34, 0xa0, 0x50, 0x65, 0xb6, 0x4c, 0x31, 0x57, 0x45, 0xd4, 0x09, 0xab, 0xce, 0x35, 0xe0, 0x53,
	0x33, 0x63, 0x9c, 0x0e, 0xea, 0x50, 0x38, 0x41, 0xec, 0x87, 0xc2, 0x5a, 0x1d, 0xfd, 0x0f, 0xd9,
	0xbe, 0x06, 0x9c, 0xd6, 0x68, 0x1c, 0xc8, 0xd4, 0x6a, 0x76, 0x47, 0x07, 0xcd, 0x46, 0x91, 0x19,
	0x5c, 0x40, 0x2e, 0x56, 0x85, 0xc1, 0xa8, 0x1b, 0xae, 0x62, 0x0d, 0x1b, 0x07, 0x34, 0x29, 0x0c,
	0x3e, 0xca, 0x8f, 0x6d, 0x4c, 0x36, 0xdb, 0x98, 0x7c, 0x6f, 0x63, 0xf2, 0xbe, 0x8b, 0x5b, 0x9b,
	0x5d, 0xdc, 0xfa, 0xda, 0xc5, 0xad, 0x97, 0x49, 0x0e, 0x7e, 0xf9, 0x36, 0x4f, 0x32, 0xa3, 0xf9,
	0xa1, 0xa5, 0xdb, 0x05, 0x60, 0x8a, 0x99, 0x6a, 0xfe, 0x62, 0xad, 0x50, 0x02, 0xe6, 0xbc, 0xfc,
	0xd3, 0x7d, 0xd3, 0xb9, 0xaf, 0xac, 0x2a, 0xe6, 0xdd, 0x50, 0xe7, 0xfd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0xac, 0x0f, 0xb4, 0xab, 0x01, 0x00, 0x00,
}

func (m *MarketParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExchangeConfigJson) > 0 {
		i -= len(m.ExchangeConfigJson)
		copy(dAtA[i:], m.ExchangeConfigJson)
		i = encodeVarintMarketParam(dAtA, i, uint64(len(m.ExchangeConfigJson)))
		i--
		dAtA[i] = 0x32
	}
	if m.MinPriceChangePpm != 0 {
		i = encodeVarintMarketParam(dAtA, i, uint64(m.MinPriceChangePpm))
		i--
		dAtA[i] = 0x28
	}
	if m.MinExchanges != 0 {
		i = encodeVarintMarketParam(dAtA, i, uint64(m.MinExchanges))
		i--
		dAtA[i] = 0x20
	}
	if m.Exponent != 0 {
		i = encodeVarintMarketParam(dAtA, i, uint64((uint32(m.Exponent)<<1)^uint32((m.Exponent>>31))))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Pair) > 0 {
		i -= len(m.Pair)
		copy(dAtA[i:], m.Pair)
		i = encodeVarintMarketParam(dAtA, i, uint64(len(m.Pair)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMarketParam(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarketParam(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarketParam(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMarketParam(uint64(m.Id))
	}
	l = len(m.Pair)
	if l > 0 {
		n += 1 + l + sovMarketParam(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sozMarketParam(uint64(m.Exponent))
	}
	if m.MinExchanges != 0 {
		n += 1 + sovMarketParam(uint64(m.MinExchanges))
	}
	if m.MinPriceChangePpm != 0 {
		n += 1 + sovMarketParam(uint64(m.MinPriceChangePpm))
	}
	l = len(m.ExchangeConfigJson)
	if l > 0 {
		n += 1 + l + sovMarketParam(uint64(l))
	}
	return n
}

func sovMarketParam(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarketParam(x uint64) (n int) {
	return sovMarketParam(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketParam
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
			return fmt.Errorf("proto: MarketParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
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
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
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
				return ErrInvalidLengthMarketParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pair = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
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
			m.Exponent = v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinExchanges", wireType)
			}
			m.MinExchanges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinExchanges |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPriceChangePpm", wireType)
			}
			m.MinPriceChangePpm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPriceChangePpm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeConfigJson", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketParam
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
				return ErrInvalidLengthMarketParam
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarketParam
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExchangeConfigJson = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarketParam(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarketParam
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
func skipMarketParam(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketParam
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
					return 0, ErrIntOverflowMarketParam
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
					return 0, ErrIntOverflowMarketParam
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
				return 0, ErrInvalidLengthMarketParam
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarketParam
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarketParam
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarketParam        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketParam          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarketParam = fmt.Errorf("proto: unexpected end of group")
)
