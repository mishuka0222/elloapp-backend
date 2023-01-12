package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql             sqlx.Config
	Cache             cache.CacheConf
	KV                kv.KvConf
	MediaClient       zrpc.RpcClientConf
	IdgenClient       zrpc.RpcClientConf
	MessageSharding   int `json:",default=1"`
	UserClient        zrpc.RpcClientConf
	AuthSessionClient zrpc.RpcClientConf
}
