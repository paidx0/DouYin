package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"

	"DouYin/server/favorite/pb"
	"DouYin/server/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 赞操作
func (l *FavoriteActionLogic) FavoriteAction(in *__.FavoriteActionReq) (*__.FavoriteActionResp, error) {
	// 检查 UserToken
	userClaim, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.FavoriteActionResp{StatusCode: 1}, err
	}

	// 处理业务
	// 点赞还是取消点赞，1-点赞，2-取消点赞
	video := new(models.Video)
	has, err := global.DBEngine.Where("id = ?", in.VideoID).Get(video)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.FavoriteActionResp{StatusCode: 1}, err
	}
	if !has {
		return &__.FavoriteActionResp{StatusCode: 1}, errors.New("视频不存在")
	}

	if in.ActionType == 1 {
		_, err := global.DBEngine.Insert(&models.UserLikeVideo{
			UserKey:  userClaim.Userkey,
			VideoKey: video.VideoKey,
		})
		if err != nil {
			global.ZAP.Error("数据库插入失败", zap.Error(err))
			return &__.FavoriteActionResp{StatusCode: 1}, err
		}
	} else if in.ActionType == 2 {
		global.DBEngine.Delete(&models.UserLikeVideo{
			UserKey:  userClaim.Userkey,
			VideoKey: video.VideoKey,
		})
	} else {
		return &__.FavoriteActionResp{StatusCode: 1}, nil
	}

	return &__.FavoriteActionResp{StatusCode: 0}, nil
}
