package relation

import (
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
	// todo: add your logic here and delete this line

	return
}
