package logic

import (
	"DouYin/global"
	"DouYin/utils"
	"context"
	"go.uber.org/zap"
	"strconv"
	"time"

	"DouYin/server/feed/pb"
	"DouYin/server/feed/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *__.FeedReq) (*__.FeedResp, error) {
	// 返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	// 以时间戳为基准逆序返回最新的视频，并以最后一个的时间作为下一次的刷新基准
	var curtime string
	if in.LatestTime == "0" || in.LatestTime == "" {
		curtime = time.Now().Format(global.DateTimeFmt)
	} else {
		parseInt, _ := strconv.ParseInt(in.LatestTime, 10, 64)
		curtime = time.Unix(parseInt, 0).Format(global.DateTimeFmt)
	}

	// 如果登录了就会查询我是否关注作者和点赞视频情况
	// 没登录就直接返回默认视频列表
	var uc *utils.UserClaim
	var err error
	if in.Token == "" {
		uc = &utils.UserClaim{}
	} else {
		uc, err = utils.CheckToken(in.Token)
		if err != nil {
			return &__.FeedResp{StatusCode: 1}, err
		}
	}

	// 返回视频单次最多30个
	videoList := make([]*__.Video, 0, 30)

	global.DBEngine.ShowSQL(true)

	// 返回视频列表：包括视频信息、发布者信息、及我是否点赞过这个视频、我是否关注这个作者(前提是我已经登录)，每次返回30个
	err = global.DBEngine.Table("video").Where("video.updated_at < ?", curtime).Desc("video.updated_at").Limit(30).
		Join("LEFT", "userpulishvideo", "video.video_key = userpulishvideo.video_key").Where("userpulishvideo.deleted_at IS NULL").
		Join("LEFT", "user", "user.user_key = userpulishvideo.user_key").Where("user.deleted_at IS NULL").
		Join("LEFT", "userfocuson", "userfocuson.to_user_key = user.user_key").Where("(userfocuson.user_key = ?  or 1 ) and userfocuson.deleted_at IS NULL", uc.Userkey).
		Join("LEFT", "userlikevideo", "userlikevideo.video_key = video.video_key").Where("(userlikevideo.user_key = ? or 1 ) and userlikevideo.deleted_at IS NULL", uc.Userkey).
		Select("video.vid ,video.updated_at ,video.title ,video.play_url ,video.cover_url ,video.favorite_count ,video.comment_count ,user.uid ,user.username ,user.follow_count ,user.follower_count ,userlikevideo.is_favorite,userfocuson.is_follow").
		Find(&videoList)

	if err != nil {
		global.ZAP.Error("数据库查询失败", zap.Error(err))
		return &__.FeedResp{StatusCode: 1}, err
	}

	parse, _ := time.Parse(global.DateTimeFmt, videoList[len(videoList)-1].UpdatedAt)

	return &__.FeedResp{
		StatusCode: 0,
		VideoList:  videoList,
		Cnt:        30,
		NextTime:   parse.Unix() - 28800,
	}, nil
}
