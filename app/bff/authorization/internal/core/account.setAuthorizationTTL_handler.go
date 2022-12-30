package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountSetAuthorizationTTL
// account.setAuthorizationTTL#bf899aa0 authorization_ttl_days:int = Bool;
func (c *AuthorizationCore) AccountSetAuthorizationTTL(in *mtproto.TLAccountSetAuthorizationTTL) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("account.setAuthorizationTTL blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
