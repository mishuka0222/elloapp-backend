package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.DialogFiltersDAO
	*mysql_dao.DialogsDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:               db,
		DialogFiltersDAO: mysql_dao.NewDialogFiltersDAO(db),
		DialogsDAO:       mysql_dao.NewDialogsDAO(db),
		CommonDAO:        sqlx.NewCommonDAO(db),
	}
}
