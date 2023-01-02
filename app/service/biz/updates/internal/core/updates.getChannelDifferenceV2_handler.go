package core

import (
	"gitlab.com/merehead/elloapp/backend/elloapp_backend/app/service/biz/updates/updates"
)

// UpdatesGetChannelDifferenceV2
// updates.getChannelDifferenceV2 auth_key_id:long user_id:long channel_id:long pts:int limit:int = ChannelDifference;
func (c *UpdatesCore) UpdatesGetChannelDifferenceV2(in *updates.TLUpdatesGetChannelDifferenceV2) (*updates.ChannelDifference, error) {
	c.Logger.Errorf("updates.getChannelDifferenceV2 blocked, License key from https://elloapp.com required to unlock enterprise features.")

	return updates.MakeTLChannelDifference(&updates.ChannelDifference{
		Final:        false,
		Pts:          in.Pts,
		NewMessages:  nil,
		OtherUpdates: nil,
	}).To_ChannelDifference(), nil
}
