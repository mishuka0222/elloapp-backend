package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.ChatInviteParticipantsDAO
	*mysql_dao.ChatInvitesDAO
	*mysql_dao.ChatParticipantsDAO
	*mysql_dao.ChatsDAO
	*sqlx.CommonDAO
}

func newMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                        db,
		ChatInviteParticipantsDAO: mysql_dao.NewChatInviteParticipantsDAO(db),
		ChatInvitesDAO:            mysql_dao.NewChatInvitesDAO(db),
		ChatParticipantsDAO:       mysql_dao.NewChatParticipantsDAO(db),
		ChatsDAO:                  mysql_dao.NewChatsDAO(db),
		CommonDAO:                 sqlx.NewCommonDAO(db),
	}
}
