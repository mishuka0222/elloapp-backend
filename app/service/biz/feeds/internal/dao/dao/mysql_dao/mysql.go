package mysql_dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*UserFeedsDAO
}

func NewMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:           db,
		UserFeedsDAO: NewUserFeedsDAO(db),
	}
}
