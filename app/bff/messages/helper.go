package messages_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages/plugin"
)

type (
	Config  = config.Config
	Service = service.Service
)

func New(c Config, plugin plugin.MessagesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
