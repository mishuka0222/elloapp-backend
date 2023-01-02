package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/messenger/msg/inbox/inbox"
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// InboxDeleteChatHistoryToInbox
// inbox.deleteChatHistoryToInbox from_id:long peer_chat_id:long max_id:int = Void;
func (c *InboxCore) InboxDeleteChatHistoryToInbox(in *inbox.TLInboxDeleteChatHistoryToInbox) (*mtproto.Void, error) {
	// TODO: not impl
	c.Logger.Errorf("inbox.deleteChatHistoryToInbox - error: method InboxDeleteChatHistoryToInbox not impl")

	return nil, mtproto.ErrMethodNotImpl
}
