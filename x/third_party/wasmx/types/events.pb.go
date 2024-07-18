// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: injective/wasmx/v1/events.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventContractExecution struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Response        []byte `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	OtherError      string `protobuf:"bytes,3,opt,name=other_error,json=otherError,proto3" json:"other_error,omitempty"`
	ExecutionError  string `protobuf:"bytes,4,opt,name=execution_error,json=executionError,proto3" json:"execution_error,omitempty"`
}

func (m *EventContractExecution) Reset()         { *m = EventContractExecution{} }
func (m *EventContractExecution) String() string { return proto.CompactTextString(m) }
func (*EventContractExecution) ProtoMessage()    {}
func (*EventContractExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ba06c5f04cb490, []int{0}
}

func (m *EventContractExecution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *EventContractExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventContractExecution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *EventContractExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventContractExecution.Merge(m, src)
}

func (m *EventContractExecution) XXX_Size() int {
	return m.Size()
}

func (m *EventContractExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_EventContractExecution.DiscardUnknown(m)
}

var xxx_messageInfo_EventContractExecution proto.InternalMessageInfo

func (m *EventContractExecution) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *EventContractExecution) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *EventContractExecution) GetOtherError() string {
	if m != nil {
		return m.OtherError
	}
	return ""
}

func (m *EventContractExecution) GetExecutionError() string {
	if m != nil {
		return m.ExecutionError
	}
	return ""
}

type EventContractRegistered struct {
	ContractAddress    string      `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	GasPrice           uint64      `protobuf:"varint,3,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	ShouldPinContract  bool        `protobuf:"varint,4,opt,name=should_pin_contract,json=shouldPinContract,proto3" json:"should_pin_contract,omitempty"`
	IsMigrationAllowed bool        `protobuf:"varint,5,opt,name=is_migration_allowed,json=isMigrationAllowed,proto3" json:"is_migration_allowed,omitempty"`
	CodeId             uint64      `protobuf:"varint,6,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	AdminAddress       string      `protobuf:"bytes,7,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
	GranterAddress     string      `protobuf:"bytes,8,opt,name=granter_address,json=granterAddress,proto3" json:"granter_address,omitempty"`
	FundingMode        FundingMode `protobuf:"varint,9,opt,name=funding_mode,json=fundingMode,proto3,enum=injective.wasmx.v1.FundingMode" json:"funding_mode,omitempty"`
}

func (m *EventContractRegistered) Reset()         { *m = EventContractRegistered{} }
func (m *EventContractRegistered) String() string { return proto.CompactTextString(m) }
func (*EventContractRegistered) ProtoMessage()    {}
func (*EventContractRegistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ba06c5f04cb490, []int{1}
}

func (m *EventContractRegistered) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *EventContractRegistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventContractRegistered.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *EventContractRegistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventContractRegistered.Merge(m, src)
}

func (m *EventContractRegistered) XXX_Size() int {
	return m.Size()
}

func (m *EventContractRegistered) XXX_DiscardUnknown() {
	xxx_messageInfo_EventContractRegistered.DiscardUnknown(m)
}

var xxx_messageInfo_EventContractRegistered proto.InternalMessageInfo

func (m *EventContractRegistered) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *EventContractRegistered) GetGasPrice() uint64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *EventContractRegistered) GetShouldPinContract() bool {
	if m != nil {
		return m.ShouldPinContract
	}
	return false
}

func (m *EventContractRegistered) GetIsMigrationAllowed() bool {
	if m != nil {
		return m.IsMigrationAllowed
	}
	return false
}

func (m *EventContractRegistered) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *EventContractRegistered) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

func (m *EventContractRegistered) GetGranterAddress() string {
	if m != nil {
		return m.GranterAddress
	}
	return ""
}

func (m *EventContractRegistered) GetFundingMode() FundingMode {
	if m != nil {
		return m.FundingMode
	}
	return FundingMode_Unspecified
}

type EventContractDeregistered struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *EventContractDeregistered) Reset()         { *m = EventContractDeregistered{} }
func (m *EventContractDeregistered) String() string { return proto.CompactTextString(m) }
func (*EventContractDeregistered) ProtoMessage()    {}
func (*EventContractDeregistered) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ba06c5f04cb490, []int{2}
}

func (m *EventContractDeregistered) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *EventContractDeregistered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventContractDeregistered.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *EventContractDeregistered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventContractDeregistered.Merge(m, src)
}

func (m *EventContractDeregistered) XXX_Size() int {
	return m.Size()
}

func (m *EventContractDeregistered) XXX_DiscardUnknown() {
	xxx_messageInfo_EventContractDeregistered.DiscardUnknown(m)
}

