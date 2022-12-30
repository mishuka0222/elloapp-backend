package core

import (
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// AccountGetContactSignUpNotification
// account.getContactSignUpNotification#9f07c728 = Bool;
func (c *ContactsCore) AccountGetContactSignUpNotification(in *mtproto.TLAccountGetContactSignUpNotification) (*mtproto.Bool, error) {
	rValue, err := c.svcCtx.Dao.UserClient.UserGetContactSignUpNotification(c.ctx, &userpb.TLUserGetContactSignUpNotification{
		UserId: c.MD.UserId,
	})
	if err != nil {
		c.Logger.Errorf("account.getContactSignUpNotification - error: %v", err)
		return mtproto.BoolFalse, nil
	}

	return rValue, nil
}
