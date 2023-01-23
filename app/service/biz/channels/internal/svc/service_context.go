package svc

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao"
	user_helper "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user"
	idgen_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/idgen/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type ServiceContext struct {
	Config config.Config
	Dao    *dao.Dao
	idgen_client.IDGenClient2
	*user_helper.UserService
}

func NewServiceContext(c config.Config, userService *user_helper.UserService) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Dao:          dao.New(c),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
		UserService:  userService,
	}
}
