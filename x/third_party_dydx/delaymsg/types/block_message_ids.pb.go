// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/delaymsg/block_message_ids.proto

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

// BlockMessageIds stores the id of each message that should be processed at a
// given block height.
type BlockMessageIds struct {
	// ids stores a list of DelayedMessage ids that should be processed at a given
	// block height.
	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (m *BlockMessageIds) Reset()         { *m = BlockMessageIds{} }
func (m *BlockMessageIds) String() string { return proto.CompactTextString(m) }
func (*BlockMessageIds) ProtoMessage()    {}
func (*BlockMessageIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_14c7b543b4bc309e, []int{0}
}
func (m *BlockMessageIds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockMessageIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockMessageIds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockMessageIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMessageIds.Merge(m, src)
}
func (m *BlockMessageIds) XXX_Size() int {
	return m.Size()
}
func (m *BlockMessageIds) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMessageIds.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMessageIds proto.InternalMessageInfo

func (m *BlockMessageIds) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockMessageIds)(nil), "joltify.third_party.dydxprotocol.delaymsg.BlockMessageIds")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/delaymsg/block_message_ids.proto", fileDescriptor_14c7b543b4bc309e)
}

var fileDescriptor_14c7b543b4bc309e = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcc, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4,
	0x4f, 0xa9, 0x4c, 0xa9, 0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x49, 0xcd,
	0x49, 0xac, 0xcc, 0x2d, 0x4e, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x8e, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0xcf, 0x4c, 0x29, 0xd6, 0x03, 0xab, 0x11, 0xd2, 0x84, 0x1a, 0xa1, 0x87,
	0x64, 0x84, 0x1e, 0xb2, 0x11, 0x7a, 0x30, 0x23, 0x94, 0x94, 0xb9, 0xf8, 0x9d, 0x40, 0xa6, 0xf8,
	0x42, 0x0c, 0xf1, 0x4c, 0x29, 0x16, 0x12, 0xe0, 0x62, 0xce, 0x4c, 0x29, 0x96, 0x60, 0x54, 0x60,
	0xd6, 0xe0, 0x0d, 0x02, 0x31, 0x9d, 0xd2, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e,
	0x21, 0xca, 0x27, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x6a, 0xa9,
	0x6e, 0x5a, 0x66, 0x5e, 0x62, 0x5e, 0x72, 0x2a, 0x8c, 0x1f, 0x9f, 0x93, 0x9a, 0x97, 0x92, 0x99,
	0x97, 0xae, 0x5f, 0x81, 0xec, 0xa3, 0x78, 0x90, 0x73, 0x10, 0x3e, 0x29, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0xbb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x8b, 0x3f, 0x39, 0x03,
	0x01, 0x00, 0x00,
}

func (m *BlockMessageIds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockMessageIds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockMessageIds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintBlockMessageIds(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockMessageIds(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockMessageIds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockMessageIds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovBlockMessageIds(uint64(e))
		}
		n += 1 + sovBlockMessageIds(uint64(l)) + l
	}
	return n
}

func sovBlockMessageIds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockMessageIds(x uint64) (n int) {
	return sovBlockMessageIds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockMessageIds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockMessageIds
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
			return fmt.Errorf("proto: BlockMessageIds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockMessageIds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlockMessageIds
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlockMessageIds
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBlockMessageIds
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBlockMessageIds
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockMessageIds
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlockMessageIds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockMessageIds
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
func skipBlockMessageIds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockMessageIds
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
					return 0, ErrIntOverflowBlockMessageIds
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
					return 0, ErrIntOverflowBlockMessageIds
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
				return 0, ErrInvalidLengthBlockMessageIds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockMessageIds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockMessageIds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockMessageIds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockMessageIds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockMessageIds = fmt.Errorf("proto: unexpected end of group")
)
