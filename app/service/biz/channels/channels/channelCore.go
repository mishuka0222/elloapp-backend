package channels

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (ch *ChannelCoreData) MakeMessageService(fromId int64, action *mtproto.MessageAction) *mtproto.Message {
	peerFrom := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_USER,
		PeerId:   fromId,
	}
	peerTo := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   ch.Channel.Id,
	}

	message := &mtproto.Message{
		Date:   ch.Channel.Date,
		FromId: peerFrom.ToPeer(),
		PeerId: peerTo.ToPeer(),
		Post:   true,
		Action: action,
	}
	return message.To_MessageService().To_Message()
}

func (ch *ChannelCoreData) MakeCreateChannelMessage(creatorId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChannelCreate{Data2: &mtproto.MessageAction{
		Title: ch.Channel.Title,
	}}
	return ch.MakeMessageService(creatorId, action.To_MessageAction())
}

func (ch *ChannelCoreData) MakeAddUserMessage(inviterId, channelUserId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChatAddUser{Data2: &mtproto.MessageAction{
		Title: ch.Channel.Title,
		Users: []int64{channelUserId},
	}}

	return ch.MakeMessageService(inviterId, action.To_MessageAction())
}

func (ch *ChannelCoreData) MakeDeleteUserMessage(operatorId, channelUserId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction{
		Title:  ch.Channel.Title,
		UserId: channelUserId,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelCoreData) MakeChannelEditTitleMessage(operatorId int64, title string) *mtproto.Message {
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction{
		Title: title,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelCoreData) FindChatParticipant(selfUserId int64) (int, *ChannelParticipant) {
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].UserId == selfUserId {
			return i, ch.Participants[i]
		}
	}
	return -1, nil
}

func (ch *ChannelCoreData) ToChannel(selfUserId int64) *mtproto.Chat {

	var forbidden = false
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].UserId == selfUserId && ch.Participants[i].State == 1 {
			forbidden = true
			break
		}
	}

	if forbidden {
		channel := &mtproto.TLChannelForbidden{Data2: &mtproto.Chat{
			Id:    ch.Channel.Id,
			Title: ch.Channel.Title,
		}}
		return channel.To_Chat()
	} else {
		channel := &mtproto.TLChannel{Data2: &mtproto.Chat{
			Creator:     ch.Channel.CreatorUserId == selfUserId,
			Id:          ch.Channel.Id,
			Title:       ch.Channel.Title,
			AdminRights: &mtproto.ChatAdminRights{AddAdmins: ch.Channel.AdminsEnabled == 1},
			// Photo:             mtproto.NewTLChatPhotoEmpty().To_ChatPhoto(),
			ParticipantsCount_INT32: ch.Channel.ParticipantCount,
			Date:                    ch.Channel.Date,
			Version:                 ch.Channel.Version,
		}}

		// TODO: write logic
		/*if ch.Channel.PhotoId == 0 {
			channel.SetPhoto(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
		} else {


			sizeList, _ := nbfs_client.GetPhotoSizeList(ch.Channel.PhotoId)
			channel.SetPhoto(photo2.MakeChatPhoto(sizeList))
		}*/
		channel.SetPhoto(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
		return channel.To_Chat()
	}
}
