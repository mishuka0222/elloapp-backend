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

		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      false,
		//Post:        false,
		Legacy: false,
		Id:     0,
		//FromId:      mtproto.MakePeerUser(fromId),
		//PeerId:      mtproto.MakePeerChat(m.Chat.Id),
		ReplyTo: nil,
		//Date:        int32(time.Now().Unix()),
		//Action:      action,
		TtlPeriod: nil,
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
