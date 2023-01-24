package comment

import (
	"DouYin/global"
	"DouYin/server/comment/rpc/commentrpc"
	"context"
	"time"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActionLogic {
	return &ActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActionLogic) Action(req *types.CommentActionReq) (resp *types.CommentActionResq, err error) {
	// 交给CommentRPC处理
	actionResp, err := l.svcCtx.CommentRpc.CommentAction(l.ctx, &commentrpc.CommentActionReq{
		Token:       req.Token,
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	})

	if actionResp.StatusCode == 1 {
		resp = &types.CommentActionResq{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
	} else if actionResp.StatusCode == 3 {
		resp = &types.CommentActionResq{
			StatusCode: global.Success,
			StatusMsg:  "操作成功",
		}
	} else {
		comment := types.Comment{
			Id: actionResp.Comment.Cid,
			User: types.User{
				Id:            actionResp.Comment.Uid,
				Username:      actionResp.Comment.Username,
				FollowCount:   actionResp.Comment.FollowCount,
				FollowerCount: actionResp.Comment.FollowerCount,
				IsFollow:      true,
			},
			Content:    req.CommentText,
			CreateDate: time.Now().Format("01-02"),
		}
		resp = &types.CommentActionResq{
			StatusCode: global.Success,
			StatusMsg:  "操作成功",
			Comment:    comment,
		}
	}
	return
}
