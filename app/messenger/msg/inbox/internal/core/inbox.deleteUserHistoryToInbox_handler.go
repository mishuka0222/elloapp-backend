package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxDeleteUserHistoryToInbox
// inbox.deleteUserHistoryToInbox flags:# from_id:long peer_user_id:long just_clear:flags.1?true max_id:int = Void;
func (c *InboxCore) InboxDeleteUserHistoryToInbox(in *inbox.TLInboxDeleteUserHistoryToInbox) (*mtproto.Void, error) {
	var (
		pts, ptsCount int32
		peer          = mtproto.MakePeerUtil(mtproto.PEER_USER, in.FromId)
	)

	lastMessage, deleteIds := c.svcCtx.Dao.GetLastMessageAndIdListByDialog(c.ctx, in.PeerUserId, peer)
	if len(deleteIds) == 0 ||
		len(deleteIds) == 1 &&
			lastMessage.GetPredicateName() == mtproto.Predicate_messageService &&
			lastMessage.GetAction().GetPredicateName() == mtproto.Predicate_messageActionHistoryClear {
		return mtproto.EmptyVoid, nil
	}

	// TODO(@benqi): chat
	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.PeerUserId, len(deleteIds)+1)
	ptsCount = int32(len(deleteIds) + 1)

	if in.JustClear {
		deleteIds = deleteIds[1:]
		if _, err := c.svcCtx.Dao.DeleteByMessageIdList(c.ctx, in.PeerUserId, deleteIds); err != nil {
			return nil, err
		}

		clearHistoryMessage := mtproto.MakeTLMessageService(&mtproto.Message{
			Out:         true,
			Mentioned:   false,
			MediaUnread: false,
			Silent:      false,
			Post:        false,
			Legacy:      false,
			Id:          lastMessage.GetId(),
			FromId:      mtproto.MakePeerUser(in.PeerUserId),
			PeerId:      mtproto.MakePeerUser(in.FromId),
			ReplyTo:     nil,
			Date:        lastMessage.GetDate(),
			Action:      mtproto.MakeTLMessageActionHistoryClear(nil).To_MessageAction(),
			TtlPeriod:   nil,
		}).To_Message()

		c.svcCtx.Dao.EditUserOutboxMessage(c.ctx, in.PeerUserId, in.FromId, clearHistoryMessage)

		pushUpdates := mtproto.MakeUpdatesByUpdates(
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
		c.svcCtx.Dao.SyncClient.SyncPushUpdates(
			c.ctx,
			&sync.TLSyncPushUpdates{
				UserId:  in.PeerUserId,
				Updates: pushUpdates,
			})
	} else {
		if _, err := c.svcCtx.Dao.DeleteByMessageIdList(c.ctx, in.PeerUserId, deleteIds); err != nil {
			return nil, err
		}

		pushUpdates := mtproto.MakeUpdatesByUpdates(
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
		c.svcCtx.Dao.SyncClient.SyncPushUpdates(
			c.ctx,
			&sync.TLSyncPushUpdates{
				UserId:  in.PeerUserId,
				Updates: pushUpdates,
			})
	}

	return mtproto.EmptyVoid, nil
}
