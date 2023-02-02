package core

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/accounts/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountsCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func NewAccounts(ctx context.Context, svcCtx *svc.ServiceContext) *AccountsCore {
	return &AccountsCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
