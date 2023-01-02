package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/bff/authorization/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
	}
}
