// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/pricefeed/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the pricefeed module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params       Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PostedPrices PostedPrices `protobuf:"bytes,2,rep,name=posted_prices,json=postedPrices,proto3,castrepeated=PostedPrices" json:"posted_prices"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc4210a18c39c1e5, []int{0}
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

func (m *GenesisState) GetPostedPrices() PostedPrices {
	if m != nil {
		return m.PostedPrices
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "joltify.third_party.pricefeed.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("joltify/third_party/pricefeed/v1beta1/genesis.proto", fileDescriptor_cc4210a18c39c1e5)
}

var fileDescriptor_cc4210a18c39c1e5 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0x2a, 0x1d, 0xd2, 0xba, 0x94, 0x0e, 0xa5, 0xc3, 0xb5, 0x08, 0x42, 0x97, 0xde,
	0xd1, 0xf4, 0x1b, 0x64, 0x51, 0x70, 0x29, 0x75, 0x73, 0x09, 0x97, 0xe4, 0x4d, 0x7a, 0xa5, 0xc9,
	0x1d, 0x77, 0xa7, 0x98, 0x6f, 0xe1, 0xc7, 0x10, 0x3f, 0x49, 0x71, 0xea, 0xe8, 0xa4, 0x35, 0xf9,
	0x22, 0x92, 0x3f, 0x4a, 0x16, 0x21, 0xdb, 0xfb, 0x3e, 0xf0, 0xfb, 0xf1, 0xf0, 0xd8, 0xab, 0x9d,
	0xd8, 0x1b, 0x1e, 0x65, 0xd4, 0x6c, 0xb9, 0x0a, 0x3d, 0xc9, 0x94, 0xc9, 0xa8, 0x54, 0x3c, 0x80,
	0x08, 0x20, 0xa4, 0x4f, 0x4b, 0x1f, 0x0c, 0x5b, 0xd2, 0x18, 0x52, 0xd0, 0x5c, 0x13, 0xa9, 0x84,
	0x11, 0xc3, 0xeb, 0x06, 0x22, 0x2d, 0x88, 0xfc, 0x41, 0xa4, 0x81, 0x26, 0xcb, 0x6e, 0x6e, 0x6d,
	0x84, 0x82, 0xda, 0x3c, 0x19, 0xc5, 0x22, 0x16, 0xd5, 0x49, 0xcb, 0xab, 0x4e, 0xaf, 0xde, 0x91,
	0x3d, 0xb8, 0xa9, 0x1b, 0xdc, 0x1b, 0x66, 0x60, 0x78, 0x67, 0xf7, 0x24, 0x53, 0x2c, 0xd1, 0x63,
	0x34, 0x43, 0xf3, 0xbe, 0xb3, 0x20, 0x9d, 0x1a, 0x91, 0x75, 0x05, 0xb9, 0x17, 0x87, 0xcf, 0xa9,
	0xb5, 0x69, 0x14, 0xc3, 0xc4, 0xbe, 0x94, 0x42, 0x1b, 0x08, 0xbd, 0x0a, 0xd0, 0xe3, 0xb3, 0xd9,
	0xf9, 0xbc, 0xef, 0x38, 0x5d, 0x9d, 0x15, 0xbb, 0x2e, 0x73, 0x77, 0x54, 0x8a, 0xdf, 0xbe, 0xa6,
	0x83, 0x56, 0xa8, 0x37, 0x03, 0xd9, 0xfa, 0xdc, 0xdd, 0xe9, 0x1b, 0xa3, 0xd7, 0x1c, 0xa3, 0x43,
	0x8e, 0xd1, 0x31, 0xc7, 0xe8, 0x94, 0x63, 0xf4, 0x52, 0x60, 0xeb, 0x58, 0x60, 0xeb, 0xa3, 0xc0,
	0xd6, 0xc3, 0x6d, 0xcc, 0xcd, 0xf6, 0xd1, 0x27, 0x81, 0x48, 0x68, 0xd3, 0x61, 0x11, 0xf1, 0x94,
	0xa5, 0x01, 0xfc, 0xfe, 0xde, 0x1e, 0xd2, 0x90, 0xa7, 0x31, 0x7d, 0xfe, 0x67, 0x5c, 0x93, 0x49,
	0xd0, 0x7e, 0xaf, 0xda, 0x6f, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x07, 0xf1, 0x21, 0xe6,
	0x01, 0x00, 0x00,
}

func (this *GenesisState) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GenesisState")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GenesisState but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GenesisState but is not nil && this == nil")
	}
	if !this.Params.Equal(&that1.Params) {
		return fmt.Errorf("Params this(%v) Not Equal that(%v)", this.Params, that1.Params)
	}
	if len(this.PostedPrices) != len(that1.PostedPrices) {
		return fmt.Errorf("PostedPrices this(%v) Not Equal that(%v)", len(this.PostedPrices), len(that1.PostedPrices))
	}
	for i := range this.PostedPrices {
		if !this.PostedPrices[i].Equal(&that1.PostedPrices[i]) {
			return fmt.Errorf("PostedPrices this[%v](%v) Not Equal that[%v](%v)", i, this.PostedPrices[i], i, that1.PostedPrices[i])
		}
	}
	return nil
}
func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	if len(this.PostedPrices) != len(that1.PostedPrices) {
		return false
	}
	for i := range this.PostedPrices {
		if !this.PostedPrices[i].Equal(&that1.PostedPrices[i]) {
			return false
		}
	}
	return true
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
	if len(m.PostedPrices) > 0 {
		for iNdEx := len(m.PostedPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostedPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PostedPrices) > 0 {
		for _, e := range m.PostedPrices {
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
				return fmt.Errorf("proto: wrong wireType = %d for field PostedPrices", wireType)
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
			m.PostedPrices = append(m.PostedPrices, PostedPrice{})
			if err := m.PostedPrices[len(m.PostedPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
