package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/authorization/internal/dao/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
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
