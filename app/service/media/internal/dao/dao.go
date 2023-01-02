package dao

import (
	"github.com/zeromicro/go-zero/zrpc"
	dfs_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Dao struct {
	*Mysql
	sqlc.CachedConn
	dfs_client.DfsClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:      newMysqlDao(db),
		CachedConn: sqlc.NewConn(db, c.Cache),
		DfsClient:  dfs_client.NewDfsClient(zrpc.MustNewClient(c.Dfs)),
	}
}
