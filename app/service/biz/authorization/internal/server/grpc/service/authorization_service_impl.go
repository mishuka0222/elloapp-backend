package service

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/authorization"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/authorization/internal/core"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthSingUP
func (s *Service) AuthSignUp(ctx context.Context, in *authorization.AuthSignUpRequest) (*authorization.AuthSignUpRsp, error) {
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
func (s *Service) AuthSignIn(ctx context.Context, in *authorization.AuthSignInRequest) (*mtproto.Auth_Authorization, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.AuthSingIN - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.AuthSingIN(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.AuthSingUP - reply: %+v", r)
	return r, err
}

func (s *Service) Confirmation(ctx context.Context, in *authorization.ConfirmationRequest) (*authorization.ConfirmationResponse, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.Confirmation(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}

func (s *Service) ForgotPassword(ctx context.Context, in *authorization.ForgotPasswordReq) (*authorization.ForgotPasswordRsp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.ForgotPassword(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}

func (s *Service) ForgotPasswordConfirm(ctx context.Context, in *authorization.ForgotPasswordConfirmReq) (*authorization.ForgotPasswordConfirmRsp, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("auth.Confirmation - metadata: %s, request: %s", c.MD.DebugString(), in.String())

	r, err := c.ForgotPasswordConfirm(in)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("auth.Confirmation - reply: %+v", r)
	return r, err
}
