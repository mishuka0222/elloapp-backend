package bizraw_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
