// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: pkg/proto/account.proto

package account

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	CheckIsAuth(ctx context.Context, in *CheckIsAuthRequest, opts ...grpc.CallOption) (*CheckIsAuthResponse, error)
	GetAddressByUserId(ctx context.Context, in *GetAddressByUserIdRequest, opts ...grpc.CallOption) (*GetAddressByUserIdResponse, error)
	GetUserByUserId(ctx context.Context, in *GetUserByUserIdRequest, opts ...grpc.CallOption) (*GetUserByUserIdResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CheckIsAuth(ctx context.Context, in *CheckIsAuthRequest, opts ...grpc.CallOption) (*CheckIsAuthResponse, error) {
	out := new(CheckIsAuthResponse)
	err := c.cc.Invoke(ctx, "/AccountService/CheckIsAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAddressByUserId(ctx context.Context, in *GetAddressByUserIdRequest, opts ...grpc.CallOption) (*GetAddressByUserIdResponse, error) {
	out := new(GetAddressByUserIdResponse)
	err := c.cc.Invoke(ctx, "/AccountService/GetAddressByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetUserByUserId(ctx context.Context, in *GetUserByUserIdRequest, opts ...grpc.CallOption) (*GetUserByUserIdResponse, error) {
	out := new(GetUserByUserIdResponse)
	err := c.cc.Invoke(ctx, "/AccountService/GetUserByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	CheckIsAuth(context.Context, *CheckIsAuthRequest) (*CheckIsAuthResponse, error)
	GetAddressByUserId(context.Context, *GetAddressByUserIdRequest) (*GetAddressByUserIdResponse, error)
	GetUserByUserId(context.Context, *GetUserByUserIdRequest) (*GetUserByUserIdResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CheckIsAuth(context.Context, *CheckIsAuthRequest) (*CheckIsAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIsAuth not implemented")
}
func (UnimplementedAccountServiceServer) GetAddressByUserId(context.Context, *GetAddressByUserIdRequest) (*GetAddressByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressByUserId not implemented")
}
func (UnimplementedAccountServiceServer) GetUserByUserId(context.Context, *GetUserByUserIdRequest) (*GetUserByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserId not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CheckIsAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIsAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckIsAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/CheckIsAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckIsAuth(ctx, req.(*CheckIsAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAddressByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAddressByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/GetAddressByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAddressByUserId(ctx, req.(*GetAddressByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetUserByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetUserByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AccountService/GetUserByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetUserByUserId(ctx, req.(*GetUserByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckIsAuth",
			Handler:    _AccountService_CheckIsAuth_Handler,
		},
		{
			MethodName: "GetAddressByUserId",
			Handler:    _AccountService_GetAddressByUserId_Handler,
		},
		{
			MethodName: "GetUserByUserId",
			Handler:    _AccountService_GetUserByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/account.proto",
}
