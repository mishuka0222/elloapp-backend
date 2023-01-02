package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/nsfw/internal/config"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	user_client.UserClient
}

func New(c config.Config) *Dao {
	return &Dao{
		UserClient: user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
	}
}
