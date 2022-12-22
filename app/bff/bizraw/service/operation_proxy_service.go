package operation_service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/teamgram/proto/mtproto"
	log "github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	register map[ServiceID]OperationServer
}

func New(register map[ServiceID]OperationServer) *Service {
	if len(ServerList) < len(register) {
		panic("Operation service error, incorrect register length")
	}
	validRegister := make(map[ServiceID]OperationServer, len(ServerList))
	for _, it := range ServerList {
		serv, ok := register[it]
		if ok {
			validRegister[it] = serv
		}
	}
	if len(register) != len(validRegister) {
		panic("Operation service error, register include unknown server id")
	}

	return &Service{validRegister}
}

func (s *Service) Handle(ctx context.Context, servID int32, opID int32, data json.RawMessage) (*mtproto.BizDataRaw, error) {
	srv, ok := s.register[ServiceID(servID)]
	if !ok {
		err := "Operation service error, unknown service id: " + string(servID)
		log.Error(err)
		return nil, errors.New(err)
	}
	return srv.GetHandler(ctx, opID, data)
}
