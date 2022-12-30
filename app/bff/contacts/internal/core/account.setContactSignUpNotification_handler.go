package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountSetContactSignUpNotification
// account.setContactSignUpNotification#cff43f61 silent:Bool = Bool;
func (c *ContactsCore) AccountSetContactSignUpNotification(in *mtproto.TLAccountSetContactSignUpNotification) (*mtproto.Bool, error) {
	rValue, err := c.svcCtx.Dao.UserClient.UserSetContactSignUpNotification(c.ctx, &userpb.TLUserSetContactSignUpNotification{
		UserId: c.MD.UserId,
		Silent: in.Silent,
	})
	if err != nil {
		c.Logger.Errorf("account.setContactSignUpNotification - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return rValue, nil
}
