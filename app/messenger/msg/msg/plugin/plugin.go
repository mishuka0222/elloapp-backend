package plugin

import (
	"context"
)

type MsgPlugin interface {
	ReadReactionUnreadMessage(ctx context.Context, userId int64, msgId int32) error
}
