package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/premium/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
