package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthsessionGetLangPack
// authsession.getLangPack auth_key_id:long = String;
func (c *AuthsessionCore) AuthsessionGetLangPack(in *authsession.TLAuthsessionGetLangPack) (*mtproto.String, error) {
	keyData, err := c.svcCtx.Dao.QueryAuthKeyV2(c.ctx, in.GetAuthKeyId())
	if err != nil {
		c.Logger.Errorf("session.getLangPack - error: %v", err)
		return nil, err
	} else if keyData == nil || keyData.PermAuthKeyId == 0 {
		return nil, fmt.Errorf("not found keyId")
	}

	langPack := c.svcCtx.GetLangPack(c.ctx, keyData.PermAuthKeyId)

	return mtproto.MakeTLString(&mtproto.String{
		V: langPack,
	}).To_String(), nil
}
