package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.UsernameDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:          db,
		UsernameDAO: mysql_dao.NewUsernameDAO(db),
		CommonDAO:   sqlx.NewCommonDAO(db),
	}
}
