package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/favorite/pb"
	"DouYin/server/favorite/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
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
	// 验证 Token 是否登录userClaim
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return nil, err
	}

	// 处理业务
	// 用户是否存在
	user := new(models.User)
	has, err := global.DBEngine.Where("uid = ?", in.UserId).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在")
	}

	// 喜爱视频数
	cnt, err := global.DBEngine.Where("userlikevideo.user_key = ? and userlikevideo.deleted_at IS NULL", user.UserKey).Count(&models.UserLikeVideo{})
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}
	if cnt < 1 {
		return nil, errors.New("点赞视频列表空")
	}
	videoList := make([]*__.Video, cnt, 2*cnt)

	// 返回用户喜爱列表：包括视频信息、发布者信息、及我是否点赞过这个视频、我是否关注这个作者
	err = global.DBEngine.Table("userlikevideo").Where("userlikevideo.user_key = ? and userlikevideo.deleted_at IS NULL", user.UserKey).
		Join("LEFT", "video", "video.video_key = userlikevideo.video_key").Where("video.deleted_at IS NULL").
		Join("LEFT", "userpulishvideo", "video.video_key = userpulishvideo.video_key").Where("userpulishvideo.deleted_at IS NULL").
		Join("LEFT", "user", "user.user_key = userpulishvideo.user_key").Where("user.deleted_at IS NULL").
		Join("LEFT", "userfocuson", "userfocuson.to_user_key = user.user_key").Where("userfocuson.user_key = ? and userfocuson.deleted_at IS NULL", uc.Userkey).
		Join("LEFT", fmt.Sprintf(`(SELECT is_favorite,video_key from userlikevideo where user_key = "%s" AND deleted_at IS NULL) myuserlikevideo`, uc.Userkey), "myuserlikevideo.video_key = userlikevideo.video_key").
		Select("video.vid ,video.title ,video.play_url ,video.cover_url ,video.favorite_count ,video.comment_count ,user.uid ,user.username ,user.follow_count ,user.follower_count ,myuserlikevideo.is_favorite,userfocuson.is_follow").
		Find(&videoList)

	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return nil, err
	}

	return &__.FavoriteListResp{
		VideoList: videoList[cnt:],
		Cnt:       cnt,
	}, nil
}
