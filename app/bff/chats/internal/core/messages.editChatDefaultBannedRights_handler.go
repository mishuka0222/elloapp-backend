package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"time"
)

// MessagesEditChatDefaultBannedRights
// messages.editChatDefaultBannedRights#a5866b41 peer:InputPeer banned_rights:ChatBannedRights = Updates;
func (c *ChatsCore) MessagesEditChatDefaultBannedRights(in *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		rUpdates *mtproto.Updates
		date     = time.Now().Unix()
	)

	switch peer.PeerType {
	case mtproto.PEER_CHAT:
		chat, err := c.svcCtx.Dao.ChatClient.Client().ChatEditChatDefaultBannedRights(c.ctx, &chatpb.TLChatEditChatDefaultBannedRights{
			ChatId:       peer.PeerId,
			OperatorId:   c.MD.UserId,
			BannedRights: in.BannedRights,
		})
		if err != nil {
			c.Logger.Errorf("messages.editChatDefaultBannedRights - error: %v", err)
			return nil, err
		}

		defaultBannedUpdates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateChatDefaultBannedRights(&mtproto.Update{
				Peer_PEER:           mtproto.MakePeerChat(peer.PeerId),
				DefaultBannedRights: chat.Chat.DefaultBannedRights,
				Version:             chat.Chat.Version,
			}).To_Update(),
			Date: int32(date),
		}).To_Updates()

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
			UserId:    c.MD.UserId,
			AuthKeyId: c.MD.AuthId,
			Updates:   defaultBannedUpdates,
		})
		chat.Walk(func(userId int64, participant *chatpb.ImmutableChatParticipant) error {
			if userId != c.MD.UserId {
				c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
					UserId:  userId,
					Updates: defaultBannedUpdates,
				})
			}
			return nil
		})

		rUpdates = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{chat.ToUnsafeChat(c.MD.UserId)},
			Date:    int32(date),
			Seq:     0,
		}).To_Updates()
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("messages.editChatDefaultBannedRights blocked, License key from https://elloapp.com required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	return rUpdates, nil
}
