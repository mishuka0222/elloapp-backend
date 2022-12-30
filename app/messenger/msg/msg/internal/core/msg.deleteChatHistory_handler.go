package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/msg/msg"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MsgDeleteChatHistory
// msg.deleteChatHistory chat_id:long delete_user_id:long = Bool;
func (c *MsgCore) MsgDeleteChatHistory(in *msg.TLMsgDeleteChatHistory) (*mtproto.Bool, error) {
	// TODO: not impl
	c.Logger.Errorf("msg.deleteChatHistory - error: method MsgDeleteChatHistory not impl")

	return nil, mtproto.ErrMethodNotImpl
}
