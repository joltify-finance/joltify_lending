// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/spv/deposit.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type DepositorInfo struct {
	InvestorId              string                                        `protobuf:"bytes,1,opt,name=investor_id,json=investorId,proto3" json:"investor_id,omitempty"`
	DepositorAddress        github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=depositor_address,json=depositorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor_address,omitempty"`
	PoolIndex               string                                        `protobuf:"bytes,3,opt,name=pool_index,json=poolIndex,proto3" json:"pool_index,omitempty"`
	LockedAmount            types.Coin                                    `protobuf:"bytes,4,opt,name=locked_amount,json=lockedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"locked_amount"`
	WithdrawableAmount      types.Coin                                    `protobuf:"bytes,5,opt,name=withdrawable_amount,json=withdrawableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawable_amount"`
	ClaimableInterestAmount types.Coin                                    `protobuf:"bytes,6,opt,name=claimable_interest_amount,json=claimableInterestAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"claimable_interest_amount"`
	IncentiveAmount         types.Coin                                    `protobuf:"bytes,7,opt,name=incentive_amount,json=incentiveAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"incentive_amount"`
	LinkedNFT               []string                                      `protobuf:"bytes,8,rep,name=linkedNFT,proto3" json:"linkedNFT,omitempty"`
}

func (m *DepositorInfo) Reset()         { *m = DepositorInfo{} }
func (m *DepositorInfo) String() string { return proto.CompactTextString(m) }
func (*DepositorInfo) ProtoMessage()    {}
func (*DepositorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bb054114dab9679, []int{0}
}
func (m *DepositorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositorInfo.Merge(m, src)
}
func (m *DepositorInfo) XXX_Size() int {
	return m.Size()
}
func (m *DepositorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DepositorInfo proto.InternalMessageInfo

func (m *DepositorInfo) GetInvestorId() string {
	if m != nil {
		return m.InvestorId
	}
	return ""
}

func (m *DepositorInfo) GetDepositorAddress() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.DepositorAddress
	}
	return nil
}

func (m *DepositorInfo) GetPoolIndex() string {
	if m != nil {
		return m.PoolIndex
	}
	return ""
}

func (m *DepositorInfo) GetLockedAmount() types.Coin {
	if m != nil {
		return m.LockedAmount
	}
	return types.Coin{}
}

func (m *DepositorInfo) GetWithdrawableAmount() types.Coin {
	if m != nil {
		return m.WithdrawableAmount
	}
	return types.Coin{}
}

func (m *DepositorInfo) GetClaimableInterestAmount() types.Coin {
	if m != nil {
		return m.ClaimableInterestAmount
	}
	return types.Coin{}
}

func (m *DepositorInfo) GetIncentiveAmount() types.Coin {
	if m != nil {
		return m.IncentiveAmount
	}
	return types.Coin{}
}

func (m *DepositorInfo) GetLinkedNFT() []string {
	if m != nil {
		return m.LinkedNFT
	}
	return nil
}

func init() {
	proto.RegisterType((*DepositorInfo)(nil), "joltify.spv.DepositorInfo")
}

func init() { proto.RegisterFile("joltify/spv/deposit.proto", fileDescriptor_5bb054114dab9679) }

