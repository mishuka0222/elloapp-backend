/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package gateway_helper

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/gateway/internal/server/server"
)

func init() {
	zrpc.DontLogContentForMethod("/gateway.RPCGateway/GatewaySendDataToGateway")
}

type (
	Server = server.Server
)