var xxx_messageInfo_EventContractDeregistered proto.InternalMessageInfo

func (m *EventContractDeregistered) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*EventContractExecution)(nil), "injective.wasmx.v1.EventContractExecution")
	proto.RegisterType((*EventContractRegistered)(nil), "injective.wasmx.v1.EventContractRegistered")
	proto.RegisterType((*EventContractDeregistered)(nil), "injective.wasmx.v1.EventContractDeregistered")
}

func init() { proto.RegisterFile("injective/wasmx/v1/events.proto", fileDescriptor_f2ba06c5f04cb490) }

var fileDescriptor_f2ba06c5f04cb490 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0xe3, 0x7e, 0xfd, 0xd2, 0x64, 0x1a, 0x1a, 0x18, 0x2a, 0x6a, 0x82, 0xe4, 0x84, 0xb0,
	0x68, 0x58, 0x60, 0x53, 0x78, 0x82, 0x16, 0x52, 0xa9, 0x82, 0x4a, 0x95, 0x97, 0x6c, 0xac, 0x89,
	0xe7, 0xd6, 0x19, 0x64, 0xcf, 0xb5, 0x66, 0xc6, 0x69, 0x79, 0x0b, 0x5e, 0x83, 0x37, 0x61, 0xd9,
	0x25, 0x4b, 0x94, 0xec, 0x79, 0x06, 0xe4, 0xf1, 0x1f, 0xa8, 0xe8, 0x06, 0x76, 0x73, 0xcf, 0x39,
	0x77, 0xee, 0x4f, 0x57, 0x97, 0x8c, 0x85, 0xfc, 0x08, 0xb1, 0x11, 0x2b, 0x08, 0xae, 0x98, 0xce,
	0xae, 0x83, 0xd5, 0x51, 0x00, 0x2b, 0x90, 0x46, 0xfb, 0xb9, 0x42, 0x83, 0x94, 0xb6, 0x01, 0xdf,
	0x06, 0xfc, 0xd5, 0xd1, 0xc8, 0xbb, 0xa3, 0xa9, 0x32, 0x6d, 0xcf, 0xe8, 0xe9, 0x1d, 0x7e, 0xae,
	0x30, 0x47, 0xcd, 0xd2, 0x3a, 0xb2, 0x9f, 0x60, 0x82, 0xf6, 0x19, 0x94, 0xaf, 0x4a, 0x9d, 0x7e,
	0x71, 0xc8, 0xa3, 0x79, 0x39, 0xfd, 0x0d, 0x4a, 0xa3, 0x58, 0x6c, 0xe6, 0xd7, 0x10, 0x17, 0x46,
	0xa0, 0xa4, 0xcf, 0xc9, 0xfd, 0xb8, 0x16, 0x23, 0xc6, 0xb9, 0x02, 0xad, 0x5d, 0x67, 0xe2, 0xcc,
	0xfa, 0xe1, 0xb0, 0xd1, 0x8f, 0x2b, 0x99, 0x8e, 0x48, 0x4f, 0x81, 0xce, 0x51, 0x6a, 0x70, 0xb7,
	0x26, 0xce, 0x6c, 0x10, 0xb6, 0x35, 0x1d, 0x93, 0x5d, 0x34, 0x4b, 0x50, 0x11, 0x28, 0x85, 0xca,
	0xfd, 0xcf, 0xfe, 0x40, 0xac, 0x34, 0x2f, 0x15, 0x7a, 0x48, 0x86, 0xd0, 0x0c, 0xad, 0x43, 0xdb,
	0x36, 0xb4, 0xd7, 0xca, 0x36, 0x38, 0xfd, 0xb1, 0x45, 0x0e, 0x6e, 0xb1, 0x86, 0x90, 0x08, 0x6d,
	0x40, 0x01, 0xff, 0x1b, 0xd8, 0x27, 0xa4, 0x9f, 0x30, 0x1d, 0xe5, 0x4a, 0xc4, 0x60, 0x71, 0xb6,
	0xc3, 0x5e, 0xc2, 0xf4, 0x45, 0x59, 0x53, 0x9f, 0x3c, 0xd4, 0x4b, 0x2c, 0x52, 0x1e, 0xe5, 0x42,
	0x46, 0x4d, 0xab, 0x05, 0xea, 0x85, 0x0f, 0x2a, 0xeb, 0x42, 0xc8, 0x86, 0x80, 0xbe, 0x24, 0xfb,
	0x42, 0x47, 0x99, 0x48, 0x14, 0xb3, 0xfc, 0x2c, 0x4d, 0xf1, 0x0a, 0xb8, 0xfb, 0xbf, 0x6d, 0xa0,
	0x42, 0x9f, 0x37, 0xd6, 0x71, 0xe5, 0xd0, 0x03, 0xb2, 0x13, 0x23, 0x87, 0x48, 0x70, 0xb7, 0x6b,
	0x87, 0x77, 0xcb, 0xf2, 0x8c, 0xd3, 0x67, 0xe4, 0x1e, 0xe3, 0x99, 0x90, 0x2d, 0xff, 0x8e, 0xe5,
	0x1f, 0x58, 0xb1, 0x81, 0x3f, 0x24, 0xc3, 0x44, 0x31, 0x69, 0x40, 0xb5, 0xb1, 0x5e, 0xb5, 0xac,
	0x5a, 0x6e, 0x82, 0x27, 0x64, 0x70, 0x59, 0x48, 0x2e, 0x64, 0x12, 0x65, 0xc8, 0xc1, 0xed, 0x4f,
	0x9c, 0xd9, 0xde, 0xab, 0xb1, 0xff, 0xe7, 0x71, 0xf9, 0xa7, 0x55, 0xee, 0x1c, 0x39, 0x84, 0xbb,
	0x97, 0xbf, 0x8a, 0xe9, 0x29, 0x79, 0x7c, 0x6b, 0xdf, 0x6f, 0x41, 0xfd, 0xcb, 0xc6, 0x4f, 0xe0,
	0xeb, 0xda, 0x73, 0x6e, 0xd6, 0x9e, 0xf3, 0x7d, 0xed, 0x39, 0x9f, 0x37, 0x5e, 0xe7, 0x66, 0xe3,
	0x75, 0xbe, 0x6d, 0xbc, 0xce, 0x87, 0x77, 0x89, 0x30, 0xcb, 0x62, 0xe1, 0xc7, 0x98, 0x05, 0x67,
	0x0d, 0xd9, 0x7b, 0xb6, 0xd0, 0x41, 0xcb, 0xf9, 0x22, 0x46, 0x05, 0xbf, 0x97, 0x4b, 0x26, 0x64,
	0x90, 0x21, 0x2f, 0x52, 0xd0, 0xf5, 0xb5, 0x9b, 0x4f, 0x39, 0xe8, 0x45, 0xd7, 0x9e, 0xf4, 0xeb,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69, 0x98, 0x56, 0x6d, 0x62, 0x03, 0x00, 0x00,
}

