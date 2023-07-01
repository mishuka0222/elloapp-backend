package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// UsernameResolveUsername
// username.resolveUsername username:string = Peer;
func (c *UsernameCore) UsernameResolveUsername(in *username.TLUsernameResolveUsername) (*mtproto.Peer, error) {
	var (
		peer     *mtproto.Peer
		err      error
		username = in.Username
	)

	switch username {
	case "gif":
	case "pic":
	case "bing":
	default:
		if len(username) < 5 {
			err = mtproto.ErrUsernameInvalid
			return nil, err
		}
	}

	usernameDO, _ := c.svcCtx.Dao.UsernameDAO.SelectByUsername(c.ctx, username)
	if usernameDO == nil {
		c.Logger.Errorf("username.resolveUsername - error: %v", err)
		err = mtproto.ErrUsernameNotOccupied
		return nil, err
	}

	switch usernameDO.PeerType {
	case mtproto.PEER_USER:
		peer = mtproto.MakePeerUser(usernameDO.PeerId)
	case mtproto.PEER_CHANNEL:
		peer = mtproto.MakePeerChannel(usernameDO.PeerId)
	default:
		err = mtproto.ErrUsernameNotOccupied
		c.Logger.Errorf("username.resolveUsername - error: %v", err)
		return nil, err
	}

	return peer, nil
}
