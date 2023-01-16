package relation

import (
	"context"

	"DouYin/Enter/internal/svc"
	"DouYin/Enter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendlistLogic {
	return &FriendlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendlistLogic) Friendlist(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
