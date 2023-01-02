package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.DraftsPlugin
}

func NewServiceContext(c config.Config, plugin plugin.DraftsPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
