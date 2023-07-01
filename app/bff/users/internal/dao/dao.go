package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/users/internal/config"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	dialog_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	user_client.UserClient
	chat_client.ChatClient
	dialog_client.DialogClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient:   user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:   chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		DialogClient: dialog_client.NewDialogClient(rpcx.GetCachedRpcClient(c.DialogClient)),
	}
}
