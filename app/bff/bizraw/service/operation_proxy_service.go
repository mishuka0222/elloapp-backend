package operation_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type Service struct {
	register map[ServiceID]OperationServer
}

func New(register map[ServiceID]OperationServer) *Service {
	if len(K_SERVER_LIST) < len(register) {
		panic("Operation service error, incorrect register length")
	}
	validRegister := make(map[ServiceID]OperationServer, len(register))
	for _, it := range K_SERVER_LIST {
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
		err := fmt.Sprintf("Operation service error, unknown service id: %d", servID)
		log.Error(err)
		return nil, errors.New(err)
	}
	resp, err := srv.GetHandler(ctx, opID, data)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return (&mtproto.BizDataRaw{
		Constructor: mtproto.CRC32_bizDataRaw,
		Data:        b,
	}).To_BizDataRaw().To_BizDataRaw(), nil
}
