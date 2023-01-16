// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: user.proto

package __

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

// UserRpcClient is the client API for UserRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRpcClient interface {
	// 用户注册接口
	Register(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	// 用户登录接口
	Login(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
	// 用户信息
	Userinfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
}

type userRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRpcClient(cc grpc.ClientConnInterface) UserRpcClient {
	return &userRpcClient{cc}
}

func (c *userRpcClient) Register(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/user.userRpc/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) Login(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/user.userRpc/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRpcClient) Userinfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/user.userRpc/userinfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRpcServer is the server API for UserRpc service.
// All implementations must embed UnimplementedUserRpcServer
// for forward compatibility
type UserRpcServer interface {
	// 用户注册接口
	Register(context.Context, *Req) (*Resp, error)
	// 用户登录接口
	Login(context.Context, *Req) (*Resp, error)
	// 用户信息
	Userinfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	mustEmbedUnimplementedUserRpcServer()
}

// UnimplementedUserRpcServer must be embedded to have forward compatible implementations.
type UnimplementedUserRpcServer struct {
}

func (UnimplementedUserRpcServer) Register(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserRpcServer) Login(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserRpcServer) Userinfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Userinfo not implemented")
}
func (UnimplementedUserRpcServer) mustEmbedUnimplementedUserRpcServer() {}

// UnsafeUserRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRpcServer will
// result in compilation errors.
type UnsafeUserRpcServer interface {
	mustEmbedUnimplementedUserRpcServer()
}

func RegisterUserRpcServer(s grpc.ServiceRegistrar, srv UserRpcServer) {
	s.RegisterService(&UserRpc_ServiceDesc, srv)
}

func _UserRpc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Register(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Login(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRpc_Userinfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRpcServer).Userinfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.userRpc/userinfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRpcServer).Userinfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRpc_ServiceDesc is the grpc.ServiceDesc for UserRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.userRpc",
	HandlerType: (*UserRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _UserRpc_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _UserRpc_Login_Handler,
		},
		{
			MethodName: "userinfo",
			Handler:    _UserRpc_Userinfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
