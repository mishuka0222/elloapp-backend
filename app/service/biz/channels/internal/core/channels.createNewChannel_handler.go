package core

import (
	"math/rand"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) CreateNewChannel(in *channels.CreateNewChannelReq) (res *channels.ChannelData, err error) {
	res = &channels.ChannelData{
		Channel: &channels.Channel{
			CreatorUserId:    in.CreatorId,
			AccessHash:       rand.Int63(),
			RandomId:         c.svcCtx.IDGenClient2.NextId(c.ctx),
			ParticipantCount: int32(1 + len(in.UserIdList)),
			Title:            in.Title,
			About:            in.About,
			PhotoId:          0,
			Version:          1,
			Date:             int32(time.Now().Unix()),
		},
		Participants: make([]*channels.ChannelParticipant, 1+len(in.UserIdList)),
	}
	res.Channel.Id, err = c.svcCtx.Dao.ChannelsDAO.Insert(c.ctx, ToChannelDO(res.Channel))
	if err != nil {
		return
	}

	res.Participants[0] = &channels.ChannelParticipant{}
	res.Participants[0].ChannelId = res.Channel.Id
	res.Participants[0].UserId = in.CreatorId
	res.Participants[0].ParticipantType = channels.K_ChannelParticipantCreator
	res.Participants[0].AdminRights = MakeTLChatAdminRights()
	_, err = c.svcCtx.Dao.ChannelParticipantsDAO.Insert(c.ctx, ToChannelParticipantDO(res.Participants[0]))
	if err != nil {
		return
	}

	for i := range in.UserIdList {
		res.Participants[i+1] = &channels.ChannelParticipant{}
		res.Participants[i+1].ChannelId = res.Channel.Id
		res.Participants[i+1].UserId = in.UserIdList[i]
		res.Participants[i+1].ParticipantType = channels.K_ChannelParticipant
		res.Participants[i+1].InviterUserId = in.CreatorId
		res.Participants[i+1].InvitedAt = res.Channel.Date
		res.Participants[i+1].JoinedAt = res.Channel.Date
		_, err = c.svcCtx.Dao.ChannelParticipantsDAO.Insert(c.ctx, ToChannelParticipantDO(res.Participants[i+1]))
		if err != nil {
			return
		}
	}

	return
}
