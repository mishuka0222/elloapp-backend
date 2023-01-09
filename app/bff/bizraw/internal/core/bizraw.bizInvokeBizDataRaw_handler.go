package core

import (
	"encoding/json"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type Operation struct {
	Service int32           `json:"service"`
	Method  int32           `json:"method"`
	Data    json.RawMessage `json:"data"`
}

func (c BizCore) OperationProxyService(in *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {

	var op Operation
	err := json.Unmarshal(in.BizData.Data, &op)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("OperationProxyService service - %d, operation - %d", op.Service, op.Method)
	return c.svcCtx.OpSrv.Handle(c.ctx, op.Service, op.Method, op.Data)
}
