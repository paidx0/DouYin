package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/user/pb"
	"DouYin/server/user/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录接口
func (l *LoginLogic) Login(in *__.Req) (*__.Resp, error) {
	user := new(models.User)
	has, err := global.DBEngine.Where("username = ? and password = ?", in.Username, in.Password).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := utils.CreateToken(user.Uid, user.UserKey, user.Username)
	if err != nil {
		global.ZAP.Error("token生成失败", zap.Error(err))
		return nil, err
	}

	return &__.Resp{
		UserID: int64(user.Uid),
		Token:  token,
	}, nil
}
