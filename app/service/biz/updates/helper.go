package updates_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
