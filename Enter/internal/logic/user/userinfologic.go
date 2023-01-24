package user

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/server/user/rpc/userrpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 交给UserRpc处理
	userinfo, err := l.svcCtx.UserRpc.Userinfo(l.ctx, &userrpc.UserInfoReq{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil || userinfo.StatusCode != 0 {
		resp = &types.UserInfoResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	resp = &types.UserInfoResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		UserInfo: types.User{
			Id:            userinfo.UserInfo.ID,
			Username:      userinfo.UserInfo.Username,
			FollowCount:   userinfo.UserInfo.FollowCount,
			FollowerCount: userinfo.UserInfo.FollowerCount,
			IsFollow:      userinfo.UserInfo.IsFollow,
		},
	}
	return
}
