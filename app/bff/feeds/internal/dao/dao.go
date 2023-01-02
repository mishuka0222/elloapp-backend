package dao

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/feeds/internal/config"
	feeds_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/feeds/client"
	message_client "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net/rpcx"
)

type Dao struct {
	message_client.MessageClient
	feeds_client.FeedsClient
}

func New(c config.Config) *Dao {
	return &Dao{
		MessageClient: message_client.NewMessageClient(rpcx.GetCachedRpcClient(c.MessageClient)),
		FeedsClient:   feeds_client.NewFeedsClient(rpcx.GetCachedRpcClient(c.FeedsClient)),
	}
}
