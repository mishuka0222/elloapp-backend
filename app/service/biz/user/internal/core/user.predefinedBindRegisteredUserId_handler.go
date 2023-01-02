package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserPredefinedBindRegisteredUserId
// user.predefinedBindRegisteredUserId phone:string registered_userId:int = Bool;
func (c *UserCore) UserPredefinedBindRegisteredUserId(in *user.TLUserPredefinedBindRegisteredUserId) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("user.predefinedBindRegisteredUserId - error: method UserPredefinedBindRegisteredUserId not impl")

	return nil, mtproto.ErrMethodNotImpl
}
