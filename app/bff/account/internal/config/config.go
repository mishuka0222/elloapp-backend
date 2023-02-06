package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	UserClient        zrpc.RpcClientConf
	AccountClient     zrpc.RpcClientConf
	AuthsessionClient zrpc.RpcClientConf
	ChatClient        zrpc.RpcClientConf
	SyncClient        *kafka.KafkaProducerConf
}
