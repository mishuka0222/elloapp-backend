package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// MessageGetPinnedMessageIdList
// message.getPinnedMessageIdList user_id:long peer_type:int peer_id:long = Vector<int>;
func (c *MessageCore) MessageGetPinnedMessageIdList(in *message.TLMessageGetPinnedMessageIdList) (*message.Vector_Int, error) {
	// TODO: not impl
	c.Logger.Errorf("message.getPinnedMessageIdList - error: method MessageGetPinnedMessageIdList not impl")

	return nil, mtproto.ErrMethodNotImpl
}
