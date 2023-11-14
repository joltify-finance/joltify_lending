// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/vault/staking.proto

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

type Params struct {
	BlockChurnInterval int64                                    `protobuf:"varint,1,opt,name=block_churn_interval,json=blockChurnInterval,proto3" json:"block_churn_interval,omitempty" yaml:"block_churn_interval"`
	Power              int64                                    `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty" yaml:"power"`
	Step               int64                                    `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty" yaml:"step"`
	CandidateRatio     github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,4,opt,name=candidate_ratio,json=candidateRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"candidate_ratio" yaml:"candidate_ratio"`
	TargetQuota        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=target_quota,json=targetQuota,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"target_quota" yaml:"target_quota"`
	HistoryLength      int32                                    `protobuf:"varint,6,opt,name=history_length,json=historyLength,proto3" json:"history_length,omitempty" yaml:"history_length"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e8951aac3bd30f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBlockChurnInterval() int64 {
	if m != nil {
		return m.BlockChurnInterval
	}
	return 0
}

func (m *Params) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *Params) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *Params) GetTargetQuota() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TargetQuota
	}
	return nil
}

func (m *Params) GetHistoryLength() int32 {
	if m != nil {
		return m.HistoryLength
	}
	return 0
}

type Validator struct {
	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Power  int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e8951aac3bd30f, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Validator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

type StandbyPower struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Power int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *StandbyPower) Reset()         { *m = StandbyPower{} }
func (m *StandbyPower) String() string { return proto.CompactTextString(m) }
func (*StandbyPower) ProtoMessage()    {}
func (*StandbyPower) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e8951aac3bd30f, []int{2}
}
func (m *StandbyPower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StandbyPower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StandbyPower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StandbyPower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StandbyPower.Merge(m, src)
}
func (m *StandbyPower) XXX_Size() int {
	return m.Size()
}
func (m *StandbyPower) XXX_DiscardUnknown() {
	xxx_messageInfo_StandbyPower.DiscardUnknown(m)
}

var xxx_messageInfo_StandbyPower proto.InternalMessageInfo

func (m *StandbyPower) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *StandbyPower) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

type Validators struct {
	AllValidators []*Validator `protobuf:"bytes,1,rep,name=all_validators,json=allValidators,proto3" json:"all_validators,omitempty"`
	Height        int64        `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Validators) Reset()         { *m = Validators{} }
func (m *Validators) String() string { return proto.CompactTextString(m) }
func (*Validators) ProtoMessage()    {}
func (*Validators) Descriptor() ([]byte, []int) {
	return fileDescriptor_84e8951aac3bd30f, []int{3}
}
func (m *Validators) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validators) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validators.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validators) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validators.Merge(m, src)
}
func (m *Validators) XXX_Size() int {
	return m.Size()
}
func (m *Validators) XXX_DiscardUnknown() {
	xxx_messageInfo_Validators.DiscardUnknown(m)
}

var xxx_messageInfo_Validators proto.InternalMessageInfo

func (m *Validators) GetAllValidators() []*Validator {
	if m != nil {
		return m.AllValidators
	}
	return nil
}

func (m *Validators) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "joltify.vault.Params")
	proto.RegisterType((*Validator)(nil), "joltify.vault.Validator")
	proto.RegisterType((*StandbyPower)(nil), "joltify.vault.StandbyPower")
	proto.RegisterType((*Validators)(nil), "joltify.vault.Validators")
}

func init() { proto.RegisterFile("joltify/vault/staking.proto", fileDescriptor_84e8951aac3bd30f) }

