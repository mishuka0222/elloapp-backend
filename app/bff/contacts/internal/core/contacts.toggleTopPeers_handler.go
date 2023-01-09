package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsToggleTopPeers
// contacts.toggleTopPeers#8514bdda enabled:Bool = Bool;
func (c *ContactsCore) ContactsToggleTopPeers(in *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("contacts.toggleTopPeers blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
