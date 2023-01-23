package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) GetChannelParticipantList(in *channels.ChannelCoreData) (res *channels.GetChannelParticipantListResp, err error) {
	err = c.checkOrLoadChannelParticipantList(in)
	if err != nil {
		return
	}

	var (
		participantList = make([]*mtproto.ChannelParticipant, 0, len(in.Participants))
		ptn             *mtproto.ChannelParticipant
	)

	for i := range in.Participants {
		if in.Participants[i].State == 0 {
			ptn, err = makeChannelParticipantByDO(in.Participants[i].ToChannelParticipantDO())
			if err != nil {
				return
			}
			participantList = append(participantList, ptn)
		}
	}
	res = &channels.GetChannelParticipantListResp{Participants: participantList}

	return
}
