package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"time"
)

func (c *ChannelsCore) EditChannelAbout(in *channels.EditChannelAboutReq) (res *channels.Void, err error) {
	var (
		channel *channels.ChannelData
	)

	channel, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		return
	}
	err = c.checkOrLoadChannelParticipantList(channel)
	if err != nil {
		return
	}

	_, participant := channel.FindChatParticipant(in.EditUserId)

	if participant == nil || participant.State != channels.K_ParticipantActiveState {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	if channel.Channel.About == in.About {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	channel.Channel.About = in.About
	channel.Channel.Date = int32(time.Now().Unix())
	channel.Channel.Version += 1

	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateAbout(c.ctx, in.About, channel.Channel.Date, channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.Void{}

	return
}
