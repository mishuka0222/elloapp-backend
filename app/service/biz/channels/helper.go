package channels_helper

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/svc"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
	username_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username"
)

type (
	Config = config.Config
)

func New(c Config, userService *user_helper.UserService, usernameService *username_helper.UsernameService) *service.Service {
	return service.New(svc.NewServiceContext(c, userService, usernameService))
}
