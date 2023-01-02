package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	UserClient    zrpc.RpcClientConf
	ChatClient    zrpc.RpcClientConf
	DialogClient  zrpc.RpcClientConf
	UpdatesClient zrpc.RpcClientConf
	SyncClient    *kafka.KafkaProducerConf
	MessageClient zrpc.RpcClientConf
	// ChannelClient zrpc.RpcClientConf
}
