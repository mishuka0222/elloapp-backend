package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UserSetNotifySettings
// user.setNotifySettings user_id:int peer_type:int peer_id:int settings:PeerNotifySettings = Bool;
func (c *UserCore) UserSetNotifySettings(in *user.TLUserSetNotifySettings) (*mtproto.Bool, error) {
	err := c.svcCtx.Dao.SetUserPeerNotifySettings(
		c.ctx,
		in.GetUserId(),
		in.GetPeerType(),
		in.GetPeerId(),
		in.GetSettings())

	if err != nil {
		c.Logger.Errorf("user.setNotifySettings - error: %v", err)
		return nil, err
	}

	return mtproto.BoolTrue, nil
}
