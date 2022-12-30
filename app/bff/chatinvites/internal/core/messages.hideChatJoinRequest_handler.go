package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// MessagesHideChatJoinRequest
// messages.hideChatJoinRequest#7fe7e815 flags:# approved:flags.0?true peer:InputPeer user_id:InputUser = Updates;
func (c *ChatInvitesCore) MessagesHideChatJoinRequest(in *mtproto.TLMessagesHideChatJoinRequest) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("messages.hideChatJoinRequest blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
