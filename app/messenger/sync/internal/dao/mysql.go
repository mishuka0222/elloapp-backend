package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthSeqUpdatesDAO
	*mysql_dao.UserPtsUpdatesDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                db,
		AuthSeqUpdatesDAO: mysql_dao.NewAuthSeqUpdatesDAO(db),
		UserPtsUpdatesDAO: mysql_dao.NewUserPtsUpdatesDAO(db),
		CommonDAO:         sqlx.NewCommonDAO(db),
	}
}
