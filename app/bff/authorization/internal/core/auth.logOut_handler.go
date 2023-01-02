package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/authsession/authsession"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// AuthLogOut
// auth.logOut#3e72ba19 = auth.LoggedOut;
func (c *AuthorizationCore) AuthLogOut(in *mtproto.TLAuthLogOut) (*mtproto.Auth_LoggedOut, error) {
	// unbind auth_key and user_id
	_, err := c.svcCtx.Dao.AuthsessionClient.AuthsessionUnbindAuthKeyUser(c.ctx, &authsession.TLAuthsessionUnbindAuthKeyUser{
		AuthKeyId: c.MD.AuthId,
		UserId:    c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.logOut - error: %v", err)
	} else {
		if c.svcCtx.Plugin != nil {
			c.svcCtx.Plugin.OnAuthLogout(c.ctx, c.MD.UserId, c.MD.AuthId)
		}
	}

	return mtproto.MakeTLAuthLoggedOut(&mtproto.Auth_LoggedOut{
		FutureAuthToken: nil,
	}).To_Auth_LoggedOut(), nil
}
