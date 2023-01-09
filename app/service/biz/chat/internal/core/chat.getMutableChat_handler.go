package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
)

// ChatGetMutableChat
// chat.getMutableChat chat_id:long = MutableChat;
func (c *ChatCore) ChatGetMutableChat(in *chat.TLChatGetMutableChat) (*chat.MutableChat, error) {
	chat2, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId)
	if err != nil {
		c.Logger.Errorf("chat.getMutableChat - error: %v", err)
		return nil, err
	}

	return chat2, nil
}
