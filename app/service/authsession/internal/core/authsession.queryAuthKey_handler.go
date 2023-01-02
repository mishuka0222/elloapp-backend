package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthsessionQueryAuthKey
// authsession.queryAuthKey auth_key_id:long = AuthKeyInfo;
func (c *AuthsessionCore) AuthsessionQueryAuthKey(in *authsession.TLAuthsessionQueryAuthKey) (*mtproto.AuthKeyInfo, error) {
	rValue, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("authsession.queryAuthKey - error: %v", err.Error())
		return nil, err
	}

	return rValue, nil
}
