package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetPredefinedUser
// user.getPredefinedUser phone:string = PredefinedUser;
func (c *UserCore) UserGetPredefinedUser(in *user.TLUserGetPredefinedUser) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getPredefinedUser - error: method UserGetPredefinedUser not impl")

	return nil, mtproto.ErrMethodNotImpl
}