var fileDescriptor_84e8951aac3bd30f = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0x8e, 0x49, 0x1b, 0x91, 0xcb, 0x8f, 0xa2, 0x23, 0x44, 0x6e, 0x2b, 0xf9, 0x22, 0x23, 0x55,
	0x59, 0x6a, 0xab, 0xb0, 0x94, 0x2e, 0xa0, 0x14, 0x09, 0x90, 0x18, 0x5a, 0x23, 0x31, 0xb0, 0x58,
	0x67, 0xfb, 0x6a, 0x1f, 0x71, 0x7c, 0xa9, 0x7d, 0x09, 0xf8, 0x0f, 0x60, 0x67, 0x64, 0xec, 0xcc,
	0x5f, 0xd2, 0xb1, 0x23, 0x62, 0x30, 0x28, 0x59, 0xd8, 0x90, 0xfc, 0x17, 0x20, 0xdf, 0x1d, 0x51,
	0x82, 0x40, 0x62, 0xf2, 0xbd, 0xef, 0x7b, 0xdf, 0xf7, 0xee, 0xde, 0xf3, 0x03, 0xfb, 0x6f, 0x59,
	0xcc, 0xe9, 0x45, 0x6e, 0xcf, 0xf1, 0x2c, 0xe6, 0x76, 0xc6, 0xf1, 0x98, 0x26, 0xa1, 0x35, 0x4d,
	0x19, 0x67, 0xb0, 0xa3, 0x48, 0x4b, 0x90, 0x7b, 0xbd, 0x90, 0x85, 0x4c, 0x30, 0x76, 0x75, 0x92,
	0x49, 0x7b, 0x86, 0xcf, 0xb2, 0x09, 0xcb, 0x6c, 0x0f, 0x67, 0xc4, 0x9e, 0x1f, 0x79, 0x84, 0xe3,
	0x23, 0xdb, 0x67, 0x34, 0x91, 0xbc, 0xf9, 0xb3, 0x0e, 0x1a, 0x67, 0x38, 0xc5, 0x93, 0x0c, 0x9e,
	0x83, 0x9e, 0x17, 0x33, 0x7f, 0xec, 0xfa, 0xd1, 0x2c, 0x4d, 0x5c, 0x9a, 0x70, 0x92, 0xce, 0x71,
	0xac, 0x6b, 0x03, 0x6d, 0x58, 0x1f, 0xa1, 0xb2, 0x40, 0xfb, 0x39, 0x9e, 0xc4, 0x27, 0xe6, 0xdf,
	0xb2, 0x4c, 0x07, 0x0a, 0xf8, 0xb4, 0x42, 0x5f, 0x28, 0x10, 0x1e, 0x80, 0xed, 0x29, 0x7b, 0x47,
	0x52, 0xfd, 0x96, 0xf0, 0xb8, 0x53, 0x16, 0xa8, 0x2d, 0x3d, 0x04, 0x6c, 0x3a, 0x92, 0x86, 0xf7,
	0xc1, 0x56, 0xc6, 0xc9, 0x54, 0xaf, 0x8b, 0xb4, 0x9d, 0xb2, 0x40, 0x2d, 0x99, 0x56, 0xa1, 0xa6,
	0x23, 0x48, 0x78, 0x09, 0x76, 0x7c, 0x9c, 0x04, 0x34, 0xc0, 0x9c, 0xb8, 0x29, 0xe6, 0x94, 0xe9,
	0x5b, 0x03, 0x6d, 0xd8, 0x1c, 0x3d, 0xbf, 0x2e, 0x50, 0xed, 0x6b, 0x81, 0x0e, 0x42, 0xca, 0xa3,
	0x99, 0x67, 0xf9, 0x6c, 0x62, 0xab, 0x67, 0xcb, 0xcf, 0x61, 0x16, 0x8c, 0x6d, 0x9e, 0x4f, 0x49,
	0x66, 0x3d, 0x25, 0x7e, 0x59, 0xa0, 0xbe, 0x74, 0xff, 0xc3, 0xce, 0x74, 0xba, 0x2b, 0xc4, 0xa9,
	0x00, 0xf8, 0x41, 0x03, 0x6d, 0x8e, 0xd3, 0x90, 0x70, 0xf7, 0x72, 0xc6, 0x38, 0xd6, 0xb7, 0x07,
	0xf5, 0x61, 0xeb, 0xc1, 0xae, 0x25, 0x7d, 0xad, 0xaa, 0xab, 0x96, 0xea, 0xaa, 0x75, 0xca, 0x68,
	0x32, 0x7a, 0x56, 0xdd, 0xa5, 0x2c, 0xd0, 0x5d, 0x59, 0x61, 0x5d, 0x6c, 0x7e, 0xfe, 0x86, 0x86,
	0xff, 0x71, 0xc5, 0xca, 0x27, 0x73, 0x5a, 0x52, 0x7a, 0x5e, 0x29, 0xe1, 0x13, 0xd0, 0x8d, 0x68,
	0xc6, 0x59, 0x9a, 0xbb, 0x31, 0x49, 0x42, 0x1e, 0xe9, 0x8d, 0x81, 0x36, 0xdc, 0x1e, 0xed, 0x96,
	0x05, 0xba, 0x27, 0x2b, 0x6d, 0xf2, 0xa6, 0xd3, 0x51, 0xc0, 0x4b, 0x11, 0x9f, 0xdc, 0xfe, 0x74,
	0x85, 0x6a, 0x3f, 0xae, 0x90, 0x66, 0x3e, 0x02, 0xcd, 0xd7, 0x38, 0xae, 0x1e, 0xc9, 0x52, 0xd8,
	0x07, 0x8d, 0xe9, 0xcc, 0x1b, 0x93, 0x5c, 0x4c, 0xb9, 0xed, 0xa8, 0x08, 0xf6, 0x36, 0x06, 0xa7,
	0xc6, 0x64, 0x1e, 0x83, 0xf6, 0x2b, 0x8e, 0x93, 0xc0, 0xcb, 0xcf, 0xc4, 0xd8, 0x20, 0xd8, 0xc2,
	0x41, 0x90, 0x0a, 0x6d, 0xd3, 0x11, 0xe7, 0x7f, 0x28, 0x09, 0x00, 0xab, 0xa2, 0x19, 0x7c, 0x0c,
	0xba, 0x38, 0x8e, 0xdd, 0xf9, 0x0a, 0xd1, 0x35, 0xd1, 0x57, 0xdd, 0xda, 0xf8, 0xa5, 0xad, 0x95,
	0xc4, 0xe9, 0xe0, 0x38, 0x5e, 0x33, 0xe8, 0x83, 0x46, 0x44, 0x68, 0x18, 0x71, 0x55, 0x45, 0x45,
	0x23, 0xe7, 0x7a, 0x61, 0x68, 0x37, 0x0b, 0x43, 0xfb, 0xbe, 0x30, 0xb4, 0x8f, 0x4b, 0xa3, 0x76,
	0xb3, 0x34, 0x6a, 0x5f, 0x96, 0x46, 0xed, 0xcd, 0xf1, 0x5a, 0xe3, 0x55, 0x91, 0xc3, 0x0b, 0x9a,
	0xe0, 0xc4, 0x27, 0xbf, 0xe3, 0xaa, 0x75, 0x01, 0x4d, 0x42, 0xfb, 0xbd, 0x5a, 0x37, 0x31, 0x0e,
	0xaf, 0x21, 0x16, 0xe5, 0xe1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xbb, 0x8b, 0xfa, 0x8c,
	0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BlockChurnInterval != that1.BlockChurnInterval {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	if !this.CandidateRatio.Equal(that1.CandidateRatio) {
		return false
	}
	if len(this.TargetQuota) != len(that1.TargetQuota) {
		return false
	}
	for i := range this.TargetQuota {
		if !this.TargetQuota[i].Equal(&that1.TargetQuota[i]) {
			return false
		}
	}
	if this.HistoryLength != that1.HistoryLength {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HistoryLength != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.HistoryLength))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TargetQuota) > 0 {
		for iNdEx := len(m.TargetQuota) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TargetQuota[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size := m.CandidateRatio.Size()
		i -= size
		if _, err := m.CandidateRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Step != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Power != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockChurnInterval != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.BlockChurnInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StandbyPower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StandbyPower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StandbyPower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Validators) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validators) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validators) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AllValidators) > 0 {
		for iNdEx := len(m.AllValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStaking(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockChurnInterval != 0 {
		n += 1 + sovStaking(uint64(m.BlockChurnInterval))
	}
	if m.Power != 0 {
		n += 1 + sovStaking(uint64(m.Power))
	}
	if m.Step != 0 {
		n += 1 + sovStaking(uint64(m.Step))
	}
	l = m.CandidateRatio.Size()
	n += 1 + l + sovStaking(uint64(l))
	if len(m.TargetQuota) > 0 {
		for _, e := range m.TargetQuota {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	if m.HistoryLength != 0 {
		n += 1 + sovStaking(uint64(m.HistoryLength))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovStaking(uint64(m.Power))
	}
	return n
}

func (m *StandbyPower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovStaking(uint64(m.Power))
	}
	return n
}

func (m *Validators) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllValidators) > 0 {
		for _, e := range m.AllValidators {
			l = e.Size()
			n += 1 + l + sovStaking(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovStaking(uint64(m.Height))
	}
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockChurnInterval", wireType)
			}
			m.BlockChurnInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockChurnInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CandidateRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetQuota", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetQuota = append(m.TargetQuota, types.Coin{})
			if err := m.TargetQuota[len(m.TargetQuota)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoryLength", wireType)
			}
			m.HistoryLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoryLength |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *StandbyPower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: StandbyPower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StandbyPower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func (m *Validators) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: Validators: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validators: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllValidators = append(m.AllValidators, &Validator{})
			if err := m.AllValidators[len(m.AllValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
