package dao

import (
	inbox_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/plugin"
	sync_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/client"
	// channel_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channel/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	*Mysql
	KV kv.Store
	idgen_client.IDGenClient2
	user_client.UserClient
	chat_client.ChatClient
	inbox_client.InboxClient
	SyncClient    sync_client.SyncClient
	BotSyncClient sync_client.SyncClient
	dialog_client.DialogClient
	plugin.MsgPlugin
}
