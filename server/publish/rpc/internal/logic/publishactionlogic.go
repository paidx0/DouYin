package logic

import (
	"context"

	"DouYin/server/publish/pb"
	"DouYin/server/publish/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 视频投稿
func (l *PublishActionLogic) PublishAction(in *__.PublishActionReq) (*__.PublishActionResp, error) {
	// todo: add your logic here and delete this line

	return &__.PublishActionResp{}, nil
}
