package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UsersGetMe
// users.getMe id:long token:string = User;
func (c *UsersCore) UsersGetMe(in *mtproto.TLUsersGetMe) (*mtproto.User, error) {
	// TODO: not impl
	c.Logger.Errorf("users.getMe blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
