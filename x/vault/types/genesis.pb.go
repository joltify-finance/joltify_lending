// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the vault module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OutboundTxList []OutboundTx `protobuf:"bytes,5,rep,name=outbound_tx_list,json=outboundTxList,proto3" json:"outbound_tx_list"`
	// this line is used by starport scaffolding # genesis/proto/state
	IssueTokenList []*IssueToken `protobuf:"bytes,2,rep,name=issue_token_list,json=issueTokenList,proto3" json:"issue_token_list,omitempty"`
	CreatePoolList []*CreatePool `protobuf:"bytes,3,rep,name=create_pool_list,json=createPoolList,proto3" json:"create_pool_list,omitempty"`
	// this line is used by starport scaffolding # ibc/genesis/proto
	ValidatorinfoList []*Validators                            `protobuf:"bytes,6,rep,name=validatorinfo_list,json=validatorinfoList,proto3" json:"validatorinfo_list,omitempty"`
	LatestTwoPool     []*CreatePool                            `protobuf:"bytes,10,rep,name=latest_twoPool,json=latestTwoPool,proto3" json:"latest_twoPool,omitempty"`
	StandbypowerList  []*StandbyPower                          `protobuf:"bytes,7,rep,name=standbypower_list,json=standbypowerList,proto3" json:"standbypower_list,omitempty"`
	FeeCollectedList  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=feeCollected_list,json=feeCollectedList,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"feeCollected_list"`
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
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0xfa, 0x87, 0xb2, 0xa5, 0x51, 0x6a, 0x81, 0x94, 0xb6, 0x92, 0x13, 0x71, 0xca,
	0xa5, 0x36, 0x6d, 0x2f, 0xdc, 0x40, 0xc9, 0x81, 0x56, 0x42, 0x22, 0xb8, 0x11, 0x07, 0x2e, 0xd6,
	0xda, 0xde, 0x98, 0x25, 0xce, 0x8e, 0xf1, 0x8e, 0xd3, 0xe4, 0x2d, 0x78, 0x0e, 0xde, 0x03, 0xa9,
	0xc7, 0x1e, 0x39, 0x01, 0x4a, 0x5e, 0x04, 0x79, 0x77, 0x13, 0x39, 0x21, 0x12, 0x27, 0xcf, 0xee,
	0xfc, 0xe6, 0xdb, 0xd9, 0xf1, 0xb7, 0xe4, 0xec, 0x0b, 0xa4, 0xc8, 0x87, 0x33, 0x6f, 0x42, 0x8b,
	0x14, 0xbd, 0x84, 0x09, 0x26, 0xb9, 0x74, 0xb3, 0x1c, 0x10, 0xec, 0x23, 0x93, 0x74, 0x55, 0xf2,
	0xb4, 0xb5, 0xce, 0x42, 0x81, 0x21, 0x14, 0x22, 0x0e, 0x70, 0xaa, 0xf9, 0x4d, 0x80, 0x4b, 0x59,
	0xb0, 0x00, 0x61, 0xc4, 0xc4, 0x76, 0x20, 0xca, 0x19, 0x45, 0x16, 0x64, 0x00, 0xa9, 0x01, 0x4e,
	0xd6, 0x81, 0xaf, 0x05, 0x20, 0x35, 0xa9, 0x8d, 0x4e, 0x25, 0xd2, 0x11, 0x17, 0x89, 0x49, 0x3e,
	0x4b, 0x20, 0x01, 0x15, 0x7a, 0x65, 0x64, 0x76, 0x9d, 0x08, 0xe4, 0x18, 0xa4, 0x17, 0x52, 0xc9,
	0xbc, 0xc9, 0x45, 0xc8, 0x90, 0x5e, 0x78, 0x11, 0x70, 0xd3, 0xce, 0x8b, 0x1f, 0x7b, 0xe4, 0xe9,
	0x5b, 0x7d, 0xe3, 0x5b, 0xa4, 0xc8, 0xec, 0x2b, 0xb2, 0x9f, 0xd1, 0x9c, 0x8e, 0x65, 0xd3, 0x6a,
	0x5b, 0x9d, 0xc3, 0xcb, 0xe7, 0xee, 0xda, 0x04, 0xdc, 0xbe, 0x4a, 0x76, 0x77, 0xef, 0x7f, 0xb5,
	0x6a, 0xbe, 0x41, 0xed, 0x1b, 0xd2, 0xa8, 0x8c, 0x22, 0x48, 0xb9, 0xc4, 0xe6, 0x5e, 0x7b, 0xa7,
	0x73, 0x78, 0x79, 0xb2, 0x51, 0xfe, 0xde, 0x60, 0x83, 0xa9, 0x91, 0xa8, 0xc3, 0x6a, 0xe7, 0x1d,
	0x97, 0x68, 0xf7, 0x48, 0xa3, 0x32, 0x34, 0x2d, 0xf5, 0x68, 0xab, 0xd4, 0x4d, 0x89, 0x0d, 0x4a,
	0xca, 0xaf, 0xf3, 0x55, 0xbc, 0x14, 0xa9, 0x0c, 0x56, 0x8b, 0xec, 0x6c, 0x15, 0xe9, 0x29, 0xac,
	0x0f, 0x90, 0xfa, 0xf5, 0x68, 0x15, 0x2b, 0x91, 0x6b, 0x62, 0x4f, 0x68, 0xca, 0x63, 0x8a, 0x90,
	0x73, 0x31, 0x04, 0x2d, 0xb3, 0xbf, 0x55, 0xe6, 0xe3, 0x12, 0x94, 0xfe, 0xf1, 0x5a, 0x91, 0x52,
	0x7a, 0x43, 0xea, 0x29, 0x45, 0x26, 0x31, 0xc0, 0x3b, 0x28, 0xf5, 0x9b, 0xe4, 0x7f, 0xcd, 0x1c,
	0xe9, 0x82, 0x81, 0xe6, 0xed, 0x6b, 0x72, 0x2c, 0x91, 0x8a, 0x38, 0x9c, 0x65, 0x70, 0xc7, 0x72,
	0xdd, 0xca, 0x63, 0x25, 0x72, 0xb6, 0x21, 0x72, 0xab, 0xb9, 0x7e, 0xc9, 0xf9, 0x8d, 0x6a, 0x95,
	0xea, 0x65, 0x4a, 0x8e, 0x87, 0x8c, 0xf5, 0x20, 0x4d, 0x59, 0x84, 0x2c, 0xd6, 0x4a, 0x07, 0xa6,
	0x1d, 0x6d, 0x16, 0xb7, 0x34, 0x8b, 0x6b, 0xcc, 0xe2, 0xf6, 0x80, 0x8b, 0xee, 0xcb, 0xf2, 0x5f,
	0x7d, 0xff, 0xdd, 0xea, 0x24, 0x1c, 0x3f, 0x17, 0xa1, 0x1b, 0xc1, 0xd8, 0x33, 0xce, 0xd2, 0x9f,
	0x73, 0x19, 0x8f, 0x3c, 0x9c, 0x65, 0x4c, 0xaa, 0x02, 0xe9, 0x37, 0xaa, 0xa7, 0xa8, 0x93, 0x5f,
	0x13, 0x52, 0x1a, 0x4f, 0x7e, 0x28, 0x1d, 0xdd, 0x7c, 0xa2, 0xdc, 0xf5, 0xcf, 0x04, 0x56, 0x80,
	0xb1, 0x47, 0xa5, 0xc4, 0x3e, 0x25, 0x07, 0x6c, 0x9a, 0x41, 0x8e, 0x2c, 0x6e, 0xee, 0xb6, 0xad,
	0xce, 0x81, 0xbf, 0x5a, 0x77, 0xfd, 0xfb, 0xb9, 0x63, 0x3d, 0xcc, 0x1d, 0xeb, 0xcf, 0xdc, 0xb1,
	0xbe, 0x2d, 0x9c, 0xda, 0xc3, 0xc2, 0xa9, 0xfd, 0x5c, 0x38, 0xb5, 0x4f, 0xaf, 0x2a, 0x2d, 0x9b,
	0xc3, 0xce, 0x87, 0x5c, 0x50, 0x11, 0xb1, 0xe5, 0x3a, 0x48, 0x99, 0x88, 0xb9, 0x48, 0xbc, 0xa9,
	0x79, 0x59, 0xea, 0x22, 0xe1, 0xbe, 0x7a, 0x22, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x76, 0x64, 0x10, 0x21, 0x04, 0x00, 0x00,
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
