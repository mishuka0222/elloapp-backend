package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization_customize/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"

	authorization_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/authorization/client"
)

type Dao struct {
	AuthorizationClient authorization_client.RPCAuthorizationClient
}

func New(c config.Config) *Dao {

	return &Dao{
		AuthorizationClient: authorization_client.NewAuthorizationClient(rpcx.GetCachedRpcClient(c.AuthorizationClient)),
	}
}
