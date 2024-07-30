// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/pricefeed/v1beta1/tx.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// MsgPostPrice represents a method for creating a new post price
type MsgPostPrice struct {
	// address of client
	FromAddress string                      `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	MarketID    string                      `protobuf:"bytes,2,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
	Price       cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=price,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"price"`
	Expiry      time.Time                   `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *MsgPostPrice) Reset()         { *m = MsgPostPrice{} }
func (m *MsgPostPrice) String() string { return proto.CompactTextString(m) }
func (*MsgPostPrice) ProtoMessage()    {}
func (*MsgPostPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d42ed7adf11289b, []int{0}
}
func (m *MsgPostPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPrice.Merge(m, src)
}
func (m *MsgPostPrice) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPrice.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPrice proto.InternalMessageInfo

// MsgPostPriceResponse defines the Msg/PostPrice response type.
type MsgPostPriceResponse struct {
}

func (m *MsgPostPriceResponse) Reset()         { *m = MsgPostPriceResponse{} }
func (m *MsgPostPriceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPostPriceResponse) ProtoMessage()    {}
func (*MsgPostPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d42ed7adf11289b, []int{1}
}
func (m *MsgPostPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPostPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPostPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPostPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPostPriceResponse.Merge(m, src)
}
func (m *MsgPostPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPostPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPostPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPostPriceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPostPrice)(nil), "joltify.third_party.pricefeed.v1beta1.MsgPostPrice")
	proto.RegisterType((*MsgPostPriceResponse)(nil), "joltify.third_party.pricefeed.v1beta1.MsgPostPriceResponse")
}

func init() {
	proto.RegisterFile("joltify/third_party/pricefeed/v1beta1/tx.proto", fileDescriptor_4d42ed7adf11289b)
}

var fileDescriptor_4d42ed7adf11289b = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xb5, 0x50, 0x25, 0xd7, 0x2c, 0x58, 0x15, 0x44, 0x41, 0xb2, 0x4b, 0x11, 0x52, 0xa9,
	0xc4, 0x9d, 0x92, 0x4e, 0xfc, 0x58, 0x88, 0x3a, 0x50, 0xa9, 0x91, 0x2a, 0x8b, 0x89, 0x25, 0xba,
	0xd8, 0xcf, 0x97, 0x6b, 0x73, 0x3e, 0xeb, 0xee, 0x5a, 0xc5, 0x03, 0x0b, 0x62, 0x40, 0x4c, 0xfd,
	0x13, 0x90, 0x58, 0x18, 0xfb, 0x67, 0x74, 0xec, 0x88, 0x18, 0x42, 0x71, 0x86, 0xfe, 0x1b, 0xc8,
	0x3f, 0x02, 0x61, 0x40, 0x82, 0x2e, 0xf6, 0xbd, 0xef, 0x7d, 0xdf, 0xbd, 0xf7, 0x7d, 0x3a, 0x4c,
	0x8e, 0xd4, 0xc4, 0x8a, 0x38, 0xa3, 0x76, 0x2c, 0x74, 0x34, 0x4c, 0x99, 0xb6, 0x19, 0x4d, 0xb5,
	0x08, 0x21, 0x06, 0x88, 0xe8, 0x69, 0x77, 0x04, 0x96, 0x75, 0xa9, 0x9d, 0x92, 0x54, 0x2b, 0xab,
	0xdc, 0x47, 0x35, 0x9f, 0x2c, 0xf1, 0xc9, 0x2f, 0x3e, 0xa9, 0xf9, 0x9d, 0x0d, 0xae, 0xb8, 0x2a,
	0x15, 0xb4, 0x38, 0x55, 0xe2, 0x8e, 0xcf, 0x95, 0xe2, 0x13, 0xa0, 0x65, 0x35, 0x3a, 0x89, 0xa9,
	0x15, 0x12, 0x8c, 0x65, 0x32, 0xad, 0x09, 0xf7, 0x42, 0x65, 0xa4, 0x32, 0x54, 0x1a, 0x4e, 0x4f,
	0xbb, 0xc5, 0xaf, 0x6e, 0xdc, 0x61, 0x52, 0x24, 0x8a, 0x96, 0xdf, 0x0a, 0xda, 0xfa, 0xbc, 0x82,
	0x5b, 0x03, 0xc3, 0x0f, 0x95, 0xb1, 0x87, 0xc5, 0x7c, 0xf7, 0x01, 0x6e, 0xc5, 0x5a, 0xc9, 0x21,
	0x8b, 0x22, 0x0d, 0xc6, 0xb4, 0xd1, 0x26, 0xda, 0x6e, 0x06, 0xeb, 0x05, 0xf6, 0xb2, 0x82, 0xdc,
	0xc7, 0xb8, 0x29, 0x99, 0x3e, 0x06, 0x3b, 0x14, 0x51, 0x7b, 0xa5, 0xe8, 0xf7, 0x5b, 0xf9, 0xcc,
	0x6f, 0x0c, 0x4a, 0x70, 0x7f, 0x2f, 0x68, 0x54, 0xed, 0xfd, 0xc8, 0x7d, 0x8a, 0x6f, 0x97, 0xb6,
	0xda, 0xab, 0x25, 0xed, 0xe1, 0xc5, 0xcc, 0x77, 0xbe, 0xcd, 0xfc, 0xfb, 0xd5, 0x86, 0x26, 0x3a,
	0x26, 0x42, 0x51, 0xc9, 0xec, 0x98, 0x1c, 0x00, 0x67, 0x61, 0xb6, 0x07, 0x61, 0x50, 0x29, 0xdc,
	0x17, 0x78, 0x0d, 0xa6, 0xa9, 0xd0, 0x59, 0xfb, 0xd6, 0x26, 0xda, 0x5e, 0xef, 0x75, 0x48, 0xe5,
	0x9b, 0x2c, 0x7c, 0x93, 0xd7, 0x0b, 0xdf, 0xfd, 0x46, 0x71, 0xef, 0xd9, 0x77, 0x1f, 0x05, 0xb5,
	0xe6, 0xd9, 0xc1, 0x87, 0x4f, 0xbe, 0xf3, 0xee, 0xfa, 0x7c, 0xe7, 0x0f, 0x37, 0x1f, 0xaf, 0xcf,
	0x77, 0x7a, 0xff, 0x94, 0x3d, 0x59, 0x0e, 0x65, 0xeb, 0x2e, 0xde, 0x58, 0xae, 0x03, 0x30, 0xa9,
	0x4a, 0x0c, 0xf4, 0xde, 0x23, 0xbc, 0x3a, 0x30, 0xdc, 0x7d, 0x8b, 0x9b, 0xbf, 0x13, 0xdc, 0x25,
	0xff, 0x3f, 0xa1, 0xf3, 0xfc, 0x06, 0xa2, 0xc5, 0x1a, 0xfd, 0xa3, 0xab, 0x1f, 0x1e, 0xfa, 0x92,
	0x7b, 0xe8, 0x22, 0xf7, 0xd0, 0x65, 0xee, 0xa1, 0xab, 0xdc, 0x43, 0x67, 0x73, 0xcf, 0xb9, 0x9c,
	0x7b, 0xce, 0xd7, 0xb9, 0xe7, 0xbc, 0x79, 0xc5, 0x85, 0x1d, 0x9f, 0x8c, 0x48, 0xa8, 0x24, 0xad,
	0x07, 0x3d, 0x89, 0x45, 0xc2, 0x92, 0x10, 0x16, 0xf5, 0x70, 0x02, 0x49, 0x24, 0x12, 0x4e, 0xa7,
	0x7f, 0x79, 0xc5, 0x36, 0x4b, 0xc1, 0x8c, 0xd6, 0xca, 0xf8, 0x77, 0x7f, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0xa8, 0x73, 0xc9, 0xf3, 0x02, 0x00, 0x00,
}

