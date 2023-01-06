package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/config"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	idgen_client.IDGenClient2
}

func New(c config.Config) *Dao {
	return &Dao{
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
	}
}
