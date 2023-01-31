package user

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/user/rpc/userrpc"
	"DouYin/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Req) (resp *types.Resp, err error) {
	// 检查用户名或密码的格式
	err = utils.CheckUserLayout(req.Username, req.Password)
	if err != nil {
		resp = &types.Resp{
			StatusCode: global.Error,
			StatusMsg:  err.Error(),
		}
		return
	}

	// 检查用户名唯一性
	count, err := global.DBEngine.Where("username = ?", req.Username).Count(&models.User{})
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		resp = &types.Resp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}
	if count > 0 {
		resp = &types.Resp{
			StatusCode: global.Error,
			StatusMsg:  "用户名已存在",
		}
		return
	}

	// 交给UserRpc处理
	register, err := l.svcCtx.UserRpc.Register(l.ctx, &userrpc.Req{
		Username: req.Username,
		Password: utils.Md5(req.Password),
	})
	if err != nil || register.StatusCode != 0 {
		resp = &types.Resp{
			StatusCode: global.Error,
			StatusMsg:  "用户注册失败",
		}
		return
	}

	resp = &types.Resp{
		StatusCode: global.Success,
		StatusMsg:  "注册成功",
		UserID:     register.UserID,
		Token:      register.Token,
	}
	return
}
