// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/school.proto

package proto

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

// WebServerClient is the client API for WebServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebServerClient interface {
	GetStudentData(ctx context.Context, in *GetStudentDataReq, opts ...grpc.CallOption) (*GetStudentDataRes, error)
}

type webServerClient struct {
	cc grpc.ClientConnInterface
}

func NewWebServerClient(cc grpc.ClientConnInterface) WebServerClient {
	return &webServerClient{cc}
}

func (c *webServerClient) GetStudentData(ctx context.Context, in *GetStudentDataReq, opts ...grpc.CallOption) (*GetStudentDataRes, error) {
	out := new(GetStudentDataRes)
	err := c.cc.Invoke(ctx, "/WebServer/GetStudentData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServerServer is the server API for WebServer service.
// All implementations must embed UnimplementedWebServerServer
// for forward compatibility
type WebServerServer interface {
	GetStudentData(context.Context, *GetStudentDataReq) (*GetStudentDataRes, error)
	mustEmbedUnimplementedWebServerServer()
}

// UnimplementedWebServerServer must be embedded to have forward compatible implementations.
type UnimplementedWebServerServer struct {
}

func (UnimplementedWebServerServer) GetStudentData(context.Context, *GetStudentDataReq) (*GetStudentDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentData not implemented")
}
func (UnimplementedWebServerServer) mustEmbedUnimplementedWebServerServer() {}

// UnsafeWebServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebServerServer will
// result in compilation errors.
type UnsafeWebServerServer interface {
	mustEmbedUnimplementedWebServerServer()
}

func RegisterWebServerServer(s grpc.ServiceRegistrar, srv WebServerServer) {
	s.RegisterService(&WebServer_ServiceDesc, srv)
}

func _WebServer_GetStudentData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServerServer).GetStudentData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WebServer/GetStudentData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServerServer).GetStudentData(ctx, req.(*GetStudentDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WebServer_ServiceDesc is the grpc.ServiceDesc for WebServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WebServer",
	HandlerType: (*WebServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudentData",
			Handler:    _WebServer_GetStudentData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/school.proto",
}
