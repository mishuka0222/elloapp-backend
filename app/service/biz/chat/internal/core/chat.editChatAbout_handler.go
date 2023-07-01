package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChatEditChatAbout
// chat.editChatAbout chat_id:long edit_user_id:long about:string = MutableChat;
func (c *ChatCore) ChatEditChatAbout(in *chat.TLChatEditChatAbout) (*chat.MutableChat, error) {
	var (
		now        = time.Now().Unix()
		chat2      *chat.MutableChat
		me         *chat.ImmutableChatParticipant
		err        error
		chatId     = in.ChatId
		editUserId = in.EditUserId
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, editUserId)
	if err != nil {
		return nil, err
	}
	if chat2.Chat.About == in.About {
		err = mtproto.ErrChatNotModified
		return nil, err
	}

	me, _ = chat2.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != mtproto.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	// TODO(@benqi): check
	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	_, _, err = c.svcCtx.Dao.CachedConn.Exec(
		c.ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rowsAffected, err2 := c.svcCtx.Dao.ChatsDAO.UpdateAbout(c.ctx, in.About, chat2.Id())
			return 0, rowsAffected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(chat2.Id()))
	if err != nil {
		return nil, err
	}

	chat2.Chat.About = in.About
	chat2.Chat.Version += 1
	chat2.Chat.Date = now

	return chat2, nil
}
