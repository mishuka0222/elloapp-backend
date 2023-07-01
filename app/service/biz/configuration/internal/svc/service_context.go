package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/configuration/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
