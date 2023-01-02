package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountChangePhone
// account.changePhone#70c32edb phone_number:string phone_code_hash:string phone_code:string = User;
func (c *AccountCore) AccountChangePhone(in *mtproto.TLAccountChangePhone) (*mtproto.User, error) {
	// TODO: not impl
	c.Logger.Errorf("account.changePhone blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
