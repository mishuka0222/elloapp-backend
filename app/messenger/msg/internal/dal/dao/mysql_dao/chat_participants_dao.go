package mysql_dao

import (
	chat_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat"
)

type (
	ChatParticipantsDAO = chat_helper.ChatParticipantsDAO
)

var (
	NewChatParticipantsDAO = chat_helper.NewChatParticipantsDAO
)
