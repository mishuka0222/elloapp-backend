package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/message/message"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessageGetPeerUserMessageId
// message.getPeerUserMessageId user_id:long peer_user_id:long msg_id:int = Int32;
func (c *MessageCore) MessageGetPeerUserMessageId(in *message.TLMessageGetPeerUserMessageId) (*mtproto.Int32, error) {
	// TODO: not impl
	c.Logger.Errorf("message.getPeerUserMessageId - error: method MessageGetPeerUserMessageId not impl")

	return nil, mtproto.ErrMethodNotImpl
}
