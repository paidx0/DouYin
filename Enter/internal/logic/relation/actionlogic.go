package relation

import (
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
	// todo: add your logic here and delete this line

	return
}
