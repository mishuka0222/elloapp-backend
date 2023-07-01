package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files/internal/config"
	user_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/client"
	dfs_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/dfs/client"
	media_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/media/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	dfs_client.DfsClient
	media_client.MediaClient
	user_client.UserClient
}

func New(c config.Config) *Dao {
	return &Dao{
		DfsClient:   dfs_client.NewDfsClient(rpcx.GetCachedRpcClient(c.DfsClient)),
		MediaClient: media_client.NewMediaClient(rpcx.GetCachedRpcClient(c.MediaClient)),
		UserClient:  user_client.NewUserClient(rpcx.GetCachedRpcClient(c.UserClient)),
	}
}
