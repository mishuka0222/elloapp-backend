package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	PhonecallClient zrpc.RpcClientConf
	UserClient      zrpc.RpcClientConf
	SyncClient      *kafka.KafkaProducerConf
}
