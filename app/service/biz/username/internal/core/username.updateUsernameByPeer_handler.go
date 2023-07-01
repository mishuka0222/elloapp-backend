package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/pkg2/stores/sqlx"
)

// UsernameUpdateUsernameByPeer
// username.updateUsernameByPeer peer_type:int peer_id:int username:string = Bool;
func (c *UsernameCore) UsernameUpdateUsernameByPeer(in *username.TLUsernameUpdateUsernameByPeer) (*mtproto.Bool, error) {
	_, _, err := c.svcCtx.Dao.UsernameDAO.Insert(c.ctx, &dataobject.UsernameDO{
		Username: in.Username,
		PeerType: in.PeerType,
		PeerId:   in.PeerId,
	})

	if err != nil {
		if sqlx.IsDuplicate(err) {
			return mtproto.BoolFalse, nil
		} else {
			c.Logger.Errorf("username.updateUsername - error: %v", err)
			return nil, err
		}
	}

	return mtproto.BoolTrue, nil
}
