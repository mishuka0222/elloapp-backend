package core

import (
	"errors"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) DeleteChannelUser(in *channels.DeleteChannelUserReq) (res *channels.Void, err error) {

	var (
		channel      *channels.ChannelData
		deletedUser  *channels.ChannelParticipant
		operatorUser *channels.ChannelParticipant
		now          = int32(time.Now().Unix())
	)

	channel, err = c.GetChannelDataById(&channels.ChannelDataByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		return
	}

	err = c.checkOrLoadChannelParticipantList(channel)
	if err != nil {
		return
	}

	_, deletedUser = channel.FindChatParticipant(in.DeleteUserId)
	_, operatorUser = channel.FindChatParticipant(in.OperatorId)

	if deletedUser.State != channels.K_ParticipantActiveState {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if !(in.OperatorId == channel.Channel.CreatorUserId || in.OperatorId == in.DeleteUserId ||
		operatorUser.ParticipantType == channels.K_ChannelParticipantAdmin) {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	if in.Reason == 0 {
		in.Reason = channels.K_ParticipantLeftState
	}

	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.DeleteChannelUser(c.ctx, in.Reason, deletedUser.Id, now)
	if err != nil {
		return
	}

	// delete found.
	// this.participants = append(this.participants[:found], this.participants[found+1:]...)

	channel.Channel.ParticipantCount = int32(len(channel.Participants) - 1)
	channel.Channel.Version += 1
	channel.Channel.Date = now
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateParticipantCount(c.ctx, channel.Channel.ParticipantCount, now, channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.Void{}

	return
}
