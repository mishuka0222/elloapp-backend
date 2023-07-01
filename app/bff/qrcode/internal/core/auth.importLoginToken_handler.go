package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthImportLoginToken
// auth.importLoginToken#95ac5ce4 token:bytes = auth.LoginToken;
func (c *QrCodeCore) AuthImportLoginToken(in *mtproto.TLAuthImportLoginToken) (*mtproto.Auth_LoginToken, error) {
	// TODO: not impl
	c.Logger.Errorf("auth.importLoginToken blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrAuthTokenInvalid
}
