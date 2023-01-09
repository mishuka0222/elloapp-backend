package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MsgDeleteHistory
// msg.deleteHistory user_id:long auth_key_id:long peer_type:int peer_id:long just_clear:Bool revoke:Bool = messages.AffectedHistory;
func (c *MsgCore) MsgDeleteHistory(in *msg.TLMsgDeleteHistory) (*mtproto.Messages_AffectedHistory, error) {
	var (
		rValue *mtproto.Messages_AffectedHistory
		err    error
	)

	switch in.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		rValue, err = c.deleteUserHistory(in)

	case mtproto.PEER_CHANNEL:
	default:
		err = mtproto.ErrPeerIdInvalid
		c.Logger.Errorf("msg.deleteHistory - error: %v", err)
		return nil, err
	}

	return rValue, nil
}

func (c *MsgCore) deleteUserHistory(in *msg.TLMsgDeleteHistory) (reply *mtproto.Messages_AffectedHistory, err error) {
	var (
		pts, ptsCount int32
		peer          = mtproto.MakePeerUtil(in.PeerType, in.PeerId)
	)

	lastMessage, deleteIds := c.svcCtx.Dao.GetLastMessageAndIdListByDialog(c.ctx, in.UserId, peer)
	if len(deleteIds) == 0 ||
		len(deleteIds) == 1 &&
			in.GetJustClear() &&
			lastMessage.GetPredicateName() == mtproto.Predicate_messageService &&
			lastMessage.GetAction().GetPredicateName() == mtproto.Predicate_messageActionHistoryClear {
		pts = c.svcCtx.Dao.IDGenClient2.CurrentPtsId(c.ctx, in.UserId)
		ptsCount = 0

		return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
			Pts:      pts,
			PtsCount: ptsCount,
			Offset:   0,
		}).To_Messages_AffectedHistory(), nil
	}

	// TODO(@benqi): chat
	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, len(deleteIds)+1)
	ptsCount = int32(len(deleteIds) + 1)
	if in.JustClear {
		deleteIds = deleteIds[1:]
		if _, err = c.svcCtx.Dao.DeleteByMessageIdList(c.ctx, in.UserId, deleteIds); err != nil {
			return nil, err
		}

		clearHistoryMessage := mtproto.MakeTLMessageService(&mtproto.Message{
			Out:         true,
			Mentioned:   false,
			MediaUnread: false,
			Silent:      false,
			Post:        false,
			Legacy:      false,
			Id:          lastMessage.Id,
			FromId:      mtproto.MakePeerUser(in.UserId),
			PeerId:      peer.ToPeer(),
			ReplyTo:     nil,
			Date:        lastMessage.GetDate(),
			Action:      mtproto.MakeTLMessageActionHistoryClear(nil).To_MessageAction(),
			TtlPeriod:   nil,
		}).To_Message()
		c.svcCtx.Dao.EditUserOutboxMessage(c.ctx, in.UserId, in.PeerId, clearHistoryMessage)

		syncUpdates := mtproto.MakeUpdatesByUpdates(
			mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  deleteIds,
				Pts_INT32: pts - 2,
				PtsCount:  ptsCount - 2,
			}).To_Update(),
			mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
				Peer_PEER: peer.ToPeer(),
				MaxId:     lastMessage.Id,
				Pts_INT32: pts - 1,
				PtsCount:  1,
			}).To_Update(),
			mtproto.MakeTLUpdateEditMessage(&mtproto.Update{
				Message_MESSAGE: clearHistoryMessage,
				Pts_INT32:       pts,
				PtsCount:        1,
			}).To_Update(),
		)

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
			c.ctx,
			&sync.TLSyncUpdatesNotMe{
				UserId:    in.UserId,
				AuthKeyId: in.AuthKeyId,
				Updates:   syncUpdates,
			})
	} else {
		if _, err = c.svcCtx.Dao.DeleteByMessageIdList(c.ctx, in.UserId, deleteIds); err != nil {
			return nil, err
		}
		// s.PrivateFacade

		syncUpdates := mtproto.MakeUpdatesByUpdates(
			mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  deleteIds,
				Pts_INT32: pts - 2,
				PtsCount:  ptsCount - 2,
			}).To_Update(),
			mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
				Peer_PEER: peer.ToPeer(),
				MaxId:     lastMessage.Id,
				Pts_INT32: pts - 1,
				PtsCount:  1,
			}).To_Update(),
		)

		c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
			c.ctx,
			&sync.TLSyncUpdatesNotMe{
				UserId:    in.UserId,
				AuthKeyId: in.AuthKeyId,
				Updates:   syncUpdates,
			})
	}

	reply = mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
	}).To_Messages_AffectedHistory()

	if in.Revoke {
		switch peer.PeerType {
		case mtproto.PEER_USER:
			if in.UserId == in.PeerId {
				return
			}

			c.svcCtx.Dao.InboxClient.InboxDeleteUserHistoryToInbox(
				c.ctx,
				&inbox.TLInboxDeleteUserHistoryToInbox{
					FromId:     in.UserId,
					PeerUserId: in.PeerId,
					JustClear:  in.JustClear,
					MaxId:      in.MaxId,
				})
		case mtproto.PEER_CHAT:
			c.svcCtx.Dao.InboxClient.InboxDeleteChatHistoryToInbox(
				c.ctx,
				&inbox.TLInboxDeleteChatHistoryToInbox{
					FromId:     in.UserId,
					PeerChatId: in.PeerId,
					MaxId:      in.MaxId,
				})
		}
	}

	return
}
