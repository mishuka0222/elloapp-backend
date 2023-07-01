package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserCreateNewPredefinedUser
// user.createNewPredefinedUser flags:# phone:string first_name:string last_name:flags.0?string username:string code:string verified:flags.1?true = PredefinedUser;
func (c *UserCore) UserCreateNewPredefinedUser(in *user.TLUserCreateNewPredefinedUser) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.createNewPredefinedUser - error: method UserCreateNewPredefinedUser not impl")

	return nil, mtproto.ErrMethodNotImpl
}
