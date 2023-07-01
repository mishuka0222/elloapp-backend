package miscellaneous_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/miscellaneous/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/miscellaneous/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/miscellaneous/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
