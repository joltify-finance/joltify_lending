// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/burnauction/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type MsgSubmitrequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Tokens  string `protobuf:"bytes,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
}

func (m *MsgSubmitrequest) Reset()         { *m = MsgSubmitrequest{} }
func (m *MsgSubmitrequest) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitrequest) ProtoMessage()    {}
func (*MsgSubmitrequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dbbd9d239a127a, []int{0}
}
func (m *MsgSubmitrequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitrequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitrequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitrequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitrequest.Merge(m, src)
}
func (m *MsgSubmitrequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitrequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitrequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitrequest proto.InternalMessageInfo

type MsgSubmitrequestResponse struct {
}

func (m *MsgSubmitrequestResponse) Reset()         { *m = MsgSubmitrequestResponse{} }
func (m *MsgSubmitrequestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitrequestResponse) ProtoMessage()    {}
func (*MsgSubmitrequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1dbbd9d239a127a, []int{1}
}
func (m *MsgSubmitrequestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitrequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitrequestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitrequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitrequestResponse.Merge(m, src)
}
func (m *MsgSubmitrequestResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitrequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitrequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitrequestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitrequest)(nil), "joltify.burnauction.MsgSubmitrequest")
	proto.RegisterType((*MsgSubmitrequestResponse)(nil), "joltify.burnauction.MsgSubmitrequestResponse")
}

func init() { proto.RegisterFile("joltify/burnauction/tx.proto", fileDescriptor_d1dbbd9d239a127a) }

var fileDescriptor_d1dbbd9d239a127a = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x4f, 0x2a, 0x2d, 0xca, 0x4b, 0x2c, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xca, 0xea, 0x21, 0xc9, 0x4a,
	0x49, 0x26, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0xc7, 0x83, 0x95, 0xe8, 0x43, 0x38, 0x10, 0xf5, 0x52,
	0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x10, 0x71, 0x10, 0x0b, 0x2a, 0x2a, 0x0e, 0x51, 0xa3, 0x9f, 0x5b,
	0x9c, 0xae, 0x5f, 0x66, 0x08, 0xa2, 0xa0, 0x12, 0x82, 0x89, 0xb9, 0x99, 0x79, 0xf9, 0xfa, 0x60,
	0x12, 0x22, 0xa4, 0xd4, 0xc7, 0xc8, 0x25, 0xe0, 0x5b, 0x9c, 0x1e, 0x5c, 0x9a, 0x94, 0x9b, 0x59,
	0x52, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x64, 0xc4, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58,
	0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0x71, 0x69, 0x8b, 0xae, 0x08, 0xd4,
	0x66, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0x20, 0x98,
	0x42, 0x21, 0x31, 0x2e, 0xb6, 0x92, 0xfc, 0xec, 0xd4, 0xbc, 0x62, 0x09, 0x26, 0x90, 0x96, 0x20,
	0x28, 0xcf, 0x4a, 0xaf, 0x63, 0x81, 0x3c, 0x43, 0xd3, 0xf3, 0x0d, 0x5a, 0x30, 0x95, 0x5d, 0xcf,
	0x37, 0x68, 0x49, 0xc2, 0xfc, 0x59, 0x5c, 0x50, 0xa6, 0xe7, 0x5b, 0x9c, 0xee, 0x0c, 0x92, 0x4a,
	0x0d, 0xc8, 0xcf, 0xcf, 0x51, 0x92, 0xe2, 0x92, 0x40, 0x77, 0x4f, 0x50, 0x6a, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x51, 0x31, 0x17, 0xb3, 0x6f, 0x71, 0xba, 0x50, 0x2a, 0x17, 0x2f, 0xaa, 0x7b,
	0x55, 0xf5, 0xb0, 0x84, 0x9b, 0x1e, 0xba, 0x31, 0x52, 0xba, 0x44, 0x29, 0x83, 0xd9, 0x26, 0xc5,
	0xda, 0xf0, 0x7c, 0x83, 0x16, 0xa3, 0x53, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31,
	0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb,
	0x31, 0x44, 0xd9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x4d,
	0xd6, 0x4d, 0xcb, 0xcc, 0x4b, 0xcc, 0x4b, 0x4e, 0x85, 0xf1, 0xe3, 0x73, 0x52, 0xf3, 0x52, 0x32,
	0xf3, 0xd2, 0xf5, 0x2b, 0x50, 0x23, 0xbc, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x05, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x02, 0xbd, 0x51, 0x14, 0x02, 0x00, 0x00,
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
	Submitrequest(ctx context.Context, in *MsgSubmitrequest, opts ...grpc.CallOption) (*MsgSubmitrequestResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Submitrequest(ctx context.Context, in *MsgSubmitrequest, opts ...grpc.CallOption) (*MsgSubmitrequestResponse, error) {
	out := new(MsgSubmitrequestResponse)
	err := c.cc.Invoke(ctx, "/joltify.burnauction.Msg/Submitrequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Submitrequest(context.Context, *MsgSubmitrequest) (*MsgSubmitrequestResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Submitrequest(ctx context.Context, req *MsgSubmitrequest) (*MsgSubmitrequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submitrequest not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Submitrequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Submitrequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/joltify.burnauction.Msg/Submitrequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Submitrequest(ctx, req.(*MsgSubmitrequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "joltify.burnauction.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submitrequest",
			Handler:    _Msg_Submitrequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "joltify/burnauction/tx.proto",
}

func (m *MsgSubmitrequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitrequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitrequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		i -= len(m.Tokens)
		copy(dAtA[i:], m.Tokens)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Tokens)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitrequestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitrequestResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitrequestResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSubmitrequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Tokens)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitrequestResponse) Size() (n int) {
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
func (m *MsgSubmitrequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitrequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitrequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
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
			m.Tokens = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSubmitrequestResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitrequestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitrequestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
