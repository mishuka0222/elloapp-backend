package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetImmutableUserByPhone
// user.getImmutableUserByPhone phone:string = ImmutableUser;
func (c *UserCore) UserGetImmutableUserByPhone(in *user.TLUserGetImmutableUserByPhone) (*user.ImmutableUser, error) {
	// TODO: performance optimization
	do, err := c.svcCtx.Dao.UsersDAO.SelectByPhoneNumber(c.ctx, in.Phone)
	if err != nil {
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	} else if do == nil {
		err = mtproto.ErrPhoneNumberUnoccupied
		c.Logger.Errorf("user.getImmutableUserByPhone - error: %v", err)
		return nil, err
	}

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: do.Id,
	})
}
