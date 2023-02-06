package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/dao"
)

type ServiceContext struct {
	Config         config.Config
	Dao            *dao.Dao
	AccountService *account_helper.AccountService
}

func NewServiceContext(c config.Config, accountService *account_helper.AccountService) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Dao:            dao.New(c),
		AccountService: accountService,
	}
}
