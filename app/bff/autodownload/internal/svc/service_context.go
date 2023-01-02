package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/autodownload/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
