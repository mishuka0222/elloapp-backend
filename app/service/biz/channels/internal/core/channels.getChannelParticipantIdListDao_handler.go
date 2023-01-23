package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) GetChannelParticipantIdListDao(in *channels.GetChannelParticipantIdListDaoReq) (res *channels.GetChannelParticipantIdListDaoResp, err error) {
	res = &channels.GetChannelParticipantIdListDaoResp{
		IdList: make([]int64, 0),
	}
	var (
		participants []dataobject.ChannelParticipantDO
	)

	participants, err = c.svcCtx.Dao.ChannelParticipantsDAO.SelectByChannelId(c.ctx, in.ChannelId)
	if err != nil {
		return
	}

	res.IdList = make([]int64, 0, len(participants))
	for i := range participants {
		if participants[i].State == 0 {
			res.IdList = append(res.IdList, participants[i].UserId)
		}
	}

	return
}
