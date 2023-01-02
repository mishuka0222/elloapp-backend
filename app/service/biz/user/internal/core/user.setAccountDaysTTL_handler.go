package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserSetAccountDaysTTL
// user.setAccountDaysTTL user_id:int ttl:int = Bool;
func (c *UserCore) UserSetAccountDaysTTL(in *user.TLUserSetAccountDaysTTL) (*mtproto.Bool, error) {
	c.svcCtx.Dao.UsersDAO.UpdateAccountDaysTTL(c.ctx, in.Ttl, in.UserId)

	return mtproto.BoolTrue, nil
}
