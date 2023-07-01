package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
)

// UserGetImmutableUser
// user.getImmutableUser id:long = ImmutableUser;
func (c *UserCore) UserGetImmutableUser(in *user.TLUserGetImmutableUser) (*user.ImmutableUser, error) {
	imUser, err := c.svcCtx.Dao.GetImmutableUser(c.ctx, in.GetId(), false)
	if err != nil {
		return nil, err
	}

	return imUser, nil
}
