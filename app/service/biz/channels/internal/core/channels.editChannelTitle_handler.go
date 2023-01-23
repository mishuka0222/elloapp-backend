package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"time"
)

func (c *ChannelsCore) EditChannelTitle(in *channels.EditChannelTitleReq) (res *channels.EditChannelTitleResp, err error) {
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
	if in.Channel.Channel.AdminsEnabled != 0 && participant.ParticipantType == kChannelParticipant {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	if in.Channel.Channel.Title == in.Title {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	in.Channel.Channel.Title = in.Title
	in.Channel.Channel.Date = int32(time.Now().Unix())
	in.Channel.Channel.Version += 1

	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateTitle(c.ctx, in.Title, in.Channel.Channel.Date, in.Channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.EditChannelTitleResp{}

	return
}
