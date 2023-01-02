package core

import (
	"math/rand"

	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesAddChatUser
// messages.addChatUser#f24753e3 chat_id:long user_id:InputUser fwd_limit:int = Updates;
func (c *ChatsCore) MessagesAddChatUser(in *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	var (
		err     error
		addUser = mtproto.FromInputUser(c.MD.UserId, in.UserId)
		chat    *chatpb.MutableChat
	)

	switch addUser.PeerType {
	case mtproto.PEER_USER:
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.addChatUser - error: %v", err)
		return nil, err
	}

	// 400	USERS_TOO_MUCH	The maximum number of users has been exceeded (to create a chat, for example)
	// 400	USER_ALREADY_PARTICIPANT	The user is already in the group
	// 400	USER_ID_INVALID	The provided user ID is invalid
	// 403	USER_NOT_MUTUAL_CONTACT	The provided user is not a mutual contact
	// 403	USER_PRIVACY_RESTRICTED	The user's privacy settings do not allow you to do this
	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, addUser.PeerId},
	})

	me, _ := users.GetImmutableUser(c.MD.UserId)
	added, _ := users.GetImmutableUser(addUser.PeerId)
	if me == nil || added == nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("messages.addChatUser - error: %v", err)
		return nil, err
	}

	rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
		UserId:  addUser.PeerId,
		KeyType: userpb.CHAT_INVITE,
	})
	if len(rules.Datas) > 0 {
		// return true
		allowAddChat := userpb.CheckPrivacyIsAllow(
			addUser.PeerId,
			rules.Datas,
			c.MD.UserId,
			func(id, checkId int64) bool {
				contact, _ := c.svcCtx.Dao.UserClient.UserCheckContact(c.ctx, &userpb.TLUserCheckContact{
					UserId: id,
					Id:     checkId,
				})
				return mtproto.FromBool(contact)
			},
			func(checkId int64, idList []int64) bool {
				chatIdList, _ := mtproto.SplitChatAndChannelIdList(idList)
				return c.svcCtx.Dao.ChatClient.CheckParticipantIsExist(c.ctx, checkId, chatIdList)
			})
		if !allowAddChat {
			err = mtproto.ErrUserPrivacyRestricted
			c.Logger.Errorf("not allow addChat: %v", err)
			return nil, err
		}
	}

	chat, err = c.svcCtx.Dao.ChatClient.Client().ChatAddChatUser(c.ctx, &chatpb.TLChatAddChatUser{
		ChatId:    in.ChatId,
		InviterId: c.MD.UserId,
		UserId:    addUser.PeerId,
	})
	//request.ChatId, md.UserId, peer.PeerId)
	if err != nil {
		c.Logger.Errorf("addChatUser error: %v", err)
		return nil, err
	}

	rUpdates, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		PeerType:  mtproto.PEER_CHAT,
		PeerId:    in.ChatId,
		Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(c.MD.UserId, mtproto.MakeMessageActionChatAddUser(addUser.PeerId)),
			ScheduleDate: nil,
		}).To_OutboxMessage(),
	})

	if err != nil {
		c.Logger.Errorf("addChatUser error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}
