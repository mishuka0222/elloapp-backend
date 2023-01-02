package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserUpdateProfilePhoto
// user.updateProfilePhoto user_id:long id:long = Int64;
func (c *UserCore) UserUpdateProfilePhoto(in *user.TLUserUpdateProfilePhoto) (*mtproto.Int64, error) {
	rV := c.svcCtx.Dao.UpdateProfilePhoto(
		c.ctx,
		in.GetUserId(),
		in.GetId())

	return &mtproto.Int64{
		V: rV,
	}, nil
}
