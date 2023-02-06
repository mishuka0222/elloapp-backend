package service

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/zeromicro/go-zero/core/logx"
)

const (
	AuthSingUP int32 = iota*100 + 100100
	AuthSingIN
	Confirmation
	ForgotPassword
	ForgotPasswordConfirm
)

func (s *Service) initOperationRegister() {
	operationRegister := map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error){
		AuthSingUP:   s.AuthSingUP,
		AuthSingIN:   s.AuthSingIN,
		Confirmation: s.Confirmation,
		ForgotPassword: s.ForgotPassword,
		ForgotPasswordConfirm: s.ForgotPasswordConfirm,
	}
	s.operationRegister = operationRegister
}

func (s *Service) GetHandler(ctx context.Context, id int32, data json.RawMessage) (interface{}, error) {
	h, ok := s.operationRegister[id]
	if !ok {
		err := "AuthorizationCustomize service error, unknown operation id " + string(id)
		log.Error(err)
		return nil, errors.New(err)
	}
	return h(ctx, data)
}
