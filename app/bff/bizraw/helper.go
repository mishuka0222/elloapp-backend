package bizraw_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/internal/svc"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bizraw/service"
)

type (
	Config = config.Config
)

func New(c Config, register map[op_srv.ServiceID]op_srv.OperationServer) *service.Service {
	return service.New(svc.NewServiceContext(c, register))
}
