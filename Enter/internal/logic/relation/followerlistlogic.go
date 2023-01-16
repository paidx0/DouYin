package relation

import (
	"context"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
