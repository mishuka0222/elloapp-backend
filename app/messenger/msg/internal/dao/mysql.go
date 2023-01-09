package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.MessagesDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.HashTagsDAO
	*mysql_dao.DialogsDAO
	*sqlx.CommonDAO
}

func NewMysqlDao(db *sqlx.DB, shardingSize int) *Mysql {
	return &Mysql{
		DB:                  db,
		MessagesDAO:         mysql_dao.NewMessagesDAO(db, shardingSize),
		ChatParticipantsDAO: mysql_dao.NewChatParticipantsDAO(db),
		HashTagsDAO:         mysql_dao.NewHashTagsDAO(db),
		DialogsDAO:          mysql_dao.NewDialogsDAO(db),
		CommonDAO:           sqlx.NewCommonDAO(db),
	}
}
