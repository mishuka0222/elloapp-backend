package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserGetPeerSettings
// user.getPeerSettings user_id:int peer_type:int peer_id:int = PeerSettings;
func (c *UserCore) UserGetPeerSettings(in *user.TLUserGetPeerSettings) (*mtproto.PeerSettings, error) {
	peerSettings, err := c.svcCtx.Dao.GetUserPeerSettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId())

	if err != nil {
		c.Logger.Errorf("user.getPeerSettings - error: %v", err)
		return nil, err
	}

	return peerSettings, nil
}
