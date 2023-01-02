package core

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthsessionBindAuthKeyUser
// authsession.bindAuthKeyUser auth_key_id:long user_id:long = Int64;
func (c *AuthsessionCore) AuthsessionBindAuthKeyUser(in *authsession.TLAuthsessionBindAuthKeyUser) (*mtproto.Int64, error) {
	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.bindAuthKeyUser - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	hash := c.svcCtx.Dao.BindAuthKeyUser(c.ctx, keyData.PermAuthKeyId, in.GetUserId())

	return &mtproto.Int64{V: hash}, nil
}
