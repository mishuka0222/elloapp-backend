package msg_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/plugin"
)

type (
	Config = config.Config
)

func New(c Config, plugin plugin.MsgPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
