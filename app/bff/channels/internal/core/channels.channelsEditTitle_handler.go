package core

import (
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"math/rand"
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsEditTitle(in *mtproto.TLChannelsEditTitle) (upd *mtproto.Updates, err error) {
	if in.Title == "" {
		err := mtproto.ErrChatTitleEmpty
		c.Logger.Errorf("messages.editChannelTitle - error: ", err)
		return nil, err
	}
	var (
		peer = mtproto.MakeChannelPeerUtil(in.Channel.GetChannelId())
	)

	_, err = c.svcCtx.Dao.ChannelsClient.EditChannelTitle(c.ctx, &channels.EditChannelTitleReq{
		ChannelId:  peer.PeerId,
		EditUserId: c.MD.UserId,
		Title:      in.Title,
	})
	if err != nil {
		return
	}

	upd, err = c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  peer.PeerType,
		PeerId:    peer.PeerId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      channels.MakeMessageService(peer.PeerId, c.MD.UserId, int32(time.Now().Unix()), mtproto.MakeMessageActionChatEditTitle(in.Title)),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})
	if err != nil {
		c.Logger.Errorf("channels.editChannelTitle - error: %v", err)
		return nil, err
	}

	return
}
