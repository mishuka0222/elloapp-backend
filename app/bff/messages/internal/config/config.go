package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf

	UserClient     zrpc.RpcClientConf
	ChatClient     zrpc.RpcClientConf
	MsgClient      zrpc.RpcClientConf
	DialogClient   zrpc.RpcClientConf
	IdgenClient    zrpc.RpcClientConf
	MessageClient  zrpc.RpcClientConf
	MediaClient    zrpc.RpcClientConf
	UsernameClient zrpc.RpcClientConf
	SyncClient     *kafka.KafkaProducerConf
}
