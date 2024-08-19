// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/subaccounts/perpetual_position.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_joltify_finance_joltify_lending_dydx_helper_dtypes "github.com/joltify-finance/joltify_lending/dydx_helper/dtypes"
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

// PerpetualPositions are an account’s positions of a `Perpetual`.
// Therefore they hold any information needed to trade perpetuals.
type PerpetualPosition struct {
	// The `Id` of the `Perpetual`.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
	// The size of the position in base quantums.
	Quantums github_com_joltify_finance_joltify_lending_dydx_helper_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=quantums,proto3,customtype=github.com/joltify-finance/joltify_lending/dydx_helper/dtypes.SerializableInt" json:"quantums"`
	// The funding_index of the `Perpetual` the last time this position was
	// settled.
	FundingIndex github_com_joltify_finance_joltify_lending_dydx_helper_dtypes.SerializableInt `protobuf:"bytes,3,opt,name=funding_index,json=fundingIndex,proto3,customtype=github.com/joltify-finance/joltify_lending/dydx_helper/dtypes.SerializableInt" json:"funding_index"`
	// The quote_balance of the `Perpetual`.
	QuoteBalance github_com_joltify_finance_joltify_lending_dydx_helper_dtypes.SerializableInt `protobuf:"bytes,4,opt,name=quote_balance,json=quoteBalance,proto3,customtype=github.com/joltify-finance/joltify_lending/dydx_helper/dtypes.SerializableInt" json:"quote_balance"`
}

func (m *PerpetualPosition) Reset()         { *m = PerpetualPosition{} }
func (m *PerpetualPosition) String() string { return proto.CompactTextString(m) }
func (*PerpetualPosition) ProtoMessage()    {}
func (*PerpetualPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_8be148b283b0825d, []int{0}
}
func (m *PerpetualPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpetualPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpetualPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpetualPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpetualPosition.Merge(m, src)
}
func (m *PerpetualPosition) XXX_Size() int {
	return m.Size()
}
func (m *PerpetualPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpetualPosition.DiscardUnknown(m)
}

var xxx_messageInfo_PerpetualPosition proto.InternalMessageInfo

func (m *PerpetualPosition) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

func init() {
	proto.RegisterType((*PerpetualPosition)(nil), "joltify.third_party.dydxprotocol.subaccounts.PerpetualPosition")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/subaccounts/perpetual_position.proto", fileDescriptor_8be148b283b0825d)
}

