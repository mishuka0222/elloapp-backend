package channels

import (
	"errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

const (
	K_ChannelParticipant     = 0
	K_ChannelParticipantSelf = iota
	K_ChannelParticipantCreator
	K_ChannelParticipantAdmin
	K_ChannelParticipantBanned
	K_ChannelParticipantLeft
	K_ChannelParticipantKicked
)

const (
	K_ParticipantActiveState = 0
	K_ParticipantLeftState   = 1
)

func (ptc *ChannelParticipant) MakeChannelParticipant(selfId int64) (participant *mtproto.ChannelParticipant, err error) {

	participant = &mtproto.ChannelParticipant{
		UserId:          ptc.UserId,
		InviterId_INT64: ptc.InviterUserId,
		Date:            ptc.JoinedAt,
	}

	switch ptc.ParticipantType {
	case K_ChannelParticipant:
		if ptc.UserId == selfId {
			participant.Constructor = mtproto.CRC32_channelParticipantSelf
		} else {
			participant.Constructor = mtproto.CRC32_channelParticipant
		}
	case K_ChannelParticipantCreator:
		participant.Constructor = mtproto.CRC32_channelParticipantCreator
	case K_ChannelParticipantAdmin:
		participant.Constructor = mtproto.CRC32_channelParticipantAdmin
	case K_ChannelParticipantBanned | K_ChannelParticipantKicked:
		participant.Constructor = mtproto.CRC32_channelParticipantBanned
	case K_ChannelParticipantLeft:
		participant.Constructor = mtproto.CRC32_channelParticipantLeft
	default:
		err = errors.New(" channelParticipant type error.")
		return
	}

	return
}
func (ptc *ChannelParticipant) ToChatParticipant() *mtproto.ChatParticipant {
	switch ptc.ParticipantType {
	case K_ChannelParticipantCreator:
		return mtproto.MakeTLChatParticipantCreator(&mtproto.ChatParticipant{
			UserId: ptc.UserId,
		}).To_ChatParticipant()
	case K_ChannelParticipantAdmin:
		return mtproto.MakeTLChatParticipantAdmin(&mtproto.ChatParticipant{
			UserId:    ptc.UserId,
			InviterId: ptc.InviterUserId,
			Date:      ptc.InvitedAt,
		}).To_ChatParticipant()
	default:
		return mtproto.MakeTLChatParticipant(&mtproto.ChatParticipant{
			UserId:    ptc.UserId,
			InviterId: ptc.InviterUserId,
			Date:      ptc.InvitedAt,
		}).To_ChatParticipant()
	}
}

func (ptc *ChannelParticipant) IsChatMemberStateNormal() bool {
	if ptc.ParticipantType == K_ChannelParticipantCreator ||
		ptc.ParticipantType == K_ChannelParticipantAdmin ||
		ptc.ParticipantType == K_ChannelParticipant {
		return true
	}
	return false
}

func (ptc *ChannelParticipant) AdminRightsToStr() (res string) {
	if ptc.AdminRights != nil {
		res, _ = jsonx.MarshalToString(ptc.AdminRights)
	}
	if res == "" {
		return "{}"
	}
	return
}

func (ch *ChannelData) MakeMessageService(fromId int64, action *mtproto.MessageAction) *mtproto.Message {
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

func MakeMessageService(channelId, fromId int64, date int32, action *mtproto.MessageAction) *mtproto.Message {
	peerFrom := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_USER,
		PeerId:   fromId,
	}
	peerTo := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   channelId,
	}

	message := &mtproto.Message{
		Date:   date,
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

func (ch *ChannelData) MakeCreateChannelMessage(creatorId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChannelCreate{Data2: &mtproto.MessageAction{
		Title: ch.Channel.Title,
	}}
	return ch.MakeMessageService(creatorId, action.To_MessageAction())
}

func (ch *ChannelData) MakeAddUserMessage(inviterId, channelUserId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChatAddUser{Data2: &mtproto.MessageAction{
		Title: ch.Channel.Title,
		Users: []int64{channelUserId},
	}}

	return ch.MakeMessageService(inviterId, action.To_MessageAction())
}

func (ch *ChannelData) MakeDeleteUserMessage(operatorId, channelUserId int64) *mtproto.Message {
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction{
		Title:  ch.Channel.Title,
		UserId: channelUserId,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelData) MakeChannelEditTitleMessage(operatorId int64, title string) *mtproto.Message {
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction{
		Title: title,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelData) GetNormalParticipants() (participants *mtproto.ChatParticipants, idList []int64) {
	participants = mtproto.MakeTLChatParticipants(&mtproto.ChatParticipants{
		ChatId:       ch.GetChannelId(),
		Participants: make([]*mtproto.ChatParticipant, 0, len(ch.Participants)),
		Version:      ch.Channel.Version,
	}).To_ChatParticipants()

	for _, cp := range ch.Participants {
		if cp.IsChatMemberStateNormal() {
			participants.Participants = append(participants.Participants, cp.ToChatParticipant())
			idList = append(idList, cp.Id)
		}
	}

	return
}

func (ch *ChannelData) FindChatParticipant(selfUserId int64) (int, *ChannelParticipant) {
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].UserId == selfUserId {
			return i, ch.Participants[i]
		}
	}
	return -1, nil
}

func (ch *ChannelData) AdminCount() (adminCount int32) {
	for i := range ch.Participants {
		if ch.Participants[i].ParticipantType == K_ChannelParticipantCreator ||
			ch.Participants[i].ParticipantType == K_ChannelParticipantAdmin {
			adminCount++
		}
	}
	return
}

func (ch *ChannelData) ParticipantCount() (participantCount int32) {
	participantCount = int32(len(ch.Participants))
	return
}

func (ch *ChannelData) GetChannelId() int64 {
	return ch.Channel.Id
}

func (ch *ChannelData) GetPhotoId() int64 {
	return ch.Channel.PhotoId
}

func (ch *ChannelData) GetTitle() string {
	return ch.Channel.Title
}
