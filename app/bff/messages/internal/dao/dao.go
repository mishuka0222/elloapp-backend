package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/messages/internal/config"
	msg_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/client"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	channels_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	message_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	username_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	msg_client.MsgClient
	user_client.UserClient
	ChatClient *chat_client.ChatClientHelper
	media_client.MediaClient
	username_client.UsernameClient
	message_client.MessageClient
	idgen_client.IDGenClient2
	dialog_client.DialogClient
	sync_client.SyncClient
	channels_client.ChannelsClient
}

func New(c config.Config) *Dao {
	return &Dao{
		MsgClient:      msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
		UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:     chat_client.NewChatClientHelper(rpcx.GetCachedRpcClient(c.ChatClient)),
		MediaClient:    media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		DialogClient:   dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
		IDGenClient2:   idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		MessageClient:  message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		UsernameClient: username_client.NewUsernameClient(rpcx.GetCachedRpcClient(c.UsernameClient)),
		SyncClient:     sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.SyncClient)),
		ChannelsClient: channels_client.NewChannelsClient(rpcx.GetCachedRpcClient(c.ChannelsClient)),
	}
}
