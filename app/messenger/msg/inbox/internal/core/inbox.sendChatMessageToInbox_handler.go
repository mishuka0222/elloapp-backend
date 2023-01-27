package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// InboxSendChatMessageToInbox
// inbox.sendChatMessageToInbox from_id:long peer_chat_id:long message:InboxMessageData = Void;
func (c *InboxCore) InboxSendChatMessageToInbox(in *inbox.TLInboxSendChatMessageToInbox) (*mtproto.Void, error) {

	if mtproto.PeerIsChannel(in.Message.Message.PeerId) {
		res, err := c.svcCtx.Dao.ChannelsClient.GetChannelParticipantList(c.ctx,
			&channels.ChannelParticipantListReq{ChannelId: in.PeerChatId})
		if err != nil {
			c.Logger.Errorf("inbox.sendChannelMessageToInbox - error: %v", err)
			return nil, err
		}
		for _, v := range res.Participants {
			func(v *mtproto.ChannelParticipant) {
				if v.UserId == in.FromId {
					return
				}

				inBox, err := c.svcCtx.Dao.SendChannelMessageToInbox(
					c.ctx,
					in.FromId,
					in.PeerChatId,
					v.UserId,
					in.Message.DialogMessageId,
					in.Message.RandomId,
					in.Message.GetMessage())
				if err != nil {
					c.Logger.Errorf(err.Error())
					return
				}

				var (
					updates = []*mtproto.Update{mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
						Message_MESSAGE: inBox.Message,
						Pts_INT32:       inBox.Pts,
						PtsCount:        inBox.PtsCount,
					}).To_Update()}
				)

				if inBox.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChannelMigrateFrom {
					c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(
						c.ctx,
						inBox.MessageId,
						v.UserId,
						mtproto.MakePeerDialogId(mtproto.PEER_CHANNEL, in.PeerChatId))

					updates = append(updates, mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
						FolderId:         nil,
						Peer_PEER:        mtproto.MakePeerChannel(in.PeerChatId),
						MaxId:            inBox.MessageId,
						StillUnreadCount: 0,
						Pts_INT32:        c.svcCtx.Dao.NextChannelPtsId(c.ctx, v.UserId),
						PtsCount:         1,
					}).To_Update())
				}

				pushUpdates := mtproto.MakePushUpdates(
					func(idList []int64) []*mtproto.User {
						users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
							&userpb.TLUserGetMutableUsers{
								Id: idList,
							})
						return users.GetUserListByIdList(v.UserId, idList...)
					},
					func(idList []int64) []*mtproto.Chat {
						chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
							&chatpb.TLChatGetChatListByIdList{
								IdList: idList,
							})
						return chats.GetChatListByIdList(v.UserId, idList...)
					},
					func(idList []int64) []*mtproto.Chat {
						res, _ := c.svcCtx.Dao.ChannelsClient.GetChatsListBySelfAndIDList(c.ctx, &channels.GetChatsListBySelfAndIDListReq{
							SelfUserId: v.UserId,
							IdList:     idList,
						})
						return res.Chats
					},
					updates...)

				c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
					UserId:  v.UserId,
					Updates: pushUpdates,
				})
			}(v)

		}
		return mtproto.EmptyVoid, nil
	}

	_, err := c.svcCtx.Dao.ChatParticipantsDAO.SelectListWithCB(
		c.ctx,
		in.PeerChatId,
		func(i int, v *dataobject.ChatParticipantsDO) {
			if v.UserId == in.FromId {
				return
			}

			switch v.State {
			case mtproto.ChatMemberStateNormal:
			case mtproto.ChatMemberStateLeft:
				return
			case mtproto.ChatMemberStateKicked:
				if in.Message.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatDeleteUser &&
					in.Message.GetMessage().GetAction().GetUserId() == v.UserId {
					// send messageActionChatDeleteUser message to deleteUser
				} else {
					return
				}
			case mtproto.ChatMemberStateMigrated:
				if in.Message.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
					// messageActionChatMigrateTo to user
				} else {
					return
				}
			default:
				return
			}

			inBox, err := c.svcCtx.Dao.SendChatMessageToInbox(
				c.ctx,
				in.FromId,
				in.PeerChatId,
				v.UserId,
				in.Message.DialogMessageId,
				in.Message.RandomId,
				in.Message.GetMessage())
			if err != nil {
				c.Logger.Errorf(err.Error())
				return
			}

			var (
				updates = []*mtproto.Update{mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
					Message_MESSAGE: inBox.Message,
					Pts_INT32:       inBox.Pts,
					PtsCount:        inBox.PtsCount,
				}).To_Update()}
			)

			if inBox.GetMessage().GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
				c.svcCtx.Dao.DialogsDAO.UpdateReadInboxMaxId(
					c.ctx,
					inBox.MessageId,
					v.UserId,
					mtproto.MakePeerDialogId(mtproto.PEER_CHAT, in.PeerChatId))

				updates = append(updates, mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
					FolderId:         nil,
					Peer_PEER:        mtproto.MakePeerChat(in.PeerChatId),
					MaxId:            inBox.MessageId,
					StillUnreadCount: 0,
					Pts_INT32:        c.svcCtx.Dao.NextPtsId(c.ctx, v.UserId),
					PtsCount:         1,
				}).To_Update())
			}

			pushUpdates := mtproto.MakePushUpdates(
				func(idList []int64) []*mtproto.User {
					users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx,
						&userpb.TLUserGetMutableUsers{
							Id: idList,
						})
					return users.GetUserListByIdList(v.UserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					chats, _ := c.svcCtx.Dao.ChatClient.ChatGetChatListByIdList(c.ctx,
						&chatpb.TLChatGetChatListByIdList{
							IdList: idList,
						})
					return chats.GetChatListByIdList(v.UserId, idList...)
				},
				func(idList []int64) []*mtproto.Chat {
					res, _ := c.svcCtx.Dao.ChannelsClient.GetChatsListBySelfAndIDList(c.ctx, &channels.GetChatsListBySelfAndIDListReq{
						SelfUserId: v.UserId,
						IdList:     idList,
					})
					return res.Chats
				},
				updates...)

			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId:  v.UserId,
				Updates: pushUpdates,
			})
		})
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMessageToInbox - error: %v", err)
	}

	return mtproto.EmptyVoid, nil
}
