package plugin

import (
	"context"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

type DialogsPlugin interface {
	GetChannelListByIdList(ctx context.Context, selfId int64, id ...int64) []*mtproto.Chat
	GetChannelDialogById(ctx context.Context, selfId int64, id int64) (*dialog.DialogExt, error)
	GetChannelMessage(ctx context.Context, selfId, channelId int64, id int32) (*mtproto.MessageBox, error)
}
