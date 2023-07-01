package core

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/svc"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/rpc/metadata"
)

var (
	usernameNotExisted   = username.MakeTLUsernameNotExisted(nil).To_UsernameExist()
	usernameExisted      = username.MakeTLUsernameExisted(nil).To_UsernameExist()
	usernameExistedNotMe = username.MakeTLUsernameExistedNotMe(nil).To_UsernameExist()
	usernameExistedIsMe  = username.MakeTLUsernameExistedIsMe(nil).To_UsernameExist()
)

type UsernameCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *UsernameCore {
	return &UsernameCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
