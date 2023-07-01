package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChatImportChatInvite
// chat.importChatInvite self_id:long hash:string = MutableChat;
func (c *ChatCore) ChatImportChatInvite(in *chat.TLChatImportChatInvite) (*chat.MutableChat, error) {
	chatInviteDO, err := c.svcCtx.Dao.ChatInvitesDAO.SelectByLink(c.ctx, in.Hash)
	if err != nil {
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	} else if chatInviteDO == nil {
		err = mtproto.ErrInviteHashInvalid
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	if chatInviteDO.ExpireDate != 0 && time.Now().Unix() > chatInviteDO.ExpireDate {
		err = mtproto.ErrInviteHashExpired
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}
	if chatInviteDO.UsageLimit > 0 {
		sz := c.svcCtx.Dao.CommonDAO.CalcSize(c.ctx, "chat_invite_participants", map[string]interface{}{
			"link": chatInviteDO.Link,
		})

		if sz >= int(chatInviteDO.UsageLimit) {
			err = mtproto.ErrInviteHashExpired
			c.Logger.Errorf("chat.importChatInvite - error: %v", err)
			return nil, err
		}
	}

	chat2, err := c.ChatAddChatUser(&chat.TLChatAddChatUser{
		ChatId:    chatInviteDO.ChatId,
		InviterId: chatInviteDO.AdminId,
		UserId:    in.SelfId,
	})
	if err != nil {
		c.Logger.Errorf("chat.importChatInvite - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.ChatInviteParticipantsDAO.Insert(c.ctx, &dataobject.ChatInviteParticipantsDO{
		ChatId: chatInviteDO.ChatId,
		Link:   in.Hash,
		UserId: in.SelfId,
		Date2:  time.Now().Unix(),
	})

	return chat2, nil
}
