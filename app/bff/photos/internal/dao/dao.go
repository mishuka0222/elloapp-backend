package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/photos/internal/config"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

type Dao struct {
	media_client.MediaClient
	user_client.UserClient
	sync_client.SyncClient
}

func New(c config.Config) *Dao {
	return &Dao{
		MediaClient: media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		UserClient:  user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		SyncClient:  sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
