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
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql sqlx.Config
	Cache cache.CacheConf
}
