package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
)

func (c *ChannelsCore) CheckUserIsAdministrator(in *channels.CheckUserIsAdministratorReq) (res *channels.CheckUserIsAdministratorResp, err error) {
	var (
		participants []dataobject.ChannelParticipantDO
	)
	participants, err = c.loadChannelParticipantDoList(in.ChannelId)
	if err != nil {
		return
	}

	for i := range participants {
		if participants[i].UserId == in.UserId {
			if participants[i].ParticipantType == channels.K_ChannelParticipantCreator ||
				participants[i].ParticipantType == channels.K_ChannelParticipantAdmin {
				res = &channels.CheckUserIsAdministratorResp{Status: true}
				return
			}
		}

	}

	res = &channels.CheckUserIsAdministratorResp{Status: false}
	return
}
