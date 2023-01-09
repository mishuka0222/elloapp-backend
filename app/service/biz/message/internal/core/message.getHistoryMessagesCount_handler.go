package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessageGetHistoryMessagesCount
// message.getHistoryMessagesCount user_id:long peer_type:int peer_id:long = Int32;
func (c *MessageCore) MessageGetHistoryMessagesCount(in *message.TLMessageGetHistoryMessagesCount) (*mtproto.Int32, error) {
	var (
		count    int
		dialogId = mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
	)

	switch in.PeerType {
	case mtproto.PEER_SELF, mtproto.PEER_USER, mtproto.PEER_CHAT:
		count = c.svcCtx.CommonDAO.CalcSize(
			c.ctx,
			c.svcCtx.Dao.MessagesDAO.CalcTableName(in.UserId),
			map[string]interface{}{
				"user_id":    in.UserId,
				"dialog_id1": dialogId.A,
				"dialog_id2": dialogId.B,
				"deleted":    0,
			})
	case mtproto.PEER_CHANNEL:
		count = c.svcCtx.Dao.CommonDAO.CalcSize(c.ctx, "channel_messages", map[string]interface{}{
			"channel_id": in.PeerId,
			"deleted":    0,
		})
	default:
		c.Logger.Errorf("invalid peer: (%d, %d, %d)", in.UserId, in.PeerType, in.PeerId)
	}

	return &mtproto.Int32{
		V: int32(count),
	}, nil
}
