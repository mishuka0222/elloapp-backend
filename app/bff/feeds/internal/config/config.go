package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	MessageClient zrpc.RpcClientConf
	FeedsClient   zrpc.RpcClientConf
}
