package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/updates/internal/config"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	*Mysql
	kv kv.Store
	idgen_client.IDGenClient2
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:        newMysqlDao(db),
		kv:           kv.NewStore(c.KV),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
	}
}
