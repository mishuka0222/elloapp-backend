package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ContactsGetContactIDs
// contacts.getContactIDs#7adc669d hash:long = Vector<int>;
func (c *ContactsCore) ContactsGetContactIDs(in *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error) {
	// @benqi: clients not use

	return &mtproto.Vector_Int{
		Datas: []int32{},
	}, nil
}
