package usernames_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/usernames/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.UsernamesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
