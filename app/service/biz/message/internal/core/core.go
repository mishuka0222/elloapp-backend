package core

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
)

type MessageCore struct {
	ctx context.Context
	logx.Logger
	svcCtx *svc.ServiceContext
	MD     *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *MessageCore {
	return &MessageCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
