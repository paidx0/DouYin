package relation

import (
	"DouYin/global"
	"DouYin/server/relation/rpc/relationrpc"
	"context"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowlistLogic {
	return &FollowlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowlistLogic) Followlist(req *types.ListReq) (resp *types.ListResp, err error) {
	// 交给RelationRpc处理
	listResp, err := l.svcCtx.RelationRpc.FollowList(l.ctx, &relationrpc.ListReq{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil || listResp.StatusCode != 0 {
		resp = &types.ListResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	userList := make([]types.User, 0, listResp.Cnt)
	for _, user := range listResp.UserList {
		userList = append(userList, types.User{
			Id:            user.Uid,
			Username:      user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		})
	}

	resp = &types.ListResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
		UserList:   userList,
	}
	return
}
