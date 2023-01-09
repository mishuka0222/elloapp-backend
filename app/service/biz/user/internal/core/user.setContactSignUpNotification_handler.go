package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserSetContactSignUpNotification
// user.setContactSignUpNotification user_id:long silent:Bool = Bool;
func (c *UserCore) UserSetContactSignUpNotification(in *user.TLUserSetContactSignUpNotification) (*mtproto.Bool, error) {
	var (
		k, v string
	)

	k = "contactSignUpNotification"

	if mtproto.FromBool(in.Silent) {
		v = "false"
	} else {
		v = "true"
	}

	_, _, err := c.svcCtx.Dao.UserSettingsDAO.InsertOrUpdate(c.ctx, &dataobject.UserSettingsDO{
		UserId: in.UserId,
		Key2:   k,
		Value:  v,
	})
	if err != nil {
		c.Logger.Errorf("user.setContactSignUpNotification - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
