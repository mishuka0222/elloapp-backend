package authorization_customize_helper

import (
	authorization_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization_customize/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, authService *authorization_helper.AuthorizationService) *service.Service {
	return service.New(svc.NewServiceContext(c, authService))
}
