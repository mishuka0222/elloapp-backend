package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthKeysDAO
	*mysql_dao.AuthUsersDAO
	*mysql_dao.AuthsDAO
	*mysql_dao.DevicesDAO
	*mysql_dao.AuthKeyInfosDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:              db,
		AuthKeysDAO:     mysql_dao.NewAuthKeysDAO(db),
		AuthUsersDAO:    mysql_dao.NewAuthUsersDAO(db),
		AuthsDAO:        mysql_dao.NewAuthsDAO(db),
		DevicesDAO:      mysql_dao.NewDevicesDAO(db),
		AuthKeyInfosDAO: mysql_dao.NewAuthKeyInfosDAO(db),
		CommonDAO:       sqlx.NewCommonDAO(db),
	}
}
