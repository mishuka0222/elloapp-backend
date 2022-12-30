package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// Dao dao.
type Dao struct {
	*Mysql
	sqlc.CachedConn
}

// New new a dao and return.
func New(db *sqlx.DB, c sqlc.CachedConn) (dao *Dao) {
	dao = &Dao{
		Mysql:      newMysqlDao(db),
		CachedConn: c,
	}
	return
}
