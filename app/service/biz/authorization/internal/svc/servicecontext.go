package svc

import (
	messages_helper "github.com/teamgram/teamgram-server/app/bff/messages"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/authorization/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	MessagesCore *messages_helper.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
