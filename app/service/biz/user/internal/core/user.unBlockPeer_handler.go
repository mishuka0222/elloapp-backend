package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserUnBlockPeer
// user.unBlockPeer user_id:long peer_type:int peer_id:long = Bool;
func (c *UserCore) UserUnBlockPeer(in *user.TLUserUnBlockPeer) (*mtproto.Bool, error) {
	unblocked := c.svcCtx.Dao.UnBlockUser(c.ctx, in.GetUserId(), in.GetPeerId())

	return mtproto.ToBool(unblocked), nil
}
