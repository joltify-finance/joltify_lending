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
	InvestorId       string                                        `protobuf:"bytes,1,opt,name=investor_id,json=investorId,proto3" json:"investor_id,omitempty"`
	DepositorAddress github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=depositor_address,json=depositorAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"depositor_address,omitempty"`
	PoolIndex        string                                        `protobuf:"bytes,3,opt,name=pool_index,json=poolIndex,proto3" json:"pool_index,omitempty"`
	LockedAmount     types.Coin                                    `protobuf:"bytes,4,opt,name=locked_amount,json=lockedAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"locked_amount"`
	WithdrawalAmount types.Coin                                    `protobuf:"bytes,5,opt,name=withdrawal_amount,json=withdrawalAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawal_amount"`
	IncentiveAmount  types.Coin                                    `protobuf:"bytes,6,opt,name=incentive_amount,json=incentiveAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"incentive_amount"`
	LinkedNFT        []string                                      `protobuf:"bytes,7,rep,name=linkedNFT,proto3" json:"linkedNFT,omitempty"`
	WithdrawProposal bool                                          `protobuf:"varint,8,opt,name=withdraw_proposal,json=withdrawProposal,proto3" json:"withdraw_proposal,omitempty"`
	TransferRequest  bool                                          `protobuf:"varint,9,opt,name=transfer_request,json=transferRequest,proto3" json:"transfer_request,omitempty"`
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

