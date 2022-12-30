package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountUpdateVerified
// account.updateVerified flags:# id:long verified:flags.0?true = User;
func (c *AccountCore) AccountUpdateVerified(in *mtproto.TLAccountUpdateVerified) (*mtproto.User, error) {
	// TODO: not impl
	c.Logger.Errorf("account.updateVerified blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
