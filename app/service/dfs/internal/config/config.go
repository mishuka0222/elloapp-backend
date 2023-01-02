package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/rest"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/dfs/internal/minio_util"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MiniHttp rest.RestConf
	Cache    cache.CacheConf
	Minio    minio_util.MinioConfig
	IdGen    zrpc.RpcClientConf
	SSDB     kv.KvConf
}
