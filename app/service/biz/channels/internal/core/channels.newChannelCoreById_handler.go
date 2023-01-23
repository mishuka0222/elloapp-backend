package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/internal/dao/dataobject"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) NewChannelCoreById(in *channels.ChannelCoreByIdReq) (res *channels.ChannelCoreData, err error) {
	var (
		channelDO *dataobject.ChannelDO
	)

	channelDO, err = c.svcCtx.Dao.ChannelsDAO.Select(c.ctx, in.ChannelId)
	if err != nil {
		return
	}

	res = &channels.ChannelCoreData{Channel: channelDO.ToChannel()}

	return
}
