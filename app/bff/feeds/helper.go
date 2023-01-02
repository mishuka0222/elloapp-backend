package feeds_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/feeds/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/feeds/internal/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/feeds/internal/svc"
	messages_helper "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/messages"
)

type (
	Config = config.Config
)

func New(c Config, messagesCore *messages_helper.Service) *service.Service {
	return service.New(svc.NewServiceContext(c, messagesCore))
}
