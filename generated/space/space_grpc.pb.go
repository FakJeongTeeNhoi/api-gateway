// Code generated by protoc-gen-go-grpcClient. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpcClient v1.5.1
// - protoc             v5.28.0
// source: proto/space.proto

package space

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpcClient package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SpaceService_CreateSpace_FullMethodName     = "/space.SpaceService/CreateSpace"
	SpaceService_GetAllSpace_FullMethodName     = "/space.SpaceService/GetAllSpace"
	SpaceService_EditSpaceDetail_FullMethodName = "/space.SpaceService/EditSpaceDetail"
	SpaceService_DeleteSpace_FullMethodName     = "/space.SpaceService/DeleteSpace"
	SpaceService_GetSpace_FullMethodName        = "/space.SpaceService/GetSpace"
)

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpaceServiceClient interface {
	CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error)
	GetAllSpace(ctx context.Context, in *GetAllSpaceRequest, opts ...grpc.CallOption) (*GetAllSpaceResponse, error)
	EditSpaceDetail(ctx context.Context, in *EditSpaceRequest, opts ...grpc.CallOption) (*SpaceServiceResponse, error)
	DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*SpaceServiceResponse, error)
	GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) CreateSpace(ctx context.Context, in *CreateSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_CreateSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetAllSpace(ctx context.Context, in *GetAllSpaceRequest, opts ...grpc.CallOption) (*GetAllSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetAllSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) EditSpaceDetail(ctx context.Context, in *EditSpaceRequest, opts ...grpc.CallOption) (*SpaceServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceServiceResponse)
	err := c.cc.Invoke(ctx, SpaceService_EditSpaceDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) DeleteSpace(ctx context.Context, in *DeleteSpaceRequest, opts ...grpc.CallOption) (*SpaceServiceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SpaceServiceResponse)
	err := c.cc.Invoke(ctx, SpaceService_DeleteSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceServiceClient) GetSpace(ctx context.Context, in *GetSpaceRequest, opts ...grpc.CallOption) (*GetSpaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSpaceResponse)
	err := c.cc.Invoke(ctx, SpaceService_GetSpace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
// All implementations must embed UnimplementedSpaceServiceServer
// for forward compatibility.
type SpaceServiceServer interface {
	CreateSpace(context.Context, *CreateSpaceRequest) (*GetSpaceResponse, error)
	GetAllSpace(context.Context, *GetAllSpaceRequest) (*GetAllSpaceResponse, error)
	EditSpaceDetail(context.Context, *EditSpaceRequest) (*SpaceServiceResponse, error)
	DeleteSpace(context.Context, *DeleteSpaceRequest) (*SpaceServiceResponse, error)
	GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error)
	mustEmbedUnimplementedSpaceServiceServer()
}

// UnimplementedSpaceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSpaceServiceServer struct{}

func (UnimplementedSpaceServiceServer) CreateSpace(context.Context, *CreateSpaceRequest) (*GetSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (UnimplementedSpaceServiceServer) GetAllSpace(context.Context, *GetAllSpaceRequest) (*GetAllSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSpace not implemented")
}
func (UnimplementedSpaceServiceServer) EditSpaceDetail(context.Context, *EditSpaceRequest) (*SpaceServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSpaceDetail not implemented")
}
func (UnimplementedSpaceServiceServer) DeleteSpace(context.Context, *DeleteSpaceRequest) (*SpaceServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (UnimplementedSpaceServiceServer) GetSpace(context.Context, *GetSpaceRequest) (*GetSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpace not implemented")
}
func (UnimplementedSpaceServiceServer) mustEmbedUnimplementedSpaceServiceServer() {}
func (UnimplementedSpaceServiceServer) testEmbeddedByValue()                      {}

// UnsafeSpaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpaceServiceServer will
// result in compilation errors.
type UnsafeSpaceServiceServer interface {
	mustEmbedUnimplementedSpaceServiceServer()
}

func RegisterSpaceServiceServer(s grpc.ServiceRegistrar, srv SpaceServiceServer) {
	// If the following call pancis, it indicates UnimplementedSpaceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SpaceService_ServiceDesc, srv)
}

func _SpaceService_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_CreateSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).CreateSpace(ctx, req.(*CreateSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetAllSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetAllSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetAllSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetAllSpace(ctx, req.(*GetAllSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_EditSpaceDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).EditSpaceDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_EditSpaceDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).EditSpaceDetail(ctx, req.(*EditSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_DeleteSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).DeleteSpace(ctx, req.(*DeleteSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceService_GetSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpaceService_GetSpace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetSpace(ctx, req.(*GetSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpaceService_ServiceDesc is the grpc.ServiceDesc for SpaceService service.
// It's only intended for direct use with grpcClient.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSpace",
			Handler:    _SpaceService_CreateSpace_Handler,
		},
		{
			MethodName: "GetAllSpace",
			Handler:    _SpaceService_GetAllSpace_Handler,
		},
		{
			MethodName: "EditSpaceDetail",
			Handler:    _SpaceService_EditSpaceDetail_Handler,
		},
		{
			MethodName: "DeleteSpace",
			Handler:    _SpaceService_DeleteSpace_Handler,
		},
		{
			MethodName: "GetSpace",
			Handler:    _SpaceService_GetSpace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/space.proto",
}