// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/create_pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_x_staking_types "github.com/cosmos/cosmos-sdk/x/staking/types"
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

type PoolProposal struct {
	PoolPubKey string                                          `protobuf:"bytes,1,opt,name=poolPubKey,proto3" json:"poolPubKey,omitempty"`
	PoolAddr   github_com_cosmos_cosmos_sdk_types.AccAddress   `protobuf:"bytes,2,opt,name=poolAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"poolAddr,omitempty"`
	Nodes      []github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,rep,name=nodes,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"nodes,omitempty"`
}

func (m *PoolProposal) Reset()         { *m = PoolProposal{} }
func (m *PoolProposal) String() string { return proto.CompactTextString(m) }
func (*PoolProposal) ProtoMessage()    {}
func (*PoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea4a3439134470a, []int{0}
}
func (m *PoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolProposal.Merge(m, src)
}
func (m *PoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *PoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PoolProposal proto.InternalMessageInfo

func (m *PoolProposal) GetPoolPubKey() string {
	if m != nil {
		return m.PoolPubKey
	}
	return ""
}

func (m *PoolProposal) GetPoolAddr() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.PoolAddr
	}
	return nil
}

func (m *PoolProposal) GetNodes() []github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type CreatePool struct {
	BlockHeight string                                                   `protobuf:"bytes,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Validators  []github_com_cosmos_cosmos_sdk_x_staking_types.Validator `protobuf:"bytes,2,rep,name=validators,proto3,customtype=github.com/cosmos/cosmos-sdk/x/staking/types.Validator" json:"validators"`
	Proposal    []*PoolProposal                                          `protobuf:"bytes,3,rep,name=proposal,proto3" json:"proposal,omitempty"`
}

func (m *CreatePool) Reset()         { *m = CreatePool{} }
func (m *CreatePool) String() string { return proto.CompactTextString(m) }
func (*CreatePool) ProtoMessage()    {}
func (*CreatePool) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ea4a3439134470a, []int{1}
}
func (m *CreatePool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePool.Merge(m, src)
}
func (m *CreatePool) XXX_Size() int {
	return m.Size()
}
func (m *CreatePool) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePool.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePool proto.InternalMessageInfo

func (m *CreatePool) GetBlockHeight() string {
	if m != nil {
		return m.BlockHeight
	}
	return ""
}

func (m *CreatePool) GetProposal() []*PoolProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func init() {
	proto.RegisterType((*PoolProposal)(nil), "joltify.vault.PoolProposal")
	proto.RegisterType((*CreatePool)(nil), "joltify.vault.CreatePool")
}

func init() { proto.RegisterFile("joltify/vault/create_pool.proto", fileDescriptor_7ea4a3439134470a) }

var fileDescriptor_7ea4a3439134470a = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0xeb, 0x30,
	0x10, 0xc7, 0xe3, 0x56, 0xef, 0xa9, 0xcf, 0xed, 0x5b, 0x22, 0x86, 0x08, 0xa4, 0x24, 0xea, 0xd4,
	0xa5, 0x89, 0x00, 0x09, 0x98, 0x90, 0x5a, 0x06, 0x90, 0x10, 0x52, 0x95, 0x81, 0x81, 0x81, 0xca,
	0x71, 0xdc, 0x34, 0xd4, 0xcd, 0x45, 0xb1, 0x5b, 0xb5, 0xdf, 0x82, 0x4f, 0x05, 0x1d, 0x3b, 0x22,
	0x86, 0x0a, 0xb5, 0xdf, 0x82, 0x09, 0xc5, 0x75, 0xab, 0xb0, 0x30, 0x30, 0x25, 0x77, 0xfe, 0xdf,
	0xcf, 0xff, 0xbb, 0x33, 0x76, 0x9e, 0x80, 0xcb, 0x64, 0x30, 0xf7, 0xa7, 0x64, 0xc2, 0xa5, 0x4f,
	0x73, 0x46, 0x24, 0xeb, 0x67, 0x00, 0xdc, 0xcb, 0x72, 0x90, 0x60, 0xfe, 0xd7, 0x02, 0x4f, 0x09,
	0x0e, 0x0f, 0x62, 0x88, 0x41, 0x9d, 0xf8, 0xc5, 0xdf, 0x56, 0xd4, 0x7c, 0x41, 0xb8, 0xd1, 0x03,
	0xe0, 0xbd, 0x1c, 0x32, 0x10, 0x84, 0x9b, 0x36, 0xc6, 0x05, 0xa3, 0x37, 0x09, 0x6f, 0xd9, 0xdc,
	0x42, 0x2e, 0x6a, 0xfd, 0x0b, 0x4a, 0x19, 0xf3, 0x0e, 0xd7, 0x8a, 0xa8, 0x13, 0x45, 0xb9, 0x55,
	0x71, 0x51, 0xab, 0xd1, 0x3d, 0xfe, 0x5c, 0x39, 0xed, 0x38, 0x91, 0xc3, 0x49, 0xe8, 0x51, 0x18,
	0xfb, 0x14, 0xc4, 0x18, 0x84, 0xfe, 0xb4, 0x45, 0x34, 0xf2, 0xe5, 0x3c, 0x63, 0xc2, 0xeb, 0x50,
	0x5a, 0x54, 0x31, 0x21, 0x82, 0x3d, 0xc2, 0xbc, 0xc6, 0x7f, 0x52, 0x88, 0x98, 0xb0, 0xaa, 0x6e,
	0xf5, 0x77, 0xac, 0x6d, 0x7d, 0xf3, 0x15, 0x61, 0x7c, 0xa5, 0x66, 0x50, 0xb4, 0x63, 0xba, 0xb8,
	0x1e, 0x72, 0xa0, 0xa3, 0x1b, 0x96, 0xc4, 0x43, 0xa9, 0xfb, 0x28, 0xa7, 0xcc, 0x47, 0x8c, 0xa7,
	0x84, 0x27, 0x11, 0x91, 0x90, 0x0b, 0xab, 0xa2, 0xae, 0xbf, 0x5c, 0xac, 0x1c, 0xe3, 0x7d, 0xe5,
	0x9c, 0xfd, 0x68, 0x61, 0xe6, 0x0b, 0x49, 0x46, 0x49, 0x1a, 0x6b, 0x33, 0xf7, 0x3b, 0x4c, 0x50,
	0x22, 0x9a, 0xe7, 0xb8, 0x96, 0xe9, 0xa1, 0xaa, 0xe6, 0xea, 0x27, 0x47, 0xde, 0xb7, 0x8d, 0x78,
	0xe5, 0xb9, 0x07, 0x7b, 0x71, 0x37, 0x58, 0xac, 0x6d, 0xb4, 0x5c, 0xdb, 0xe8, 0x63, 0x6d, 0xa3,
	0xe7, 0x8d, 0x6d, 0x2c, 0x37, 0xb6, 0xf1, 0xb6, 0xb1, 0x8d, 0x87, 0x8b, 0x92, 0x2d, 0x8d, 0x6a,
	0x0f, 0x92, 0x94, 0xa4, 0x94, 0xed, 0xe2, 0x3e, 0x67, 0x69, 0x54, 0x58, 0x9b, 0xe9, 0x77, 0xa1,
	0x2c, 0x86, 0x7f, 0xd5, 0xb6, 0x4f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x71, 0x4b, 0xd0,
	0x35, 0x02, 0x00, 0x00,
}

