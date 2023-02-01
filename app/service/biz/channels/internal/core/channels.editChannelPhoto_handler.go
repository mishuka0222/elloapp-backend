package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) EditChannelPhoto(in *channels.EditChannelPhotoReq) (channel *channels.ChannelData, err error) {
	var (
		strPhoto string
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

	if participant == nil || participant.State == 1 {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	// check editUserId is creator or admin
	if channel.Channel.AdminsEnabled && participant.ParticipantType == channels.K_ChannelParticipant {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	if channel.Channel.Photo.Id == in.Photo.Id {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	channel.Channel.Photo = in.Photo
	channel.Channel.Date = int32(time.Now().Unix())
	channel.Channel.Version += 1

	strPhoto, err = jsonx.MarshalToString(in.Photo)

	_, err = c.svcCtx.Dao.ChannelsDAO.UpdatePhoto(c.ctx, strPhoto, channel.Channel.Date, channel.Channel.Id)
	if err != nil {
		return
	}

	return
}
