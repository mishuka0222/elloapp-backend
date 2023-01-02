package core

import (
	"github.com/zeromicro/go-zero/core/mr"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionSetAuthKey
// authsession.setAuthKey auth_key:AuthKeyInfo future_salt:FutureSalt = Bool;
func (c *AuthsessionCore) AuthsessionSetAuthKey(in *authsession.TLAuthsessionSetAuthKey) (*mtproto.Bool, error) {
	var (
		keyInfo = in.GetAuthKey()
		salt    *mtproto.TLFutureSalt
		err     error
	)

	if in.FutureSalt != nil {
		salt = in.FutureSalt.To_FutureSalt()
	}
	if salt == nil {
		err = c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
	} else {
		err = mr.Finish(
			func() error {
				return c.svcCtx.Dao.SetAuthKeyV2(c.ctx, keyInfo, in.ExpiresIn)
			},
			func() error {
				return c.svcCtx.Dao.PutSaltCache(c.ctx, keyInfo.AuthKeyId, salt)
			})
	}

	if err != nil {
		c.Logger.Errorf("authsession.setAuthKey - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
