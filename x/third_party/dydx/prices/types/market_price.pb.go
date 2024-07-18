// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/prices/market_price.proto

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

// MarketPrice is used by the application to store/retrieve oracle price.
type MarketPrice struct {
	// Unique, sequentially-generated value that matches `MarketParam`.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Static value. The exponent of the price. See the comment on the duplicate
	// MarketParam field for more information.
	Exponent int32 `protobuf:"zigzag32,2,opt,name=exponent,proto3" json:"exponent,omitempty"`
	// The variable value that is updated by oracle price updates. `0` if it has
	// never been updated, `>0` otherwise.
	Price uint64 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *MarketPrice) Reset()         { *m = MarketPrice{} }
func (m *MarketPrice) String() string { return proto.CompactTextString(m) }
func (*MarketPrice) ProtoMessage()    {}
func (*MarketPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d6aa4f2bbfea6d9, []int{0}
}
func (m *MarketPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketPrice.Merge(m, src)
}
func (m *MarketPrice) XXX_Size() int {
	return m.Size()
}
func (m *MarketPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MarketPrice proto.InternalMessageInfo

func (m *MarketPrice) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MarketPrice) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *MarketPrice) GetPrice() uint64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*MarketPrice)(nil), "joltify.third_party.dydxprotocol.prices.MarketPrice")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/prices/market_price.proto", fileDescriptor_7d6aa4f2bbfea6d9)
}

var fileDescriptor_7d6aa4f2bbfea6d9 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xca, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4,
	0x4f, 0xa9, 0x4c, 0xa9, 0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x28, 0xca,
	0x4c, 0x4e, 0x2d, 0xd6, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x89, 0x07, 0xf3, 0xf4, 0xc0, 0x92,
	0x42, 0xea, 0x50, 0xbd, 0x7a, 0x48, 0x7a, 0xf5, 0x90, 0xf5, 0xea, 0x41, 0xf4, 0x2a, 0xf9, 0x73,
	0x71, 0xfb, 0x82, 0xb5, 0x07, 0x80, 0xf8, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0xbc, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x52, 0x5c, 0x1c, 0xa9, 0x15, 0x05, 0xf9, 0x79,
	0xa9, 0x79, 0x25, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x82, 0x41, 0x70, 0xbe, 0x90, 0x08, 0x17, 0x2b,
	0xd8, 0x10, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0xc7, 0x29, 0xe5, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xbc, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xa1, 0xce, 0xd3, 0x4d, 0xcb, 0xcc, 0x4b, 0xcc, 0x4b, 0x4e, 0x85, 0xf1, 0xe3,
	0x73, 0x52, 0xf3, 0x52, 0x32, 0xf3, 0xd2, 0xf5, 0x2b, 0x30, 0x3c, 0x0d, 0xf3, 0x6c, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x1f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x83,
	0x5c, 0xd1, 0x24, 0x01, 0x00, 0x00,
}

func (m *MarketPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != 0 {
		i = encodeVarintMarketPrice(dAtA, i, uint64(m.Price))
		i--
		dAtA[i] = 0x18
	}
	if m.Exponent != 0 {
		i = encodeVarintMarketPrice(dAtA, i, uint64((uint32(m.Exponent)<<1)^uint32((m.Exponent>>31))))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintMarketPrice(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarketPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarketPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MarketPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMarketPrice(uint64(m.Id))
	}
	if m.Exponent != 0 {
		n += 1 + sozMarketPrice(uint64(m.Exponent))
	}
	if m.Price != 0 {
		n += 1 + sovMarketPrice(uint64(m.Price))
	}
	return n
}

func sovMarketPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarketPrice(x uint64) (n int) {
	return sovMarketPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MarketPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarketPrice
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
			return fmt.Errorf("proto: MarketPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPrice
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPrice
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			m.Price = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarketPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Price |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMarketPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarketPrice
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
func skipMarketPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarketPrice
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
					return 0, ErrIntOverflowMarketPrice
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
					return 0, ErrIntOverflowMarketPrice
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
				return 0, ErrInvalidLengthMarketPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarketPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarketPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarketPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarketPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarketPrice = fmt.Errorf("proto: unexpected end of group")
)
