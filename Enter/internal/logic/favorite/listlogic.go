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
	listResp, err := l.svcCtx.FavoriteRpc.FavoriteList(l.ctx, &favoriterpc.FavoriteListReq{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil || listResp.StatusCode != 0 {
		resp = &types.FavoriteListResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	videoList := make([]types.Video, 0, listResp.Cnt)
	for _, video := range listResp.VideoList {
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

	resp = &types.FavoriteListResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		VideoList:  videoList,
	}
	return
}
