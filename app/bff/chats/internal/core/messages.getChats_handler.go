package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetChats
// messages.getChats#49e9528f id:Vector<long> = messages.Chats;
func (c *ChatsCore) MessagesGetChats(in *mtproto.TLMessagesGetChats) (*mtproto.Messages_Chats, error) {
	chats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: make([]*mtproto.Chat, 0, len(in.Id)),
	}).To_Messages_Chats()

	for _, id := range in.Id {
		chat, _ := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
			ChatId: id,
		})
		if chat != nil {
			chats.Chats = append(chats.Chats, chat.ToUnsafeChat(c.MD.UserId))
		}
	}

	return chats, nil
}
