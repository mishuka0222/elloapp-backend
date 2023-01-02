package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthResetAuthorizations
// auth.resetAuthorizations#9fab0d1a = Bool;
func (c *AuthorizationCore) AuthResetAuthorizations(in *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.resetAuthorizations blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
