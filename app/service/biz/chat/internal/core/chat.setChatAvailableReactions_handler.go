package core

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/pkg2/stores/sqlx"
)

// ChatSetChatAvailableReactions
// chat.setChatAvailableReactions self_id:long chat_id:long available_reactions:Vector<string> = Bool;
func (c *ChatCore) ChatSetChatAvailableReactions(in *chat.TLChatSetChatAvailableReactions) (*chat.MutableChat, error) {
	var (
		chat2 *chat.MutableChat
		me    *chat.ImmutableChatParticipant
		err   error
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.SelfId)
	if err != nil {
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(in.SelfId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrParticipantIdInvalid
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminRequired
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	var (
		availableReactions string
	)

	if len(in.AvailableReactions) > 0 {
		availableReactionsData, _ := jsonx.Marshal(in.AvailableReactions)
		if availableReactionsData != nil {
			availableReactions = string(availableReactionsData)
		}
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			affected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateAvailableReactions(c.ctx, in.AvailableReactionsType, availableReactions, in.ChatId)
			return 0, affected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(in.ChatId))
	if err != nil {
		c.Logger.Errorf("chat.setChatAvailableReactions - error: %v")
		return nil, err
	}

	chat2.Chat.AvailableReactions = in.AvailableReactions
	return chat2, nil
}
