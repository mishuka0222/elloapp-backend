package message_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/dal/dao/mysql_dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.MessagePlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}

type (
	MessagesDAO = mysql_dao.MessagesDAO
	MessagesDO  = dataobject.MessagesDO
	// ChannelMessagesDAO   = mysql_dao.ChannelMessagesDAO
	// ChannelMessagesDO    = dataobject.ChannelMessagesDO
	// ScheduledMessagesDAO = mysql_dao.ScheduledMessagesDAO
	// ScheduledMessagesDO  = dataobject.ScheduledMessagesDO
)

var (
	NewMessagesDAO = mysql_dao.NewMessagesDAO
	// NewChannelMessagesDAO   = mysql_dao.NewChannelMessagesDAO
	// NewScheduledMessagesDAO = mysql_dao.NewScheduledMessagesDAO
)
