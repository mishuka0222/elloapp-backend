package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthBindTempAuthKey
// auth.bindTempAuthKey#cdd42a05 perm_auth_key_id:long nonce:long expires_at:int encrypted_message:bytes = Bool;
func (c *AuthorizationCore) AuthBindTempAuthKey(in *mtproto.TLAuthBindTempAuthKey) (*mtproto.Bool, error) {
	// TODO: not impl
	// c.Logger.Errorf("auth.bindTempAuthKey blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
