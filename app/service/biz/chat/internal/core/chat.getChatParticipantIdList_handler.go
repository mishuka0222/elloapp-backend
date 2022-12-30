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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/chat"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/chat/internal/dal/dataobject"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ChatGetChatParticipantIdList
// chat.getChatParticipantIdList chat_id:long = Vector<long>;
func (c *ChatCore) ChatGetChatParticipantIdList(in *chat.TLChatGetChatParticipantIdList) (*chat.Vector_Long, error) {
	var (
		idList = make([]int64, 0)
	)

	c.svcCtx.Dao.ChatParticipantsDAO.SelectListWithCB(
		c.ctx,
		in.ChatId,
		func(i int, v *dataobject.ChatParticipantsDO) {
			if v.State != mtproto.ChatMemberStateNormal {
				return
			}
			idList = append(idList, v.UserId)
		})

	return &chat.Vector_Long{
		Datas: idList,
	}, nil
}
