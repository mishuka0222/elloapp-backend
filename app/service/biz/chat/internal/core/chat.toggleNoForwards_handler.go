package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
	"time"
)

// ChatToggleNoForwards
// chat.toggleNoForwards chat_id:long operator_id:long enabled:Bool = MutableChat;
func (c *ChatCore) ChatToggleNoForwards(in *chat.TLChatToggleNoForwards) (*chat.MutableChat, error) {
	var (
		now   = time.Now().Unix()
		chat2 *chat.MutableChat
		me    *chat.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.OperatorId)
	if err != nil {
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.OperatorId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	if !me.IsChatMemberCreator() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateNoforwards(c.ctx, mtproto.FromBool(in.Enabled), in.ChatId)
			return 0, affected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId))
	if err != nil {
		c.Logger.Errorf("chat.toggleNoForwards - error: %v", err)
		return nil, err
	}

	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	chat2.Chat.Noforwards = mtproto.FromBool(in.Enabled)

	return chat2, nil
}
