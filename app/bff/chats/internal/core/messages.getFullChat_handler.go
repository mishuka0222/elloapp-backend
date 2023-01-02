package core

import (
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/dialog/dialog"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessagesGetFullChat
// messages.getFullChat#aeb00b34 chat_id:long = messages.ChatFull;
func (c *ChatsCore) MessagesGetFullChat(in *mtproto.TLMessagesGetFullChat) (*mtproto.Messages_ChatFull, error) {
	chat, err := c.svcCtx.Dao.ChatClient.Client().ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: in.ChatId,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: %v", err)
		return nil, err
	}

	me, ok := chat.GetImmutableChatParticipant(c.MD.UserId)
	if !ok {
		c.Logger.Errorf("messages.getFullChat - error: not participant{chat_id: %d, chat_participant_id: %d}", in.ChatId, c.MD.UserId)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	dialog2, err := c.svcCtx.Dao.DialogClient.DialogGetDialogsByIdList(c.ctx, &dialog.TLDialogGetDialogsByIdList{
		UserId: c.MD.UserId,
		IdList: []int64{mtproto.MakePeerDialogId(mtproto.PEER_CHAT, in.ChatId)},
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: %v", err)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	} else if len(dialog2.Datas) == 0 {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	dlg := dialog2.Datas[0].GetDialog()
	chatFull := mtproto.MakeTLChatFull(&mtproto.ChatFull{
		CanSetUsername:                       true,
		HasScheduled:                         false, // TODO
		Id:                                   chat.Id(),
		About:                                chat.About(),
		Participants:                         chat.ToChatParticipants(c.MD.UserId),
		ChatPhoto:                            chat.Photo(),
		NotifySettings:                       nil,
		ExportedInvite:                       nil, // TODO
		BotInfo:                              nil, // TODO
		PinnedMsgId:                          nil, // TODO
		FolderId:                             dlg.FolderId,
		Call:                                 chat.Call(),
		TtlPeriod:                            mtproto.MakeFlagsInt32(chat.TTLPeriod()), // TODO
		GroupcallDefaultJoinAs:               nil,                                      // TODO
		ThemeEmoticon:                        nil,                                      // TODO
		RequestsPending:                      nil,                                      // TODO
		RecentRequesters:                     nil,                                      // TODO
		AvailableReactions_FLAGVECTORSTRING:  chat.GetChat().GetAvailableReactions(),
		AvailableReactions_FLAGCHATREACTIONS: chat.AvailableReactions(),
	}).To_ChatFull()

	var (
		idList []int64
	)

	// NotifySettings
	if settings, _ := c.svcCtx.Dao.UserClient.UserGetNotifySettings(c.ctx, &userpb.TLUserGetNotifySettings{
		UserId:   c.MD.UserId,
		PeerType: mtproto.PEER_CHAT,
		PeerId:   in.ChatId,
	}); settings != nil {
		chatFull.NotifySettings = settings
	}

	if me.GetAdminRights().GetInviteUsers() {
		if me.Link != "" {
			chatFull.ExportedInvite, _ = c.svcCtx.Dao.ChatClient.Client().ChatGetExportedChatInvite(c.ctx, &chatpb.TLChatGetExportedChatInvite{
				ChatId: in.ChatId,
				Link:   me.Link,
			})
		}
	}

	rValue := mtproto.MakeTLMessagesChatFull(&mtproto.Messages_ChatFull{
		FullChat: chatFull,
		Chats:    []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
		Users:    nil,
	}).To_Messages_ChatFull()

	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if participant.IsChatMemberStateNormal() {
			idList = append(idList, participant.UserId)
		}
		return nil
	})

	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("messages.getFullChat - error: not found dialog")
	}
	rValue.Users = mUsers.GetUserListByIdList(c.MD.UserId, idList...)

	return rValue, nil
}
