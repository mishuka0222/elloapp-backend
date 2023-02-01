package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/username/username"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChannelsUpdateUsername
// channels.updateUsername#3514b3de channel:InputChannel username:string = Bool;
func (c *UsernamesCore) ChannelsUpdateUsername(in *mtproto.TLChannelsUpdateUsername) (res *mtproto.Bool, err error) {
	var (
		rOk *channels.CheckUserIsAdministratorResp
	)

	rOk, err = c.svcCtx.Dao.CheckUserIsAdministrator(c.ctx, &channels.CheckUserIsAdministratorReq{
		ChannelId: in.Channel.GetChannelId(),
		UserId:    c.MD.UserId,
	})
	if err != nil {
		return
	} else if !rOk.Status {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	if in.Username == "" {
		var channel *channels.ChannelData
		channel, err = c.svcCtx.ChannelsClient.GetChannelDataById(c.ctx, &channels.ChannelDataByIdReq{
			ChannelId: in.Channel.GetChannelId(),
		})
		if err != nil {
			return
		}
		c.svcCtx.Dao.UsernameClient.UsernameDeleteUsername(c.ctx, &username.TLUsernameDeleteUsername{
			Username: channel.Channel.Username,
		})
	}

	_, err = c.svcCtx.Dao.ChannelsClient.UpdateChannelLink(c.ctx, &channels.UpdateChannelLinkReq{
		ChannelId: in.Channel.GetChannelId(),
		Link:      in.Username,
	})
	if err != nil {
		return
	}

	res = mtproto.BoolTrue
	return
}
