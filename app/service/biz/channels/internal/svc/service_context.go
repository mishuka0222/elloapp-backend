package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
	username_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
	idgen_client.IDGenClient2
	*user_helper.UserService
	*username_helper.UsernameService
	media_client.MediaClient
}

func NewServiceContext(c config.Config, userService *user_helper.UserService, usernameService *username_helper.UsernameService) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		Dao:             dao.New(c),
		IDGenClient2:    idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		UserService:     userService,
		UsernameService: usernameService,
		MediaClient:     media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
	}
}
