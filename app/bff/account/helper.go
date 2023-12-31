package account_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account/internal/svc"
)

type (
	Config         = config.Config
	AccountService = service.Service
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
