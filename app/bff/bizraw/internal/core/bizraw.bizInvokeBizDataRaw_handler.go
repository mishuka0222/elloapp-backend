package core

import (
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
)

type Operation struct {
	Service int32           `json:"service"`
	Method  int32           `json:"method"`
	Data    json.RawMessage `json:"data"`
}

func (c BizCore) BizInvokeBizDataRaw(in *mtproto.TLBizInvokeBizDataRaw) (*mtproto.BizDataRaw, error) {

	var op Operation
	err := json.Unmarshal(in.BizData.Data, &op)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BizInvokeBizDataRaw service - %d, operation - %d", op.Service, op.Method)
	return c.svcCtx.OpSrv.Handle(c.ctx, op.Service, op.Method, op.Data)
}
