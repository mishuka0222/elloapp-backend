package plugin

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

type MessagePlugin interface {
	GetMessageMediaPoll(ctx context.Context, userId int64, pollId int64) (*mtproto.MessageMedia, error)
}