func (m *DepositorInfo) GetWithdrawalAmount() types.Coin {
	if m != nil {
		return m.WithdrawalAmount
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

func (m *DepositorInfo) GetWithdrawProposal() bool {
	if m != nil {
		return m.WithdrawProposal
	}
	return false
}

func (m *DepositorInfo) GetTransferRequest() bool {
	if m != nil {
		return m.TransferRequest
	}
	return false
}

func init() {
	proto.RegisterType((*DepositorInfo)(nil), "joltify.spv.DepositorInfo")
}

func init() { proto.RegisterFile("joltify/spv/deposit.proto", fileDescriptor_5bb054114dab9679) }

var fileDescriptor_5bb054114dab9679 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x6f, 0xd3, 0x4e,
	0x14, 0xc0, 0xe3, 0x7f, 0xff, 0x2d, 0xcd, 0xa5, 0x55, 0x53, 0xc3, 0xe0, 0x54, 0x60, 0x5b, 0x4c,
	0x46, 0x28, 0x3e, 0x0a, 0x12, 0x7b, 0x02, 0x42, 0x64, 0x41, 0x55, 0xc4, 0xc4, 0x62, 0x9d, 0x7d,
	0x67, 0xf7, 0x88, 0x7d, 0xcf, 0xf8, 0xce, 0x6e, 0xf2, 0x2d, 0xf8, 0x08, 0xcc, 0xcc, 0x7c, 0x88,
	0x8e, 0x15, 0x13, 0x53, 0x41, 0xc9, 0xb7, 0x60, 0x42, 0xb6, 0xcf, 0x09, 0x23, 0x43, 0x27, 0xfb,
	0x7e, 0xf7, 0x9e, 0x7f, 0xef, 0x3d, 0xdf, 0xa1, 0xd1, 0x47, 0x48, 0x15, 0x8f, 0x57, 0x58, 0xe6,
	0x15, 0xa6, 0x2c, 0x07, 0xc9, 0x95, 0x9f, 0x17, 0xa0, 0xc0, 0x1c, 0xe8, 0x2d, 0x5f, 0xe6, 0xd5,
	0xd9, 0x28, 0x02, 0x99, 0x81, 0x0c, 0x9a, 0x2d, 0xdc, 0x2e, 0xda, 0xb8, 0xb3, 0x07, 0x09, 0x24,
	0xd0, 0xf2, 0xfa, 0x4d, 0x53, 0xbb, 0x8d, 0xc1, 0x21, 0x91, 0x0c, 0x57, 0xe7, 0x21, 0x53, 0xe4,
	0x1c, 0x47, 0xc0, 0x85, 0xde, 0x77, 0x12, 0x80, 0x24, 0x65, 0xb8, 0x59, 0x85, 0x65, 0x8c, 0x15,
	0xcf, 0x98, 0x54, 0x24, 0xcb, 0xdb, 0x80, 0xc7, 0x5f, 0xf6, 0xd1, 0xf1, 0xeb, 0xb6, 0x20, 0x28,
	0x66, 0x22, 0x06, 0xd3, 0x41, 0x03, 0x2e, 0x2a, 0x26, 0x15, 0x14, 0x01, 0xa7, 0x96, 0xe1, 0x1a,
	0x5e, 0x7f, 0x8e, 0x3a, 0x34, 0xa3, 0x66, 0x89, 0x4e, 0x69, 0x97, 0x11, 0x10, 0x4a, 0x0b, 0x26,
	0xa5, 0xf5, 0x9f, 0x6b, 0x78, 0x47, 0xd3, 0xb7, 0xbf, 0x6f, 0x9d, 0x71, 0xc2, 0xd5, 0x65, 0x19,
	0xfa, 0x11, 0x64, 0xba, 0x03, 0xfd, 0x18, 0x4b, 0xba, 0xc0, 0x6a, 0x95, 0x33, 0xe9, 0x4f, 0xa2,
	0x68, 0xd2, 0x26, 0x7e, 0xff, 0x36, 0xbe, 0xaf, 0xfb, 0xd4, 0x64, 0xba, 0x52, 0x4c, 0xce, 0x87,
	0x5b, 0x85, 0xc6, 0xe6, 0x23, 0x84, 0x72, 0x80, 0x34, 0xe0, 0x82, 0xb2, 0xa5, 0xb5, 0xd7, 0x94,
	0xd5, 0xaf, 0xc9, 0xac, 0x06, 0x66, 0x8e, 0x8e, 0x53, 0x88, 0x16, 0x8c, 0x06, 0x24, 0x83, 0x52,
	0x28, 0xeb, 0x7f, 0xd7, 0xf0, 0x06, 0xcf, 0x47, 0xbe, 0xfe, 0x7a, 0x3d, 0x21, 0x5f, 0x4f, 0xc8,
	0x7f, 0x05, 0x5c, 0x4c, 0x9f, 0x5d, 0xdf, 0x3a, 0xbd, 0xaf, 0x3f, 0x1d, 0xef, 0x1f, 0x0a, 0xae,
	0x13, 0xe4, 0xfc, 0xa8, 0x35, 0x4c, 0x1a, 0x81, 0xb9, 0x44, 0xa7, 0x57, 0x5c, 0x5d, 0xd2, 0x82,
	0x5c, 0x91, 0xb4, 0xb3, 0xee, 0xdf, 0xbd, 0x75, 0xb8, 0xb3, 0x68, 0x73, 0x85, 0x86, 0x5c, 0x44,
	0x4c, 0x28, 0x5e, 0xb1, 0x4e, 0x7c, 0x70, 0xf7, 0xe2, 0x93, 0xad, 0x44, 0x7b, 0x1f, 0xa2, 0x7e,
	0xca, 0xc5, 0x82, 0xd1, 0x77, 0x6f, 0xde, 0x5b, 0xf7, 0xdc, 0xbd, 0xfa, 0x0f, 0x6c, 0x81, 0xf9,
	0x74, 0x37, 0x8f, 0xfa, 0x00, 0xe7, 0x20, 0x49, 0x6a, 0x1d, 0xba, 0x86, 0x77, 0xb8, 0x6b, 0xe1,
	0x42, 0x73, 0xf3, 0x09, 0x1a, 0xaa, 0x82, 0x08, 0x19, 0xb3, 0x22, 0x28, 0xd8, 0xa7, 0x92, 0x49,
	0x65, 0xf5, 0x9b, 0xd8, 0x93, 0x8e, 0xcf, 0x5b, 0x3c, 0xbd, 0xb8, 0x5e, 0xdb, 0xc6, 0xcd, 0xda,
	0x36, 0x7e, 0xad, 0x6d, 0xe3, 0xf3, 0xc6, 0xee, 0xdd, 0x6c, 0xec, 0xde, 0x8f, 0x8d, 0xdd, 0xfb,
	0xf0, 0xf2, 0xaf, 0x56, 0xf4, 0x35, 0x1a, 0xc7, 0x5c, 0x10, 0x11, 0xb1, 0x6e, 0x1d, 0xa4, 0x4c,
	0x50, 0x2e, 0x12, 0xbc, 0x6c, 0xee, 0x5e, 0xd3, 0x5e, 0x78, 0xd0, 0x9c, 0xfd, 0x17, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x47, 0xac, 0xda, 0xdf, 0x97, 0x03, 0x00, 0x00,
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
	if m.TransferRequest {
		i--
		if m.TransferRequest {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.WithdrawProposal {
		i--
		if m.WithdrawProposal {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.LinkedNFT) > 0 {
		for iNdEx := len(m.LinkedNFT) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.LinkedNFT[iNdEx])
			copy(dAtA[i:], m.LinkedNFT[iNdEx])
			i = encodeVarintDeposit(dAtA, i, uint64(len(m.LinkedNFT[iNdEx])))
			i--
			dAtA[i] = 0x3a
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
	dAtA[i] = 0x32
	{
		size, err := m.WithdrawalAmount.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.WithdrawalAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	l = m.IncentiveAmount.Size()
	n += 1 + l + sovDeposit(uint64(l))
	if len(m.LinkedNFT) > 0 {
		for _, s := range m.LinkedNFT {
			l = len(s)
			n += 1 + l + sovDeposit(uint64(l))
		}
	}
	if m.WithdrawProposal {
		n += 2
	}
	if m.TransferRequest {
		n += 2
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
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawalAmount", wireType)
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
			if err := m.WithdrawalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
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
		case 7:
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
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawProposal", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
			m.WithdrawProposal = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferRequest", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeposit
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
			m.TransferRequest = bool(v != 0)
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
