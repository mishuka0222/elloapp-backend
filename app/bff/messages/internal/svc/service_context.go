package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.MessagesPlugin
}

func NewServiceContext(c config.Config, plugin plugin.MessagesPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
