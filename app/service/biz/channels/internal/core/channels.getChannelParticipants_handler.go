package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) GetChannelParticipants(in *channels.ChannelParticipantsReq) (res *channels.GetChannelParticipantsResp, err error) {
	var (
		r *channels.ChannelParticipantListResp
	)

	r, err = c.GetChannelParticipantList(&channels.ChannelParticipantListReq{
		ChannelId: in.ChannelId,
		Types:     in.Types,
	})
	if err != nil {
		return
	}

	res = &channels.GetChannelParticipantsResp{
		Participants: (&mtproto.Channels_ChannelParticipants{
			Count:        int32(len(r.Participants)),
			Participants: r.Participants,
			Chats:        []*mtproto.Chat{},
			Users:        []*mtproto.User{},
		}).To_ChannelsChannelParticipants()}

	return
}
