package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Config struct {
	zrpc.RpcServerConf
	InboxConsumer    kafka.KafkaConsumerConf
	Mysql            sqlx.Config
	KV               kv.KvConf
	IdgenClient      zrpc.RpcClientConf
	BizServiceClient zrpc.RpcClientConf
	InboxClient      *kafka.KafkaProducerConf
	SyncClient       *kafka.KafkaProducerConf
	BotSyncClient    *kafka.KafkaProducerConf `json:",optional"`
	MessageSharding  int                      `json:",default=1"`
}
