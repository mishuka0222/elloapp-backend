package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/dialog/dialog"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// DialogGetChannelMessageReadParticipants
// dialog.getChannelMessageReadParticipants user_id:long channel_id:long msg_id:int = Vector<long>;
func (c *DialogCore) DialogGetChannelMessageReadParticipants(in *dialog.TLDialogGetChannelMessageReadParticipants) (*dialog.Vector_Long, error) {
	idList, _ := c.svcCtx.Dao.DialogsDAO.SelectDialogsByGTReadInboxMaxId(
		c.ctx,
		mtproto.PEER_CHANNEL,
		in.ChannelId,
		in.MsgId,
		in.UserId)

	if idList == nil {
		idList = []int64{}
	}

	return &dialog.Vector_Long{
		Datas: idList,
	}, nil
}
