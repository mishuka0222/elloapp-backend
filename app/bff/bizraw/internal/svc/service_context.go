package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/internal/config"
	op_srv "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/bizraw/service"
)

type ServiceContext struct {
	Config config.Config
	OpSrv  *op_srv.Service
}

func NewServiceContext(c config.Config, register map[op_srv.ServiceID]op_srv.OperationServer) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OpSrv:  op_srv.New(register),
	}
}
