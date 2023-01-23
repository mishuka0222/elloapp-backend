package core

import (
	"encoding/base64"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) ExportedChatInvite(in *channels.ChannelCoreData) (res *channels.ExportedChatInviteResp, err error) {
	if in.Channel.Link == "" {
		in.Channel.Link = "https://elloapp.org/joinchat/" + base64.StdEncoding.EncodeToString(crypto.GenerateNonce(16))
		_, err = c.svcCtx.Dao.ChannelsDAO.UpdateLink(c.ctx, in.Channel.Link, int32(time.Now().Unix()), in.Channel.Id)
		if err != nil {
			return
		}
	}
	res = &channels.ExportedChatInviteResp{Link: in.Channel.Link}
	return
}
