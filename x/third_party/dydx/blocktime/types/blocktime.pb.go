// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/blocktime/blocktime.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// BlockInfo stores information about a block
type BlockInfo struct {
	Height    uint32    `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a36c561d9e264b, []int{0}
}
func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(m, src)
}
func (m *BlockInfo) XXX_Size() int {
	return m.Size()
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockInfo) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

// AllDowntimeInfo stores information for all downtime durations.
type AllDowntimeInfo struct {
	// The downtime information for each tracked duration. Sorted by duration,
	// ascending. (i.e. the same order as they appear in DowntimeParams).
	Infos []*AllDowntimeInfo_DowntimeInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
}

func (m *AllDowntimeInfo) Reset()         { *m = AllDowntimeInfo{} }
func (m *AllDowntimeInfo) String() string { return proto.CompactTextString(m) }
func (*AllDowntimeInfo) ProtoMessage()    {}
func (*AllDowntimeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a36c561d9e264b, []int{1}
}
func (m *AllDowntimeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllDowntimeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllDowntimeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllDowntimeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllDowntimeInfo.Merge(m, src)
}
func (m *AllDowntimeInfo) XXX_Size() int {
	return m.Size()
}
func (m *AllDowntimeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AllDowntimeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AllDowntimeInfo proto.InternalMessageInfo

func (m *AllDowntimeInfo) GetInfos() []*AllDowntimeInfo_DowntimeInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

// Stores information about downtime. block_info corresponds to the most
// recent block at which a downtime occurred.
type AllDowntimeInfo_DowntimeInfo struct {
	Duration  time.Duration `protobuf:"bytes,1,opt,name=duration,proto3,stdduration" json:"duration"`
	BlockInfo BlockInfo     `protobuf:"bytes,2,opt,name=block_info,json=blockInfo,proto3" json:"block_info"`
}

func (m *AllDowntimeInfo_DowntimeInfo) Reset()         { *m = AllDowntimeInfo_DowntimeInfo{} }
func (m *AllDowntimeInfo_DowntimeInfo) String() string { return proto.CompactTextString(m) }
func (*AllDowntimeInfo_DowntimeInfo) ProtoMessage()    {}
func (*AllDowntimeInfo_DowntimeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a36c561d9e264b, []int{1, 0}
}
func (m *AllDowntimeInfo_DowntimeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllDowntimeInfo_DowntimeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllDowntimeInfo_DowntimeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllDowntimeInfo_DowntimeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllDowntimeInfo_DowntimeInfo.Merge(m, src)
}
func (m *AllDowntimeInfo_DowntimeInfo) XXX_Size() int {
	return m.Size()
}
func (m *AllDowntimeInfo_DowntimeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AllDowntimeInfo_DowntimeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AllDowntimeInfo_DowntimeInfo proto.InternalMessageInfo

func (m *AllDowntimeInfo_DowntimeInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *AllDowntimeInfo_DowntimeInfo) GetBlockInfo() BlockInfo {
	if m != nil {
		return m.BlockInfo
	}
	return BlockInfo{}
}

func init() {
	proto.RegisterType((*BlockInfo)(nil), "joltify.third_party.dydxprotocol.blocktime.BlockInfo")
	proto.RegisterType((*AllDowntimeInfo)(nil), "joltify.third_party.dydxprotocol.blocktime.AllDowntimeInfo")
	proto.RegisterType((*AllDowntimeInfo_DowntimeInfo)(nil), "joltify.third_party.dydxprotocol.blocktime.AllDowntimeInfo.DowntimeInfo")
}

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/blocktime/blocktime.proto", fileDescriptor_07a36c561d9e264b)
}

var fileDescriptor_07a36c561d9e264b = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0x8d, 0x7b, 0xef, 0xad, 0x5a, 0xf7, 0x5e, 0x5d, 0x29, 0x42, 0xa8, 0x64, 0x48, 0xab, 0x4e,
	0x15, 0x12, 0xb6, 0x54, 0xc4, 0xc2, 0x82, 0x88, 0x3a, 0xc0, 0xc0, 0x12, 0x31, 0x75, 0xa0, 0xca,
	0xaf, 0x63, 0x48, 0xed, 0x28, 0x75, 0x45, 0xf3, 0x16, 0x1d, 0xd9, 0x78, 0x00, 0x5e, 0xa4, 0x63,
	0x47, 0x26, 0x40, 0xed, 0x8b, 0xa0, 0x38, 0x3f, 0x0d, 0xed, 0x02, 0x9b, 0xbf, 0x9f, 0x73, 0xce,
	0x77, 0x8e, 0x0c, 0xcf, 0xef, 0x79, 0x28, 0xa8, 0x9f, 0x60, 0x11, 0xd0, 0xd8, 0x1d, 0x47, 0x56,
	0x2c, 0x12, 0xec, 0x26, 0xee, 0x3c, 0x8a, 0xb9, 0xe0, 0x0e, 0x0f, 0xb1, 0x1d, 0x72, 0xe7, 0x41,
	0xd0, 0x89, 0xb7, 0x7d, 0x21, 0x39, 0x54, 0x8f, 0x73, 0x2c, 0xaa, 0x60, 0x51, 0x15, 0x8b, 0x4a,
	0x84, 0x76, 0x40, 0x38, 0xe1, 0xb2, 0x8f, 0xd3, 0x57, 0xc6, 0xa0, 0xe9, 0x84, 0x73, 0x12, 0x7a,
	0x58, 0x56, 0xf6, 0xcc, 0xc7, 0xee, 0x2c, 0xb6, 0x04, 0xe5, 0x2c, 0x9f, 0x77, 0x76, 0xe7, 0x29,
	0xd7, 0x54, 0x58, 0x93, 0x28, 0x5b, 0xe8, 0x11, 0xd8, 0x34, 0x52, 0x8d, 0x6b, 0xe6, 0x73, 0xf5,
	0x10, 0xd6, 0x03, 0x8f, 0x92, 0x40, 0xb4, 0x41, 0x17, 0xf4, 0xff, 0x99, 0x79, 0xa5, 0x1a, 0xb0,
	0x59, 0xe2, 0xda, 0xb5, 0x2e, 0xe8, 0xb7, 0x06, 0x1a, 0xca, 0x98, 0x51, 0xc1, 0x8c, 0x6e, 0x8b,
	0x0d, 0xa3, 0xb1, 0x7c, 0xeb, 0x28, 0x8b, 0xf7, 0x0e, 0x30, 0xb7, 0xb0, 0xde, 0x73, 0x0d, 0xfe,
	0xbf, 0x0c, 0xc3, 0x21, 0x7f, 0x64, 0x69, 0x53, 0xea, 0xdd, 0xc1, 0x3f, 0x94, 0xf9, 0x7c, 0xda,
	0x06, 0xdd, 0x5f, 0xfd, 0xd6, 0xe0, 0x0a, 0x7d, 0x3f, 0x0f, 0xb4, 0xc3, 0x85, 0xaa, 0x85, 0x99,
	0xd1, 0x6a, 0x2f, 0x00, 0xfe, 0xfd, 0x22, 0x78, 0x01, 0x1b, 0x45, 0x40, 0xd2, 0x62, 0x6b, 0x70,
	0xb4, 0xe7, 0x63, 0x98, 0x2f, 0x64, 0x36, 0x9e, 0x52, 0x1b, 0x25, 0x48, 0x1d, 0x41, 0x28, 0x4f,
	0x18, 0xa7, 0x02, 0x79, 0x14, 0x67, 0x3f, 0x39, 0xbb, 0x0c, 0xdb, 0xf8, 0x9d, 0xd2, 0x9b, 0x4d,
	0xbb, 0x6c, 0x90, 0xe5, 0x5a, 0x07, 0xab, 0xb5, 0x0e, 0x3e, 0xd6, 0x3a, 0x58, 0x6c, 0x74, 0x65,
	0xb5, 0xd1, 0x95, 0xd7, 0x8d, 0xae, 0x8c, 0x6e, 0x08, 0x15, 0xc1, 0xcc, 0x46, 0x0e, 0x9f, 0xe0,
	0x5c, 0xeb, 0xc4, 0xa7, 0xcc, 0x62, 0x8e, 0x57, 0xd4, 0xe3, 0xd0, 0x63, 0x2e, 0x65, 0x04, 0xcf,
	0xf7, 0x3e, 0x62, 0xe5, 0x03, 0x8a, 0x24, 0xf2, 0xa6, 0x76, 0x5d, 0xde, 0x75, 0xfa, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0xc2, 0x86, 0xbe, 0xbb, 0x02, 0x00, 0x00,
}

func (m *BlockInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBlocktime(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintBlocktime(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AllDowntimeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllDowntimeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllDowntimeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Infos) > 0 {
		for iNdEx := len(m.Infos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Infos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlocktime(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AllDowntimeInfo_DowntimeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllDowntimeInfo_DowntimeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllDowntimeInfo_DowntimeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BlockInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBlocktime(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintBlocktime(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBlocktime(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlocktime(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlocktime(uint64(m.Height))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovBlocktime(uint64(l))
	return n
}

func (m *AllDowntimeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Infos) > 0 {
		for _, e := range m.Infos {
			l = e.Size()
			n += 1 + l + sovBlocktime(uint64(l))
		}
	}
	return n
}

func (m *AllDowntimeInfo_DowntimeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovBlocktime(uint64(l))
	l = m.BlockInfo.Size()
	n += 1 + l + sovBlocktime(uint64(l))
	return n
}

func sovBlocktime(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlocktime(x uint64) (n int) {
	return sovBlocktime(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlocktime
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
			return fmt.Errorf("proto: BlockInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlocktime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlocktime
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
				return ErrInvalidLengthBlocktime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlocktime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlocktime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlocktime
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
func (m *AllDowntimeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlocktime
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
			return fmt.Errorf("proto: AllDowntimeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllDowntimeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Infos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlocktime
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
				return ErrInvalidLengthBlocktime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlocktime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Infos = append(m.Infos, &AllDowntimeInfo_DowntimeInfo{})
			if err := m.Infos[len(m.Infos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlocktime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlocktime
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
func (m *AllDowntimeInfo_DowntimeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlocktime
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
			return fmt.Errorf("proto: DowntimeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DowntimeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlocktime
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
				return ErrInvalidLengthBlocktime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlocktime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlocktime
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
				return ErrInvalidLengthBlocktime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlocktime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlocktime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlocktime
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
func skipBlocktime(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlocktime
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
					return 0, ErrIntOverflowBlocktime
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
					return 0, ErrIntOverflowBlocktime
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
				return 0, ErrInvalidLengthBlocktime
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlocktime
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlocktime
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlocktime        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlocktime          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlocktime = fmt.Errorf("proto: unexpected end of group")
)
