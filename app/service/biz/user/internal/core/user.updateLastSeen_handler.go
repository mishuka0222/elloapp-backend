package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserUpdateLastSeen
// user.updateLastSeen id:long last_seen_at:long expires:int = Bool;
func (c *UserCore) UserUpdateLastSeen(in *user.TLUserUpdateLastSeen) (*mtproto.Bool, error) {
	if in.GetId() > 0 {
		c.svcCtx.Dao.PutLastSeenAt(
			c.ctx,
			in.GetId(),
			in.GetLastSeenAt(),
			in.GetExpires())
	}

	return mtproto.BoolTrue, nil
}
