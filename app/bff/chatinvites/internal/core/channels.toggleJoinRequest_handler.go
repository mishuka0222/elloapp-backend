package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChannelsToggleJoinRequest
// channels.toggleJoinRequest#4c2985b6 channel:InputChannel enabled:Bool = Updates;
func (c *ChatInvitesCore) ChannelsToggleJoinRequest(in *mtproto.TLChannelsToggleJoinRequest) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("channels.toggleJoinRequest blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
