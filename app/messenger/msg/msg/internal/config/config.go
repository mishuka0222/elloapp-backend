package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql           sqlx.Config
	KV              kv.KvConf
	IdgenClient     zrpc.RpcClientConf
	UserClient      zrpc.RpcClientConf
	InboxClient     *kafka.KafkaProducerConf
	ChatClient      zrpc.RpcClientConf
	SyncClient      *kafka.KafkaProducerConf
	ChannelClient   zrpc.RpcClientConf
	DialogClient    zrpc.RpcClientConf
	MessageSharding int `json:",default=1"`
}
