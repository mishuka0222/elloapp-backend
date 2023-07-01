package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/updates/internal/config"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/client"
	updates_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	updates_client.UpdatesClient
	user_client.UserClient
	chat_client.ChatClient
	authsession_client.AuthsessionClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UpdatesClient:     updates_client.NewUpdatesClient(rpcx.GetCachedRpcClient(c.UpdatesClient)),
		UserClient:        user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient:        chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		AuthsessionClient: authsession_client.NewAuthsessionClient(rpcx.GetCachedRpcClient(c.AuthsessionClient)),
	}
}
