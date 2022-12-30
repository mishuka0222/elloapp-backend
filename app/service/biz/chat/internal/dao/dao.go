package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/net/rpcx"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/config"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/media/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
	media_client.MediaClient
	Plugin plugin.ChatPlugin
}

// New new a dao and return.
func New(c config.Config, plugin plugin.ChatPlugin) (dao *Dao) {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:       newMysqlDao(db),
		CachedConn:  sqlc.NewConn(db, c.Cache),
		MediaClient: media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		Plugin:      plugin,
	}
}
