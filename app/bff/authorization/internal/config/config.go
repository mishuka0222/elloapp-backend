package config

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code/conf"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
)

type Config struct {
	zrpc.RpcServerConf
	KV                        kv.KvConf
	Code                      *conf.SmsVerifyCodeConfig
	UserClient                zrpc.RpcClientConf
	UsernameClient            zrpc.RpcClientConf
	AuthsessionClient         zrpc.RpcClientConf
	ChatClient                zrpc.RpcClientConf
	StatusClient              zrpc.RpcClientConf
	SyncClient                *kafka.KafkaProducerConf
	MsgClient                 zrpc.RpcClientConf
	SignInServiceNotification []conf.MessageEntityConfig `json:",optional"`
	SignInMessage             []conf.MessageEntityConfig `json:",optional"`
}