func (m *PoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Nodes[iNdEx])
			copy(dAtA[i:], m.Nodes[iNdEx])
			i = encodeVarintCreatePool(dAtA, i, uint64(len(m.Nodes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PoolAddr) > 0 {
		i -= len(m.PoolAddr)
		copy(dAtA[i:], m.PoolAddr)
		i = encodeVarintCreatePool(dAtA, i, uint64(len(m.PoolAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolPubKey) > 0 {
		i -= len(m.PoolPubKey)
		copy(dAtA[i:], m.PoolPubKey)
		i = encodeVarintCreatePool(dAtA, i, uint64(len(m.PoolPubKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreatePool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proposal) > 0 {
		for iNdEx := len(m.Proposal) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposal[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCreatePool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Validators[iNdEx].Size()
				i -= size
				if _, err := m.Validators[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintCreatePool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BlockHeight) > 0 {
		i -= len(m.BlockHeight)
		copy(dAtA[i:], m.BlockHeight)
		i = encodeVarintCreatePool(dAtA, i, uint64(len(m.BlockHeight)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCreatePool(dAtA []byte, offset int, v uint64) int {
	offset -= sovCreatePool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolPubKey)
	if l > 0 {
		n += 1 + l + sovCreatePool(uint64(l))
	}
	l = len(m.PoolAddr)
	if l > 0 {
		n += 1 + l + sovCreatePool(uint64(l))
	}
	if len(m.Nodes) > 0 {
		for _, b := range m.Nodes {
			l = len(b)
			n += 1 + l + sovCreatePool(uint64(l))
		}
	}
	return n
}

func (m *CreatePool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockHeight)
	if l > 0 {
		n += 1 + l + sovCreatePool(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovCreatePool(uint64(l))
		}
	}
	if len(m.Proposal) > 0 {
		for _, e := range m.Proposal {
			l = e.Size()
			n += 1 + l + sovCreatePool(uint64(l))
		}
	}
	return n
}

func sovCreatePool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCreatePool(x uint64) (n int) {
	return sovCreatePool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCreatePool
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
			return fmt.Errorf("proto: PoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolPubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAddr = append(m.PoolAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.PoolAddr == nil {
				m.PoolAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, make([]byte, postIndex-iNdEx))
			copy(m.Nodes[len(m.Nodes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCreatePool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCreatePool
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
func (m *CreatePool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCreatePool
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
			return fmt.Errorf("proto: CreatePool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_x_staking_types.Validator
			m.Validators = append(m.Validators, v)
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCreatePool
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
				return ErrInvalidLengthCreatePool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCreatePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposal = append(m.Proposal, &PoolProposal{})
			if err := m.Proposal[len(m.Proposal)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCreatePool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCreatePool
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
func skipCreatePool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCreatePool
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
					return 0, ErrIntOverflowCreatePool
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
					return 0, ErrIntOverflowCreatePool
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
				return 0, ErrInvalidLengthCreatePool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCreatePool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCreatePool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCreatePool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCreatePool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCreatePool = fmt.Errorf("proto: unexpected end of group")
)
