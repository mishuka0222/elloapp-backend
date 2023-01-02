package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/contacts/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/contacts/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/contacts/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	Plugin plugin.ContactsPlugin
}

func NewServiceContext(c config.Config, plugin plugin.ContactsPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
		Plugin: plugin,
	}
}