var fileDescriptor_5bb054114dab9679 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x4a, 0x0b, 0xb9, 0xb4, 0xa2, 0xb8, 0x48, 0x38, 0x15, 0xd8, 0x11, 0x53, 0x96,
	0xf8, 0x28, 0x48, 0xec, 0x09, 0x08, 0x91, 0x05, 0xa1, 0x88, 0x89, 0xc5, 0xb2, 0xef, 0x5e, 0xdc,
	0x23, 0xf6, 0x3d, 0xcb, 0x77, 0x71, 0x1b, 0x89, 0x9d, 0x95, 0xcf, 0xc1, 0xcc, 0x87, 0xe8, 0x58,
	0x31, 0x20, 0xa6, 0x82, 0x92, 0x6f, 0xc1, 0x84, 0x6c, 0x9f, 0xad, 0x8c, 0x1d, 0x32, 0xd9, 0xef,
	0xff, 0xde, 0xf3, 0xef, 0xff, 0x9e, 0xfc, 0x48, 0xff, 0x33, 0x26, 0x5a, 0xcc, 0x57, 0x54, 0x65,
	0x05, 0xe5, 0x90, 0xa1, 0x12, 0xda, 0xcf, 0x72, 0xd4, 0x68, 0xf7, 0x4c, 0xca, 0x57, 0x59, 0x71,
	0xda, 0x67, 0xa8, 0x52, 0x54, 0x41, 0x95, 0xa2, 0x75, 0x50, 0xd7, 0x9d, 0x3e, 0x8a, 0x31, 0xc6,
	0x5a, 0x2f, 0xdf, 0x8c, 0xea, 0xd6, 0x35, 0x34, 0x0a, 0x15, 0xd0, 0xe2, 0x2c, 0x02, 0x1d, 0x9e,
	0x51, 0x86, 0x42, 0x9a, 0xbc, 0x17, 0x23, 0xc6, 0x09, 0xd0, 0x2a, 0x8a, 0x96, 0x73, 0xaa, 0x45,
	0x0a, 0x4a, 0x87, 0x69, 0x56, 0x17, 0x3c, 0xfb, 0xb5, 0x4f, 0x8e, 0xde, 0xd4, 0x86, 0x30, 0x9f,
	0xca, 0x39, 0xda, 0x1e, 0xe9, 0x09, 0x59, 0x80, 0xd2, 0x98, 0x07, 0x82, 0x3b, 0xd6, 0xc0, 0x1a,
	0x76, 0x67, 0xa4, 0x91, 0xa6, 0xdc, 0x5e, 0x92, 0x87, 0xbc, 0xe9, 0x08, 0x42, 0xce, 0x73, 0x50,
	0xca, 0xb9, 0x33, 0xb0, 0x86, 0x87, 0x93, 0x77, 0xff, 0x6e, 0xbc, 0x51, 0x2c, 0xf4, 0xf9, 0x32,
	0xf2, 0x19, 0xa6, 0x66, 0x02, 0xf3, 0x18, 0x29, 0xbe, 0xa0, 0x7a, 0x95, 0x81, 0xf2, 0xc7, 0x8c,
	0x8d, 0xeb, 0xc6, 0x9f, 0x3f, 0x46, 0x27, 0x66, 0x4e, 0xa3, 0x4c, 0x56, 0x1a, 0xd4, 0xec, 0xb8,
	0x45, 0x18, 0xd9, 0x7e, 0x4a, 0x48, 0x86, 0x98, 0x04, 0x42, 0x72, 0xb8, 0x74, 0xf6, 0x2a, 0x5b,
	0xdd, 0x52, 0x99, 0x96, 0x82, 0x9d, 0x91, 0xa3, 0x04, 0xd9, 0x02, 0x78, 0x10, 0xa6, 0xb8, 0x94,
	0xda, 0xb9, 0x3b, 0xb0, 0x86, 0xbd, 0x17, 0x7d, 0xdf, 0x7c, 0xbd, 0xdc, 0x90, 0x6f, 0x36, 0xe4,
	0xbf, 0x46, 0x21, 0x27, 0xcf, 0xaf, 0x6e, 0xbc, 0xce, 0xf7, 0x3f, 0xde, 0xf0, 0x16, 0x86, 0xcb,
	0x06, 0x35, 0x3b, 0xac, 0x09, 0xe3, 0x0a, 0x60, 0x7f, 0x21, 0x27, 0x17, 0x42, 0x9f, 0xf3, 0x3c,
	0xbc, 0x08, 0xa3, 0x04, 0x1a, 0xee, 0xfe, 0xee, 0xb9, 0xf6, 0x36, 0xc7, 0xd0, 0xbf, 0x5a, 0xa4,
	0xcf, 0x92, 0x50, 0xa4, 0x15, 0x5b, 0x48, 0x0d, 0x39, 0x28, 0xdd, 0x98, 0x38, 0xd8, 0xbd, 0x89,
	0xc7, 0x2d, 0x6d, 0x6a, 0x60, 0xc6, 0x49, 0x41, 0x8e, 0x85, 0x64, 0x20, 0xb5, 0x28, 0xda, 0x25,
	0xdc, 0xdb, 0x3d, 0xff, 0x41, 0x0b, 0x31, 0xdc, 0x27, 0xa4, 0x9b, 0x08, 0xb9, 0x00, 0xfe, 0xfe,
	0xed, 0x47, 0xe7, 0xfe, 0x60, 0xaf, 0xfc, 0x1f, 0x5a, 0x61, 0xf2, 0xe1, 0x6a, 0xed, 0x5a, 0xd7,
	0x6b, 0xd7, 0xfa, 0xbb, 0x76, 0xad, 0x6f, 0x1b, 0xb7, 0x73, 0xbd, 0x71, 0x3b, 0xbf, 0x37, 0x6e,
	0xe7, 0xd3, 0xab, 0x2d, 0xa4, 0x39, 0xbe, 0xd1, 0x5c, 0xc8, 0x50, 0x32, 0x68, 0xe2, 0x20, 0x01,
	0xc9, 0x85, 0x8c, 0xe9, 0x65, 0x75, 0xb1, 0x95, 0x8d, 0xe8, 0xa0, 0xba, 0x98, 0x97, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0xc3, 0x32, 0xd4, 0xcd, 0x03, 0x00, 0x00,
}

