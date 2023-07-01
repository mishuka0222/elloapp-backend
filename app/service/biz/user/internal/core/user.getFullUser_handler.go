package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetFullUser
// user.getFullUser self_user_id:long id:long = users.UserFull;
func (c *UserCore) UserGetFullUser(in *user.TLUserGetFullUser) (*mtproto.Users_UserFull, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getFullUser - error: method UserGetFullUser not impl")

	return nil, mtproto.ErrMethodNotImpl
}
