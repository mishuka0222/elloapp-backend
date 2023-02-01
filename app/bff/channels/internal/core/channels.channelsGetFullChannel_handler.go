package core

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsGetFullChannel(in *mtproto.TLChannelsGetFullChannel) (res *mtproto.Messages_ChatFull, err error) {

	if in.Channel.Constructor == mtproto.CRC32_inputChannelEmpty {
		err = errors.New("BAD_REQUEST")
		return
	}

	var (
		inputChannel = in.GetChannel().To_InputChannel()
		channel      *channels.ChannelData
		resFull      *channels.GetChannelFullBySelfIdResp
	)

	channel, err = c.svcCtx.Dao.ChannelsClient.GetAllChannelDataById(c.ctx, &channels.ChannelDataByIdReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	resFull, err = c.svcCtx.Dao.ChannelsClient.GetChannelFullBySelfId(c.ctx, &channels.GetChannelFullBySelfIdReq{
		ChannelData: channel,
		SelfUserId:  c.MD.UserId,
	})
	if err != nil {
		return
	}

	res = mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: resFull.Channel.To_ChatFull(),
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_ChatFull()

	str, _ := jsonx.MarshalToString(res)
	c.Logger.Debug("MY_DEBUG", str)

	return
}
