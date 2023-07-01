package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	// "errors"
	"math/rand"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (upd *mtproto.Updates, err error) {
	var (
		channelUserIdList = []int64{
			//c.MD.UserId,
		}
		channel *channels.ChannelData
	)

	channel, err = c.svcCtx.Dao.ChannelsClient.CreateNewChannel(c.ctx, &channels.CreateNewChannelReq{
		CreatorId:  c.MD.UserId,
		UserIdList: channelUserIdList,
		Title:      in.Title,
		About:      in.About,
	})
	if err != nil {
		return
	}

	peer := mtproto.MakeChannelPeerUtil(channel.Channel.Id)

	// TODO: add attach_data (chat and chat_participants)
	upd, err = c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      channel.MakeCreateChannelMessage(c.MD.UserId),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("channels.createChannel - error: %v", err)
		return
	}

	return
}
