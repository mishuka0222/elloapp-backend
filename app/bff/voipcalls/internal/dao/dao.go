package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls/internal/config"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	phonecall_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	phonecall_client.PhonecallClient
	user_client.UserClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		PhonecallClient: phonecall_client.NewPhonecallClient(rpcx.GetCachedRpcClient(c.PhonecallClient)),
		UserClient:      user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		SyncClient:      sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
