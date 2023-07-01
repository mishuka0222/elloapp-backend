package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsResetTopPeerRating
// contacts.resetTopPeerRating#1ae373ac category:TopPeerCategory peer:InputPeer = Bool;
func (c *ContactsCore) ContactsResetTopPeerRating(in *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("contacts.resetTopPeerRating blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
