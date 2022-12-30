package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/files/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/files/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/files/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.FilesPlugin
}

func NewServiceContext(c config.Config, plugin plugin.FilesPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
