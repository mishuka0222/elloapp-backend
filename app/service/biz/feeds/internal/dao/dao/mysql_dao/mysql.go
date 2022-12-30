package mysql_dao

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
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
