package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthDropTempAuthKeys
// auth.dropTempAuthKeys#8e48a188 except_auth_keys:Vector<long> = Bool;
func (c *AuthorizationCore) AuthDropTempAuthKeys(in *mtproto.TLAuthDropTempAuthKeys) (*mtproto.Bool, error) {
	// TODO: not impl
	// c.Logger.Errorf("auth.dropTempAuthKeys blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return mtproto.BoolTrue, nil
}
