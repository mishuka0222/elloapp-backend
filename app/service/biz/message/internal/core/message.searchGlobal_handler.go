package core

import (
	"math"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessageSearchGlobal
// message.searchGlobal user_id:long q:string offset:int limit:int = Vector<MessageBox>;
func (c *MessageCore) MessageSearchGlobal(in *message.TLMessageSearchGlobal) (*message.Vector_MessageBox, error) {
	var (
		offset  = in.Offset
		rValues []*mtproto.MessageBox
	)

	if offset == 0 {
		offset = math.MaxInt32
	}

	rList, _ := c.svcCtx.Dao.MessagesDAO.SearchGlobalWithCB(
		c.ctx,
		in.UserId,
		offset, "%"+in.Q+"%",
		in.Limit,
		func(i int, v *dataobject.MessagesDO) {
			rValues = append(rValues, c.svcCtx.Dao.MakeMessageBox(c.ctx, in.UserId, v))
		})
	_ = rList

	if rValues == nil {
		rValues = []*mtproto.MessageBox{}
	}

	return &message.Vector_MessageBox{
		Datas: rValues,
	}, nil
}
