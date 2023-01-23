package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) CheckUserIsAdministrator(in *channels.CheckUserIsAdministratorReq) (res *channels.CheckUserIsAdministratorResp, err error) {
	err = c.checkOrLoadChannelParticipantList(in.Channel)
	if err != nil {
		return
	}

	for i := range in.Channel.Participants {
		if in.Channel.Participants[i].UserId == in.UserId {
			if in.Channel.Participants[i].ParticipantType == kChannelParticipantCreator ||
				in.Channel.Participants[i].ParticipantType == kChannelParticipantAdmin {
				res = &channels.CheckUserIsAdministratorResp{Status: true}
				return
			}
		}

	}

	res = &channels.CheckUserIsAdministratorResp{Status: false}
	return
}
