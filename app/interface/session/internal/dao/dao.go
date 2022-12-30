/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import (
	bff_proxy_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/bff/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/internal/config"
	authsession_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/client"
	status_client "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/status/client"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto/rpc/metadata"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/cache"

	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	cache *cache.LRUCache
	authsession_client.AuthsessionClient
	status_client.StatusClient
	*bff_proxy_client.BFFProxyClient
}

func New(c config.Config) *Dao {
	return &Dao{
		cache:             cache.NewLRUCache(1024 * 1024 * 1024),
		AuthsessionClient: authsession_client.NewAuthsessionClient(zrpc.MustNewClient(c.AuthSession)),
		BFFProxyClient:    bff_proxy_client.NewBFFProxyClients(c.BFFProxyClients.Clients, c.BFFProxyClients.IDMap),
		StatusClient:      status_client.NewStatusClient(zrpc.MustNewClient(c.StatusClient)),
	}
}

func (d *Dao) Invoke(rpcMetaData *metadata.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	return d.BFFProxyClient.Invoke(rpcMetaData, object)
}
