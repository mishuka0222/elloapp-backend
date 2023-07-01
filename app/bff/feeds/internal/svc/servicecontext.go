package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/feeds/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/feeds/internal/dao"
	messages_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/messages"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	MessagesCore *messages_helper.Service
}

func NewServiceContext(c config.Config, messagesCore *messages_helper.Service) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Dao:          dao.New(c),
		MessagesCore: messagesCore,
	}
}
