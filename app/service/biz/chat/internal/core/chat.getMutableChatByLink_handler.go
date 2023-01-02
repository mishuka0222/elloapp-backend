package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ChatGetMutableChatByLink
// chat.getMutableChatByLink link:string = MutableChat;
func (c *ChatCore) ChatGetMutableChatByLink(in *chat.TLChatGetMutableChatByLink) (*chat.MutableChat, error) {
	// TODO: not impl
	c.Logger.Errorf("chat.getMutableChatByLink - error: method ChatGetMutableChatByLink not impl")

	return nil, mtproto.ErrMethodNotImpl
}
