package relation

import (
	"DouYin/global"
	"DouYin/server/relation/rpc/relationrpc"
	"context"
	"strconv"

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

func (l *ActionLogic) Action(req *types.ActionReq) (resp *types.ActionResp, err error) {
	// string to int
	actonType, err := strconv.Atoi(req.ActionType)
	if err != nil {
		resp = &types.ActionResp{
			StatusCode: global.Error,
			StatusMsg:  "参数有误",
		}
		return
	}
	toUserId, err := strconv.Atoi(req.ToUserId)
	if err != nil {
		resp = &types.ActionResp{
			StatusCode: global.Error,
			StatusMsg:  "参数有误",
		}
		return
	}

	// action_type =1 关注 ，=2 取消关注，不能为其他数
	if actonType != 1 && actonType != 2 {
		resp = &types.ActionResp{
			StatusCode: global.Error,
			StatusMsg:  "参数有误",
		}
		return
	}

	// 交给 RelationRpc
	relationAction, err := l.svcCtx.RelationRpc.RelationAction(l.ctx, &relationrpc.ActionReq{
		Token:      req.Token,
		ToUserId:   int64(toUserId),
		ActionType: int32(actonType),
	})
	if err != nil {
		resp = &types.ActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}
	resp = &types.ActionResp{
		StatusCode: relationAction.StatusCode,
		StatusMsg:  relationAction.StatusMsg,
	}

	return
}
