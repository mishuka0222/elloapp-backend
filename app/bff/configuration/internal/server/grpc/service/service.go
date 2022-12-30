package service

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/bff/configuration/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
	}
}
