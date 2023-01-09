package plugin

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type NotificationPlugin interface {
	GetChannelListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.Chat
	GetChannelById(ctx context.Context, selfId int64, id int64) (*mtproto.Chat, error)
}
