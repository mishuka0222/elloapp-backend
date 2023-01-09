package core

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetAuthorizations
// authsession.getAuthorizations user_id:long exclude_auth_keyId:long = account.Authorizations;
func (c *AuthsessionCore) AuthsessionGetAuthorizations(in *authsession.TLAuthsessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	myKeyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetExcludeAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getAuthorizations - error: %v", err)
		return nil, err
	} else if myKeyData == nil || myKeyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	authorizationList := c.svcCtx.Dao.GetAuthorizations(c.ctx, in.GetUserId(), myKeyData.PermAuthKeyId)

	return mtproto.MakeTLAccountAuthorizations(&mtproto.Account_Authorizations{
		Authorizations: authorizationList,
	}).To_Account_Authorizations(), nil
}
