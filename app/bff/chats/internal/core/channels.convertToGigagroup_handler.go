package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

// ChannelsConvertToGigagroup
// channels.convertToGigagroup#b290c69 channel:InputChannel = Updates;
func (c *ChatsCore) ChannelsConvertToGigagroup(in *mtproto.TLChannelsConvertToGigagroup) (*mtproto.Updates, error) {
	// TODO: not impl
	c.Logger.Errorf("channels.convertToGigagroup blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return nil, mtproto.ErrEnterpriseIsBlocked
}
