package feeds_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/service"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/svc"
	messages_helper "github.com/teamgram/teamgram-server/app/bff/messages"
)

type (
	Config = config.Config
)

func New(c Config, messagesCore *messages_helper.Service) *service.Service {
	return service.New(svc.NewServiceContext(c, messagesCore))
}
