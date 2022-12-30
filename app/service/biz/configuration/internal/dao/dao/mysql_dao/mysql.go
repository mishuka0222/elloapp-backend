package mysql_dao

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
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
