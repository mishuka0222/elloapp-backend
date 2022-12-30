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
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessageUpdatePinnedMessageId
// message.updatePinnedMessageId user_id:long peer_type:int peer_id:long id:int pinned:Bool = Bool;
func (c *MessageCore) MessageUpdatePinnedMessageId(in *message.TLMessageUpdatePinnedMessageId) (*mtproto.Bool, error) {
	switch in.PeerType {
	case mtproto.PEER_SELF,
		mtproto.PEER_USER,
		mtproto.PEER_CHAT:
		_, err := c.svcCtx.Dao.MessagesDAO.UpdatePinned(c.ctx, mtproto.FromBool(in.Pinned), in.UserId, in.Id)
		if err != nil {
			c.Logger.Errorf("message.updatePinnedMessageId - error: %v", err)
		}
	case mtproto.PEER_CHANNEL:
		c.Logger.Errorf("message.updatePinnedMessageId blocked, License key from https://teamgram.net required to unlock enterprise features.")

		return nil, mtproto.ErrEnterpriseIsBlocked
	}

	return mtproto.BoolTrue, nil
}
