package mysql_dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

type Mysql struct {
	*sqlx.DB
	ChannelsDAO            *ChannelsDAO
	ChannelParticipantsDAO *ChannelParticipantsDAO
	UsersDAO               *UsersDAO
}

func NewMysqlDao(db *sqlx.DB) *Mysql {
	return &Mysql{
		DB:                     db,
		ChannelsDAO:            NewChannelsDAO(db),
		ChannelParticipantsDAO: NewChannelParticipantsDAO(db),
		UsersDAO:               NewUsersDAO(db),
	}
}
