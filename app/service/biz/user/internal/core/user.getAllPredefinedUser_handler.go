package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetAllPredefinedUser
// user.getAllPredefinedUser = Vector<PredefinedUser>;
func (c *UserCore) UserGetAllPredefinedUser(in *user.TLUserGetAllPredefinedUser) (*user.Vector_PredefinedUser, error) {
	// TODO: not impl
	c.Logger.Errorf("user.getAllPredefinedUser - error: method UserGetAllPredefinedUser not impl")

	return &user.Vector_PredefinedUser{
		Datas: []*mtproto.PredefinedUser{},
	}, nil
}
