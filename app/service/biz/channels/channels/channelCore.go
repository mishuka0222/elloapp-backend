package channels

import (
	"encoding/base64"
	"errors"
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto/crypto"
	"time"
)

const (
	kChannelParticipant        = 0
	kChannelParticipantSelf    = 1
	kChannelParticipantCreator = 2
	kChannelParticipantAdmin   = 3
	kChannelParticipantBanned  = 4
)

func makeChannelParticipantByDO(do *ChannelParticipant) (participant *mtproto.ChannelParticipant, err error) {
	participant = (&mtproto.ChannelParticipant{
		UserId:          do.UserId,
		InviterId_INT64: do.InviterUserId,
		Date:            do.JoinedAt,
	}).To_ChannelParticipant().To_ChannelParticipant()

	switch do.ParticipantType {
	case kChannelParticipant:
		participant.Constructor = mtproto.CRC32_channelParticipant
	case kChannelParticipantCreator:
		participant.Constructor = mtproto.CRC32_channelParticipantCreator
	case kChannelParticipantAdmin:
		participant.Constructor = mtproto.CRC32_channelParticipantAdmin
	default:
		err = errors.New(" channelParticipant type error")
	}

	return
}

func (ch *ChannelCoreData) GetPhotoId() int64 {
	return ch.Channel.PhotoId
}

func (ch *ChannelCoreData) GetChannelId() int32 {
	return ch.Channel.Id
}

func (ch *ChannelCoreData) GetVersion() int32 {
	return ch.Channel.Version
}

func (ch *ChannelCoreData) ExportedChatInvite() string {
	if ch.Channel.Link == "" {
		// TODO(@benqi): 检查唯一性
		ch.Channel.Link = "https://nebula.im/joinchat/" + base64.StdEncoding.EncodeToString(crypto.GenerateNonce(16))
		dao.GetChannelsDAO(dao.DB_MASTER).UpdateLink(ch.Channel.Link, int32(time.Now().Unix()), ch.Channel.Id)
	}
	return ch.Channel.Link
}

// TODO(@benqi): 性能优化
func (ch *ChannelCoreData) checkUserIsAdministrator(userId int32) bool {
	ch.checkOrLoadChannelParticipantList()
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].ParticipantType == kChannelParticipantCreator ||
			ch.Participants[i].ParticipantType == kChannelParticipantAdmin {
			return true
		}
	}
	return false
}

//func (ch *ChannelCoreData) checkOrLoadChannelParticipantList() {
//	if len(ch.Participants) == 0 {
//		ch.Participants = dao.GetChannelParticipantsDAO(dao.DB_SLAVE).SelectByChannelId(ch.Channel.Id)
//	}
//}

func (ch *ChannelCoreData) MakeMessageService(fromId int32, action *mtproto.MessageAction) *mtproto.Message {
	peer := &mtproto.PeerUtil{
		PeerType: mtproto.PEER_CHANNEL,
		PeerId:   ch.Channel.Id,
	}

	message := &mtproto.TLMessageService{Data2: &mtproto.Message{
		Date:   ch.Channel.Date,
		FromId: fromId,
		ToId:   peer.ToPeer(),
		Post:   true,
		Action: action,
	}}
	return message.To_Message()
}

func (ch *ChannelCoreData) MakeCreateChannelMessage(creatorId int32) *mtproto.Message {
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

func (ch *ChannelCoreData) MakeDeleteUserMessage(operatorId, channelUserId int32) *mtproto.Message {
	action := &mtproto.TLMessageActionChatDeleteUser{Data2: &mtproto.MessageAction{
		Title:  ch.Channel.Title,
		UserId: channelUserId,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelCoreData) MakeChannelEditTitleMessage(operatorId int32, title string) *mtproto.Message {
	action := &mtproto.TLMessageActionChatEditTitle{Data2: &mtproto.MessageAction{
		Title: title,
	}}

	return ch.MakeMessageService(operatorId, action.To_MessageAction())
}

func (ch *ChannelCoreData) GetChannelParticipantList() ([]*mtproto.ChannelParticipant, error) {
	ch.checkOrLoadChannelParticipantList()

	participantList := make([]*mtproto.ChannelParticipant, 0, len(ch.Participants))
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].State == 0 {
			pnt, err := makeChannelParticipantByDO(ch.Participants[i])
			if err != nil {
				return nil, err
			}
			participantList = append(participantList, pnt)
		}
	}
	return participantList, nil
}

func (ch *ChannelCoreData) GetChannelParticipantIdList() []int32 {
	ch.checkOrLoadChannelParticipantList()

	idList := make([]int32, 0, len(ch.Participants))
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].State == 0 {
			idList = append(idList, ch.Participants[i].UserId)
		}
	}
	return idList
}

