package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserGetImmutableUserByToken
// user.getImmutableUserByToken token:string = ImmutableUser;
func (c *UserCore) UserGetImmutableUserByToken(in *user.TLUserGetImmutableUserByToken) (*user.ImmutableUser, error) {
	// TODO: performance optimization
	botId, err := c.svcCtx.Dao.BotsDAO.SelectByToken(c.ctx, in.Token)
	if err != nil || botId == 0 {
		err = mtproto.ErrTokenInvalid
		return nil, err
	}

	return c.UserGetImmutableUser(&user.TLUserGetImmutableUser{
		Id: botId,
	})
}
