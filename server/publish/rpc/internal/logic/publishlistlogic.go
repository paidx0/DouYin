package logic

import (
	"DouYin/global"
	"DouYin/models"
	"DouYin/server/publish/pb"
	"DouYin/server/publish/rpc/internal/svc"
	"DouYin/utils"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布列表
func (l *PublishListLogic) PublishList(in *__.PublishListReq) (*__.PublishListResp, error) {
	// 验证 Token 是否登录
	uc, err := utils.CheckToken(in.Token)
	if err != nil {
		return &__.PublishListResp{StatusCode: 1}, err
	}

	// 处理业务
	user := new(models.User)
	has, err := global.DBEngine.Where("uid = ?", in.UserId).Get(user)
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.PublishListResp{StatusCode: 1}, err
	}
	if !has {
		return &__.PublishListResp{StatusCode: 1}, errors.New("列表该用户不存在")
	}

	// 发布视频数
	cnt, err := global.DBEngine.Where("userpulishvideo.user_key = ? and userpulishvideo.deleted_at IS NULL", user.UserKey).Count(&models.UserPulishVideo{})
	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.PublishListResp{StatusCode: 1}, err
	}
	if cnt < 1 {
		return &__.PublishListResp{StatusCode: 1}, errors.New("该用户还未投稿过视频")
	}
	videoList := make([]*__.Video, cnt, 2*cnt)

	// 返回用户投稿列表：包括用户投稿过的视频信息，作者信息，及我是否关注作者和点赞视频
	err = global.DBEngine.Table("userpulishvideo").Where("userpulishvideo.user_key = ? and userpulishvideo.deleted_at IS NULL", user.UserKey).
		Join("LEFT", "video", "video.video_key = userpulishvideo.video_key").Where("video.deleted_at IS NULL").
		Join("LEFT", "user", "user.user_key = userpulishvideo.user_key").Where("user.deleted_at IS NULL").
		Join("LEFT", "userfocuson", "userfocuson.to_user_key = user.user_key").Where("(userfocuson.user_key = ? and userfocuson.deleted_at IS NULL) or 1", uc.Userkey).
		Join("LEFT", fmt.Sprintf(`(SELECT is_favorite,video_key from userlikevideo where user_key = "%s" AND userlikevideo.deleted_at IS NULL) myuserlikevideo`, uc.Userkey), "myuserlikevideo.video_key = userpulishvideo.video_key").
		Select("video.vid ,video.title ,video.play_url ,video.cover_url ,video.favorite_count ,video.comment_count ,user.uid ,user.username ,user.follow_count ,user.follower_count ,myuserlikevideo.is_favorite,userfocuson.is_follow").
		Find(&videoList)

	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.PublishListResp{StatusCode: 1}, err
	}

	return &__.PublishListResp{
		VideoList:  videoList[cnt:],
		Cnt:        cnt,
		StatusCode: 0,
	}, nil
}
