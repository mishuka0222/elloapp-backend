package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
)

// ChatGetChatBySelfId
// chat.getChatBySelfId self_id:long chat_id:long = MutableChat;
func (c *ChatCore) ChatGetChatBySelfId(in *chat.TLChatGetChatBySelfId) (*chat.MutableChat, error) {
	chat2, err := c.svcCtx.Dao.GetMutableChat(c.ctx, in.ChatId, in.SelfId)
	if err != nil {
		c.Logger.Errorf("chat.getMutableChat - error: %v", err)
		return nil, err
	}

	return chat2, nil
}
