package drafts_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/drafts/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.DraftsPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
