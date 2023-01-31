package comment

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/server/comment/rpc/commentrpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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

func (l *ListLogic) List(req *types.CommentListReq) (resp *types.CommentListResp, err error) {
	// 交给CommentRPC处理
	listResp, err := l.svcCtx.CommentRpc.CommentList(l.ctx, &commentrpc.CommentListReq{
		Token:   req.Token,
		VideoId: req.VideoId,
	})
	if err != nil || listResp.StatusCode != 0 {
		resp = &types.CommentListResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	commentList := make([]types.Comment, 0, listResp.Cnt)
	for _, comment := range listResp.CommentList {
		date, _ := time.Parse(global.DateTimeFmt, comment.CreatedAt)
		commentList = append(commentList, types.Comment{
			Id: comment.Cid,
			User: types.User{
				Id:            comment.Uid,
				Username:      comment.Username,
				FollowCount:   comment.FollowCount,
				FollowerCount: comment.FollowerCount,
				IsFollow:      comment.IsFollow,
			},
			Content:    comment.CommentText,
			CreateDate: date.Format("01-02"),
		})
	}

	resp = &types.CommentListResp{
		StatusCode:  global.Success,
		StatusMsg:   "操作成功",
		CommentList: commentList,
	}
	return
}
