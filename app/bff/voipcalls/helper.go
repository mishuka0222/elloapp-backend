package voipcalls_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
