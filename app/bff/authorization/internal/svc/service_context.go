package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/dao"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/internal/logic"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/authorization/plugin"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg/code"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	*logic.AuthLogic
	Plugin plugin.AuthorizationPlugin
}

func NewServiceContext(c config.Config, code2 code.VerifyCodeInterface, plugin plugin.AuthorizationPlugin) *ServiceContext {
	d := dao.New(c)
	if code2 == nil {
		code2 = code.NewVerifyCode(c.Code)
	}
	return &ServiceContext{
		Config:    c,
		Dao:       d,
		AuthLogic: logic.NewAuthSignLogic(d, code2),
		Plugin:    plugin,
	}
}
