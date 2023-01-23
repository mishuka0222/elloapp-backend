package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) GetChannelParticipants(in *channels.ChannelCoreData) (res *channels.GetChannelParticipantsResp, err error) {
	err = c.checkOrLoadChannelParticipantList(in)
	if err != nil {
		return
	}

	var (
		r *channels.GetChannelParticipantListResp
	)

	r, err = c.GetChannelParticipantList(in)
	if err != nil {
		return
	}

	res = &channels.GetChannelParticipantsResp{Participants: (&mtproto.Channels_ChannelParticipants{
		Participants: r.Participants,
	}).To_ChannelsChannelParticipants()}

	return
}
