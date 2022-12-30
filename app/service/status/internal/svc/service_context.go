package svc

import (
	"github.com/zeromicro/go-zero/core/stores/kv"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/internal/config"
)

type ServiceContext struct {
	Config config.Config
	KV     kv.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		KV:     kv.NewStore(c.Status),
	}
}
