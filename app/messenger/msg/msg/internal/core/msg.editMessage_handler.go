package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"

	"github.com/gogo/protobuf/types"
)

// MsgEditMessage
// msg.editMessage user_id:long auth_key_id:long peer_type:int peer_id:long message:OutboxMessage = Updates;
func (c *MsgCore) MsgEditMessage(in *msg.TLMsgEditMessage) (*mtproto.Updates, error) {
	var (
		err      error
		rUpdates *mtproto.Updates
	)

	if in.Message.Message.EditDate == nil {
		in.Message.Message.EditDate = &types.Int32Value{Value: int32(time.Now().Unix())}
	}

	switch in.PeerType {
	case mtproto.PEER_USER:
		rUpdates, err = c.editUserOutgoingMessage(in)
	case mtproto.PEER_CHAT:
		rUpdates, err = c.editChatOutgoingMessage(in)
	case mtproto.PEER_CHANNEL:
		// rUpdates, err = c.editChannelOutgoingMessage(in)
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}

func (c *MsgCore) editUserOutgoingMessage(in *msg.TLMsgEditMessage) (*mtproto.Updates, error) {
	outBox, err := c.svcCtx.Dao.EditUserOutboxMessage(c.ctx, in.UserId, in.PeerId, in.Message.Message)
	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	if in.UserId != in.PeerId {
		c.svcCtx.Dao.InboxClient.InboxEditUserMessageToInbox(c.ctx, &inbox.TLInboxEditUserMessageToInbox{
			FromId:     in.UserId,
			PeerUserId: in.PeerId,
			Message:    outBox.Message,
		})
	}

	updateEditMessage := mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		Message_MESSAGE: outBox.Message,
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(in.UserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(in.UserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateEditMessage)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:    in.UserId,
			AuthKeyId: in.AuthKeyId,
			Updates: mtproto.MakeSyncNotMeUpdates(
				func(idList []int64) []*mtproto.User {
					return rUpdates.Users
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return rUpdates.Chats
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return rUpdates.Chats
				},
				updateEditMessage),
		})

	return rUpdates, nil
}

func (c *MsgCore) editChatOutgoingMessage(in *msg.TLMsgEditMessage) (*mtproto.Updates, error) {
	outBox, err := c.svcCtx.Dao.EditChatOutboxMessage(c.ctx, in.UserId, in.PeerId, in.Message.Message)
	if err != nil {
		c.Logger.Errorf("msg.editMessage - error: %v", err)
		return nil, err
	}

	c.svcCtx.Dao.InboxClient.InboxEditChatMessageToInbox(c.ctx, &inbox.TLInboxEditChatMessageToInbox{
		FromId:     in.UserId,
		PeerChatId: in.PeerId,
		Message:    outBox.Message,
	})

	updateEditMessage := mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		Message_MESSAGE: outBox.Message,
	}).To_Update()

	rUpdates := mtproto.MakeReplyUpdates(
		func(idList []int64) []*mtproto.User {
			users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
				&userpb.TLUserGetMutableUsers{
					Id: idList,
				})
			return users.GetUserListByIdList(in.UserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
				&chatpb.TLChatGetChatListByIdList{
					IdList: idList,
				})
			return chats.GetChatListByIdList(in.UserId, idList...)
		},
		func(idList []int64) []*mtproto.Chat {
			// TODO
			return nil
		},
		updateEditMessage)

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:    in.UserId,
			AuthKeyId: in.AuthKeyId,
			Updates: mtproto.MakeSyncNotMeUpdates(
				func(idList []int64) []*mtproto.User {
					return rUpdates.Users
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return rUpdates.Chats
				},
				func(idList []int64) []*mtproto.Chat {
					// TODO
					return rUpdates.Chats
				},
				updateEditMessage),
		})

	return rUpdates, nil
}