func (m *EventContractExecution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventContractExecution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventContractExecution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExecutionError) > 0 {
		i -= len(m.ExecutionError)
		copy(dAtA[i:], m.ExecutionError)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ExecutionError)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OtherError) > 0 {
		i -= len(m.OtherError)
		copy(dAtA[i:], m.OtherError)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.OtherError)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Response) > 0 {
		i -= len(m.Response)
		copy(dAtA[i:], m.Response)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Response)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventContractRegistered) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventContractRegistered) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventContractRegistered) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FundingMode != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.FundingMode))
		i--
		dAtA[i] = 0x48
	}
	if len(m.GranterAddress) > 0 {
		i -= len(m.GranterAddress)
		copy(dAtA[i:], m.GranterAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.GranterAddress)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if m.CodeId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x30
	}
	if m.IsMigrationAllowed {
		i--
		if m.IsMigrationAllowed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.ShouldPinContract {
		i--
		if m.ShouldPinContract {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.GasPrice != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.GasPrice))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventContractDeregistered) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventContractDeregistered) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventContractDeregistered) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *EventContractExecution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.OtherError)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ExecutionError)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventContractRegistered) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.GasPrice != 0 {
		n += 1 + sovEvents(uint64(m.GasPrice))
	}
	if m.ShouldPinContract {
		n += 2
	}
	if m.IsMigrationAllowed {
		n += 2
	}
	if m.CodeId != 0 {
		n += 1 + sovEvents(uint64(m.CodeId))
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.GranterAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.FundingMode != 0 {
		n += 1 + sovEvents(uint64(m.FundingMode))
	}
	return n
}

func (m *EventContractDeregistered) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *EventContractExecution) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventContractExecution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventContractExecution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response[:0], dAtA[iNdEx:postIndex]...)
			if m.Response == nil {
				m.Response = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OtherError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecutionError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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

func (m *EventContractRegistered) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventContractRegistered: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventContractRegistered: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShouldPinContract", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
			m.ShouldPinContract = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsMigrationAllowed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
			m.IsMigrationAllowed = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GranterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GranterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundingMode", wireType)
			}
			m.FundingMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FundingMode |= FundingMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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

func (m *EventContractDeregistered) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventContractDeregistered: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventContractDeregistered: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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

func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
