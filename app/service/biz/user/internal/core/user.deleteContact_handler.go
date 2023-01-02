package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserDeleteContact
// user.deleteContact user_id:long id:long = Bool;
func (c *UserCore) UserDeleteContact(in *user.TLUserDeleteContact) (*mtproto.Bool, error) {
	c.svcCtx.Dao.DeleteUserContact(c.ctx, in.GetUserId(), in.GetId())

	return mtproto.BoolTrue, nil
}
