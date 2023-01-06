package core

import (
	"github.com/gogo/protobuf/types"
	msgpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"math/rand"
	"time"
)

func (c *VoipcallsCore) PhoneDiscardCall(in *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	var (
		peer = in.GetPeer().To_InputPhoneCall()
	)
	callSession, err := c.svcCtx.Dao.PhonecallClient.MakePhoneCallSessionByLoad(c.ctx, &phonecall.TLMakePhoneCallSessionByLoad{SessionId: peer.GetId()})
	if err != nil {
		c.Logger.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}

	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{callSession.AdminId, callSession.ParticipantId},
	})
	me, _ := users.GetImmutableUser(c.MD.UserId)
	added, _ := users.GetImmutableUser(callSession.AdminId)
	if me == nil || added == nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("voipcalls.phoneRequestCall - error: %v", err)
		return nil, err
	}

	phoneCallDiscarded := mtproto.MakeTLPhoneCallDiscarded(&mtproto.PhoneCall{
		Constructor: mtproto.CRC32_phoneCallDiscarded,
		Id:          callSession.Id,
		NeedDebug:   false,
		Reason:      in.GetReason(),
		Duration:    &types.Int32Value{Value: in.GetDuration()},
	})
	// TODO: delete this line
	c.Logger.Errorf("phoneCallDiscarded, err: %v", phoneCallDiscarded)

	//TODO: write logic
	action := mtproto.MakeTLMessageActionPhoneCall(&mtproto.MessageAction{
		CallId:   callSession.Id,
		Reason:   in.GetReason(),
		Duration: &types.Int32Value{Value: in.GetDuration()},
	}).To_MessageAction()

	if c.MD.UserId == callSession.ParticipantId {
		_, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
			UserId: callSession.AdminId,
			//AuthKeyId: c.MD.AuthId,
			PeerType: mtproto.PEER_USER,
			PeerId:   callSession.ParticipantId,
			Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
				NoWebpage:  true,
				Background: false,
				RandomId:   rand.Int63(),
				Message: mtproto.MakeTLMessageService(&mtproto.Message{
					Out:    true,
					FromId: mtproto.MakePeerUser(callSession.AdminId),
					PeerId: mtproto.MakePeerUser(callSession.ParticipantId),
					Date:   int32(time.Now().Unix()),
					Action: action,
				}).To_Message(),
			}).To_OutboxMessage(),
		})
		if err != nil {
			c.Logger.Errorf("PhoneDiscardCall error: %v", err)
		}
	} else {
		_, err := c.svcCtx.Dao.MsgClient.MsgSendMessage(c.ctx, &msgpb.TLMsgSendMessage{
			UserId: callSession.ParticipantId,
			//AuthKeyId: c.MD.AuthId,
			PeerType: mtproto.PEER_USER,
			PeerId:   callSession.AdminId,
			Message: msgpb.MakeTLOutboxMessage(&msgpb.OutboxMessage{
				NoWebpage:  true,
				Background: false,
				RandomId:   rand.Int63(),
				Message: mtproto.MakeTLMessageService(&mtproto.Message{
					Out:    true,
					FromId: mtproto.MakePeerUser(callSession.AdminId),
					PeerId: mtproto.MakePeerUser(callSession.ParticipantId),
					Date:   int32(time.Now().Unix()),
					Action: action,
				}).To_Message(),
			}).To_OutboxMessage(),
		})
		if err != nil {
			c.Logger.Errorf("PhoneDiscardCall error: %v", err)
		}
	}

	adminUpdatesData := (&mtproto.TLUpdatePhoneCall{
		Data2: &mtproto.Update{
			UserId:    callSession.AdminId,
			PhoneCall: phoneCallDiscarded.To_PhoneCall(),
		},
	}).To_Update()
	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			// TODO: check
			//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
			//	&userpb.TLUserGetMutableUsers{
			//		Id: idList,
			//	})
			return users.GetUserListByIdList(callSession.AdminId, idList...)
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
		adminUpdatesData)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		// MARK: From Participant to Admin
		UserId: callSession.AdminId,
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
			adminUpdatesData),
	})

	participantUpdatesData := (&mtproto.TLUpdatePhoneCall{
		Data2: &mtproto.Update{
			UserId:    callSession.ParticipantId,
			PhoneCall: phoneCallDiscarded.To_PhoneCall(),
		},
	}).To_Update()
	rUpdates = mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			// TODO: check
			//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
			//	&userpb.TLUserGetMutableUsers{
			//		Id: idList,
			//	})
			return users.GetUserListByIdList(callSession.ParticipantId, idList...)
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
		participantUpdatesData)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		// MARK: From Admin to Participant
		UserId: callSession.ParticipantId,
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
			participantUpdatesData),
	})

	updatePhoneCall := mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
		PhoneCall: phoneCallDiscarded.To_PhoneCall(),
	}).To_Update()

	return &mtproto.Updates{
		Updates: []*mtproto.Update{updatePhoneCall},
		Users:   []*mtproto.User{me.ToSelfUser(), added.ToSelfUser()},
	}, nil
}
