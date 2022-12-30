package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Cache           cache.CacheConf
	AuthSession     zrpc.RpcClientConf
	StatusClient    zrpc.RpcClientConf
	GatewayClient   zrpc.RpcClientConf
	BFFProxyClients BFFProxyClients
}

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}

type BFFProxyClients struct {
	Clients []zrpc.RpcClientConf
	IDMap   map[string]string
}
