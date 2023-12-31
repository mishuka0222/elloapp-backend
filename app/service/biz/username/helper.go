package username_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/svc"
)

type (
	Config          = config.Config
	UsernameService = service.Service
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
