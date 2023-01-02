package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	KV                kv.KvConf
	Code              *conf.SmsVerifyCodeConfig
	BizServiceClient  zrpc.RpcClientConf
	AuthSessionClient zrpc.RpcClientConf
	MediaClient       zrpc.RpcClientConf
	IdgenClient       zrpc.RpcClientConf
	MsgClient         zrpc.RpcClientConf
	SyncClient        *kafka.KafkaProducerConf
	DfsClient         zrpc.RpcClientConf
	StatusClient      zrpc.RpcClientConf
	Mysql             sqlx.Config
}
