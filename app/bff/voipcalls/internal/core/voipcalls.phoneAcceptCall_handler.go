package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *VoipcallsCore) PhoneAcceptCall(in *mtproto.TLPhoneAcceptCall) (*mtproto.Phone_PhoneCall, error) {
	var (
		peer = in.GetPeer().To_InputPhoneCall()
	)
	callSession, err := c.svcCtx.Dao.PhonecallClient.MakePhoneCallSessionByLoad(c.ctx, &phonecall.TLMakePhoneCallSessionByLoad{SessionId: peer.GetId()})
	if err != nil {
		c.Logger.Errorf("invalid peer: {%v}, err: %v", peer, err)
		return nil, err
	}

	if peer.GetAccessHash() != callSession.ParticipantAccessHash {
		err = fmt.Errorf("invalid peer: {%v}", peer)
		c.Logger.Errorf("invalid peer: {%v}", peer)
		return nil, err
	}

	users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: []int64{c.MD.UserId, callSession.AdminId},
	})
	me, _ := users.GetImmutableUser(c.MD.UserId)
	added, _ := users.GetImmutableUser(callSession.AdminId)
	if me == nil || added == nil {
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("voipcalls.phoneRequestCall - error: %v", err)
		return nil, err
	}

	// cache g_b
	//callSession.SetGB(c.svcCtx.Dao ,in.GetGB())
	// TODO: write logic
	callSession.GB = in.GetGB()

	phoneCallAccepted := mtproto.MakeTLPhoneCallAccepted(&mtproto.PhoneCall{
		Constructor:   mtproto.CRC32_phoneCallAccepted,
		Id:            callSession.Id,
		Video:         true,
		AccessHash:    callSession.ParticipantAccessHash,
		Date:          int32(callSession.Date),
		AdminId:       callSession.AdminId,
		ParticipantId: callSession.ParticipantId,
		Protocol:      callSession.ToPhoneCallProtocol(),
		GB:            callSession.GB,
	})

	updatePhoneCall := (&mtproto.TLUpdatePhoneCall{
		Data2: &mtproto.Update{
			PhoneCall: phoneCallAccepted.To_PhoneCall(),
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
			updatePhoneCall),
	})

	return mtproto.MakeTLPhonePhoneCall(&mtproto.Phone_PhoneCall{
		PhoneCall: callSession.ToPhoneCallWaiting(c.MD.UserId, 0).To_PhoneCall(),
		Users:     []*mtproto.User{me.ToSelfUser(), added.ToSelfUser()},
	}).To_Phone_PhoneCall(), nil
}
