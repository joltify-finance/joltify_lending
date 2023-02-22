// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/spv/nft.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NftInfo struct {
	Issuer      string                                 `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Receiver    string                                 `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Ratio       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=ratio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ratio"`
	LastPayment time.Time                              `protobuf:"bytes,4,opt,name=last_payment,json=lastPayment,proto3,stdtime" json:"last_payment"`
}

func (m *NftInfo) Reset()         { *m = NftInfo{} }
func (m *NftInfo) String() string { return proto.CompactTextString(m) }
func (*NftInfo) ProtoMessage()    {}
func (*NftInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e522259423be67f7, []int{0}
}
func (m *NftInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NftInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NftInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NftInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NftInfo.Merge(m, src)
}
func (m *NftInfo) XXX_Size() int {
	return m.Size()
}
func (m *NftInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NftInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NftInfo proto.InternalMessageInfo

func (m *NftInfo) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *NftInfo) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *NftInfo) GetLastPayment() time.Time {
	if m != nil {
		return m.LastPayment
	}
	return time.Time{}
}

type PaymentItem struct {
	PaymentTime   time.Time   `protobuf:"bytes,1,opt,name=payment_time,json=paymentTime,proto3,stdtime" json:"payment_time"`
	PaymentAmount types1.Coin `protobuf:"bytes,2,opt,name=payment_amount,json=paymentAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"payment_amount"`
}