var fileDescriptor_8be148b283b0825d = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4e, 0xf3, 0x30,
	0x18, 0xc6, 0x93, 0xaf, 0x9f, 0x10, 0x0a, 0xed, 0x40, 0xc4, 0x10, 0x31, 0xa4, 0x85, 0xa9, 0x03,
	0xc4, 0x03, 0x37, 0xa8, 0xc4, 0xd0, 0x01, 0x51, 0x15, 0xb1, 0xb0, 0x58, 0x4e, 0xec, 0xa6, 0x46,
	0xae, 0xed, 0x26, 0xaf, 0xa5, 0xa6, 0xa7, 0xe0, 0x0a, 0xdc, 0xa6, 0x63, 0x47, 0xc4, 0x50, 0xa1,
	0xe6, 0x22, 0x28, 0x6e, 0x5a, 0xc2, 0xc8, 0xd0, 0xcd, 0x7f, 0x7f, 0xbf, 0xe7, 0x95, 0x1e, 0xef,
	0xfe, 0x55, 0x09, 0xe0, 0x93, 0x02, 0xc1, 0x94, 0x67, 0x14, 0x6b, 0x92, 0x41, 0x81, 0x68, 0x41,
	0x17, 0x3a, 0x53, 0xa0, 0x12, 0x25, 0x50, 0x6e, 0x62, 0x92, 0x24, 0xca, 0x48, 0xc8, 0x91, 0x66,
	0x99, 0x66, 0x60, 0x88, 0xc0, 0x5a, 0xe5, 0x1c, 0xb8, 0x92, 0x91, 0x7d, 0xe7, 0xdf, 0xd4, 0x98,
	0xa8, 0x81, 0x89, 0x9a, 0x98, 0xa8, 0x81, 0xb9, 0xbc, 0x48, 0x55, 0xaa, 0xec, 0x0d, 0xaa, 0x56,
	0x3b, 0xc6, 0xf5, 0x7b, 0xcb, 0x3b, 0x1f, 0xed, 0x05, 0xa3, 0x9a, 0xef, 0x5f, 0x79, 0xed, 0x1f,
	0x2b, 0xa7, 0x81, 0xdb, 0x73, 0xfb, 0x9d, 0xf1, 0xd9, 0xe1, 0x6c, 0x48, 0xfd, 0xb9, 0x77, 0x3a,
	0x37, 0x44, 0x82, 0x99, 0xe5, 0xc1, 0xbf, 0x9e, 0xdb, 0x6f, 0x0f, 0x9e, 0x57, 0x9b, 0xae, 0xf3,
	0xb9, 0xe9, 0x3e, 0xa4, 0x1c, 0xa6, 0x26, 0x8e, 0x12, 0x35, 0x43, 0x75, 0xc2, 0xdb, 0x09, 0x97,
	0x44, 0x26, 0x6c, 0xbf, 0xc7, 0x82, 0x49, 0xca, 0x65, 0x6a, 0x87, 0xc6, 0x53, 0x26, 0x34, 0xcb,
	0x10, 0x85, 0x42, 0xb3, 0x3c, 0x7a, 0x62, 0x19, 0x27, 0x82, 0x2f, 0x49, 0x2c, 0xd8, 0x50, 0xc2,
	0xf8, 0xa0, 0xf1, 0x97, 0x5e, 0x67, 0x62, 0xec, 0x3f, 0xcc, 0x25, 0x65, 0x8b, 0xa0, 0x75, 0x4c,
	0x6f, 0xbb, 0x76, 0x0d, 0x2b, 0x55, 0xe5, 0x9e, 0x1b, 0x05, 0x0c, 0xc7, 0x44, 0x54, 0xc4, 0xe0,
	0xff, 0x51, 0xdd, 0xd6, 0x35, 0xd8, 0xa9, 0x06, 0x7c, 0xb5, 0x0d, 0xdd, 0xf5, 0x36, 0x74, 0xbf,
	0xb6, 0xa1, 0xfb, 0x56, 0x86, 0xce, 0xba, 0x0c, 0x9d, 0x8f, 0x32, 0x74, 0x5e, 0x1e, 0xff, 0xa0,
	0x5d, 0x34, 0xdb, 0x86, 0xab, 0x10, 0xbf, 0x5a, 0x66, 0x83, 0xc4, 0x27, 0xb6, 0x15, 0x77, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xcb, 0x22, 0xe8, 0xa2, 0x02, 0x00, 0x00,
}

func (m *PerpetualPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpetualPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpetualPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.QuoteBalance.Size()
		i -= size
		if _, err := m.QuoteBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.FundingIndex.Size()
		i -= size
		if _, err := m.FundingIndex.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Quantums.Size()
		i -= size
		if _, err := m.Quantums.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PerpetualId != 0 {
		i = encodeVarintPerpetualPosition(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPerpetualPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovPerpetualPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerpetualPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovPerpetualPosition(uint64(m.PerpetualId))
	}
	l = m.Quantums.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	l = m.FundingIndex.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	l = m.QuoteBalance.Size()
	n += 1 + l + sovPerpetualPosition(uint64(l))
	return n
}

func sovPerpetualPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPerpetualPosition(x uint64) (n int) {
	return sovPerpetualPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpetualPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerpetualPosition
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
			return fmt.Errorf("proto: PerpetualPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpetualPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerpetualId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantums.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingIndex", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FundingIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteBalance", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerpetualPosition
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
				return ErrInvalidLengthPerpetualPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPerpetualPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QuoteBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerpetualPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerpetualPosition
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
func skipPerpetualPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerpetualPosition
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
					return 0, ErrIntOverflowPerpetualPosition
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
					return 0, ErrIntOverflowPerpetualPosition
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
				return 0, ErrInvalidLengthPerpetualPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPerpetualPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPerpetualPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPerpetualPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerpetualPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPerpetualPosition = fmt.Errorf("proto: unexpected end of group")
)