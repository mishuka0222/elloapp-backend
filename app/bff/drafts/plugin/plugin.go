package plugin

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type DraftsPlugin interface {
	GetChannelListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.Chat
}