func (this *MsgPostPrice) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MsgPostPrice)
	if !ok {
		that2, ok := that.(MsgPostPrice)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MsgPostPrice")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MsgPostPrice but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MsgPostPrice but is not nil && this == nil")
	}
	if this.FromAddress != that1.FromAddress {
		return fmt.Errorf("FromAddress this(%v) Not Equal that(%v)", this.FromAddress, that1.FromAddress)
	}
	if this.MarketID != that1.MarketID {
		return fmt.Errorf("MarketID this(%v) Not Equal that(%v)", this.MarketID, that1.MarketID)
	}
	if !this.Price.Equal(that1.Price) {
		return fmt.Errorf("Price this(%v) Not Equal that(%v)", this.Price, that1.Price)
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return fmt.Errorf("Expiry this(%v) Not Equal that(%v)", this.Expiry, that1.Expiry)
	}
	return nil
}
func (this *MsgPostPrice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPostPrice)
	if !ok {
		that2, ok := that.(MsgPostPrice)
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
	if this.FromAddress != that1.FromAddress {
		return false
	}
	if this.MarketID != that1.MarketID {
		return false
	}
	if !this.Price.Equal(that1.Price) {
		return false
	}
	if !this.Expiry.Equal(that1.Expiry) {
		return false
	}
	return true
}
func (this *MsgPostPriceResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*MsgPostPriceResponse)
	if !ok {
		that2, ok := that.(MsgPostPriceResponse)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *MsgPostPriceResponse")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *MsgPostPriceResponse but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *MsgPostPriceResponse but is not nil && this == nil")
	}
	return nil
}
func (this *MsgPostPriceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPostPriceResponse)
	if !ok {
		that2, ok := that.(MsgPostPriceResponse)
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
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PostPrice defines a method for creating a new post price
	PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PostPrice(ctx context.Context, in *MsgPostPrice, opts ...grpc.CallOption) (*MsgPostPriceResponse, error) {
	out := new(MsgPostPriceResponse)
	err := c.cc.Invoke(ctx, "/joltify.third_party.pricefeed.v1beta1.Msg/PostPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PostPrice defines a method for creating a new post price
	PostPrice(context.Context, *MsgPostPrice) (*MsgPostPriceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PostPrice(ctx context.Context, req *MsgPostPrice) (*MsgPostPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPrice not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PostPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPostPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PostPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/joltify.third_party.pricefeed.v1beta1.Msg/PostPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PostPrice(ctx, req.(*MsgPostPrice))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "joltify.third_party.pricefeed.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostPrice",
			Handler:    _Msg_PostPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "joltify/third_party/pricefeed/v1beta1/tx.proto",
}

func (m *MsgPostPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTx(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.MarketID) > 0 {
		i -= len(m.MarketID)
		copy(dAtA[i:], m.MarketID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MarketID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPostPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPostPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPostPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPostPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MarketID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Price.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPostPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPostPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPostPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarketID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPostPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPostPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPostPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
