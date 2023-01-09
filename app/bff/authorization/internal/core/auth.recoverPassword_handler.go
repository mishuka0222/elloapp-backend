package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthRecoverPassword
// auth.recoverPassword#37096c70 flags:# code:string new_settings:flags.0?account.PasswordInputSettings = auth.Authorization;
func (c *AuthorizationCore) AuthRecoverPassword(in *mtproto.TLAuthRecoverPassword) (*mtproto.Auth_Authorization, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.recoverPassword blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
