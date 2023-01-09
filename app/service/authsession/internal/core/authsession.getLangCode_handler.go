package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetLangCode
// authsession.getLangCode auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetLangCode(in *authsession.TLAuthsessionGetLangCode) (*mtproto.String, error) {
	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getLangCode - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		c.Logger.Errorf("session.getLangCode - not found keyId %d", in.GetAuthKeyId())
		return nil, fmt.Errorf("not found keyId")
	}

	langCode := c.svcCtx.Dao.GetLangCode(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: langCode,
	}).To_String(), nil
}
