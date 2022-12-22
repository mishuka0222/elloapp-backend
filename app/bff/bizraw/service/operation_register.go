package operation_service

import (
	"context"
	"encoding/json"
	"github.com/teamgram/proto/mtproto"
)

type OperationServer interface {
	GetHandler(ctx context.Context, id int32, data json.RawMessage) (*mtproto.BizDataRaw, error)
}

type ServiceID int32

const (
	Feeds ServiceID = iota*100 + 100100
	OtherServer
)

var ServerList = []ServiceID{
	Feeds,
}
