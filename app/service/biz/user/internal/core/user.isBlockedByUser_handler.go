package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserIsBlockedByUser
// user.isBlockedByUser user_id:long peer_user_id:long = Bool;
func (c *UserCore) UserIsBlockedByUser(in *user.TLUserIsBlockedByUser) (*mtproto.Bool, error) {
	blocked := c.svcCtx.Dao.CheckBlocked(c.ctx, in.GetUserId(), in.GetPeerUserId())

	return mtproto.ToBool(blocked), nil
}
