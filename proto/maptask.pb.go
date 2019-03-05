// Code generated by protoc-gen-go.
// source: maptask.proto
// DO NOT EDIT!

/*
Package maptask is a generated protocol buffer package.

It is generated from these files:
	maptask.proto

It has these top-level messages:
	HeartbeatRequest
	HeartbeatResponse
*/
package maptask

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HeartbeatRequest struct {
}

func (m *HeartbeatRequest) Reset()                    { *m = HeartbeatRequest{} }
func (m *HeartbeatRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatRequest) ProtoMessage()               {}
func (*HeartbeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HeartbeatResponse struct {
}

func (m *HeartbeatResponse) Reset()                    { *m = HeartbeatResponse{} }
func (m *HeartbeatResponse) String() string            { return proto.CompactTextString(m) }
func (*HeartbeatResponse) ProtoMessage()               {}
func (*HeartbeatResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HeartbeatRequest)(nil), "maptask.HeartbeatRequest")
	proto.RegisterType((*HeartbeatResponse)(nil), "maptask.HeartbeatResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for MapTask service

type MapTaskClient interface {
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type mapTaskClient struct {
	cc *grpc.ClientConn
}

func NewMapTaskClient(cc *grpc.ClientConn) MapTaskClient {
	return &mapTaskClient{cc}
}

func (c *mapTaskClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := grpc.Invoke(ctx, "/maptask.MapTask/Heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MapTask service

type MapTaskServer interface {
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
}

func RegisterMapTaskServer(s *grpc.Server, srv MapTaskServer) {
	s.RegisterService(&_MapTask_serviceDesc, srv)
}

func _MapTask_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapTaskServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maptask.MapTask/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapTaskServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MapTask_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maptask.MapTask",
	HandlerType: (*MapTaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _MapTask_Heartbeat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2c, 0x28,
	0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x84, 0xb8,
	0x04, 0x3c, 0x52, 0x13, 0x8b, 0x4a, 0x92, 0x52, 0x13, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x94, 0x84, 0xb9, 0x04, 0x91, 0xc4, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x8d, 0x7c, 0xb9,
	0xd8, 0x7d, 0x13, 0x0b, 0x42, 0x80, 0x7a, 0x84, 0x9c, 0xb8, 0x38, 0xe1, 0xf2, 0x42, 0x92, 0x7a,
	0x30, 0x93, 0xd1, 0xcd, 0x91, 0x92, 0xc2, 0x26, 0x05, 0x31, 0x2e, 0x89, 0x0d, 0xec, 0x0e, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x26, 0x05, 0xed, 0xbe, 0x98, 0x00, 0x00, 0x00,
}