func (m *PaymentItem) Reset()         { *m = PaymentItem{} }
func (m *PaymentItem) String() string { return proto.CompactTextString(m) }
func (*PaymentItem) ProtoMessage()    {}
func (*PaymentItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e522259423be67f7, []int{1}
}
func (m *PaymentItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentItem.Merge(m, src)
}
func (m *PaymentItem) XXX_Size() int {
	return m.Size()
}
func (m *PaymentItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentItem.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentItem proto.InternalMessageInfo

func (m *PaymentItem) GetPaymentTime() time.Time {
	if m != nil {
		return m.PaymentTime
	}
	return time.Time{}
}

func (m *PaymentItem) GetPaymentAmount() types1.Coin {
	if m != nil {
		return m.PaymentAmount
	}
	return types1.Coin{}
}

type BorrowInterest struct {
	PoolIndex    string                                 `protobuf:"bytes,1,opt,name=pool_index,json=poolIndex,proto3" json:"pool_index,omitempty"`
	Apy          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=apy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"apy"`
	PayFreq      int32                                  `protobuf:"varint,3,opt,name=pay_freq,json=payFreq,proto3" json:"pay_freq,omitempty"`
	IssueTime    time.Time                              `protobuf:"bytes,4,opt,name=issue_time,json=issueTime,proto3,stdtime" json:"issue_time"`
	Borrowed     types1.Coin                            `protobuf:"bytes,5,opt,name=borrowed,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"borrowed"`
	BorrowedLast types1.Coin                            `protobuf:"bytes,6,opt,name=borrowed_last,json=borrowedLast,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"borrowed_last"`
	MonthlyRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=monthly_ratio,json=monthlyRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"monthly_ratio"`
	InterestSPY  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=interest_sPY,json=interestSPY,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_sPY"`
	Payments     []*PaymentItem                         `protobuf:"bytes,9,rep,name=payments,proto3" json:"payments,omitempty"`
}

func (m *BorrowInterest) Reset()         { *m = BorrowInterest{} }
func (m *BorrowInterest) String() string { return proto.CompactTextString(m) }
func (*BorrowInterest) ProtoMessage()    {}
func (*BorrowInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e522259423be67f7, []int{2}
}
func (m *BorrowInterest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BorrowInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BorrowInterest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BorrowInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BorrowInterest.Merge(m, src)
}
func (m *BorrowInterest) XXX_Size() int {
	return m.Size()
}
func (m *BorrowInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_BorrowInterest.DiscardUnknown(m)
}

var xxx_messageInfo_BorrowInterest proto.InternalMessageInfo

func (m *BorrowInterest) GetPoolIndex() string {
	if m != nil {
		return m.PoolIndex
	}
	return ""
}

func (m *BorrowInterest) GetPayFreq() int32 {
	if m != nil {
		return m.PayFreq
	}
	return 0
}

func (m *BorrowInterest) GetIssueTime() time.Time {
	if m != nil {
		return m.IssueTime
	}
	return time.Time{}
}

func (m *BorrowInterest) GetBorrowed() types1.Coin {
	if m != nil {
		return m.Borrowed
	}
	return types1.Coin{}
}

func (m *BorrowInterest) GetBorrowedLast() types1.Coin {
	if m != nil {
		return m.BorrowedLast
	}
	return types1.Coin{}
}

func (m *BorrowInterest) GetPayments() []*PaymentItem {
	if m != nil {
		return m.Payments
	}
	return nil
}

func init() {
	proto.RegisterType((*NftInfo)(nil), "joltify.spv.NftInfo")
	proto.RegisterType((*PaymentItem)(nil), "joltify.spv.PaymentItem")
	proto.RegisterType((*BorrowInterest)(nil), "joltify.spv.BorrowInterest")
}

func init() { proto.RegisterFile("joltify/spv/nft.proto", fileDescriptor_e522259423be67f7) }

var fileDescriptor_e522259423be67f7 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x18, 0xdd, 0x5a, 0xa7, 0xdb, 0xc1, 0x02, 0x94, 0x55, 0x22, 0x9d, 0x76, 0x40, 0xbd,
	0xcc, 0x61, 0x05, 0x71, 0xe2, 0x42, 0x37, 0x81, 0x2a, 0xa1, 0xa9, 0x0a, 0x5c, 0xc6, 0x25, 0x72,
	0x52, 0x27, 0x33, 0x24, 0x76, 0x16, 0xbb, 0x65, 0xf9, 0x17, 0xfb, 0x1d, 0x9c, 0xf7, 0x23, 0x26,
	0x4e, 0x13, 0x27, 0xb4, 0xc3, 0x86, 0xda, 0x3f, 0x82, 0xec, 0x38, 0x55, 0x8f, 0x43, 0xf4, 0x14,
	0xbf, 0xf7, 0xf2, 0xbe, 0xcf, 0xfe, 0xbe, 0x67, 0x83, 0xa7, 0x5f, 0x79, 0x2a, 0x69, 0x5c, 0x7a,
	0x22, 0x9f, 0x79, 0x2c, 0x96, 0x28, 0x2f, 0xb8, 0xe4, 0xd0, 0x36, 0x69, 0x24, 0xf2, 0x59, 0xb7,
	0x97, 0x70, 0x9e, 0xa4, 0xc4, 0xd3, 0xa5, 0x70, 0x1a, 0x7b, 0x92, 0x66, 0x44, 0x48, 0x9c, 0xe5,
	0xd5, 0xdf, 0xdd, 0x27, 0x09, 0x4f, 0xb8, 0x5e, 0x7a, 0x6a, 0x65, 0xb2, 0x6e, 0xc4, 0x45, 0xc6,
	0x85, 0x17, 0x62, 0x41, 0xbc, 0xd9, 0x61, 0x48, 0x24, 0x3e, 0xf4, 0x22, 0x4e, 0x99, 0xa9, 0xef,
	0x56, 0xf5, 0xa0, 0x6a, 0xac, 0x82, 0xaa, 0xb4, 0x7f, 0x6b, 0x81, 0xad, 0x93, 0x58, 0x8e, 0x58,
	0xcc, 0xe1, 0x33, 0xb0, 0x49, 0x85, 0x98, 0x92, 0xc2, 0xb1, 0xf6, 0xac, 0x7e, 0xdb, 0x37, 0x11,
	0xec, 0x82, 0x56, 0x41, 0x22, 0x42, 0x67, 0xa4, 0x70, 0x1e, 0xe9, 0xca, 0x32, 0x86, 0x3e, 0x68,
	0x16, 0x58, 0x52, 0xee, 0x6c, 0xa8, 0xc2, 0xf0, 0xed, 0xf5, 0x5d, 0xaf, 0x71, 0x7b, 0xd7, 0x7b,
	0x91, 0x50, 0x79, 0x36, 0x0d, 0x51, 0xc4, 0x33, 0xc3, 0x67, 0x3e, 0x07, 0x62, 0xf2, 0xcd, 0x93,
	0x65, 0x4e, 0x04, 0x3a, 0x26, 0xd1, 0xaf, 0xab, 0x03, 0x60, 0xb6, 0x73, 0x4c, 0x22, 0xbf, 0x82,
	0x82, 0x1f, 0x40, 0x27, 0xc5, 0x42, 0x06, 0x39, 0x2e, 0x33, 0xc2, 0xa4, 0xf3, 0x78, 0xcf, 0xea,
	0xdb, 0x83, 0x2e, 0xaa, 0xc4, 0x41, 0xb5, 0x38, 0xe8, 0x73, 0x2d, 0xce, 0xb0, 0xa5, 0x68, 0x2f,
	0xef, 0x7b, 0x96, 0x6f, 0xab, 0xce, 0x71, 0xd5, 0xb8, 0xff, 0xd3, 0x02, 0xb6, 0x59, 0x8f, 0x24,
	0xc9, 0x14, 0xb0, 0xc1, 0x0c, 0x94, 0xb0, 0xfa, 0x98, 0x0f, 0x06, 0x36, 0x9d, 0xaa, 0x06, 0x0b,
	0xb0, 0x53, 0x03, 0xe1, 0x8c, 0x4f, 0x99, 0xd4, 0xba, 0xd8, 0x83, 0x5d, 0x64, 0x4e, 0xa3, 0x9c,
	0x40, 0xc6, 0x09, 0x74, 0xc4, 0x29, 0x1b, 0xbe, 0x54, 0x48, 0x3f, 0xee, 0x7b, 0xfd, 0x07, 0x28,
	0xa3, 0x1a, 0x84, 0xbf, 0x6d, 0x28, 0xde, 0x69, 0x86, 0xfd, 0xab, 0x26, 0xd8, 0x19, 0xf2, 0xa2,
	0xe0, 0xdf, 0x47, 0x4c, 0x92, 0x82, 0x08, 0x09, 0x9f, 0x03, 0x90, 0x73, 0x9e, 0x06, 0x94, 0x4d,
	0xc8, 0x85, 0x31, 0xad, 0xad, 0x32, 0x23, 0x95, 0x80, 0x27, 0x60, 0x03, 0xe7, 0x65, 0x65, 0xd9,
	0x7f, 0x3a, 0xa3, 0x80, 0xe0, 0x2e, 0x68, 0xe5, 0xb8, 0x0c, 0xe2, 0x82, 0x9c, 0x6b, 0xbb, 0x9b,
	0xfe, 0x56, 0x8e, 0xcb, 0xf7, 0x05, 0x39, 0x87, 0x47, 0x00, 0xe8, 0x61, 0xa9, 0x74, 0xfd, 0x17,
	0xc3, 0xda, 0xba, 0x4f, 0xab, 0x9a, 0x80, 0x56, 0xa8, 0x0f, 0x48, 0x26, 0x4e, 0x73, 0xfd, 0x7a,
	0x2e, 0xc1, 0x61, 0x0e, 0xb6, 0xeb, 0x75, 0xa0, 0xe6, 0xc5, 0xd9, 0x5c, 0x3f, 0x5b, 0xa7, 0x66,
	0xf8, 0x88, 0x85, 0x84, 0x18, 0x6c, 0x67, 0x9c, 0xc9, 0xb3, 0xb4, 0x0c, 0xaa, 0xeb, 0xb2, 0xb5,
	0x06, 0x53, 0x3a, 0x06, 0xd2, 0xd7, 0xb7, 0x26, 0x00, 0x1d, 0x6a, 0x06, 0x23, 0x10, 0xe3, 0x53,
	0xa7, 0xb5, 0x06, 0x06, 0xbb, 0x46, 0xfc, 0x34, 0x3e, 0x85, 0xaf, 0xb5, 0xfd, 0x6a, 0x22, 0x85,
	0xd3, 0xde, 0xdb, 0xe8, 0xdb, 0x03, 0x07, 0xad, 0x3c, 0x5e, 0x68, 0xe5, 0xa6, 0xf9, 0xcb, 0x3f,
	0x87, 0xe3, 0xeb, 0xb9, 0x6b, 0xdd, 0xcc, 0x5d, 0xeb, 0xcf, 0xdc, 0xb5, 0x2e, 0x17, 0x6e, 0xe3,
	0x66, 0xe1, 0x36, 0x7e, 0x2f, 0xdc, 0xc6, 0x97, 0x37, 0x2b, 0x5b, 0x32, 0x38, 0x07, 0x31, 0x65,
	0x98, 0x45, 0xa4, 0x8e, 0x83, 0x94, 0xb0, 0x09, 0x65, 0x89, 0x77, 0xa1, 0x5f, 0x4d, 0xbd, 0xcd,
	0x70, 0x53, 0xcf, 0xd3, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x7a, 0xb1, 0x7d, 0x51,
	0x05, 0x00, 0x00,
}

