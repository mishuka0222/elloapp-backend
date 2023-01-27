package core

import (
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) GetChannelParticipantList(in *channels.ChannelParticipantListReq) (res *channels.ChannelParticipantListResp, err error) {
	var (
		participantList []*mtproto.ChannelParticipant
		participants    []dataobject.ChannelParticipantDO
	)

	participants, err = c.svcCtx.Dao.ChannelParticipantsDAO.SelectByChannelId(c.ctx, in.ChannelId)
	if err != nil {
		return
	} else if len(participants) == 0 {
		err = errors.New(fmt.Sprintf("can`t load participants list for this channel: %d", in.ChannelId))
	}
	participantList = make([]*mtproto.ChannelParticipant, len(participants))
	for i := range participants {
		participantList[i], err = makeChannelParticipantByDO(&participants[i])
	}

	res = &channels.ChannelParticipantListResp{Participants: participantList}

	return
}
