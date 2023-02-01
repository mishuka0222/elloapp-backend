package core

import (
	"errors"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) DeleteChannel(in *channels.DeleteChannelReq) (channel *channels.ChannelData, err error) {
	var (
		now = int32(time.Now().Unix())
	)

	channel, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		return
	}
	err = c.checkOrLoadChannelParticipantList(channel)
	if err != nil {
		return
	}

	if in.OperatorId != channel.Channel.CreatorUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	_, err = c.svcCtx.Dao.ChannelsDAO.DeleteChannel(c.ctx, now, in.ChannelId)
	if err != nil {
		return
	}

	channel.Channel.Date = int32(time.Now().Unix())
	channel.Channel.Version += 1
	channel.Channel.Deactivated = true

	return
}
