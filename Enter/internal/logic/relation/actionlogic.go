package relation

import (
	"DouYin/global"
	"DouYin/server/relation/rpc/relationrpc"
	"context"

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
	// 交给 RelationRpc处理
	actionResp, err := l.svcCtx.RelationRpc.RelationAction(l.ctx, &relationrpc.ActionReq{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil || actionResp.StatusCode != 0 {
		resp = &types.ActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
		return
	}

	resp = &types.ActionResp{
		StatusCode: global.Success,
		StatusMsg:  "操作成功",
	}
	return
}
