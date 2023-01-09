package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UpdatesClient     zrpc.RpcClientConf
	UserClient        zrpc.RpcClientConf
	ChatClient        zrpc.RpcClientConf
	AuthsessionClient zrpc.RpcClientConf
}
