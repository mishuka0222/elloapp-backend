package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/internal/dao"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
	// channel_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channel/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewMySQL(&c.Mysql)

	dao := &dao.Dao{
		Mysql:        dao.NewMysqlDao(db, c.MessageSharding),
		KV:           kv.NewStore(c.KV),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		UserClient:   user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:   chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		SyncClient:   sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.SyncClient)),
		DialogClient: dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
	}
	if c.BotSyncClient != nil {
		dao.BotSyncClient = sync_client.NewSyncMqClient(kafka.GetCachedMQClient(c.BotSyncClient))
	}

	return &ServiceContext{
		Config: c,
		Dao:    dao,
	}
}
