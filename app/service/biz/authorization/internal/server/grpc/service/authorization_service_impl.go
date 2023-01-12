package service

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/core"
)

// AuthSingUP
func (s *Service) AuthSingUP(ctx context.Context, in *authorization.AuthSingUpRequest) (*authorization.AuthSingUpRsp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.AuthSingUP - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.AuthSingUP(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.AuthSingUP - reply: %+v", r)
	return r, err
}

// AuthSingIN
func (s *Service) AuthSingIN(ctx context.Context, in *authorization.AuthSingInRequest) (*authorization.AuthSingInRsp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.AuthSingIN - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.AuthSingIN(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.AuthSingUP - reply: %+v", r)
	return r, err
}
