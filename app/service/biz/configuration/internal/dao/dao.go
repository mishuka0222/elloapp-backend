package dao

import (
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/dao/dao/mysql_dao"
)

type Dao struct {
	*mysql_dao.Mysql
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql: mysql_dao.NewMysqlDao(db),
	}
}
