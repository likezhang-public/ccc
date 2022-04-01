// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Package1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// Test2Client is the client API for Test2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Test2Client interface {
	SayHello2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type test2Client struct {
	cc grpc.ClientConnInterface
}

func NewTest2Client(cc grpc.ClientConnInterface) Test2Client {
	return &test2Client{cc}
}

func (c *test2Client) SayHello2(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/ModuleTest.Package1.Test2/SayHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Test2Server is the server API for Test2 service.
// All implementations must embed UnimplementedTest2Server
// for forward compatibility
type Test2Server interface {
	SayHello2(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTest2Server()
}

// UnimplementedTest2Server must be embedded to have forward compatible implementations.
type UnimplementedTest2Server struct {
}

func (UnimplementedTest2Server) SayHello2(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello2 not implemented")
}
func (UnimplementedTest2Server) mustEmbedUnimplementedTest2Server() {}

// UnsafeTest2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Test2Server will
// result in compilation errors.
type UnsafeTest2Server interface {
	mustEmbedUnimplementedTest2Server()
}

func RegisterTest2Server(s grpc.ServiceRegistrar, srv Test2Server) {
	s.RegisterService(&Test2_ServiceDesc, srv)
}

func _Test2_SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Test2Server).SayHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ModuleTest.Package1.Test2/SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Test2Server).SayHello2(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Test2_ServiceDesc is the grpc.ServiceDesc for Test2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ModuleTest.Package1.Test2",
	HandlerType: (*Test2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello2",
			Handler:    _Test2_SayHello2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Test2.proto",
}
