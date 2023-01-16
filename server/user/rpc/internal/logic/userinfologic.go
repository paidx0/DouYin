package logic

import (
	"DouYin/server/user/pb"
	"DouYin/server/user/rpc/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserinfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户信息
func (l *UserinfoLogic) Userinfo(in *__.UserInfoReq) (*__.UserInfoResp, error) {
	// todo: add your logic here and delete this line

	return &__.UserInfoResp{}, nil
}
