package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetContact
// user.getContact user_id:long id:long = ContactData;
func (c *UserCore) UserGetContact(in *user.TLUserGetContact) (*user.ContactData, error) {
	contact := c.svcCtx.Dao.GetUserContact(c.ctx, in.GetUserId(), in.GetId())
	if contact == nil {
		err := mtproto.ErrContactIdInvalid
		c.Logger.Errorf("user.getContact - error: %v", err)
		return nil, err
	}

	return contact, nil
}
