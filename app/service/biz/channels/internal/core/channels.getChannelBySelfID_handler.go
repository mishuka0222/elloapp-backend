package core

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) GetChannelBySelfID(in *channels.GetChannelBySelfIDReq) (res *channels.GetChannelBySelfIDResp, err error) {
	res = &channels.GetChannelBySelfIDResp{}
	var (
		channelData *channels.ChannelCoreData
		r           *channels.ToChannelResp
	)
	channelData, err = c.NewChannelCoreById(&channels.ChannelCoreByIdReq{ChannelId: in.ChannelId})
	if err != nil {
		logx.Error("GetChannelBySelfID - not find chat_id: ", in.ChannelId, err.Error())
		channelEmpty := (&mtproto.Chat{
			Id: in.ChannelId,
		}).To_ChatEmpty()
		res.Chat = channelEmpty.To_Chat()
	} else {
		r, _ = c.ToChannel(&channels.ToChannelReq{Channel: channelData, SelfUserId: in.SelfUserId})
		res.Chat = r.Chat
	}

	return
}
