package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// UsernameDeleteUsernameByPeer
// username.deleteUsernameByPeer peer_type:int peer_id:int = Bool;
func (c *UsernameCore) UsernameDeleteUsernameByPeer(in *username.TLUsernameDeleteUsernameByPeer) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.UsernameDAO.Delete2(c.ctx, in.PeerType, in.PeerId)
	if err != nil {
		c.Logger.Errorf("username.deleteUsernameByPeer - error: %v")
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