func (m *DepositorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LinkedNFT) > 0 {
		for iNdEx := len(m.LinkedNFT) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LinkedNFT[iNdEx])
			copy(dAtA[i:], m.LinkedNFT[iNdEx])
			i = encodeVarintDeposit(dAtA, i, uint64(len(m.LinkedNFT[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.IncentiveAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.ClaimableInterestAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.WithdrawableAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.LockedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDeposit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.PoolIndex) > 0 {
		i -= len(m.PoolIndex)
		copy(dAtA[i:], m.PoolIndex)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.PoolIndex)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DepositorAddress) > 0 {
		i -= len(m.DepositorAddress)
		copy(dAtA[i:], m.DepositorAddress)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.DepositorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.InvestorId) > 0 {
		i -= len(m.InvestorId)
		copy(dAtA[i:], m.InvestorId)
		i = encodeVarintDeposit(dAtA, i, uint64(len(m.InvestorId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeposit(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeposit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InvestorId)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.DepositorAddress)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = len(m.PoolIndex)
	if l > 0 {
		n += 1 + l + sovDeposit(uint64(l))
	}
	l = m.LockedAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.WithdrawableAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.ClaimableInterestAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.IncentiveAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	if len(m.LinkedNFT) > 0 {
		for _, s := range m.LinkedNFT {
			l = len(s)
			n += 1 + l + sovDeposit(uint64(l))
		}
	}
	return n
}

func sovDeposit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeposit(x uint64) (n int) {
	return sovDeposit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeposit
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
			return fmt.Errorf("proto: DepositorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InvestorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InvestorId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositorAddress = append(m.DepositorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DepositorAddress == nil {
				m.DepositorAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LockedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WithdrawableAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimableInterestAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimableInterestAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncentiveAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IncentiveAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkedNFT", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
				return ErrInvalidLengthDeposit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeposit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkedNFT = append(m.LinkedNFT, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeposit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeposit
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
func skipDeposit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
					return 0, ErrIntOverflowDeposit
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
				return 0, ErrInvalidLengthDeposit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeposit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeposit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeposit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeposit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeposit = fmt.Errorf("proto: unexpected end of group")
)
