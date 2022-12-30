package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.NotificationPlugin
}

func NewServiceContext(c config.Config, plugin plugin.NotificationPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
