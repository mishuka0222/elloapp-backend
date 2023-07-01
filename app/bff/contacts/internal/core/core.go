package core

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/contacts/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
)

type ContactsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *ContactsCore {
	return &ContactsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
