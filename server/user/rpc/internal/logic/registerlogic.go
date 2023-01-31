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
	"time"
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
	uuid := utils.UUID()
	uc := &models.User{
		Username: in.Username,
		Password: in.Password,
		UserKey:  uuid,
	}
	_, err := global.DBEngine.Insert(uc)
	if err != nil {
		global.ZAP.Error("数据库插入失败", zap.Error(err))
		return &__.Resp{StatusCode: 1}, err
	}

	global.DBEngine.Where("user_key = ?", uc.UserKey).Get(uc)
	token, err := utils.CreateToken(uc.Uid, uuid, uc.Username)
	if err != nil {
		global.ZAP.Error("token生成失败", zap.Error(err))
		return &__.Resp{StatusCode: 1}, err
	}
	// token 存入Redis
	err = global.REDIS.Set(l.ctx, "token", token, time.Second*time.Duration(global.CONFIG.JWT.ExpiresTime)).Err()
	if err != nil {
		global.ZAP.Error("redis保存token失败", zap.Error(err))
		return &__.Resp{StatusCode: 1}, err
	}

	return &__.Resp{
		UserID:     int64(uc.Uid),
		Token:      token,
		StatusCode: 0,
	}, nil
}
