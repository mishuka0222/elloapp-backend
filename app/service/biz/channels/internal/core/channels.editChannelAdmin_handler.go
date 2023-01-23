package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"time"
)

func (c *ChannelsCore) EditChannelAdmin(in *channels.EditChannelAdminReq) (res *channels.EditChannelAdminResp, err error) {
	// operatorId is creator
	if in.OperatorId != in.Channel.Channel.CreatorUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	// editChatAdminId not creator
	if in.EditChannelAdminId == in.OperatorId {
		err = errors.New("BAD_REQUEST")
		return
	}

	err = c.checkOrLoadChannelParticipantList(in.Channel)
	if err != nil {
		return
	}

	// check exists
	_, participant := in.Channel.FindChatParticipant(in.EditChannelAdminId)
	if participant == nil || participant.State == 1 {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	if in.IsAdmin && participant.ParticipantType == kChannelParticipantAdmin || !in.IsAdmin && participant.ParticipantType == kChannelParticipant {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	if in.IsAdmin {
		participant.ParticipantType = kChannelParticipantAdmin
	} else {
		participant.ParticipantType = kChannelParticipant
	}
	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.UpdateParticipantType(c.ctx, int8(participant.ParticipantType), participant.Id)
	if err != nil {
		return
	}

	// update version
	in.Channel.Channel.Date = int32(time.Now().Unix())
	in.Channel.Channel.Version += 1
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateVersion(c.ctx, in.Channel.Channel.Date, in.Channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.EditChannelAdminResp{}

	return
}
