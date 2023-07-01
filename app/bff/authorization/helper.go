package authorization_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg/code"
)

type (
	Config               = config.Config
	AuthorizationService = service.Service
)

func New(c Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, code2, plugin))
}
