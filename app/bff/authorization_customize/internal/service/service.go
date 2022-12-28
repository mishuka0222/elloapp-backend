package service

import (
	"context"
	"encoding/json"

	"github.com/teamgram/teamgram-server/app/bff/authorization_customize/internal/svc"
)

type Service struct {
	svcCtx            *svc.ServiceContext
	operationRegister map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error)
}

func New(svcCtx *svc.ServiceContext) *Service {
	srv := &Service{
		svcCtx: svcCtx,
	}
	srv.initOperationRegister()
	return srv
}
