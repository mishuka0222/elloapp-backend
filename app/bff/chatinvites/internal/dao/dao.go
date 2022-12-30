package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/chatinvites/internal/config"
	msg_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/client"
	chat_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/client"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"
)

type Dao struct {
	user_client.UserClient
	chat_client.ChatClient
	msg_client.MsgClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient: user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
		ChatClient: chat_client.NewChatClient(rpcx.GetCachedRpcClient(c.ChatClient)),
		MsgClient:  msg_client.NewMsgClient(rpcx.GetCachedRpcClient(c.MsgClient)),
	}
}
