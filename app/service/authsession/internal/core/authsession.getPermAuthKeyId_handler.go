package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetPermAuthKeyId
// authsession.getPermAuthKeyId auth_key_id:long= Int64;
func (c *AuthsessionCore) AuthsessionGetPermAuthKeyId(in *authsession.TLAuthsessionGetPermAuthKeyId) (*mtproto.Int64, error) {
	v := c.svcCtx.Dao.GetPermAuthKeyId(c.ctx, in.AuthKeyId)

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: v,
	}).To_Int64(), nil
}
