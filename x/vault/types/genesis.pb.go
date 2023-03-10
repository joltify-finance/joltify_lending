// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the vault module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OutboundTxList []OutboundTx `protobuf:"bytes,5,rep,name=outboundTxList,proto3" json:"outboundTxList"`
	// this line is used by starport scaffolding # genesis/proto/state
	IssueTokenList []*IssueToken `protobuf:"bytes,2,rep,name=issueTokenList,proto3" json:"issueTokenList,omitempty"`
	CreatePoolList []*CreatePool `protobuf:"bytes,3,rep,name=createPoolList,proto3" json:"createPoolList,omitempty"`
	// this line is used by starport scaffolding # ibc/genesis/proto
	ValidatorinfoList []*Validators                            `protobuf:"bytes,6,rep,name=validatorinfoList,proto3" json:"validatorinfoList,omitempty"`
	LatestTwoPool     []*CreatePool                            `protobuf:"bytes,10,rep,name=latestTwoPool,proto3" json:"latestTwoPool,omitempty"`
	StandbypowerList  []*StandbyPower                          `protobuf:"bytes,7,rep,name=standbypowerList,proto3" json:"standbypowerList,omitempty"`
	FeeCollectedList  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=feeCollectedList,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"feeCollectedList"`
	CoinsQuota        CoinsQuota                               `protobuf:"bytes,9,opt,name=coinsQuota,proto3" json:"coinsQuota"`
	Exported          bool                                     `protobuf:"varint,4,opt,name=exported,proto3" json:"exported,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3de5765bc18ded2, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOutboundTxList() []OutboundTx {
	if m != nil {
		return m.OutboundTxList
	}
	return nil
}

func (m *GenesisState) GetIssueTokenList() []*IssueToken {
	if m != nil {
		return m.IssueTokenList
	}
	return nil
}

func (m *GenesisState) GetCreatePoolList() []*CreatePool {
	if m != nil {
		return m.CreatePoolList
	}
	return nil
}

func (m *GenesisState) GetValidatorinfoList() []*Validators {
	if m != nil {
		return m.ValidatorinfoList
	}
	return nil
}

func (m *GenesisState) GetLatestTwoPool() []*CreatePool {
	if m != nil {
		return m.LatestTwoPool
	}
	return nil
}

func (m *GenesisState) GetStandbypowerList() []*StandbyPower {
	if m != nil {
		return m.StandbypowerList
	}
	return nil
}

func (m *GenesisState) GetFeeCollectedList() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeCollectedList
	}
	return nil
}

func (m *GenesisState) GetCoinsQuota() CoinsQuota {
	if m != nil {
		return m.CoinsQuota
	}
	return CoinsQuota{}
}

func (m *GenesisState) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "joltify.vault.GenesisState")
}

func init() { proto.RegisterFile("joltify/vault/genesis.proto", fileDescriptor_e3de5765bc18ded2) }

