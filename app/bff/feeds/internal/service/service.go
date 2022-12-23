package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/teamgram/teamgram-server/app/bff/feeds/internal/svc"
	log "github.com/zeromicro/go-zero/core/logx"
)

const (
	GetFeedList int32 = iota*100 + 100100
	UpdateFeedList
	GetHistoryCounter
	ReadHistory
)

type Service struct {
	svcCtx            *svc.ServiceContext
	operationRegister map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error)
}

func New(svcCtx *svc.ServiceContext) *Service {
	srv := &Service{
		svcCtx: svcCtx,
	}
	operationRegister := map[int32]func(ctx context.Context, data json.RawMessage) (interface{}, error){
		GetFeedList:       srv.GetFeedList,
		UpdateFeedList:    srv.UpdateFeedList,
		GetHistoryCounter: srv.GetHistoryCounter,
		ReadHistory:       srv.ReadHistory,
	}
	srv.operationRegister = operationRegister
	return srv
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
