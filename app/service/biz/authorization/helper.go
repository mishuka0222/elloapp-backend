package authorization_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/svc"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
)

type (
	Config = config.Config
)

func New(c Config, userService *user_helper.UserService) *service.Service {
	return service.New(svc.NewServiceContext(c, userService))
}
