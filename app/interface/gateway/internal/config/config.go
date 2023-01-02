package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/net2"
)

type Config struct {
	zrpc.RpcServerConf
	MaxProc        int
	KeyFile        string
	KeyFingerprint string
	Server         *net2.TcpServerConfig
	Session        zrpc.RpcClientConf
}

//type TcpServerConfig struct {
//	Addrs      []string
//	Multicore  bool
//	SendBuf    int
//	ReceiveBuf int
//}
