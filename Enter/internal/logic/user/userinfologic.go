package user

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"go.uber.org/zap"

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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq, claim *utils.UserClaim) (resp *types.UserInfoResp, err error) {
	user := new(models.User)
	has, err := global.DBEngine.Where("id = ?", req.UserId).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return
	}
	if !has {
		resp = &types.UserInfoResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}
	has, err = global.DBEngine.Where("userkey = ? and touserkey = ?", claim.UserKey, user.UserKey).Get(&models.UserFocusOn{})
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return
	}
	isfollow := false
	if has {
		isfollow = true
	}

	resp = &types.UserInfoResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		UserInfo: types.User{
			Id:            int64(user.Id),
			Username:      user.Username,
			FollowCount:   int64(user.FollowCount),
			FollowerCount: int64(user.FollowerCount),
			IsFollow:      isfollow,
		},
	}
	return
}
