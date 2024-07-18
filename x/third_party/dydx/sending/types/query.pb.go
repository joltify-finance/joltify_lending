// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/sending/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() {
	proto.RegisterFile("joltify/third_party/dydxprotocol/sending/query.proto", fileDescriptor_e4cc79fa377a88ee)
}

var fileDescriptor_e4cc79fa377a88ee = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4,
	0x4f, 0xa9, 0x4c, 0xa9, 0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4e, 0xcd,
	0x4b, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x03, 0x0b, 0x0b, 0x69, 0x40,
	0x75, 0xe9, 0x21, 0xe9, 0xd2, 0x43, 0xd6, 0xa5, 0x07, 0xd5, 0x65, 0xc4, 0xce, 0xc5, 0x1a, 0x08,
	0xd2, 0xe8, 0x94, 0x7a, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xde, 0xe9,
	0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x50, 0x73, 0x75, 0xd3, 0x32, 0xf3,
	0x12, 0xf3, 0x92, 0x53, 0x61, 0xfc, 0xf8, 0x1c, 0xa8, 0x23, 0x2a, 0x30, 0xdc, 0x09, 0x77, 0x5f,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x96, 0x7c, 0x62, 0x3d, 0xd8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "joltify.third_party.dydxprotocol.sending.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "joltify/third_party/dydxprotocol/sending/query.proto",
}
