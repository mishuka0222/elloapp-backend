package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/tos/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
