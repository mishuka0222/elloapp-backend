package core

import (
	"encoding/base64"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
	"strings"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) UpdateChannelLink(in *channels.UpdateChannelLinkReq) (res *channels.UpdateChannelLinkResp, err error) {
	if in.Link == "" {
		in.Link = "https://elloapp.com/" + base64.StdEncoding.EncodeToString(crypto.GenerateNonce(16))
	}
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateLink(c.ctx, in.Link, int32(time.Now().Unix()), in.ChannelId)
	if err != nil {
		return
	}

	username_ := strings.Replace(in.Link, "https://elloapp.com/", "", 1)
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateUsername(c.ctx, username_, int32(time.Now().Unix()), in.ChannelId)
	if err != nil {
		return
	}

	_, err = c.svcCtx.UsernameService.UsernameUpdateUsername(c.ctx, &username.TLUsernameUpdateUsername{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   in.ChannelId,
		Username: username_,
	})
	if err != nil {
		return
	}

	res = &channels.UpdateChannelLinkResp{Link: in.Link}
	return
}
