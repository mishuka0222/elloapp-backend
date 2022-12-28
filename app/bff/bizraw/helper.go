package bizraw_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/bizraw/internal/svc"
	op_srv "github.com/teamgram/teamgram-server/app/bff/bizraw/service"
)

type (
	Config = config.Config
)

func New(c Config, register map[op_srv.ServiceID]op_srv.OperationServer) *service.Service {
	return service.New(svc.NewServiceContext(c, register))
}
