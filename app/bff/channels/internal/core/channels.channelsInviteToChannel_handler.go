package core

import (
	"errors"

	"gitlab.com/merehead/elloapp/backend/elloapp_tg_backend/mtproto"
)

func (c *ChannelsCore) ChannelsInviteToChannel(in *mtproto.TLChannelsInviteToChannel) (*mtproto.Updates, error) {
	// todo: add your logic here and delete this line
	return nil, errors.New("Unimplemented")
	// return &mtproto.Updates{}, nil
}
