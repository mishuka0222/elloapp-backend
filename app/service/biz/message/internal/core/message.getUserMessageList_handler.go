package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessageGetUserMessageList
// message.getUserMessageList user_id:long id_list:Vector<int> = Vector<MessageBox>;
func (c *MessageCore) MessageGetUserMessageList(in *message.TLMessageGetUserMessageList) (*message.Vector_MessageBox, error) {
	rValueList := &message.Vector_MessageBox{
		Datas: make([]*mtproto.MessageBox, 0, len(in.IdList)),
	}

	c.svcCtx.Dao.MessagesDAO.SelectByMessageIdListWithCB(
		c.ctx,
		in.UserId,
		in.IdList,
		func(i int, v *dataobject.MessagesDO) {
			rValueList.Datas = append(rValueList.GetDatas(), c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})

	return rValueList, nil
}
