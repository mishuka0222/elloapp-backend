package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DfsClient   zrpc.RpcClientConf
	UserClient  zrpc.RpcClientConf
	MediaClient zrpc.RpcClientConf
}
