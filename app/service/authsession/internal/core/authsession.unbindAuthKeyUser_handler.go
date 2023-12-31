package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionUnbindAuthKeyUser
// authsession.unbindAuthKeyUser auth_key_id:long user_id:long = Bool;
func (c *AuthsessionCore) AuthsessionUnbindAuthKeyUser(in *authsession.TLAuthsessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	var (
		unBindKeyId = in.AuthKeyId
	)

	if unBindKeyId != 0 {
		keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, unBindKeyId)
		if err != nil {
			c.Logger.Errorf("session.unbindAuthKeyUser - error: %v", err)
			return nil, err
		} else if keyData == nil || keyData.PermAuthKeyId == 0 {
			err = mtproto.ErrAuthKeyInvalid
			return nil, err
		} else {
			unBindKeyId = keyData.PermAuthKeyId
		}
	}

	c.svcCtx.Dao.UnbindAuthUser(c.ctx, unBindKeyId, in.UserId)

	return mtproto.BoolTrue, nil
}
