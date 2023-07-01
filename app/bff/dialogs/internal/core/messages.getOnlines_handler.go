package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetOnlines
// messages.getOnlines#6e2be050 peer:InputPeer = ChatOnlines;
func (c *DialogsCore) MessagesGetOnlines(in *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error) {
	// TODO:
	onlines := mtproto.MakeTLChatOnlines(&mtproto.ChatOnlines{
		Onlines: 1,
	}).To_ChatOnlines()

	// TODO: not impl
	c.Logger.Errorf("messages.getOnlines - error: method MessagesGetOnlines not impl")

	return onlines, nil
}
