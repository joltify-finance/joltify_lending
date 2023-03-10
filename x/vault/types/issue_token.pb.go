// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/issue_token.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type IssueToken struct {
	Creator  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Index    string                                        `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Coin     *github_com_cosmos_cosmos_sdk_types.Coin      `protobuf:"bytes,3,opt,name=coin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin,omitempty"`
	Receiver github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=receiver,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"receiver,omitempty"`
}

func (m *IssueToken) Reset()         { *m = IssueToken{} }
func (m *IssueToken) String() string { return proto.CompactTextString(m) }
func (*IssueToken) ProtoMessage()    {}
func (*IssueToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_299d4465e592eb2f, []int{0}
}
func (m *IssueToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IssueToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IssueToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IssueToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueToken.Merge(m, src)
}
func (m *IssueToken) XXX_Size() int {
	return m.Size()
}
func (m *IssueToken) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueToken.DiscardUnknown(m)
}

var xxx_messageInfo_IssueToken proto.InternalMessageInfo

func (m *IssueToken) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *IssueToken) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *IssueToken) GetReceiver() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func init() {
	proto.RegisterType((*IssueToken)(nil), "joltify.vault.IssueToken")
}

func init() { proto.RegisterFile("joltify/vault/issue_token.proto", fileDescriptor_299d4465e592eb2f) }

var fileDescriptor_299d4465e592eb2f = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x2f, 0x4b, 0x2c, 0xcd, 0x29, 0xd1, 0xcf, 0x2c, 0x2e, 0x2e, 0x4d, 0x8d,
	0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x2a, 0xd0,
	0x03, 0x2b, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xe8, 0x83, 0x58, 0x10, 0x45, 0x4a,
	0xcd, 0x4c, 0x5c, 0x5c, 0x9e, 0x20, 0xad, 0x21, 0x20, 0x9d, 0x42, 0xde, 0x5c, 0xec, 0xc9, 0x45,
	0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x86, 0xbf, 0xee, 0xc9,
	0xeb, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7,
	0xe6, 0x17, 0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x3d, 0xc7,
	0xe4, 0x64, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0x20, 0x98, 0x09, 0x42, 0x22, 0x5c, 0xac,
	0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x33,
	0x17, 0x4b, 0x72, 0x7e, 0x66, 0x9e, 0x04, 0x33, 0xd8, 0x7c, 0xfd, 0x13, 0xf7, 0xe4, 0x19, 0x6f,
	0xdd, 0x93, 0x57, 0x27, 0xc2, 0x0e, 0xe7, 0xfc, 0xcc, 0xbc, 0x20, 0xb0, 0x66, 0x21, 0x5f, 0x2e,
	0x8e, 0xa2, 0xd4, 0xe4, 0xd4, 0xcc, 0xb2, 0xd4, 0x22, 0x09, 0x16, 0x72, 0x1d, 0x0a, 0x37, 0xc2,
	0x29, 0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x2c, 0x90, 0x8c, 0x84,
	0x86, 0xa7, 0x6e, 0x5a, 0x66, 0x5e, 0x62, 0x5e, 0x72, 0x2a, 0x8c, 0x1f, 0x9f, 0x93, 0x9a, 0x97,
	0x92, 0x99, 0x97, 0xae, 0x5f, 0x01, 0x8d, 0x0a, 0xb0, 0x45, 0x49, 0x6c, 0xe0, 0x00, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xba, 0xf6, 0xa6, 0xe4, 0xa8, 0x01, 0x00, 0x00,
}

func (m *IssueToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IssueToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IssueToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintIssueToken(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x22
	}
	if m.Coin != nil {
		{
			size := m.Coin.Size()
			i -= size
			if _, err := m.Coin.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintIssueToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintIssueToken(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIssueToken(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssueToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssueToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IssueToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIssueToken(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovIssueToken(uint64(l))
	}
	if m.Coin != nil {
		l = m.Coin.Size()
		n += 1 + l + sovIssueToken(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovIssueToken(uint64(l))
	}
	return n
}

func sovIssueToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssueToken(x uint64) (n int) {
	return sovIssueToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IssueToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssueToken
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
			return fmt.Errorf("proto: IssueToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IssueToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssueToken
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
				return ErrInvalidLengthIssueToken
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIssueToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssueToken
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
				return ErrInvalidLengthIssueToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssueToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssueToken
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
				return ErrInvalidLengthIssueToken
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIssueToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Coin = &v
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssueToken
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
				return ErrInvalidLengthIssueToken
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIssueToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = append(m.Receiver[:0], dAtA[iNdEx:postIndex]...)
			if m.Receiver == nil {
				m.Receiver = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssueToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIssueToken
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
func skipIssueToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssueToken
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
					return 0, ErrIntOverflowIssueToken
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
					return 0, ErrIntOverflowIssueToken
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
				return 0, ErrInvalidLengthIssueToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssueToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssueToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssueToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssueToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssueToken = fmt.Errorf("proto: unexpected end of group")
)
