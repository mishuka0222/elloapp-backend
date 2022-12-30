package service

import (
	"github.com/teamgram/teamgram-server/app/service/biz/configuration/internal/svc"
)

type Service struct {
	svcCtx *svc.ServiceContext
}

func New(ctx *svc.ServiceContext) *Service {
	return &Service{
		svcCtx: ctx,
	}
}
