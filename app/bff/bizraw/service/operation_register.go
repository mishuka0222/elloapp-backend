package operation_service

import (
	"context"
	"encoding/json"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type OperationServer interface {
	GetHandler(ctx context.Context, id int32, data json.RawMessage) (interface{}, error)
}

type ServiceID int32

const (
	Feeds ServiceID = iota*100 + 100100
	Accounts
	AuthorizationCustomize
)

var K_WITHOUT_LOGIN = []ServiceID{
	AuthorizationCustomize,
}

var K_SERVER_LIST = []ServiceID{
	Feeds,
	Accounts,
	AuthorizationCustomize,
}

// Operation
// Do simplify tests
type Operation struct {
	Service ServiceID   `json:"service"`
	Method  int32       `json:"method"`
	Data    interface{} `json:"data"`
}

// NewOperation
// Do simplify tests
func NewOperation(op Operation) (*mtproto.TLBizInvokeBizDataRaw, error) {
	b, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}
	return &mtproto.TLBizInvokeBizDataRaw{
		Constructor: mtproto.CRC32_biz_invokeBizDataRaw,
		BizData: (&mtproto.BizDataRaw{
			Constructor: mtproto.CRC32_bizDataRaw,
			Data:        b,
		}).To_BizDataRaw().To_BizDataRaw(),
	}, nil
}
