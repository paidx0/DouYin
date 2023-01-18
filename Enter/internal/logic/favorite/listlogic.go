package favorite

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/server/favorite/rpc/favoriterpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	// 交给 FavoriteRPC处理
	listResp, err := l.svcCtx.FavoriteRpc.FavoriteList(l.ctx, &favoriterpc.FavoriteListReq{})
	if err != nil {
		return nil, err
	}

	videoList := make([]types.Video, len(listResp.VideoList))
	for _, video := range listResp.VideoList {
		videoList = append(videoList, types.Video{
			Id: video.ID,
			Author: types.User{
				Id:            video.Author.ID,
				Username:      video.Author.Username,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      video.Author.IsFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		})
	}

	resp = &types.FavoriteListResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		VideoList:  videoList,
	}

	return
}
