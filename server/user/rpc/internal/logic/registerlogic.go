package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/user/pb"
	"DouYin/server/user/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册接口
func (l *RegisterLogic) Register(in *__.Req) (*__.Resp, error) {
	// 插入用户数据
	uc := &models.User{
		Username: in.Username,
		Password: in.Password,
		UserKey:  utils.UUID(),
	}
	_, err := global.DBEngine.Insert(uc)
	if err != nil {
		global.ZAP.Error("数据库插入失败", zap.Error(err))
		return &__.Resp{}, err
	}
	token, err := utils.CreateToken(uc.Id, uc.UserKey, uc.Username)
	if err != nil {
		global.ZAP.Error("token生成失败", zap.Error(err))
		return &__.Resp{}, err
	}
	return &__.Resp{
		UserID: int64(uc.Id),
		Token:  token,
	}, nil
}
