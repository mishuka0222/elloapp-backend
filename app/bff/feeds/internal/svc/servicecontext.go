package svc

import (
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/dao"
	messages_helper "github.com/teamgram/teamgram-server/app/bff/messages"
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
