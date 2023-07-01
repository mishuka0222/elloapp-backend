package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/messenger/sync/sync"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MsgDeletePhoneCallHistory
// msg.deletePhoneCallHistory user_id:long auth_key_id:long revoke:Bool = messages.AffectedFoundMessages;
func (c *MsgCore) MsgDeletePhoneCallHistory(in *msg.TLMsgDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	var (
		pts, ptsCount int32
		msgIdList     []int32
		msgDataIdList []int64
		err           error
	)

	if msgIdList, msgDataIdList, err = c.svcCtx.Dao.DeletePhoneCallHistory(c.ctx, in.UserId); err != nil {
		c.Logger.Errorf("DeleteMessages - %v", err)
		return nil, err
	}

	pts = c.svcCtx.Dao.IDGenClient2.NextNPtsId(c.ctx, in.UserId, len(msgDataIdList))
	ptsCount = int32(len(msgDataIdList))

	c.svcCtx.Dao.SyncClient.SyncUpdatesNotMe(
		c.ctx,
		&sync.TLSyncUpdatesNotMe{
			UserId:    in.UserId,
			AuthKeyId: in.AuthKeyId,
			Updates: mtproto.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteMessages(&mtproto.Update{
				Messages:  msgIdList,
				Pts_INT32: pts,
				PtsCount:  ptsCount,
			}).To_Update()),
		})

	if in.Revoke {
		c.svcCtx.Dao.InboxClient.InboxDeleteMessagesToInbox(
			c.ctx,
			&inbox.TLInboxDeleteMessagesToInbox{
				FromId: in.UserId,
				Id:     msgDataIdList,
			})
	}

	return mtproto.MakeTLMessagesAffectedFoundMessages(&mtproto.Messages_AffectedFoundMessages{
		Pts:      pts,
		PtsCount: ptsCount,
		Offset:   0,
		Messages: msgIdList,
	}).To_Messages_AffectedFoundMessages(), nil
}
