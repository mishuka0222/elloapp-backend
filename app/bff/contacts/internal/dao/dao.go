package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/contacts/internal/config"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	username_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

type Dao struct {
	user_client.UserClient
	chat_client.ChatClient
	sync_client.SyncClient
	username_client.UsernameClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:     chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		UsernameClient: username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		SyncClient:     sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
	}
}
