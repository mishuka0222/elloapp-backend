package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	KV                kv.KvConf
	UserClient        zrpc.RpcClientConf
	AuthSessionClient zrpc.RpcClientConf
	SyncClient        *kafka.KafkaProducerConf
}
