package core

import (
	"context"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// InboxDeleteMessagesToInbox
// inbox.deleteMessagesToInbox from_id:long id:Vector<int> = Void;
func (c *InboxCore) InboxDeleteMessagesToInbox(in *inbox.TLInboxDeleteMessagesToInbox) (*mtproto.Void, error) {
	c.svcCtx.Dao.DeleteInboxMessages(
		c.ctx,
		in.FromId,
		mtproto.MakePeerUtil(in.PeerType, in.PeerId),
		in.Id,
		func(ctx context.Context, userId int64, idList []int32) {
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(ctx, &sync.TLSyncPushUpdates{
				UserId: userId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
					Messages:  idList,
					Pts_INT32: c.svcCtx.Dao.IDGenClient2.NextNPtsId(ctx, userId, len(idList)),
					PtsCount:  int32(len(idList)),
				}).To_Update()),
			})
		})
	return mtproto.EmptyVoid, nil
}
