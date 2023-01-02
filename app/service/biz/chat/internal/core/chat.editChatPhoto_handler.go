package core

import (
	"context"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

// ChatEditChatPhoto
// chat.editChatPhoto chat_id:long edit_user_id:long chat_photo:Photo = MutableChat;
func (c *ChatCore) ChatEditChatPhoto(in *chat.TLChatEditChatPhoto) (*chat.MutableChat, error) {
	var (
		err        error
		now        = time.Now().Unix()
		chatId     = in.ChatId
		editUserId = in.EditUserId
		chat2      *chat.MutableChat
		me         *chat.ImmutableChatParticipant
	)

	chat2, err = c.svcCtx.Dao.GetMutableChat(c.ctx, chatId, editUserId)
	if err != nil {
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
			affected, err2 := c.svcCtx.Dao.ChatsDAO.UpdatePhotoId(c.ctx, in.GetChatPhoto().GetId(), chatId)
			return 0, affected, err2
		},
		c.svcCtx.Dao.GetChatCacheKey(chatId))

	chat2.Chat.Version += 1
	chat2.Chat.Date = now
	return chat2, nil
}
