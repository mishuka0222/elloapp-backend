package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserAddPeerSettings
// user.addPeerSettings user_id:int peer_type:int peer_id:int settings:PeerSettings = Bool;
func (c *UserCore) UserAddPeerSettings(in *user.TLUserAddPeerSettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.SetUserPeerSettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId(),
		in.GetSettings())

	if err != nil {
		c.Logger.Errorf("user.addPeerSettings - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
