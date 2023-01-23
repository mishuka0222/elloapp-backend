package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

// channel_util
func (c *ChannelsCore) GetChannelListBySelfAndIDList(in *channels.GetChannelListBySelfAndIDListReq) (*channels.GetChannelListBySelfAndIDListResp, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &channels.GetChannelListBySelfAndIDListResp{}, nil
}
