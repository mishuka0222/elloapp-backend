package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesEditChatAdmin
// messages.editChatAdmin#a85bd1c2 chat_id:long user_id:InputUser is_admin:Bool = Bool;
func (c *ChatsCore) MessagesEditChatAdmin(in *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	var (
		adminUser = mtproto.FromInputUser(c.MD.UserId, in.UserId)
	)

	if adminUser.PeerType != mtproto.PEER_USER ||
		adminUser.PeerId == c.MD.UserId {
		err := mtproto.ErrUserIdInvalid
		c.Logger.Errorf("messages.editChatAdmin - invalid user_id, err: %v", err)
		return nil, err
	}

	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatAdmin(c.ctx, &chatpb.TLChatEditChatAdmin{
		ChatId:          in.ChatId,
		OperatorId:      c.MD.UserId,
		EditChatAdminId: adminUser.PeerId,
		IsAdmin:         in.IsAdmin,
	})
	_ = chat
	if err != nil {
		c.Logger.Errorf("messages.editChatAdmin - error: ", err)
		return nil, err
	}

	var (
		idList []int64
	)

	updateChatParticipants := mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
		Participants_CHATPARTICIPANTS: chat.ToChatParticipants(0),
	}).To_Update()

	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if participant.IsChatMemberStateNormal() {
			idList = append(idList, userId)
		}
		return nil
	})

	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
	}

	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if !participant.IsChatMemberStateNormal() {
			return nil
		}

		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx,
			&sync.TLSyncPushUpdates{
				UserId: userId,
				Updates: mtproto.MakeUpdatesByUpdatesUsersChats(
					mUsers.GetUserListByIdList(userId, idList...),
					[]*mtproto.Chat{chat.ToUnsafeChat(userId)},
					updateChatParticipants),
			})

		return nil
	})

	return mtproto.BoolTrue, nil
}
