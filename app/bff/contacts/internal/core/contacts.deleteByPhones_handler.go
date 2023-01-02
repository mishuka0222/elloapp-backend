package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ContactsDeleteByPhones
// contacts.deleteByPhones#1013fd9e phones:Vector<string> = Bool;
func (c *ContactsCore) ContactsDeleteByPhones(in *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	// @benqi: clients not use

	return mtproto.BoolTrue, nil
}
