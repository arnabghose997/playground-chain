// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: playgroundchain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

func init() { proto.RegisterFile("playgroundchain/tx.proto", fileDescriptor_8bcf5da93a95c418) }

var fileDescriptor_8bcf5da93a95c418 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xc8, 0x49, 0xac,
	0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x4d, 0x2c, 0xca, 0x4b, 0x4c, 0x4a, 0xcf, 0xc8, 0x2f, 0x4e,
	0xb5, 0xb4, 0x34, 0xd7, 0x43, 0x53, 0x87, 0xce, 0x37, 0x62, 0xe5, 0x62, 0xf6, 0x2d, 0x4e, 0x77,
	0x8a, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xc7, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x14, 0xa3, 0xf5, 0x11, 0x46, 0xe9, 0x42, 0xdc,
	0x50, 0xa1, 0x8f, 0xe1, 0xaa, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xcb, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x2a, 0xc2, 0x07, 0xb2, 0xb5, 0x00, 0x00, 0x00,
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
	ServiceName: "arnabghose997.playgroundchain.playgroundchain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "playgroundchain/tx.proto",
}
