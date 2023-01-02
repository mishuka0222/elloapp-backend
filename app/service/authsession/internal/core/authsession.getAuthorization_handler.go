package core

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetAuthorization
// authsession.getAuthorization auth_key_id:long = Authorization;
func (c *AuthsessionCore) AuthsessionGetAuthorization(in *authsession.TLAuthsessionGetAuthorization) (*mtproto.Authorization, error) {
	myKeyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.AuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	} else if myKeyData == nil || myKeyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	authorization, err := c.svcCtx.Dao.GetAuthorization(c.ctx, myKeyData.PermAuthKeyId)
	if err != nil {
		c.Logger.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	return authorization, nil
}
