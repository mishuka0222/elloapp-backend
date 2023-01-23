package core

import (
	"errors"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) DeleteChannelUser(in *channels.DeleteChannelUserReq) (res *channels.DeleteChannelUserResp, err error) {
	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if in.OperatorId != in.Channel.Channel.CreatorUserId && in.OperatorId != in.DeleteUserId {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	err = c.checkOrLoadChannelParticipantList(in.Channel)
	if err != nil {
		return
	}

	var found = -1
	for i := range in.Channel.Participants {
		if in.DeleteUserId == in.Channel.Participants[i].UserId {
			if in.Channel.Participants[i].State == 0 {
				found = i
			}
			break
		}
	}

	if found == -1 {
		err = errors.New("PARTICIPANT_NOT_EXISTS")
		return
	}

	in.Channel.Participants[found].State = 1
	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.DeleteChannelUser(c.ctx, in.Channel.Participants[found].Id)
	if err != nil {
		return
	}

	// delete found.
	// this.participants = append(this.participants[:found], this.participants[found+1:]...)

	var now = int32(time.Now().Unix())
	in.Channel.Channel.ParticipantCount = int32(len(in.Channel.Participants) - 1)
	in.Channel.Channel.Version += 1
	in.Channel.Channel.Date = now
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateParticipantCount(c.ctx, in.Channel.Channel.ParticipantCount, now, in.Channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.DeleteChannelUserResp{}

	return
}
