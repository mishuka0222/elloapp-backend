package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserUpdatePredefinedVerified
// user.updatePredefinedVerified flags:# phone:string verified:flags.1?true = PredefinedUser;
func (c *UserCore) UserUpdatePredefinedVerified(in *user.TLUserUpdatePredefinedVerified) (*mtproto.PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.updatePredefinedVerified - error: method UserUpdatePredefinedVerified not impl")

	return nil, mtproto.ErrMethodNotImpl
}
