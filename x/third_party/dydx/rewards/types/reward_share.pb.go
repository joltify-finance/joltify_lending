// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/rewards/reward_share.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_joltify_third_party_dydxprotocol_v4_chain_protocol_dtypes "github.com/joltify/third_party/dydxprotocol/v4-chain/protocol/dtypes"
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

// RewardShare stores the relative weight of rewards that each address is
// entitled to.
type RewardShare struct {
	Address string                                                                               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Weight  github_com_joltify_third_party_dydxprotocol_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/joltify/third_party/dydxprotocol/v4-chain/protocol/dtypes.SerializableInt" json:"weight"`
}

func (m *RewardShare) Reset()         { *m = RewardShare{} }
func (m *RewardShare) String() string { return proto.CompactTextString(m) }
func (*RewardShare) ProtoMessage()    {}
func (*RewardShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_b002633abe6bc335, []int{0}
}
func (m *RewardShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardShare.Merge(m, src)
}
func (m *RewardShare) XXX_Size() int {
	return m.Size()
}
func (m *RewardShare) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardShare.DiscardUnknown(m)
}

var xxx_messageInfo_RewardShare proto.InternalMessageInfo

func (m *RewardShare) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*RewardShare)(nil), "joltify.third_party.dydxprotocol.rewards.RewardShare")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/rewards/reward_share.proto", fileDescriptor_b002633abe6bc335)
}

var fileDescriptor_b002633abe6bc335 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x86, 0x22, 0x02, 0x53, 0xd5, 0xa1, 0x74, 0x70, 0x2b, 0xa6, 0x2c, 0xb1, 0x25,
	0x60, 0x63, 0x22, 0x1b, 0x62, 0x4b, 0x98, 0x10, 0x52, 0xe4, 0xc4, 0xae, 0x63, 0x94, 0xda, 0x91,
	0x6d, 0x68, 0xc3, 0x29, 0x38, 0x0c, 0x12, 0x57, 0xe8, 0x58, 0x31, 0x21, 0x86, 0x0a, 0x25, 0x17,
	0x41, 0x4d, 0xd2, 0xaa, 0x52, 0x07, 0x26, 0xfb, 0xfd, 0xbf, 0xbf, 0xf7, 0x3f, 0xeb, 0xb9, 0x37,
	0xcf, 0x2a, 0xb7, 0x62, 0x5a, 0x62, 0x9b, 0x09, 0x4d, 0xe3, 0x82, 0x68, 0x5b, 0x62, 0x5a, 0xd2,
	0x45, 0xa1, 0x95, 0x55, 0xa9, 0xca, 0xb1, 0x66, 0x73, 0xa2, 0xa9, 0xe9, 0xce, 0xd8, 0x64, 0x44,
	0x33, 0xd4, 0xb8, 0x7d, 0xaf, 0x83, 0xd1, 0x1e, 0x8c, 0xf6, 0x61, 0xd4, 0xc1, 0xa3, 0xf3, 0x54,
	0x99, 0x99, 0x32, 0x71, 0x63, 0xe0, 0xb6, 0x68, 0x9b, 0x8c, 0x06, 0x5c, 0x71, 0xd5, 0xea, 0x9b,
	0x5b, 0xab, 0x5e, 0x7c, 0x02, 0xf7, 0x34, 0x6c, 0xe0, 0x68, 0x13, 0xd8, 0xbf, 0x74, 0x8f, 0x09,
	0xa5, 0x9a, 0x19, 0x33, 0x04, 0x13, 0xe0, 0x9d, 0x04, 0xc3, 0xaf, 0x0f, 0x7f, 0xd0, 0x35, 0xba,
	0x6d, 0x9d, 0xc8, 0x6a, 0x21, 0x79, 0xb8, 0x7d, 0xd8, 0xb7, 0x6e, 0x6f, 0xce, 0x04, 0xcf, 0xec,
	0xf0, 0x68, 0x02, 0xbc, 0xb3, 0xe0, 0x69, 0xb9, 0x1e, 0x3b, 0x3f, 0xeb, 0xf1, 0x03, 0x17, 0x36,
	0x7b, 0x49, 0x50, 0xaa, 0x66, 0xf8, 0xdf, 0xef, 0xbf, 0x5e, 0xfb, 0x69, 0x46, 0x84, 0xc4, 0x3b,
	0x85, 0xda, 0xb2, 0x60, 0x06, 0x45, 0x4c, 0x0b, 0x92, 0x8b, 0x37, 0x92, 0xe4, 0xec, 0x4e, 0xda,
	0xb0, 0xcb, 0x0a, 0xd8, 0xb2, 0x82, 0x60, 0x55, 0x41, 0xf0, 0x5b, 0x41, 0xf0, 0x5e, 0x43, 0x67,
	0x55, 0x43, 0xe7, 0xbb, 0x86, 0xce, 0xe3, 0xfd, 0x61, 0xae, 0x3f, 0x15, 0x92, 0xc8, 0x94, 0x6d,
	0xeb, 0x38, 0x67, 0x92, 0x0a, 0xc9, 0xf1, 0xe2, 0x60, 0xa2, 0xdd, 0x22, 0x9a, 0xf8, 0xa4, 0xd7,
	0xcc, 0x73, 0xf5, 0x17, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xed, 0xa3, 0x37, 0xc1, 0x01, 0x00, 0x00,
}

func (m *RewardShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewardShare(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRewardShare(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRewardShare(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovRewardShare(uint64(l))
	return n
}

func sovRewardShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardShare(x uint64) (n int) {
	return sovRewardShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardShare
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
			return fmt.Errorf("proto: RewardShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardShare
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
				return ErrInvalidLengthRewardShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardShare
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
				return ErrInvalidLengthRewardShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewardShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewardShare
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
func skipRewardShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardShare
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
					return 0, ErrIntOverflowRewardShare
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
					return 0, ErrIntOverflowRewardShare
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
				return 0, ErrInvalidLengthRewardShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardShare = fmt.Errorf("proto: unexpected end of group")
)
