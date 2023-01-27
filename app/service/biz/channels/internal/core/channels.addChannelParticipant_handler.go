package core

import (
	"errors"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) AddChannelParticipant(in *channels.AddChannelParticipantReq) (res *channels.Void, err error) {
	err = c.checkOrLoadChannelParticipantList(in.Channel)
	if err != nil {
		return
	}

	var founded = -1
	for i := range in.Channel.Participants {
		if in.UserId == in.Channel.Participants[i].UserId {
			if in.Channel.Participants[i].State == 1 {
				founded = i
			} else {
				err = errors.New(" userId exists")
			}
		}
	}

	var now = int32(time.Now().Unix())

	if founded != -1 {
		in.Channel.Participants[founded].State = 0
		_, err = c.svcCtx.Dao.ChannelParticipantsDAO.Update(c.ctx, in.InviterId, now, now, in.Channel.Participants[founded].Id)
		if err != nil {
			return
		}
	} else {
		channelParticipant := &dataobject.ChannelParticipantDO{
			ChannelId:       in.Channel.Channel.Id,
			UserId:          in.UserId,
			ParticipantType: channels.K_ChannelParticipant,
			InviterUserId:   in.InviterId,
			InvitedAt:       now,
			JoinedAt:        now,
		}
		channelParticipant.Id, err = c.svcCtx.Dao.ChannelParticipantsDAO.Insert(c.ctx, channelParticipant)
		if err != nil {
			return
		}
		in.Channel.Participants = append(in.Channel.Participants, channelParticipant.ToChannelParticipant())
	}

	// update chat
	in.Channel.Channel.ParticipantCount += 1
	in.Channel.Channel.Version += 1
	in.Channel.Channel.Date = now
	_, err = c.svcCtx.Dao.ChannelsDAO.UpdateParticipantCount(c.ctx, in.Channel.Channel.ParticipantCount, now, in.Channel.Channel.Id)
	if err != nil {
		return
	}

	res = &channels.Void{}

	return
}
