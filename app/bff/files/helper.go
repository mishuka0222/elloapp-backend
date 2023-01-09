package files_helper

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files/internal/config"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files/internal/server/grpc/service"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/files/plugin"
)

type (
	Config = config.Config
)

func init() {
	zrpc.DontLogContentForMethod("/mtproto.RPCFiles/UploadSaveFilePart")
	zrpc.DontLogContentForMethod("/mtproto.RPCFiles/UploadSaveBigFilePart")
}

func New(c Config, plugin plugin.FilesPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin))
}
