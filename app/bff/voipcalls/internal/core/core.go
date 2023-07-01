package core

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/voipcalls/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type VoipcallsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func NewVoipcalls(ctx context.Context, svcCtx *svc.ServiceContext) *VoipcallsCore {
	return &VoipcallsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
