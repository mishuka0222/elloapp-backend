package mysql_dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*UserAccountsDAO
}

func NewMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:              db,
		UserAccountsDAO: NewUserAccountsDAO(db),
	}
}
