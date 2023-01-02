package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesSearchSentMedia
// messages.searchSentMedia#107e31a0 q:string filter:MessagesFilter limit:int = messages.Messages;
func (c *MessagesCore) MessagesSearchSentMedia(in *mtproto.TLMessagesSearchSentMedia) (*mtproto.Messages_Messages, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.searchSentMedia blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
