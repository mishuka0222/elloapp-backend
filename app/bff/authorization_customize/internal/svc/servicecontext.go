package svc

import (
	"github.com/teamgram/teamgram-server/app/bff/authorization"
	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/dao"
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
