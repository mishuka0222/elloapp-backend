package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneRequestCall(in *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	var (
		err           error
		participantId int64
	)

	addUser := mtproto.FromInputUser(c.MD.UserId, in.UserId)
	switch addUser.PeerType {
	case mtproto.PEER_USER:
		participantId = in.GetUserId().GetUserId()
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("voipcalls.phoneRequestCall - error: %v", err)
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
		c.Logger.Errorf("voipcalls.phoneRequestCall - error: %v", err)
		return nil, err
	}

	rules, _ := c.svcCtx.Dao.UserClient.UserGetPrivacy(c.ctx, &userpb.TLUserGetPrivacy{
		UserId:  addUser.PeerId,
		KeyType: userpb.PHONE_CALL,
	})
	if len(rules.Datas) > 0 {
		// return true
		allowDoPhoneCall := userpb.CheckPrivacyIsAllow(
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
				return true
			})
		if !allowDoPhoneCall {
			err = mtproto.ErrUserPrivacyRestricted
			c.Logger.Errorf("not allow phone call: %v", err)
			return nil, err
		}
	}

	callSession, err := c.svcCtx.Dao.PhonecallClient.MakePhoneCallSession(c.ctx, &phonecall.TLMakePhoneCallSession{
		AdminId:       c.MD.UserId,
		ParticipantId: participantId,
		Ga:            in.GetGAHash(),
		Protocol:      in.GetProtocol(),
	})
	if err != nil {
		return nil, err
	}

	updatePhoneCall := (&mtproto.TLUpdatePhoneCall{
		Data2: &mtproto.Update{
			PhoneCall: callSession.ToPhoneCallRequested().To_PhoneCall(),
		},
	}).To_Update()
	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			// TODO: check
			//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
			//	&userpb.TLUserGetMutableUsers{
			//		Id: idList,
			//	})
			return users.GetUserListByIdList(c.MD.UserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			//if len(idList) > 0 {
			//	chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
			//		&chatpb.TLChatGetChatListByIdList{
			//			IdList: idList,
			//		})
			//	return chats.GetChatListByIdList(fromUserId, idList...)
			//} else {
			return []*mtproto.Chat{}
			//}
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updatePhoneCall)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId: participantId,
		//AuthKeyId: c.MD.PermAuthKeyId,
		Updates: mtproto.MakeSyncNotMeUpdates(
			func(idList []int64) []*mtproto.User {
				return rUpdates.Users
			},
			func(idList []int64) []*mtproto.Chat {
				return rUpdates.Chats
			},
			func(idList []int64) []*mtproto.Chat {
				// rUpdates.Chats include chats
				return nil
			},
			updatePhoneCall),
	})

	return mtproto.MakeTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: callSession.ToPhoneCallWaiting(c.MD.UserId, 0).To_PhoneCall(),
		Users:     []*mtproto.User{me.ToSelfUser(), added.ToSelfUser()},
	}).To_Phone_PhoneCall(), nil
}
