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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ChatDeleteRevokedExportedChatInvites
// chat.deleteRevokedExportedChatInvites self_id:long chat_id:long admin_id:long = Bool;
func (c *ChatCore) ChatDeleteRevokedExportedChatInvites(in *chat.TLChatDeleteRevokedExportedChatInvites) (*mtproto.Bool, error) {
	_, err := c.svcCtx.Dao.ChatInvitesDAO.DeleteByRevoked(c.ctx, in.ChatId, in.AdminId)
	if err != nil {
		c.Logger.Errorf("chat.deleteRevokedExportedChatInvites - error: %v", err)
	}

	return mtproto.BoolTrue, nil
}
