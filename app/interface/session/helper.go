/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package session_helper

import (
	"github.com/teamgram/teamgram-server/app/interface/session/internal/server"
	"github.com/zeromicro/go-zero/zrpc"
)

func init() {
	zrpc.DontLogContentForMethod("/session.RPCSession/SessionSendDataToSession")
}

var (
	New = server.New
)
