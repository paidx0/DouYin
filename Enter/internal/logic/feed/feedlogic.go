package feed

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/server/feed/rpc/feedrpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	// 交给FeedRPC处理
	feedResp, err := l.svcCtx.FeedRpc.Feed(l.ctx, &feedrpc.FeedReq{
		Token:      req.Token,
		LatestTime: req.LatestTime,
	})
	if err != nil || feedResp.StatusCode != 0 {
		resp = &types.FeedResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	videoList := make([]types.Video, 0, feedResp.Cnt)
	for _, video := range feedResp.VideoList {
		videoList = append(videoList, types.Video{
			Id: video.Vid,
			Author: types.User{
				Id:            video.Uid,
				Username:      video.Username,
				FollowCount:   video.FollowCount,
				FollowerCount: video.FollowerCount,
				IsFollow:      video.IsFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	resp = &types.FeedResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		VideoList:  videoList,
		NextTime:   feedResp.NextTime,
	}
	return
}
