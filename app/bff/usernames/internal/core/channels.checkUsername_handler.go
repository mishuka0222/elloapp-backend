package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChannelsCheckUsername
// channels.checkUsername#10e6bd2c channel:InputChannel username:string = Bool;
func (c *UsernamesCore) ChannelsCheckUsername(in *mtproto.TLChannelsCheckUsername) (res *mtproto.Bool, err error) {
	var (
		rOk *channels.CheckChannelUserNameResp
	)
	rOk, err = c.svcCtx.Dao.ChannelsClient.CheckChannelUserName(c.ctx, &channels.CheckChannelUserNameReq{
		ChannelId: in.Channel.GetChannelId(),
		UserName:  in.Username,
	})
	if err != nil {
		return
	}

	res = mtproto.ToBool(rOk.Status)
	return
}
