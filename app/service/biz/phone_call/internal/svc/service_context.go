package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Dao:    dao.New(c),
	}
}
