package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// ChatSetHistoryTTL
// chat.setHistoryTTL self_id:long chat_id:long ttl_period:int = Bool;
func (c *ChatCore) ChatSetHistoryTTL(in *chat.TLChatSetHistoryTTL) (*chat.MutableChat, error) {
	mChat, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.setHistoryTTL - error: %v", err)
		return nil, err
	}
	if mChat.Creator() != in.SelfId {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.setHistoryTTL - error: %v", err)
		return nil, err
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateTTLPeriod(c.ctx, in.TtlPeriod, in.ChatId)
			return 0, rowsAffected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId))

	if err != nil {
		c.Logger.Errorf("chat.editChatAdmin - error: %v", err)
		return nil, err
	}

	return mChat, nil
}
