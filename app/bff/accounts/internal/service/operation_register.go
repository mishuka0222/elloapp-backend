package service

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/zeromicro/go-zero/core/logx"
)

const (
	AccountsDelete int32 = iota*100 + 100100
	AccountsChangePassword
	AccountsChangeEmail
	AccountsChangeProfile
	AccountsForgotPassword
)

func (s *Service) initOperationRegister() {
	operationRegister := map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error){
		AccountsDelete:       		s.AccountsDelete,
		AccountsChangePassword:    	s.AccountsChangePassword,
		AccountsChangeEmail: 		s.AccountsChangeEmail,
		AccountsChangeProfile:      s.AccountsChangeProfile,
		AccountsForgotPassword:     s.AccountsForgotPassword,
	}
	s.operationRegister = operationRegister
}

func (s *Service) GetHandler(ctx context.Context, id int32, data json.RawMessage) (interface{}, error) {
	h, ok := s.operationRegister[id]
	if !ok {
		err := "Accounts service error, unknown operation id " + string(id)
		log.Error(err)
		return nil, errors.New(err)
	}
	return h(ctx, data)
}
