package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) GetChannelParticipantIdList(in *channels.ChannelData) (res *channels.GetChannelParticipantIdListResp, err error) {
	err = c.checkOrLoadChannelParticipantList(in)
	if err != nil {
		return
	}

	var (
		participantList = make([]int64, 0, len(in.Participants))
	)

	for i := range in.Participants {
		if in.Participants[i].State == 0 {
			participantList = append(participantList, in.Participants[i].UserId)
		}
	}
	res = &channels.GetChannelParticipantIdListResp{IdList: participantList}

	return
}
