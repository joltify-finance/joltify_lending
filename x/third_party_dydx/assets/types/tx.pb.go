// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: joltify/third_party/dydxprotocol/assets/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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
	proto.RegisterFile("joltify/third_party/dydxprotocol/assets/tx.proto", fileDescriptor_f6a087697b8df44e)
}

var fileDescriptor_f6a087697b8df44e = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xc8, 0xca, 0xcf, 0x29,
	0xc9, 0x4c, 0xab, 0xd4, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4,
	0x4f, 0xa9, 0x4c, 0xa9, 0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0x2c, 0x2e,
	0x4e, 0x2d, 0x29, 0xd6, 0x2f, 0xa9, 0xd0, 0x03, 0x0b, 0x09, 0xa9, 0x43, 0x75, 0xe8, 0x21, 0xe9,
	0xd0, 0x43, 0xd6, 0xa1, 0x07, 0xd1, 0x21, 0x25, 0x9e, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f,
	0x5b, 0x9c, 0xae, 0x5f, 0x66, 0x08, 0xa2, 0x20, 0x26, 0x18, 0xf1, 0x70, 0x31, 0xfb, 0x16, 0xa7,
	0x4b, 0xb1, 0x36, 0x3c, 0xdf, 0xa0, 0xc5, 0xe8, 0x94, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x5e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x50, 0x4b, 0x75, 0xd3, 0x32, 0xf3, 0x12, 0xf3, 0x92, 0x53, 0x61, 0xfc, 0xf8, 0x9c, 0xd4, 0xbc,
	0x94, 0xcc, 0xbc, 0x74, 0xfd, 0x0a, 0x64, 0x0f, 0xc4, 0x83, 0x9c, 0x03, 0x77, 0x78, 0x65, 0x41,
	0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x17, 0xc7,
	0x55, 0xf0, 0x00, 0x00, 0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "joltify.third_party.dydxprotocol.assets.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "joltify/third_party/dydxprotocol/assets/tx.proto",
}