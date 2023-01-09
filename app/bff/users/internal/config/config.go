package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserClient   zrpc.RpcClientConf
	ChatClient   zrpc.RpcClientConf
	DialogClient zrpc.RpcClientConf
}
