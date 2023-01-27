package core

import (
	"errors"
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	"math/rand"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsInviteToChannel(in *mtproto.TLChannelsInviteToChannel) (upd *mtproto.Updates, err error) {

	if in.Channel.Constructor == mtproto.CRC32_inputChannelEmpty {
		err = errors.New("BAD_REQUEST")
		return
	}

	var (
		inputChannel = in.GetChannel().To_InputChannel()
		channel      *channels.ChannelData
		peer         *mtproto.PeerUtil
		rOk          *channels.CheckUserIsAdministratorResp
	)

	rOk, err = c.svcCtx.Dao.CheckUserIsAdministrator(c.ctx, &channels.CheckUserIsAdministratorReq{
		ChannelId: in.Channel.GetChannelId(),
		UserId:    c.MD.UserId,
	})
	if err != nil {
		return
	} else if !rOk.Status {
		err = errors.New("NO_EDIT_CHAT_PERMISSION")
		return
	}

	channel, err = c.svcCtx.Dao.ChannelsClient.GetChannelDataById(c.ctx, &channels.ChannelDataByIdReq{ChannelId: inputChannel.GetChannelId()})
	if err != nil {
		return
	}

	peer = mtproto.MakeChannelPeerUtil(channel.Channel.Id)

	for _, u := range in.Users {
		if u.GetConstructor() == mtproto.CRC32_inputUserEmpty || u.GetConstructor() == mtproto.CRC32_inputUserSelf {
			continue
		}
		_, err = c.svcCtx.Dao.ChannelsClient.AddChannelParticipant(c.ctx, &channels.AddChannelParticipantReq{
			Channel:   channel,
			InviterId: c.MD.UserId,
			UserId:    u.GetUserId(),
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
				Message:      channel.MakeAddUserMessage(c.MD.UserId, u.GetUserId()),
				ScheduleDate: nil,
			}).To_OutboxMessage(),
		})
		if err != nil {
			return
		}

	}

	return
}
