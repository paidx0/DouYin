package favorite

import (
	"DouYin/global"
	"DouYin/server/favorite/rpc/favoriterpc"
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

func (l *ActionLogic) Action(req *types.FavoriteActionReq) (resp *types.FavoriteActionResp, err error) {
	// 交给FavoriteRPC处理
	actionResp, err := l.svcCtx.FavoriteRpc.FavoriteAction(l.ctx, &favoriterpc.FavoriteActionReq{
		Token:      req.Token,
		VideoID:    req.VideoID,
		ActionType: req.ActionType,
	})

	if actionResp.StatusCode != 0 {
		resp = &types.FavoriteActionResp{
			StatusCode: global.Error,
			StatusMsg:  "操作失败",
		}
	} else {
		resp = &types.FavoriteActionResp{
			StatusCode: global.Success,
			StatusMsg:  "操作成功",
		}
	}
	return
}
