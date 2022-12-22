package feeds_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/service"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
