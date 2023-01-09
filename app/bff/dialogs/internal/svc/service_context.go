package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.DialogsPlugin
}

func NewServiceContext(c config.Config, plugin plugin.DialogsPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
