package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesGetExtendedMedia
// messages.getExtendedMedia#84f80814 peer:InputPeer id:Vector<int> = Updates;
func (c *MessagesCore) MessagesGetExtendedMedia(in *mtproto.TLMessagesGetExtendedMedia) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.getExtendedMedia blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
