// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/auction/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgPlaceBid represents a message used by bidders to place bids on auctions
type MsgPlaceBid struct {
	AuctionId uint64     `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	Bidder    string     `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgPlaceBid) Reset()         { *m = MsgPlaceBid{} }
func (m *MsgPlaceBid) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBid) ProtoMessage()    {}
func (*MsgPlaceBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_cec8c6dc3f3e47b1, []int{0}
}
func (m *MsgPlaceBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBid.Merge(m, src)
}
func (m *MsgPlaceBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBid proto.InternalMessageInfo

// MsgPlaceBidResponse defines the Msg/PlaceBid response type.
type MsgPlaceBidResponse struct {
}

func (m *MsgPlaceBidResponse) Reset()         { *m = MsgPlaceBidResponse{} }
func (m *MsgPlaceBidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBidResponse) ProtoMessage()    {}
func (*MsgPlaceBidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cec8c6dc3f3e47b1, []int{1}
}
func (m *MsgPlaceBidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBidResponse.Merge(m, src)
}
func (m *MsgPlaceBidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBidResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPlaceBid)(nil), "joltify.third_party.auction.v1beta1.MsgPlaceBid")
	proto.RegisterType((*MsgPlaceBidResponse)(nil), "joltify.third_party.auction.v1beta1.MsgPlaceBidResponse")
}

func init() {
	proto.RegisterFile("joltify/third_party/auction/v1beta1/tx.proto", fileDescriptor_cec8c6dc3f3e47b1)
}

var fileDescriptor_cec8c6dc3f3e47b1 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0x32, 0x31,
	0x1c, 0xc6, 0xaf, 0x2f, 0x84, 0x40, 0xd9, 0xee, 0x7d, 0x5f, 0x83, 0x24, 0x16, 0x82, 0x0b, 0x83,
	0xb6, 0x82, 0x83, 0xc6, 0x11, 0x13, 0x13, 0x07, 0x12, 0x73, 0xa3, 0x0b, 0xf6, 0xae, 0xe5, 0xa8,
	0x81, 0xfe, 0x2f, 0xd7, 0x42, 0x60, 0x70, 0xd6, 0xd1, 0x8f, 0xc0, 0xc7, 0x61, 0x64, 0x74, 0x32,
	0x06, 0x16, 0x3f, 0x86, 0x01, 0x0a, 0x61, 0x70, 0xd0, 0xad, 0xff, 0xf6, 0xf7, 0x7f, 0x9e, 0x27,
	0x7d, 0xf0, 0xc9, 0x23, 0xf4, 0xad, 0xea, 0x4e, 0x98, 0xed, 0xa9, 0x54, 0x74, 0x12, 0x9e, 0xda,
	0x09, 0xe3, 0xc3, 0xc8, 0x2a, 0xd0, 0x6c, 0xd4, 0x08, 0xa5, 0xe5, 0x0d, 0x66, 0xc7, 0x34, 0x49,
	0xc1, 0x82, 0x7f, 0xec, 0x68, 0xba, 0x47, 0x53, 0x47, 0x53, 0x47, 0x97, 0xff, 0xc5, 0x10, 0xc3,
	0x9a, 0x67, 0xab, 0xd3, 0x66, 0xb5, 0x4c, 0x22, 0x30, 0x03, 0x30, 0x2c, 0xe4, 0x46, 0xee, 0x84,
	0x23, 0x50, 0x7a, 0xf3, 0x5e, 0x7b, 0x46, 0xb8, 0xd8, 0x36, 0xf1, 0x5d, 0x9f, 0x47, 0xb2, 0xa5,
	0x84, 0x7f, 0x84, 0xb1, 0x13, 0xee, 0x28, 0x51, 0x42, 0x55, 0x54, 0xcf, 0x06, 0x05, 0x77, 0x73,
	0x2b, 0xfc, 0x03, 0x9c, 0x0b, 0x95, 0x10, 0x32, 0x2d, 0xfd, 0xa9, 0xa2, 0x7a, 0x21, 0x70, 0x93,
	0x7f, 0x81, 0x73, 0x7c, 0x00, 0x43, 0x6d, 0x4b, 0x99, 0x2a, 0xaa, 0x17, 0x9b, 0x87, 0x74, 0xe3,
	0x4b, 0x57, 0xbe, 0xdb, 0x88, 0xf4, 0x1a, 0x94, 0x6e, 0x65, 0x67, 0xef, 0x15, 0x2f, 0x70, 0xf8,
	0x55, 0xfe, 0x65, 0x5a, 0xf1, 0x3e, 0xa7, 0x15, 0xaf, 0xf6, 0x1f, 0xff, 0xdd, 0x0b, 0x12, 0x48,
	0x93, 0x80, 0x36, 0xb2, 0xf9, 0x84, 0x33, 0x6d, 0x13, 0xfb, 0x23, 0x9c, 0xdf, 0x65, 0x3c, 0xa3,
	0x3f, 0xf8, 0x0f, 0xba, 0x27, 0x56, 0xbe, 0xfc, 0xed, 0xc6, 0xd6, 0xbe, 0xf5, 0x30, 0x5b, 0x10,
	0x34, 0x5f, 0x10, 0xf4, 0xb1, 0x20, 0xe8, 0x75, 0x49, 0xbc, 0xf9, 0x92, 0x78, 0x6f, 0x4b, 0xe2,
	0xdd, 0xdf, 0xc4, 0xca, 0xf6, 0x86, 0x21, 0x8d, 0x60, 0xc0, 0x9c, 0xfa, 0x69, 0x57, 0x69, 0xae,
	0x23, 0xb9, 0x9d, 0x3b, 0x7d, 0xa9, 0x85, 0xd2, 0x31, 0x1b, 0x7f, 0xdb, 0xb3, 0x9d, 0x24, 0xd2,
	0x84, 0xb9, 0x75, 0x11, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x07, 0x6c, 0xa5, 0x13,
	0x02, 0x00, 0x00,
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
	// PlaceBid message type used by bidders to place bids on auctions
	PlaceBid(ctx context.Context, in *MsgPlaceBid, opts ...grpc.CallOption) (*MsgPlaceBidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlaceBid(ctx context.Context, in *MsgPlaceBid, opts ...grpc.CallOption) (*MsgPlaceBidResponse, error) {
	out := new(MsgPlaceBidResponse)
	err := c.cc.Invoke(ctx, "/joltify.third_party.auction.v1beta1.Msg/PlaceBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PlaceBid message type used by bidders to place bids on auctions
	PlaceBid(context.Context, *MsgPlaceBid) (*MsgPlaceBidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlaceBid(ctx context.Context, req *MsgPlaceBid) (*MsgPlaceBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/joltify.third_party.auction.v1beta1.Msg/PlaceBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceBid(ctx, req.(*MsgPlaceBid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "joltify.third_party.auction.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBid",
			Handler:    _Msg_PlaceBid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "joltify/third_party/auction/v1beta1/tx.proto",
}

func (m *MsgPlaceBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceBidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgPlaceBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovTx(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPlaceBidResponse) Size() (n int) {
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
func (m *MsgPlaceBid) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlaceBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
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
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgPlaceBidResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgPlaceBidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
