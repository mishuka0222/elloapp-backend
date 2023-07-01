package service

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/zeromicro/go-zero/core/logx"
)

const (
	GetFeedList int32 = iota*100 + 100100
	UpdateFeedList
	GetHistoryCounter
	ReadHistory
	GetHistory
)

func (s *Service) initOperationRegister() {
	operationRegister := map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error){
		GetFeedList:       s.GetFeedList,
		UpdateFeedList:    s.UpdateFeedList,
		GetHistoryCounter: s.GetHistoryCounter,
		ReadHistory:       s.ReadHistory,
		GetHistory:        s.GetHistory,
	}
	s.operationRegister = operationRegister
}

func (s *Service) GetHandler(ctx context.Context, id int32, data json.RawMessage) (interface{}, error) {
	h, ok := s.operationRegister[id]
	if !ok {
		err := "Feeds service error, unknown operation id " + string(id)
		log.Error(err)
		return nil, errors.New(err)
	}
	return h(ctx, data)
}
