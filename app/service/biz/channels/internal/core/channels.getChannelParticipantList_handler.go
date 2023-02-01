package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func contains(v int32, args ...int32) bool {
	for i := range args {
		if v == args[i] {
			return true
		}
	}
	return false
}

func (c *ChannelsCore) GetChannelParticipantList(in *channels.ChannelParticipantListReq) (res *channels.ChannelParticipantListResp, err error) {
	var (
		participantList []*mtproto.ChannelParticipant
		participants    []dataobject.ChannelParticipantDO
		participant     *mtproto.ChannelParticipant
	)
	if in.Types == nil || len(in.Types) == 0 {
		in.Types = []int32{
			channels.K_ChannelParticipant,
			channels.K_ChannelParticipantAdmin,
			channels.K_ChannelParticipantCreator,
		}
	}

	participants, err = c.svcCtx.Dao.ChannelParticipantsDAO.SelectByChannelId(c.ctx, in.ChannelId)
	if err != nil {
		return
	} else if len(participants) == 0 {
		err = errors.New(fmt.Sprintf("can`t load participants list for this channel: %d", in.ChannelId))
	}

	for i := range participants {
		if contains(participants[i].ParticipantType, in.Types...) {
			participant, err = makeChannelParticipantByDO(&participants[i])
			if err != nil {
				c.Logger.Error("channels.getChannelParticipantList, error: ", err.Error())
			} else {
				participantList = append(participantList, participant)
			}
		}
	}

	res = &channels.ChannelParticipantListResp{Participants: participantList}

	return
}
