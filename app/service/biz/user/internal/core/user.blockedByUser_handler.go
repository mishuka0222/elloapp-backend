package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UserBlockedByUser
// user.blockedByUser userId:long peer_user_id:long = Bool;
func (c *UserCore) UserBlockedByUser(in *user.TLUserBlockedByUser) (*mtproto.Bool, error) {
	blocked := c.svcCtx.Dao.CheckBlocked(c.ctx, in.GetUserId(), in.GetPeerUserId())

	return mtproto.ToBool(blocked), nil
}
