package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"time"
)

func (c *ChannelsCore) EditChannelAdmin(in *channels.EditChannelAdminReq) (res *channels.Void, err error) {
	var (
		channel        *channels.ChannelData
		strAdminRights string
	)

	channel, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		return
	}

	// operatorId is creator
	if in.OperatorId != channel.Channel.CreatorUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	// editChatAdminId not creator
	if in.EditChannelAdminId == in.OperatorId {
		err = errors.New("BAD_REQUEST")
		return
	}

	err = c.checkOrLoadChannelParticipantList(channel)
	if err != nil {
		return
	}

	// check exists
	_, participant := channel.FindChatParticipant(in.EditChannelAdminId)
	if participant == nil || participant.State == 1 {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	if in.IsAdmin && participant.ParticipantType == channels.K_ChannelParticipantAdmin ||
		!in.IsAdmin && participant.ParticipantType == channels.K_ChannelParticipant {
		err = errors.New("CHAT_NOT_MODIFIED")
		return
	}

	if in.IsAdmin {
		participant.ParticipantType = channels.K_ChannelParticipantAdmin
	} else {
		participant.ParticipantType = channels.K_ChannelParticipant
	}
	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.UpdateParticipantType(c.ctx, int8(participant.ParticipantType), participant.Id)
	if err != nil {
		return
	}

	if in.IsAdmin {
		strAdminRights, _ = jsonx.MarshalToString(defaultAdminRights)
	}
	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.UpdateAdminRights(c.ctx, strAdminRights, participant.Id)
	if err != nil {
		return
	}

	// update version
	channel.Channel.Date = int32(time.Now().Unix())
	channel.Channel.Version += 1
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateVersion(c.ctx, channel.Channel.Date, channel.GetChannelId())
	if err != nil {
		return
	}

	res = &channels.Void{}

	return
}
