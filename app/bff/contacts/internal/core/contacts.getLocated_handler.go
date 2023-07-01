package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsGetLocated
// contacts.getLocated#d348bc44 flags:# background:flags.1?true geo_point:InputGeoPoint self_expires:flags.0?int = Updates;
func (c *ContactsCore) ContactsGetLocated(in *mtproto.TLContactsGetLocated) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("contacts.getLocated blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.MakeEmptyUpdates(), nil
}
