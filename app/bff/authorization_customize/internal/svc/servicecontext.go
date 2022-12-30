package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization_customize/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization_customize/internal/dao"
)

type ServiceContext struct {
	Config               config.Config
	Dao                  *dao.Dao
	AuthorizationService *authorization_helper.AuthorizationService
}

func NewServiceContext(c config.Config, authService *authorization_helper.AuthorizationService) *ServiceContext {
	return &ServiceContext{
		Config:               c,
		Dao:                  dao.New(c),
		AuthorizationService: authService,
	}
}
