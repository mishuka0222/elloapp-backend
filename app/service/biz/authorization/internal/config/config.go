package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql             sqlx.Config
	AuthsessionClient zrpc.RpcClientConf
}
