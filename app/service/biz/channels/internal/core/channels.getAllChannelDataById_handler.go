package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

// USE
func (c *ChannelsCore) GetAllChannelDataById(in *channels.ChannelDataByIdReq) (res *channels.ChannelData, err error) {
	var (
		channelDO *dataobject.ChannelDO
	)

	channelDO, err = c.svcCtx.Dao.ChannelsDAO.Select(c.ctx, in.ChannelId)
	if err != nil {
		return
	}

	res = &channels.ChannelData{Channel: channelDO.ToChannel()}

	err = c.checkOrLoadChannelParticipantList(res)
	if err != nil {
		return
	}

	return
}
