package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/usernames/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/usernames/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/usernames/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.UsernamesPlugin
}

func NewServiceContext(c config.Config, plugin plugin.UsernamesPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
