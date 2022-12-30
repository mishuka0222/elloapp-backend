package plugin

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

type UsernamesPlugin interface {
	GetChannelListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.Chat
}
