package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/config"
	account_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/account/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	AccountClient account_client.AccountClient
}

func New(c config.Config) *Dao {
	return &Dao{
		AccountClient: account_client.NewAccountClient(rpcx.GetCachedRpcClient(c.AccountClient)),
	}
}
