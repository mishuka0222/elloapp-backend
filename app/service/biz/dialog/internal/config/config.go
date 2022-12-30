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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql sqlx.Config
	Cache cache.CacheConf
}
