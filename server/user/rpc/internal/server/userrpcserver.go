// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"DouYin/server/user/pb"
	logic2 "DouYin/server/user/rpc/internal/logic"
	"DouYin/server/user/rpc/internal/svc"
	"context"
)

type UserRpcServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedUserRpcServer
}

func NewUserRpcServer(svcCtx *svc.ServiceContext) *UserRpcServer {
	return &UserRpcServer{
		svcCtx: svcCtx,
	}
}

// 用户注册接口
func (s *UserRpcServer) Register(ctx context.Context, in *__.Req) (*__.Resp, error) {
	l := logic2.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 用户登录接口
func (s *UserRpcServer) Login(ctx context.Context, in *__.Req) (*__.Resp, error) {
	l := logic2.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 用户信息
func (s *UserRpcServer) Userinfo(ctx context.Context, in *__.UserInfoReq) (*__.UserInfoResp, error) {
	l := logic2.NewUserinfoLogic(ctx, s.svcCtx)
	return l.Userinfo(in)
}