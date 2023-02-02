package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/config"
	accounts_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	accounts_client.AccountsClient
}

func New(c config.Config) *Dao {
	return &Dao{
		AccountsClient: accounts_client.NewAccountsClient(rpcx.GetCachedRpcClient(c.AccountsClient)),
	}
}
