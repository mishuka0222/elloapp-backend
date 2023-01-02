package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AuthCheckPassword
// auth.checkPassword#d18b4d16 password:InputCheckPasswordSRP = auth.Authorization;
func (c *AuthorizationCore) AuthCheckPassword(in *mtproto.TLAuthCheckPassword) (*mtproto.Auth_Authorization, error) {
	// TODO: check password
	c.Logger.Errorf("auth.checkPassword blocked, License key from https://elloapp.com required to unlock enterprise features.")

	user, err := c.svcCtx.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
		Id: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("auth.checkPassword - error: %v", err)
		return nil, err
	}

	return mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user.ToSelfUser(),
	}).To_Auth_Authorization(), nil
}
