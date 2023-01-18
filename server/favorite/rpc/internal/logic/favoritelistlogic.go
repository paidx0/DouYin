package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/utils"
	"context"
	"errors"
	"go.uber.org/zap"
	"time"

	"DouYin/server/favorite/pb"
	"DouYin/server/favorite/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 喜欢列表
func (l *FavoriteListLogic) FavoriteList(in *__.FavoriteListReq) (*__.FavoriteListResp, error) {
	// 验证 Token 是否登录
	_, err := utils.CheckToken(in.Token)
	if err != nil {
		return nil, err
	}

	// 处理业务
	// 用户是否存在
	global.DBEngine.ShowSQL(true)
	videoList := make([]*__.Video, 0)
	user := new(models.User)
	has, err := global.DBEngine.Where("id = ?", in.UserID).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	global.DBEngine.Table("userlikevideo").Select("userlikevideo.video_key").Where("user_key = ?", user.UserKey).
		Join("LEFT", "video", "video.video_key = userlikevideo.video_key").Where("share_basic.deleted_at = ? OR share_basic.deleted_at IS NULL", time.Time{}.Format(global.DateTimeFmt)).


	return &__.FavoriteListResp{
		VideoList: videoList,
	}, nil
}
