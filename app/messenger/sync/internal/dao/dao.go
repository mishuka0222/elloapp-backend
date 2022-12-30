package dao

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/internal/config"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/idgen/client"
	status_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/client"
	kafka "gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/mq"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Dao struct {
	*Mysql
	kv             kv.Store
	conf           *config.Config
	sessionServers map[string]*Session
	idgen_client.IDGenClient2
	status_client.StatusClient
	chat_client.ChatClient
	PushClient sync_client.SyncClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	d := &Dao{
		Mysql:          newMysqlDao(db),
		kv:             kv.NewStore(c.KV),
		conf:           &c,
		sessionServers: make(map[string]*Session),
		IDGenClient2:   idgen_client.NewIDGenClient2(zrpc.MustNewClient(c.IdgenClient)),
		StatusClient:   status_client.NewStatusClient(zrpc.MustNewClient(c.StatusClient)),
		ChatClient:     chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
	}
	if c.PushClient != nil {
		d.PushClient = sync_client.NewSyncMqClient(kafka.MustKafkaProducer(c.PushClient))
	}

	go d.watch(c.SessionClient)
	return d
}
