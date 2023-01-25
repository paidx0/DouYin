package relation

import (
	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"
	"DouYin/global"
	"DouYin/server/relation/rpc/relationrpc"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerlistLogic {
	return &FollowerlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerlistLogic) Followerlist(req *types.ListReq) (resp *types.ListResp, err error) {
	// string to int
	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		resp = &types.ListResp{
			StatusCode: string(global.Error),
			StatusMsg:  "参数有误",
		}
		return
	}

	// 交给 RelationRpc
	followList, err := l.svcCtx.RelationRpc.FollowerList(l.ctx, &relationrpc.ListReq{
		UserID: int64(userID),
		Token:  req.Token,
	})
	if err != nil {
		resp = &types.ListResp{
			StatusCode: string(global.Error),
			StatusMsg:  "操作失败",
		}
		return
	}

	userList := make([]types.User, 0, len(followList.UserList))
	for _, user := range followList.UserList {
		userList = append(userList, types.User{
			Id:            user.Uid,
			Username:      user.Username,
			FollowCount:   *(user.FollowCount),
			FollowerCount: *(user.FollowerCount),
			IsFollow:      user.IsFollow,
		})
	}
	statusCode := strconv.Itoa(int(followList.StatusCode))
	resp = &types.ListResp{
		StatusCode: statusCode,
		StatusMsg:  followList.StatusMsg,
		UserList:   userList,
	}

	return
}
