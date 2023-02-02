package accounts_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/accounts/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
