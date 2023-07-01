package svc

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	inbox_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/plugin"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	channels_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config, plugin plugin.MsgPlugin) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)
	return &ServiceContext{
		Config: c,
		Dao: &dao.Dao{
			Mysql:          dao.NewMysqlDao(db, c.MessageSharding),
			KV:             kv.NewStore(c.KV),
			IDGenClient2:   idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
			UserClient:     user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
			InboxClient:    inbox_client.NewInboxMqClient(kafka.MustKafkaProducer(c.InboxClient)),
			ChatClient:     chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
			ChannelsClient: channels_client.NewChannelsClient(rpcx.GetCachedRpcClient(c.ChannelsClient)),
			SyncClient:     sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.SyncClient)),
			DialogClient:   dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
			MsgPlugin:      plugin,
		},
	}
}
