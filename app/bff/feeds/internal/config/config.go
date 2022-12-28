package config

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Mysql         sqlx.Config
	MessageClient zrpc.RpcClientConf
}
