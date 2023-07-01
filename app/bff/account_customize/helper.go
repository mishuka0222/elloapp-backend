package account_customize_helper

import (
	account_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/account_customize/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, accountService *account_helper.AccountService) *service.Service {
	return service.New(svc.NewServiceContext(c, accountService))
}
