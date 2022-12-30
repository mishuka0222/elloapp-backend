package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesSetTyping
// messages.setTyping#58943ee2 flags:# peer:InputPeer top_msg_id:flags.0?int action:SendMessageAction = Bool;
func (c *DialogsCore) MessagesSetTyping(in *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	var (
		peer = mtproto.FromInputPeer2(c.MD.UserId, in.GetPeer())
		date = int32(time.Now().Unix())
	)

	// TODO(@benqi): push chat/channel时机
	switch peer.PeerType {
	case mtproto.PEER_USER:
		updates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateUserTyping(&mtproto.Update{
				UserId: c.MD.UserId,
				Action: in.Action,
			}).To_Update(),
			Date: date,
		}).To_Updates()
		c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
			UserId:  peer.PeerId,
			Updates: updates,
		})
	case mtproto.PEER_CHAT:
		chat, err := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
			ChatId: peer.PeerId,
		})
		_ = err
		updates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			// updateChatUserTyping#86cadb6c chat_id:int from_id:Peer action:SendMessageAction = Update;
			Update: mtproto.MakeTLUpdateChatUserTyping(&mtproto.Update{
				ChatId_INT64: peer.PeerId,
				FromId:       mtproto.MakePeerUser(c.MD.UserId),
				Action:       in.Action,
			}).To_Update(),
			Date: date,
		}).To_Updates()
		chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
			if userId == c.MD.UserId {
				return nil
			}
			if participant.IsChatMemberStateNormal() {
				c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
					UserId:  userId,
					Updates: updates,
				})
			}
			return nil
		})
	case mtproto.PEER_CHANNEL:
		// Tips: disable updates
	}

	return mtproto.BoolTrue, nil
}
