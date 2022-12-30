package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	UserClient zrpc.RpcClientConf
	ChatClient zrpc.RpcClientConf
	SyncClient *kafka.KafkaProducerConf
}
