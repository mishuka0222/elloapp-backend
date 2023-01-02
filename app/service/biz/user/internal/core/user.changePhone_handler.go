package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserChangePhone
// user.changePhone user_id:int phone:string = Bool;
func (c *UserCore) UserChangePhone(in *user.TLUserChangePhone) (*mtproto.Bool, error) {
	c.svcCtx.Dao.UsersDAO.UpdateUser(c.ctx, map[string]interface{}{
		"phone": in.Phone, // TODO(@benqi): country_code
	}, in.UserId)

	c.svcCtx.Dao.UserContactsDAO.UpdatePhoneByContactId(c.ctx, in.Phone, in.UserId)

	return mtproto.BoolTrue, nil
}
