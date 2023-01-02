package dao

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/qrcode/internal/config"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

type Dao struct {
	kv kv.Store
	user_client.UserClient
	authsession_client.AuthsessionClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		kv:                kv.NewStore(c.KV),
		UserClient:        user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		AuthsessionClient: authsession_client.NewAuthsessionClient(rpcx.GetCachedRpcClient(c.AuthSessionClient)),
		SyncClient:        sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
