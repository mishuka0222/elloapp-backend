package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UsernameDeleteUsername
// username.deleteUsername username:string = Bool;
func (c *UsernameCore) UsernameDeleteUsername(in *username.TLUsernameDeleteUsername) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UsernameDAO.Delete(c.ctx, in.Username)
	if err != nil {
		return mtproto.BoolFalse, nil
	}

	return mtproto.BoolTrue, nil
}