func (ch *ChannelCoreData) GetChannelParticipants() *mtproto.TLChannelsChannelParticipants {
	ch.checkOrLoadChannelParticipantList()

	return &mtproto.TLChannelsChannelParticipants{Data2: &mtproto.Channels_ChannelParticipants{
		// ChatId: ch.Channel.Id,
		Participants: ch.GetChannelParticipantList(),
		// Version: ch.Channel.Version,
	}}
}

func (ch *ChannelCoreData) AddChannelUser(inviterId, userId int32) error {
	ch.checkOrLoadChannelParticipantList()

	// TODO(@benqi): check userId exisits
	var founded = -1
	for i := 0; i < len(ch.Participants); i++ {
		if userId == ch.Participants[i].UserId {
			if ch.Participants[i].State == 1 {
				founded = i
			} else {
				return fmt.Errorf("userId exisits")
			}
		}
	}

	var now = int32(time.Now().Unix())

	if founded != -1 {
		ch.Participants[founded].State = 0
		dao.GetChannelParticipantsDAO(dao.DB_MASTER).Update(inviterId, now, now, ch.Participants[founded].Id)
	} else {
		channelParticipant := &dataobject.ChannelParticipantsDO{
			ChannelId:       ch.Channel.Id,
			UserId:          userId,
			ParticipantType: kChannelParticipant,
			InviterUserId:   inviterId,
			InvitedAt:       now,
			JoinedAt:        now,
		}
		channelParticipant.Id = int32(dao.GetChannelParticipantsDAO(dao.DB_MASTER).Insert(channelParticipant))
		ch.Participants = append(ch.Participants, *channelParticipant)
	}

	// update chat
	ch.Channel.ParticipantCount += 1
	ch.Channel.Version += 1
	ch.Channel.Date = now
	dao.GetChannelsDAO(dao.DB_MASTER).UpdateParticipantCount(ch.Channel.ParticipantCount, now, ch.Channel.Id)

	return nil
}

func (ch *ChannelCoreData) findChatParticipant(selfUserId int32) (int, *dataobject.ChannelParticipantsDO) {
	for i := 0; i < len(ch.Participants); i++ {
		if ch.Participants[i].UserId == selfUserId {
			return i, &ch.Participants[i]
		}
	}
	return -1, nil
}

func (ch *ChannelCoreData) ToChannel(selfUserId int32) *mtproto.Chat {
	// TODO(@benqi): kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true

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
			Creator:       ch.Channel.CreatorUserId == selfUserId,
			Id:            ch.Channel.Id,
			Title:         ch.Channel.Title,
			AdminsEnabled: ch.Channel.AdminsEnabled == 1,
			// Photo:             mtproto.NewTLChatPhotoEmpty().To_ChatPhoto(),
			ParticipantsCount: ch.Channel.ParticipantCount,
			Date:              ch.Channel.Date,
			Version:           ch.Channel.Version,
		}}

		if ch.Channel.PhotoId == 0 {
			channel.SetPhoto(mtproto.NewTLChatPhotoEmpty().To_ChatPhoto())
		} else {
			sizeList, _ := nbfs_client.GetPhotoSizeList(ch.Channel.PhotoId)
			channel.SetPhoto(photo2.MakeChatPhoto(sizeList))
		}
		return channel.To_Chat()
	}
}

func (ch *ChannelCoreData) CheckDeleteChannelUser(operatorId, deleteUserId int32) error {
	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if operatorId != ch.Channel.CreatorUserId && operatorId != deleteUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	ch.checkOrLoadChannelParticipantList()
	var found = -1
	for i := 0; i < len(ch.Participants); i++ {
		if deleteUserId == ch.Participants[i].UserId {
			if ch.Participants[i].State == 0 {
				found = i
			}
			break
		}
	}

	if found == -1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	return nil
}

