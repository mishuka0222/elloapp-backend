package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
	"math"
)

// MessageSearchV2
// message.searchV2 user_id:long peer_type:int peer_id:long q:string from_id:long min_date:int max_date:int offset_id:int add_offset:int limit:int max_id:int min_id:int hash:long = Vector<MessageBox>;
func (c *MessageCore) MessageSearchV2(in *message.TLMessageSearchV2) (*message.Vector_MessageBox, error) {
	if in.FromId == 0 {
		return c.MessageSearch(&message.TLMessageSearch{
			UserId:   in.UserId,
			PeerType: in.PeerType,
			PeerId:   in.PeerId,
			Q:        in.Q,
			Offset:   in.OffsetId,
			Limit:    in.Limit,
		})
	}

	var (
		//offset  = in.Offset
		//q       = in.Q
		boxList []*mtproto.MessageBox
		//limit   = in.Limit
	)

	dialogId := mtproto.MakeDialogId(in.UserId, in.PeerType, in.PeerId)
	c.svcCtx.Dao.MessagesDAO.SelectBackwardBySendUserIdOffsetIdLimitWithCB(
		c.ctx,
		in.UserId,
		dialogId.A,
		dialogId.B,
		in.FromId,
		math.MaxInt32,
		in.Limit,
		func(i int, v *dataobject.MessagesDO) {
			boxList = append(boxList, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})

	if boxList == nil {
		boxList = []*mtproto.MessageBox{}
	}

	return &message.Vector_MessageBox{
		Datas: boxList,
	}, nil
}
