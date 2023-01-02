package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesSaveDefaultSendAs
// messages.saveDefaultSendAs#ccfddf96 peer:InputPeer send_as:InputPeer = Bool;
func (c *MessagesCore) MessagesSaveDefaultSendAs(in *mtproto.TLMessagesSaveDefaultSendAs) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.saveDefaultSendAs blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