func (ch *ChannelCoreData) DeleteChannelUser(operatorId, deleteUserId int32) error {
	// operatorId is creatorUserId，allow delete all user_id
	// other delete me
	if operatorId != ch.Channel.CreatorUserId && operatorId != deleteUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	ch.checkOrLoadChannelParticipantList()
	var found = -1
	for i := 0; i < len(ch.Participants); i++ {
		if deleteUserId == ch.Participants[i].UserId {
			if ch.Participants[i].State == 0 {
				found = i
			}
			break
		}
	}

	if found == -1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	ch.Participants[found].State = 1
	dao.GetChannelParticipantsDAO(dao.DB_MASTER).DeleteChannelUser(ch.Participants[found].Id)

	// delete found.
	// ch.Participants = append(ch.Participants[:found], ch.Participants[found+1:]...)

	var now = int32(time.Now().Unix())
	ch.Channel.ParticipantCount = int32(len(ch.Participants) - 1)
	ch.Channel.Version += 1
	ch.Channel.Date = now
	dao.GetChannelsDAO(dao.DB_MASTER).UpdateParticipantCount(ch.Channel.ParticipantCount, now, ch.Channel.Id)

	return nil
}

func (ch *ChannelCoreData) EditChannelTitle(editUserId int32, title string) error {
	ch.checkOrLoadChannelParticipantList()

	_, participant := ch.findChatParticipant(editUserId)

	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	// check editUserId is creator or admin
	if ch.Channel.AdminsEnabled != 0 && participant.ParticipantType == kChannelParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	if ch.Channel.Title == title {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	ch.Channel.Title = title
	ch.Channel.Date = int32(time.Now().Unix())
	ch.Channel.Version += 1

	dao.GetChannelsDAO(dao.DB_MASTER).UpdateTitle(title, ch.Channel.Date, ch.Channel.Id)
	return nil
}

func (ch *ChannelCoreData) EditChannelPhoto(editUserId int32, photoId int64) error {
	ch.checkOrLoadChannelParticipantList()

	_, participant := ch.findChatParticipant(editUserId)

	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	// check editUserId is creator or admin
	if ch.Channel.AdminsEnabled != 0 && participant.ParticipantType == kChannelParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	if ch.Channel.PhotoId == photoId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	ch.Channel.PhotoId = photoId
	ch.Channel.Date = int32(time.Now().Unix())
	ch.Channel.Version += 1

	dao.GetChannelsDAO(dao.DB_MASTER).UpdatePhotoId(photoId, ch.Channel.Date, ch.Channel.Id)
	return nil
}

func (ch *ChannelCoreData) EditChannelAdmin(operatorId, editChannelAdminId int32, isAdmin bool) error {
	// operatorId is creator
	if operatorId != ch.Channel.CreatorUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	// editChatAdminId not creator
	if editChannelAdminId == operatorId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	ch.checkOrLoadChannelParticipantList()

	// check exists
	_, participant := ch.findChatParticipant(editChannelAdminId)
	if participant == nil || participant.State == 1 {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS)
	}

	if isAdmin && participant.ParticipantType == kChannelParticipantAdmin || !isAdmin && participant.ParticipantType == kChannelParticipant {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	if isAdmin {
		participant.ParticipantType = kChannelParticipantAdmin
	} else {
		participant.ParticipantType = kChannelParticipant
	}
	dao.GetChannelParticipantsDAO(dao.DB_MASTER).UpdateParticipantType(participant.ParticipantType, participant.Id)

	// update version
	ch.Channel.Date = int32(time.Now().Unix())
	ch.Channel.Version += 1
	dao.GetChannelsDAO(dao.DB_MASTER).UpdateVersion(ch.Channel.Date, ch.Channel.Id)

	return nil
}

func (ch *ChannelCoreData) ToggleChannelAdmins(userId int32, adminsEnabled bool) error {
	// check is creator
	if userId != ch.Channel.CreatorUserId {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION)
	}

	var (
		channelAdminsEnabled = ch.Channel.AdminsEnabled == 1
	)

	// Check modified
	if channelAdminsEnabled == adminsEnabled {
		return mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CHAT_NOT_MODIFIED)
	}

	ch.Channel.AdminsEnabled = base2.BoolToInt8(adminsEnabled)
	ch.Channel.Date = int32(time.Now().Unix())
	ch.Channel.Version += 1

	dao.GetChannelsDAO(dao.DB_MASTER).UpdateAdminsEnabled(ch.Channel.AdminsEnabled, ch.Channel.Date, ch.Channel.Id)

	return nil
}
