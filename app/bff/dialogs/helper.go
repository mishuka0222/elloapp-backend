package dialogs_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/dialogs/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.DialogsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
