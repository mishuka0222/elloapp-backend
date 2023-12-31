package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserUpdateAbout
// user.updateAbout user_id:long about:string = Bool;
func (c *UserCore) UserUpdateAbout(in *user.TLUserUpdateAbout) (*mtproto.Bool, error) {
	rB := c.svcCtx.Dao.UpdateUserAbout(
		c.ctx,
		in.GetUserId(),
		in.GetAbout())

	return mtproto.ToBool(rB), nil
}
