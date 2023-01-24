package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 同样把RPC加进来
	UserRpc     zrpc.RpcClientConf
	FavoriteRpc zrpc.RpcClientConf
	PublishRpc  zrpc.RpcClientConf
	CommentRpc  zrpc.RpcClientConf
}
