package core

import (
	"fmt"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/phone_call/phonecall"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"time"
)

func (c *VoipcallsCore) PhoneReceivedCall(in *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {

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

	updatePhoneCall := (&mtproto.TLUpdatePhoneCall{
		Data2: &mtproto.Update{
			PhoneCall: callSession.ToPhoneCallWaiting(callSession.AdminId, int32(time.Now().Unix())).To_PhoneCall(),
		},
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			// TODO: check
			//users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
			//	&userpb.TLUserGetMutableUsers{
			//		Id: idList,
			//	})
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
				Id: []int64{c.MD.UserId, callSession.AdminId},
			})
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

	return mtproto.ToBool(true), nil
}
