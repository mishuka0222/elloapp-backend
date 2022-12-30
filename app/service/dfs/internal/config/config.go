/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

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
