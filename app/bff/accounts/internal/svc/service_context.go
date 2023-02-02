package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
