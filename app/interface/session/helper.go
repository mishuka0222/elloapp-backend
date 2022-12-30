package session_helper

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/interface/session/internal/server"
)

func init() {
	zrpc.DontLogContentForMethod("/session.RPCSession/SessionSendDataToSession")
}

var (
	New = server.New
)
