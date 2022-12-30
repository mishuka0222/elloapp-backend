package mysql_dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*CountriesDAO
}

func NewMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:           db,
		CountriesDAO: NewCountriesDAO(db),
	}
}
