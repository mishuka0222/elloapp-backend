package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxSendChatMultiMessageToInbox
// inbox.sendChatMultiMessageToInbox from_id:long peer_chat_id:long message:Vector<InboxMessageData> = Void;
func (c *InboxCore) InboxSendChatMultiMessageToInbox(in *inbox.TLInboxSendChatMultiMessageToInbox) (*mtproto.Void, error) {
	chat, err := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
		ChatId: in.PeerChatId,
	})
	if err != nil {
		c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
		if in.FromId == userId {
			return nil
		}

		if !participant.IsChatMemberStateNormal() {
			return nil
		}

		inBoxList, err2 := c.svcCtx.Dao.SendChatMultiMessageToInbox(
			c.ctx,
			in.FromId,
			in.PeerChatId,
			userId,
			in.Message)
		if err2 != nil {
			c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err2)
			return nil
		}

		_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId:  userId,
			Updates: c.makeUpdateNewMessageListUpdates(userId, inBoxList...),
		})
		if err != nil {
			c.Logger.Errorf("inbox.sendChatMultiMessageToInbox - error: %v", err)
		}

		return nil
	})

	return mtproto.EmptyVoid, nil
}
