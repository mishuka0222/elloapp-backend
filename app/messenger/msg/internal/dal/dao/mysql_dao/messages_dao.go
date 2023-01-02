package mysql_dao

import (
	message_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message"
)

type (
	MessagesDAO = message_helper.MessagesDAO
)

var (
	NewMessagesDAO = message_helper.NewMessagesDAO
)
