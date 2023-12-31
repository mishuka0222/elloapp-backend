package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/plugin"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config, plugin plugin.ChatPlugin) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c, plugin),
	}
}
