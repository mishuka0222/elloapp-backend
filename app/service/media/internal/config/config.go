package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql sqlx.Config
	Cache cache.CacheConf
	Dfs   zrpc.RpcClientConf
}
