package configuration_helper

import (
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/config"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
