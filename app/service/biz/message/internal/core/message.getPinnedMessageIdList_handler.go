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
	"github.com/teamgram/proto/mtproto"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
)

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (c *MessageCore) MessageGetPinnedMessageIdList(in *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error) {
	// TODO: not impl
	c.Logger.Errorf("message.getPinnedMessageIdList - error: method MessageGetPinnedMessageIdList not impl")

	return nil, mtproto.ErrMethodNotImpl
}
