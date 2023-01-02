package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

// ChannelsViewSponsoredMessage
// channels.viewSponsoredMessage#beaedb94 channel:InputChannel random_id:bytes = Bool;
func (c *SponsoredMessagesCore) ChannelsViewSponsoredMessage(in *mtproto.TLChannelsViewSponsoredMessage) (*mtproto.Bool, error) {
	// TODO: not impl

	return mtproto.BoolTrue, nil
}