var fileDescriptor_e3de5765bc18ded2 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0x63, 0xda, 0x86, 0x30, 0xa5, 0x55, 0x19, 0x81, 0x94, 0xb6, 0x92, 0x13, 0xb1, 0xca,
	0xa6, 0x36, 0x6d, 0x37, 0xec, 0x2a, 0x92, 0x45, 0x84, 0x84, 0x44, 0x70, 0x23, 0x16, 0x6c, 0xa2,
	0xb1, 0x3d, 0x31, 0x43, 0x9c, 0x79, 0xc6, 0xf3, 0x9c, 0x3f, 0xb7, 0xe0, 0x1c, 0x5c, 0x81, 0x0b,
	0x74, 0xd9, 0x25, 0x2b, 0x40, 0xc9, 0x45, 0x90, 0xc7, 0x93, 0xd0, 0x38, 0x91, 0x58, 0x65, 0x26,
	0xef, 0xf7, 0x7d, 0xf3, 0xe6, 0xf9, 0x1b, 0x72, 0xfe, 0x05, 0x62, 0x14, 0xc3, 0xb9, 0x3b, 0x61,
	0x59, 0x8c, 0x6e, 0xc4, 0x25, 0x57, 0x42, 0x39, 0x49, 0x0a, 0x08, 0xf4, 0xc8, 0x14, 0x1d, 0x5d,
	0x3c, 0x6b, 0x6c, 0xb2, 0x90, 0xa1, 0x0f, 0x99, 0x0c, 0x07, 0x38, 0x2b, 0xf8, 0x32, 0x20, 0x94,
	0xca, 0xf8, 0x00, 0x61, 0xc4, 0xe5, 0x6e, 0x20, 0x48, 0x39, 0x43, 0x3e, 0x48, 0x00, 0x62, 0x03,
	0x9c, 0x6e, 0x02, 0x5f, 0x33, 0x40, 0x66, 0x4a, 0xa5, 0x4e, 0x15, 0xb2, 0x91, 0x90, 0x91, 0x29,
	0x3e, 0x8f, 0x20, 0x02, 0xbd, 0x74, 0xf3, 0x95, 0xf9, 0xd7, 0x0e, 0x40, 0x8d, 0x41, 0xb9, 0x3e,
	0x53, 0xdc, 0x9d, 0x5c, 0xfa, 0x1c, 0xd9, 0xa5, 0x1b, 0x80, 0x30, 0xed, 0xbc, 0xfc, 0x71, 0x40,
	0x9e, 0x76, 0x8b, 0x1b, 0xdf, 0x22, 0x43, 0x4e, 0xaf, 0x49, 0x35, 0x61, 0x29, 0x1b, 0xab, 0xba,
	0xd5, 0xb4, 0x5a, 0x87, 0x57, 0x2f, 0x9c, 0x8d, 0x09, 0x38, 0x3d, 0x5d, 0x6c, 0xef, 0xdf, 0xfd,
	0x6a, 0x54, 0x3c, 0x83, 0xd2, 0x2e, 0x39, 0x5e, 0x8d, 0xa2, 0x3f, 0x7b, 0x27, 0x14, 0xd6, 0x0f,
	0x9a, 0x7b, 0xad, 0xc3, 0xab, 0xd3, 0x92, 0xf8, 0xfd, 0x1a, 0x32, 0x06, 0x25, 0x19, 0x7d, 0x43,
	0x8e, 0xf5, 0xc8, 0xfa, 0xf9, 0xc4, 0xb4, 0xd1, 0xa3, 0x9d, 0x46, 0x6f, 0xd7, 0x90, 0x57, 0x12,
	0xe4, 0x16, 0xc5, 0x50, 0x7b, 0x00, 0xb1, 0xb6, 0xd8, 0xdb, 0x69, 0xd1, 0x59, 0x43, 0x5e, 0x49,
	0x40, 0xbb, 0xe4, 0xd9, 0x84, 0xc5, 0x22, 0x64, 0x08, 0xa9, 0x90, 0x43, 0xd0, 0x2e, 0xd5, 0x9d,
	0x2e, 0x1f, 0x57, 0x9c, 0xf2, 0xb6, 0x35, 0xf4, 0x86, 0x1c, 0xc5, 0x0c, 0xb9, 0xc2, 0xfe, 0x14,
	0x72, 0xf7, 0x3a, 0xf9, 0x5f, 0x2b, 0x9b, 0x3c, 0xed, 0x92, 0x13, 0x85, 0x4c, 0x86, 0xfe, 0x3c,
	0x81, 0x29, 0x4f, 0x75, 0x23, 0x8f, 0xb5, 0xc7, 0x79, 0xc9, 0xe3, 0xb6, 0xc0, 0x7a, 0x39, 0xe6,
	0x6d, 0x89, 0xe8, 0x94, 0x9c, 0x0c, 0x39, 0xef, 0x40, 0x1c, 0xf3, 0x00, 0x79, 0xa8, 0x8d, 0x6a,
	0xa6, 0x99, 0x22, 0x22, 0x4e, 0x1e, 0x11, 0xc7, 0x44, 0xc4, 0xe9, 0x80, 0x90, 0xed, 0x57, 0xf9,
	0x37, 0xfa, 0xfe, 0xbb, 0xd1, 0x8a, 0x04, 0x7e, 0xce, 0x7c, 0x27, 0x80, 0xb1, 0x6b, 0xf2, 0x54,
	0xfc, 0x5c, 0xa8, 0x70, 0xe4, 0xe2, 0x3c, 0xe1, 0x4a, 0x0b, 0x94, 0xb7, 0x75, 0x08, 0xbd, 0x21,
	0x24, 0x8f, 0x9b, 0xfa, 0x90, 0xe7, 0xb8, 0xfe, 0x44, 0x67, 0xaa, 0x7c, 0xff, 0x7f, 0x80, 0x89,
	0xc5, 0x03, 0x09, 0x3d, 0x23, 0x35, 0x3e, 0x4b, 0x20, 0x45, 0x1e, 0xd6, 0xf7, 0x9b, 0x56, 0xab,
	0xe6, 0xad, 0xf7, 0x6d, 0xef, 0x6e, 0x61, 0x5b, 0xf7, 0x0b, 0xdb, 0xfa, 0xb3, 0xb0, 0xad, 0x6f,
	0x4b, 0xbb, 0x72, 0xbf, 0xb4, 0x2b, 0x3f, 0x97, 0x76, 0xe5, 0xd3, 0xeb, 0x07, 0x2d, 0x9b, 0xc3,
	0x2e, 0x86, 0x42, 0x32, 0x19, 0xf0, 0xd5, 0x7e, 0x10, 0x73, 0x19, 0x0a, 0x19, 0xb9, 0x33, 0xf3,
	0x9e, 0xf4, 0x45, 0xfc, 0xaa, 0x7e, 0x18, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x11,
	0xd8, 0x26, 0x17, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LatestTwoPool) > 0 {
		for iNdEx := len(m.LatestTwoPool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LatestTwoPool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	{
		size, err := m.CoinsQuota.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.FeeCollectedList) > 0 {
		for iNdEx := len(m.FeeCollectedList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeCollectedList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.StandbypowerList) > 0 {
		for iNdEx := len(m.StandbypowerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StandbypowerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ValidatorinfoList) > 0 {
		for iNdEx := len(m.ValidatorinfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorinfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.OutboundTxList) > 0 {
		for iNdEx := len(m.OutboundTxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutboundTxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Exported {
		i--
		if m.Exported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.CreatePoolList) > 0 {
		for iNdEx := len(m.CreatePoolList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreatePoolList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.IssueTokenList) > 0 {
		for iNdEx := len(m.IssueTokenList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IssueTokenList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.IssueTokenList) > 0 {
		for _, e := range m.IssueTokenList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CreatePoolList) > 0 {
		for _, e := range m.CreatePoolList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Exported {
		n += 2
	}
	if len(m.OutboundTxList) > 0 {
		for _, e := range m.OutboundTxList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorinfoList) > 0 {
		for _, e := range m.ValidatorinfoList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.StandbypowerList) > 0 {
		for _, e := range m.StandbypowerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.FeeCollectedList) > 0 {
		for _, e := range m.FeeCollectedList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CoinsQuota.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LatestTwoPool) > 0 {
		for _, e := range m.LatestTwoPool {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueTokenList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssueTokenList = append(m.IssueTokenList, &IssueToken{})
			if err := m.IssueTokenList[len(m.IssueTokenList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePoolList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatePoolList = append(m.CreatePoolList, &CreatePool{})
			if err := m.CreatePoolList[len(m.CreatePoolList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Exported = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutboundTxList = append(m.OutboundTxList, OutboundTx{})
			if err := m.OutboundTxList[len(m.OutboundTxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorinfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorinfoList = append(m.ValidatorinfoList, &Validators{})
			if err := m.ValidatorinfoList[len(m.ValidatorinfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StandbypowerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StandbypowerList = append(m.StandbypowerList, &StandbyPower{})
			if err := m.StandbypowerList[len(m.StandbypowerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollectedList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollectedList = append(m.FeeCollectedList, types.Coin{})
			if err := m.FeeCollectedList[len(m.FeeCollectedList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinsQuota", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CoinsQuota.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestTwoPool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LatestTwoPool = append(m.LatestTwoPool, &CreatePool{})
			if err := m.LatestTwoPool[len(m.LatestTwoPool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
