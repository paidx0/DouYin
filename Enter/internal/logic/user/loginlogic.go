package user

import (
	"DouYin/global"
	"DouYin/server/user/rpc/userrpc"
	"DouYin/utils"
	"context"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.Req) (resp *types.Resp, err error) {
	// 交给UserPrc
	login, err := l.svcCtx.UserRpc.Login(l.ctx, &userrpc.Req{
		Username: req.Username,
		Password: utils.Md5(req.Password),
	})

	if err != nil || login.StatusCode != 0 {
		resp = &types.Resp{
			StatusCode: global.Error,
			StatusMsg:  "登录失败",
		}
		return
	}

	resp = &types.Resp{
		StatusCode: global.Success,
		StatusMsg:  "登录成功",
		UserID:     login.UserID,
		Token:      login.Token,
	}
	return
}
