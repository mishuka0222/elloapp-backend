package core

import (
	"errors"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) EditChannelPhoto(in *channels.EditChannelPhotoReq) (res *channels.Void, err error) {
	err = c.checkOrLoadChannelParticipantList(in.Channel)
	if err != nil {
		return
	}

	_, participant := in.Channel.FindChatParticipant(in.EditUserId)

	if participant == nil || participant.State == 1 {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	// check editUserId is creator or admin
	if in.Channel.Channel.AdminsEnabled && participant.ParticipantType == channels.K_ChannelParticipant {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	if in.Channel.Channel.PhotoId == in.PhotoId {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	in.Channel.Channel.PhotoId = in.PhotoId
	in.Channel.Channel.Date = int32(time.Now().Unix())
	in.Channel.Channel.Version += 1

	_, err = c.svcCtx.Dao.ChannelsDAO.UpdatePhotoId(c.ctx, in.PhotoId, in.Channel.Channel.Date, in.Channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.Void{}

	return
}
