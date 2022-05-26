// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: apps/application/pb/application.proto

package application

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	ValidateClientCredential(ctx context.Context, in *ValidateClientCredentialRequest, opts ...grpc.CallOption) (*Application, error)
	CreateService(ctx context.Context, in *CreateMicroRequest, opts ...grpc.CallOption) (*Application, error)
	QueryService(ctx context.Context, in *QueryMicroRequest, opts ...grpc.CallOption) (*Set, error)
	DescribeService(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Application, error)
	DeleteService(ctx context.Context, in *DeleteMicroRequest, opts ...grpc.CallOption) (*Application, error)
	RefreshServiceClientSecret(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Application, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ValidateClientCredential(ctx context.Context, in *ValidateClientCredentialRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/ValidateClientCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) CreateService(ctx context.Context, in *CreateMicroRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/CreateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryService(ctx context.Context, in *QueryMicroRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/QueryService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeService(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/DescribeService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteService(ctx context.Context, in *DeleteMicroRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RefreshServiceClientSecret(ctx context.Context, in *DescribeMicroRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := c.cc.Invoke(ctx, "/infraboard.mcenter.application.Service/RefreshServiceClientSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	ValidateClientCredential(context.Context, *ValidateClientCredentialRequest) (*Application, error)
	CreateService(context.Context, *CreateMicroRequest) (*Application, error)
	QueryService(context.Context, *QueryMicroRequest) (*Set, error)
	DescribeService(context.Context, *DescribeMicroRequest) (*Application, error)
	DeleteService(context.Context, *DeleteMicroRequest) (*Application, error)
	RefreshServiceClientSecret(context.Context, *DescribeMicroRequest) (*Application, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) ValidateClientCredential(context.Context, *ValidateClientCredentialRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateClientCredential not implemented")
}
func (UnimplementedServiceServer) CreateService(context.Context, *CreateMicroRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedServiceServer) QueryService(context.Context, *QueryMicroRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryService not implemented")
}
func (UnimplementedServiceServer) DescribeService(context.Context, *DescribeMicroRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeService not implemented")
}
func (UnimplementedServiceServer) DeleteService(context.Context, *DeleteMicroRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedServiceServer) RefreshServiceClientSecret(context.Context, *DescribeMicroRequest) (*Application, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshServiceClientSecret not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_ValidateClientCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateClientCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ValidateClientCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/ValidateClientCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ValidateClientCredential(ctx, req.(*ValidateClientCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/CreateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreateService(ctx, req.(*CreateMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/QueryService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryService(ctx, req.(*QueryMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/DescribeService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeService(ctx, req.(*DescribeMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteService(ctx, req.(*DeleteMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RefreshServiceClientSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMicroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RefreshServiceClientSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mcenter.application.Service/RefreshServiceClientSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RefreshServiceClientSecret(ctx, req.(*DescribeMicroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mcenter.application.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateClientCredential",
			Handler:    _Service_ValidateClientCredential_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _Service_CreateService_Handler,
		},
		{
			MethodName: "QueryService",
			Handler:    _Service_QueryService_Handler,
		},
		{
			MethodName: "DescribeService",
			Handler:    _Service_DescribeService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _Service_DeleteService_Handler,
		},
		{
			MethodName: "RefreshServiceClientSecret",
			Handler:    _Service_RefreshServiceClientSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/application/pb/application.proto",
}
