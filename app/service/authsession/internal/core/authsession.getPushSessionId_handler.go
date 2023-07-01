package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetPushSessionId
// authsession.getPushSessionId user_id:long auth_key_id:long token_type:int = Int64;
func (c *AuthsessionCore) AuthsessionGetPushSessionId(in *authsession.TLAuthsessionGetPushSessionId) (*mtproto.Int64, error) {
	sessionId := c.svcCtx.Dao.GetPushSessionId(c.ctx, in.GetUserId(), in.GetAuthKeyId(), in.GetTokenType())

	return mtproto.MakeTLInt64(&mtproto.Int64{
		V: sessionId,
	}).To_Int64(), nil
}
