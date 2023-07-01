package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
	"math/rand"
)

// MsgUpdatePinnedMessage
// msg.updatePinnedMessage flags:# user_id:long auth_key_id:long silent:flags.0?true unpin:flags.1?true pm_oneside:flags.2?true peer_type:int peer_id:long id:int = Updates;
func (c *MsgCore) MsgUpdatePinnedMessage(in *msg.TLMsgUpdatePinnedMessage) (*mtproto.Updates, error) {
	var (
		peer     = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
		rUpdates *mtproto.Updates
	)

	switch peer.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		boxMsg, err := c.svcCtx.Dao.MessagesDAO.SelectByMessageId(c.ctx, in.UserId, in.Id)
		if err != nil {
			c.Logger.Errorf("msg.updatePinnedMessage - error: %v", err)
			return nil, mtproto.ErrMsgIdInvalid
		} else if boxMsg == nil {
			c.Logger.Errorf("msg.updatePinnedMessage - error: not found msg_id (%v, %v)", in.UserId, in.Id)
			return nil, mtproto.ErrMsgIdInvalid
		}

		if in.GetUnpin() {
			var (
				pinnedMsgId int32 = 0
			)
			idList, _ := c.svcCtx.Dao.MessagesDAO.SelectLastTwoPinnedList(c.ctx, in.UserId, boxMsg.DialogId1, boxMsg.DialogId2)
			if len(idList) == 2 {
				if in.Id == idList[0] {
					pinnedMsgId = idList[1]
				} else {
					pinnedMsgId = idList[0]
				}
			}
			c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, pinnedMsgId, in.UserId, mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId))
		} else {
			c.svcCtx.Dao.DialogsDAO.UpdatePinnedMsgId(c.ctx, in.Id, in.UserId, mtproto.MakePeerDialogId(peer.PeerType, peer.PeerId))
		}

		// pinned
		c.svcCtx.Dao.MessagesDAO.UpdatePinned(c.ctx, !in.GetUnpin(), in.UserId, in.Id)
		updatePinnedMessage := mtproto.MakeTLUpdatePinnedMessages(&mtproto.Update{
			Pinned:    !in.GetUnpin(),
			Peer_PEER: peer.ToPeer(),
			Messages:  []int32{in.Id},
			Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, in.UserId),
			PtsCount:  1,
		}).To_Update()

		if !in.GetUnpin() && !in.PmOneside && !peer.IsSelfUser(in.UserId) {
			rUpdates, err = c.MsgSendMessage(&msg.TLMsgSendMessage{
				UserId:    in.UserId,
				AuthKeyId: in.AuthKeyId,
				PeerType:  peer.PeerType,
				PeerId:    peer.PeerId,
				Message: msg.MakeTLOutboxMessage(&msg.OutboxMessage{
					NoWebpage:    false,
					Background:   false,
					RandomId:     rand.Int63(),
					Message:      mtproto.MakePinnedMessageService(in.GetSilent(), in.UserId, peer, in.GetId()),
					ScheduleDate: nil,
				}).To_OutboxMessage(),
			})
			if err != nil {
				c.Logger.Errorf("msg.updatePinnedMessage - error: err", err)
				return nil, err
			}

			// push
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(
				c.ctx,
				&sync.TLSyncPushUpdates{
					UserId:  in.UserId,
					Updates: mtproto.MakeUpdatesByUpdates(updatePinnedMessage),
				})
		} else {
			rUpdates = mtproto.MakeUpdatesByUpdates(updatePinnedMessage)

			// sync
			c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
				c.ctx,
				&sync.TLSyncUpdatesNotMe{
					UserId:    in.UserId,
					AuthKeyId: in.AuthKeyId,
					Updates:   rUpdates,
				})
		}

		if !in.PmOneside && !peer.IsSelfUser(in.UserId) {
			c.svcCtx.Dao.InboxClient.InboxUpdatePinnedMessage(c.ctx, &inbox.TLInboxUpdatePinnedMessage{
				UserId:          in.UserId,
				Unpin:           in.Unpin,
				PeerType:        in.PeerType,
				PeerId:          in.PeerId,
				Id:              in.Id,
				DialogMessageId: boxMsg.DialogMessageId,
			})
		}

	case mtproto.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.updatePinnedMessage - error: %v", err)
		return nil, err
	}

	return rUpdates, nil
}
