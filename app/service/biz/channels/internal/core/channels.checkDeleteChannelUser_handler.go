package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) CheckDeleteChannelUser(in *channels.CheckDeleteChannelUserReq) (res *channels.CheckDeleteChannelUserResp, err error) {
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

	res = &channels.CheckDeleteChannelUserResp{}

	return
}
