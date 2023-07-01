package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetAuthStateData
// authsession.getAuthStateData auth_key_id:long = AuthKeyStateData;
func (c *AuthsessionCore) AuthsessionGetAuthStateData(in *authsession.TLAuthsessionGetAuthStateData) (*authsession.AuthKeyStateData, error) {
	// TODO: not impl
	c.Logger.Errorf("authsession.getAuthStateData - error: method AuthsessionGetAuthStateData not impl")

	return nil, mtproto.ErrMethodNotImpl
}
