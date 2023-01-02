package core

import (
	"time"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	chatpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	userpb "gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/user/user"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesSaveDraft
// messages.saveDraft#bc39e14b flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int peer:InputPeer message:string entities:flags.3?Vector<MessageEntity> = Bool;
func (c *DraftsCore) MessagesSaveDraft(in *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	var (
		peer                = mtproto.FromInputPeer2(c.MD.UserId, in.Peer)
		draft               *mtproto.DraftMessage
		isDraftMessageEmpty = true
		date                = int32(time.Now().Unix())
	)

	if in.NoWebpage == true {
		isDraftMessageEmpty = false
	} else if in.ReplyToMsgId != nil {
		isDraftMessageEmpty = false
	} else if in.Message != "" {
		isDraftMessageEmpty = false
	} else if in.Entities != nil {
		isDraftMessageEmpty = false
	}

	if isDraftMessageEmpty {
		draft = mtproto.MakeTLDraftMessageEmpty(&mtproto.DraftMessage{
			Date_FLAGINT32: mtproto.MakeFlagsInt32(date),
		}).To_DraftMessage()

		c.svcCtx.Dao.DialogClient.DialogClearDraftMessage(c.ctx, &dialog.TLDialogClearDraftMessage{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
		})
	} else {
		draft = mtproto.MakeTLDraftMessage(&mtproto.DraftMessage{
			NoWebpage:    in.GetNoWebpage(),
			ReplyToMsgId: in.GetReplyToMsgId(),
			Message:      in.GetMessage(),
			Entities:     in.GetEntities(),
			Date_INT32:   date,
		}).To_DraftMessage()

		c.svcCtx.Dao.DialogClient.DialogSaveDraftMessage(c.ctx, &dialog.TLDialogSaveDraftMessage{
			UserId:   c.MD.UserId,
			PeerType: peer.PeerType,
			PeerId:   peer.PeerId,
			Message:  draft,
		})
	}

	syncUpdates := mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
		Peer_PEER: peer.ToPeer(),
		Draft:     draft,
	}).To_Update())

	switch peer.PeerType {
	case mtproto.PEER_SELF:
		user, _ := c.svcCtx.Dao.UserClient.UserGetImmutableUser(c.ctx, &userpb.TLUserGetImmutableUser{
			Id: c.MD.UserId,
		})
		if user != nil {
			syncUpdates.PushUser(user.ToSelfUser())
		}
	case mtproto.PEER_USER:
		users, _ := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
			Id: []int64{c.MD.UserId, peer.PeerId},
		})
		user, _ := users.GetUnsafeUser(c.MD.UserId, peer.PeerId)
		syncUpdates.AddSafeUser(user)
	case mtproto.PEER_CHAT:
		chat, _ := c.svcCtx.Dao.ChatClient.ChatGetMutableChat(c.ctx, &chatpb.TLChatGetMutableChat{
			ChatId: peer.PeerId,
		})
		syncUpdates.AddSafeChat(chat.ToUnsafeChat(c.MD.UserId))
	case mtproto.PEER_CHANNEL:
		if c.svcCtx.Plugin != nil {
			chats := c.svcCtx.Plugin.GetChannelListByIdList(c.ctx, c.MD.UserId, peer.PeerId)
			syncUpdates.PushChat(chats...)
		} else {
			c.Logger.Errorf("messages.saveDraft blocked, License key from https://elloapp.com required to unlock enterprise features.")
			return nil, mtproto.ErrEnterpriseIsBlocked
		}
	}

	// sync
	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
		UserId:    c.MD.UserId,
		AuthKeyId: c.MD.AuthId,
		Updates:   syncUpdates,
	})

	return mtproto.BoolTrue, nil
}
