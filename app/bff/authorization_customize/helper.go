package authorization_customize_helper

import (
	authorization_helper "github.com/teamgram/teamgram-server/app/bff/authorization"
	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/service"
	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, authService *authorization_helper.AuthorizationService) *service.Service {
	return service.New(svc.NewServiceContext(c, authService))
}