func (m *NftInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NftInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NftInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastPayment, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastPayment):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintNft(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Ratio.Size()
		i -= size
		if _, err := m.Ratio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PaymentItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PaymentAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PaymentTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PaymentTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintNft(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BorrowInterest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BorrowInterest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BorrowInterest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payments) > 0 {
		for iNdEx := len(m.Payments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size := m.InterestSPY.Size()
		i -= size
		if _, err := m.InterestSPY.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MonthlyRatio.Size()
		i -= size
		if _, err := m.MonthlyRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.BorrowedLast.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Borrowed.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n6, err6 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.IssueTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.IssueTime):])
	if err6 != nil {
		return 0, err6
	}
	i -= n6
	i = encodeVarintNft(dAtA, i, uint64(n6))
	i--
	dAtA[i] = 0x22
	if m.PayFreq != 0 {
		i = encodeVarintNft(dAtA, i, uint64(m.PayFreq))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Apy.Size()
		i -= size
		if _, err := m.Apy.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.PoolIndex) > 0 {
		i -= len(m.PoolIndex)
		copy(dAtA[i:], m.PoolIndex)
		i = encodeVarintNft(dAtA, i, uint64(len(m.PoolIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NftInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = m.Ratio.Size()
	n += 1 + l + sovNft(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastPayment)
	n += 1 + l + sovNft(uint64(l))
	return n
}

func (m *PaymentItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PaymentTime)
	n += 1 + l + sovNft(uint64(l))
	l = m.PaymentAmount.Size()
	n += 1 + l + sovNft(uint64(l))
	return n
}

func (m *BorrowInterest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolIndex)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = m.Apy.Size()
	n += 1 + l + sovNft(uint64(l))
	if m.PayFreq != 0 {
		n += 1 + sovNft(uint64(m.PayFreq))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.IssueTime)
	n += 1 + l + sovNft(uint64(l))
	l = m.Borrowed.Size()
	n += 1 + l + sovNft(uint64(l))
	l = m.BorrowedLast.Size()
	n += 1 + l + sovNft(uint64(l))
	l = m.MonthlyRatio.Size()
	n += 1 + l + sovNft(uint64(l))
	l = m.InterestSPY.Size()
	n += 1 + l + sovNft(uint64(l))
	if len(m.Payments) > 0 {
		for _, e := range m.Payments {
			l = e.Size()
			n += 1 + l + sovNft(uint64(l))
		}
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NftInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: NftInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NftInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ratio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastPayment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastPayment, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func (m *PaymentItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: PaymentItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PaymentTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PaymentAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func (m *BorrowInterest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: BorrowInterest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BorrowInterest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Apy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Apy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayFreq", wireType)
			}
			m.PayFreq = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PayFreq |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.IssueTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Borrowed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Borrowed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BorrowedLast", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BorrowedLast.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonthlyRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MonthlyRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestSPY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestSPY.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payments = append(m.Payments, &PaymentItem{})
			if err := m.Payments[len(m.Payments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)
