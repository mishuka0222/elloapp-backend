package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetAllChats
// messages.getAllChats#875f74be except_ids:Vector<long> = messages.Chats;
func (c *ChatsCore) MessagesGetAllChats(in *mtproto.TLMessagesGetAllChats) (*mtproto.Messages_Chats, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getAllChats blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats(), nil
}
