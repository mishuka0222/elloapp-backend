/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package core

import (
	"math"

	"github.com/teamgram/proto/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
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
