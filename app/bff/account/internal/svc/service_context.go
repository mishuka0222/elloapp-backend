package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/account/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/account/internal/dao"
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
