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
	// 验证 Token，拿到用户ID和Key
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.UserInfoResp{StatusCode: 1}, err
	}

	// 处理业务
	user := new(models.User)
	has, err := global.DBEngine.Where("uid = ?", in.UserId).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.UserInfoResp{StatusCode: 1}, err
	}
	if !has {
		return &__.UserInfoResp{StatusCode: 1}, errors.New("列表该用户不存在")
	}

	// 是否关注了该用户
	has, err = global.DBEngine.Where("user_key = ? and to_user_key = ?", uc.Userkey, user.UserKey).Get(&models.UserFocusOn{})
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.UserInfoResp{StatusCode: 1}, err
	}
	isfollow := false
	if has {
		isfollow = true
	}

	return &__.UserInfoResp{
		StatusCode: 0,
		UserInfo: &__.User{
			ID:            int64(user.Uid),
			Username:      user.Username,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      isfollow,
		},
	}, nil
}
