package notification_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/notification/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.NotificationPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
