package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxSendUserMultiMessageToInbox
// inbox.sendUserMultiMessageToInbox from_id:long peer_user_id:long message:Vector<InboxMessageData> = Void;
func (c *InboxCore) InboxSendUserMultiMessageToInbox(in *inbox.TLInboxSendUserMultiMessageToInbox) (*mtproto.Void, error) {
	if in.FromId == in.PeerUserId {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: sendToSelfUser")
		err := mtproto.ErrPeerIdInvalid
		return nil, err
	}

	inBoxList, err := c.svcCtx.Dao.SendUserMultiMessageToInbox(c.ctx,
		in.FromId,
		in.PeerUserId,
		in.Message)
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	pushUpdates := c.makeUpdateNewMessageListUpdates(in.PeerUserId, inBoxList...)

	_, err = c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
		UserId:  in.PeerUserId,
		Updates: pushUpdates,
	})
	if err != nil {
		c.Logger.Errorf("inbox.sendUserMultiMessageToInbox - error: %v", err)
		return nil, err
	}

	return mtproto.EmptyVoid, nil
}
