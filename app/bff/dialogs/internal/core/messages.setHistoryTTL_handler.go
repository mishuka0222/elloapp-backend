package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesSetHistoryTTL
// messages.setHistoryTTL#b80e5fe4 peer:InputPeer period:int = Updates;
func (c *DialogsCore) MessagesSetHistoryTTL(in *mtproto.TLMessagesSetHistoryTTL) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.setHistoryTTL - error: method MessagesSetHistoryTTL not impl")

	return nil, mtproto.ErrMethodNotImpl
}
