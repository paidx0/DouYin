package svc

import (
	"DouYin/Enter/internal/config"
	"DouYin/server/favorite/rpc/favoriterpc"
	"DouYin/server/user/rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// 同样把RPC加进来
	UserRpc     userrpc.UserRpc
	FavoriteRpc favoriterpc.FavoriteRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 同样把RPC加进来
		UserRpc:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpc)),
	}
}
