package service

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/zeromicro/go-zero/core/logx"
)

const (
	AccountChangeEmail int32 = iota*100 + 100100
	AccountChangePassword
	AccountChangeProfile
	AccountDelete
	AccountForgotPassword
	AccountConfirmationSend
	AccountConfirmationConfirm
)

func (s *Service) initOperationRegister() {
	operationRegister := map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error){
		AccountChangeEmail:         s.AccountChangeEmail,
		AccountChangePassword:      s.AccountChangePassword,
		AccountChangeProfile:       s.AccountChangeProfile,
		AccountDelete:              s.AccountDelete,
		AccountForgotPassword:      s.AccountForgotPassword,
		AccountConfirmationSend:    s.AccountConfirmationSend,
		AccountConfirmationConfirm: s.AccountConfirmationConfirm,
	}
	s.operationRegister = operationRegister
}

func (s *Service) GetHandler(ctx context.Context, id int32, data json.RawMessage) (interface{}, error) {
	h, ok := s.operationRegister[id]
	if !ok {
		err := "AccountCustomize service error, unknown operation id " + string(id)
		log.Error(err)
		return nil, errors.New(err)
	}
	return h(ctx, data)
}
