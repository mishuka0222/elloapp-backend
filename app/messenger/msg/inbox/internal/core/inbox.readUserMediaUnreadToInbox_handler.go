package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxReadUserMediaUnreadToInbox
// inbox.readUserMediaUnreadToInbox from_id:long id:Vector<int> = Void;
func (c *InboxCore) InboxReadUserMediaUnreadToInbox(in *inbox.TLInboxReadUserMediaUnreadToInbox) (*mtproto.Void, error) {
	idList := make([]int64, 0, len(in.GetId()))
	for _, id := range in.GetId() {
		idList = append(idList, id.DialogMessageId)
	}

	c.svcCtx.Dao.MessagesDAO.SelectByMessageDataIdListWithCB(
		c.ctx,
		c.svcCtx.Dao.MessagesDAO.CalcTableName(in.PeerUserId),
		idList,
		func(i int, v *dataobject.MessagesDO) {
			c.svcCtx.Dao.MessagesDAO.UpdateMediaUnread(c.ctx, v.UserId, v.UserMessageBoxId)

			// TODO: batch handle
			pts := c.svcCtx.Dao.IDGenClient2.NextPtsId(c.ctx, v.UserId)
			c.svcCtx.Dao.SyncClient.SyncPushUpdates(c.ctx, &sync.TLSyncPushUpdates{
				UserId: v.UserId,
				Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadMessagesContents(&mtproto.Update{
					Messages:  []int32{v.UserMessageBoxId},
					Pts_INT32: pts,
					PtsCount:  1,
				}).To_Update()),
			})
		},
	)

	return mtproto.EmptyVoid, nil
}
