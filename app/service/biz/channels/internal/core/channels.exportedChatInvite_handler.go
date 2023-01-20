package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/app/service/biz/channels/channels"
)

func (c *ChannelsCore) ExportedChatInvite(in *channels.ChannelCoreData) (*channels.ExportedChatInviteResp, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	return &channels.ExportedChatInviteResp{}, nil
}
