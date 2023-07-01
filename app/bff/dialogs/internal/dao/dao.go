package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/config"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	channels_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	message_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/client"
	updates_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	dialog_client.DialogClient
	user_client.UserClient
	chat_client.ChatClient
	sync_client.SyncClient
	updates_client.UpdatesClient
	message_client.MessageClient
	channels_client.ChannelsClient
}

func New(c config.Config) *Dao {
	return &Dao{
		DialogClient:   dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		UpdatesClient:  updates_client.NewUpdatesClient(rpcx.GetCachedRpcClient(c.UpdatesClient)),
		UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		SyncClient:     sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
		MessageClient:  message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		ChatClient:     chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		ChannelsClient: channels_client.NewChannelsClient(rpcx.GetCachedRpcClient(c.ChannelsClient)),
	}
}
