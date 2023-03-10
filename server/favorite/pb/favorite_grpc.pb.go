// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: favorite.proto

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

// FavoriteRpcClient is the client API for FavoriteRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavoriteRpcClient interface {
	// 赞操作
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error)
	// 喜欢列表
	FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error)
}

type favoriteRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewFavoriteRpcClient(cc grpc.ClientConnInterface) FavoriteRpcClient {
	return &favoriteRpcClient{cc}
}

func (c *favoriteRpcClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionResp, error) {
	out := new(FavoriteActionResp)
	err := c.cc.Invoke(ctx, "/favorite.favoriteRpc/favoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favoriteRpcClient) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListResp, error) {
	out := new(FavoriteListResp)
	err := c.cc.Invoke(ctx, "/favorite.favoriteRpc/favoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavoriteRpcServer is the server API for FavoriteRpc service.
// All implementations must embed UnimplementedFavoriteRpcServer
// for forward compatibility
type FavoriteRpcServer interface {
	// 赞操作
	FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error)
	// 喜欢列表
	FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error)
	mustEmbedUnimplementedFavoriteRpcServer()
}

// UnimplementedFavoriteRpcServer must be embedded to have forward compatible implementations.
type UnimplementedFavoriteRpcServer struct {
}

func (UnimplementedFavoriteRpcServer) FavoriteAction(context.Context, *FavoriteActionReq) (*FavoriteActionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedFavoriteRpcServer) FavoriteList(context.Context, *FavoriteListReq) (*FavoriteListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteList not implemented")
}
func (UnimplementedFavoriteRpcServer) mustEmbedUnimplementedFavoriteRpcServer() {}

// UnsafeFavoriteRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavoriteRpcServer will
// result in compilation errors.
type UnsafeFavoriteRpcServer interface {
	mustEmbedUnimplementedFavoriteRpcServer()
}

func RegisterFavoriteRpcServer(s grpc.ServiceRegistrar, srv FavoriteRpcServer) {
	s.RegisterService(&FavoriteRpc_ServiceDesc, srv)
}

func _FavoriteRpc_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.favoriteRpc/favoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavoriteRpc_FavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavoriteRpcServer).FavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/favorite.favoriteRpc/favoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavoriteRpcServer).FavoriteList(ctx, req.(*FavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// FavoriteRpc_ServiceDesc is the grpc.ServiceDesc for FavoriteRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavoriteRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "favorite.favoriteRpc",
	HandlerType: (*FavoriteRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "favoriteAction",
			Handler:    _FavoriteRpc_FavoriteAction_Handler,
		},
		{
			MethodName: "favoriteList",
			Handler:    _FavoriteRpc_FavoriteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "favorite.proto",
}
